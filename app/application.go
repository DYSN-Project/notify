package app

import (
	"context"
	"dysn/notify/config"
	"dysn/notify/internal/helper"
	"dysn/notify/internal/transport/grpc/server"
	"dysn/notify/pkg/i18n"
	"dysn/notify/pkg/log"
	"dysn/notify/pkg/sender"
	temp "dysn/notify/pkg/template"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx context.Context) {
	cfg := config.NewConfig()
	logger := log.NewLogger()

	mailSender := sender.NewGoMail(cfg.GetSmtpHost(), cfg.GetSmtpPort(), cfg.GetFrom(), cfg.GetSmtpPassword())
	template := temp.NewTemplate(cfg.GetTemplatePath())

	i18nPkg := i18n.NewI18n("locale/active", helper.LangList)

	srv := server.NewGrpc(cfg.GetGrpcPort(), mailSender, template, i18nPkg, logger)
	go srv.StartServer()
	defer srv.StopServer()

	sgn := make(chan os.Signal, 1)
	signal.Notify(sgn, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
	case <-sgn:
	}
}
