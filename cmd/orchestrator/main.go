package main

import (
	"context"
	"github.com/Yshariale/FinalTaskFirstSprint/internal/app"
	"github.com/Yshariale/FinalTaskFirstSprint/internal/config"
	"github.com/Yshariale/FinalTaskFirstSprint/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cfg, err := config.NewConfig()
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Error("error loading config: %v", zap.Error(err))
		return
	}
	ctx, _ = logger.New(ctx)
	application := app.New(cfg, ctx)
	application.MustRun()
}
