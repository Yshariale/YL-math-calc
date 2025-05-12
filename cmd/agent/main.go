package main

import (
	"context"
	"github.com/Yshariale/FinalTaskFirstSprint/internal/agent/server"
	"github.com/Yshariale/FinalTaskFirstSprint/internal/config"
	"github.com/Yshariale/FinalTaskFirstSprint/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	ctx, _ = logger.New(ctx)
	cfg, err := config.NewConfig()
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Error("error loading config: %v", zap.Error(err))
		return
	}
	agent := server.NewAgent(ctx, cfg)
	logger.GetLoggerFromCtx(ctx).Info("Agent started")
	agent.Run()
}
