package server

import (
	pb "dysn/notify/internal/transport/grpc/pb/notify"
	"dysn/notify/pkg/i18n"
	"dysn/notify/pkg/log"
	"dysn/notify/pkg/sender"
	"dysn/notify/pkg/template"
	"google.golang.org/grpc"
	"net"
)

type TransportInterface interface {
	StartServer()
	StopServer()
}

type Grpc struct {
	server *grpc.Server
	port   string
	logger *log.Logger
}

func NewGrpc(
	port string,
	authSrv AuthServiceInterface,
	sender sender.Sender,
	template template.TemplateInterface,
	i18n *i18n.I18n,
	logger *log.Logger,
) *Grpc {
	srv := grpc.NewServer()
	notifyServer := NewNotify(sender, template, i18n, authSrv, logger)
	pb.RegisterNotifyServer(srv, notifyServer)

	return &Grpc{server: srv, port: port, logger: logger}
}

func (g *Grpc) StartServer() {
	g.logger.InfoLog.Println("Server delivery starting...")

	connection, err := net.Listen("tcp", g.port)
	if err != nil {
		g.logger.ErrorLog.Panic(err)
	}

	err = g.server.Serve(connection)
	if err != nil {
		g.logger.ErrorLog.Panic(err)
	}
}

func (g *Grpc) StopServer() {
	g.logger.InfoLog.Println("Server delivery stopping...")
	g.server.Stop()
}
