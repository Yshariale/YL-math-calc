package app

import (
	"context"
	grpcapp "github.com/Yshariale/FinalTaskFirstSprint/internal/app/grpc"
	orchestratorapp "github.com/Yshariale/FinalTaskFirstSprint/internal/app/orchestrator"
	"github.com/Yshariale/FinalTaskFirstSprint/internal/config"
	"github.com/Yshariale/FinalTaskFirstSprint/internal/orchestrator/service"
	"github.com/Yshariale/FinalTaskFirstSprint/internal/orchestrator/storage/sqlite"
	"github.com/Yshariale/FinalTaskFirstSprint/internal/orchestrator/transport/http"
	"github.com/Yshariale/FinalTaskFirstSprint/pkg/logger"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type App struct {
	gRPCServer   *grpcapp.App
	Orchestrator *orchestratorapp.App
	ctx          context.Context
	wg           sync.WaitGroup
	cancel       context.CancelFunc
}

func New(
	cfg *config.Config,
	ctx context.Context,
) *App {
	storage, err := sqlite.New("./storage/expr.db")
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Error("error creating storage: %v", zap.Error(err))
		panic(err)
	}
	srv := service.NewService(ctx, storage)
	orchestrator := http.New(ctx, srv, cfg)
	orchestratorApp := orchestratorapp.New(orchestrator)
	grpcApp := grpcapp.New(cfg, srv, ctx)
	return &App{
		gRPCServer:   grpcApp,
		Orchestrator: orchestratorApp,
		ctx:          ctx,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	errCh := make(chan error, 2)
	a.wg.Add(1)
	go func() {
		logger.GetLoggerFromCtx(a.ctx).Info("Orchestrator started")
		defer a.wg.Done()
		if err := a.Orchestrator.Run(); err != nil {
			errCh <- err
			a.cancel()
		}
	}()
	a.wg.Add(1)
	go func() {
		defer a.wg.Done()
		logger.GetLoggerFromCtx(a.ctx).Info("gRPC server started")
		if err := a.gRPCServer.Run(); err != nil {
			errCh <- err
			a.cancel()
		}
	}()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	select {
	case err := <-errCh:
		logger.GetLoggerFromCtx(a.ctx).Error("error running app", zap.Error(err))
		return err
	case sig := <-sigCh:
		logger.GetLoggerFromCtx(a.ctx).Info("received signal", zap.String("signal", sig.String()))
		a.Stop()
	case <-a.ctx.Done():
		logger.GetLoggerFromCtx(a.ctx).Info("context done")
	}

	return nil
}

func (a *App) Stop() {
	logger.GetLoggerFromCtx(a.ctx).Info("stopping app")
	a.cancel()
	a.gRPCServer.GRPCSrv.GracefulStop()
	a.wg.Wait()
	logger.GetLoggerFromCtx(a.ctx).Info("app stopped")
}
