package app

import (
	"context"
	"dysn/notify/config"
	"dysn/notify/internal/consumers"
	"dysn/notify/internal/helper"
	"dysn/notify/internal/service"
	"dysn/notify/internal/transport/grpc/server"
	"dysn/notify/pkg/i18n"
	"dysn/notify/pkg/log"
	"dysn/notify/pkg/sender"
	temp "dysn/notify/pkg/template"
	"github.com/segmentio/kafka-go"
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

	authSrv := service.NewAuthService(i18nPkg, template, logger, mailSender)

	registerReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{cfg.GetKafkaBroker1()},
		GroupID: "user-register-group-id",
		Topic:   cfg.GetTopicUserRegister(),
	})

	consumerRegister := consumers.NewConsumerRegister(registerReader, authSrv, logger)
	go consumerRegister.ReadRegisters(ctx)

	srv := server.NewGrpc(cfg.GetGrpcPort(), authSrv, mailSender, template, i18nPkg, logger)
	go srv.StartServer()
	defer srv.StopServer()

	sgn := make(chan os.Signal, 1)
	signal.Notify(sgn, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
	case <-sgn:
	}
}
