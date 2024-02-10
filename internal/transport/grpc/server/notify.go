package server

import (
	"context"
	"dysn/notify/internal/model"
	"dysn/notify/internal/model/dto"
	pb "dysn/notify/internal/transport/grpc/pb/notify"
	"dysn/notify/pkg/i18n"
	"dysn/notify/pkg/log"
	"dysn/notify/pkg/sender"
	"dysn/notify/pkg/template"
	"github.com/golang/protobuf/ptypes/empty"
)

type AuthServiceInterface interface {
	ConfirmRegister(ctx context.Context, code dto.Code) error
}
type Notify struct {
	sender   sender.Sender
	template template.TemplateInterface
	logger   *log.Logger
	i18n     *i18n.I18n
	authSrv  AuthServiceInterface
	pb.UnimplementedNotifyServer
}

func NewNotify(sender sender.Sender,
	template template.TemplateInterface,
	i18n *i18n.I18n,
	authSrv AuthServiceInterface,
	logger *log.Logger) *Notify {
	return &Notify{
		sender:   sender,
		template: template,
		logger:   logger,
		i18n:     i18n,
		authSrv:  authSrv,
	}
}

func (n *Notify) Ping(_ context.Context, _ *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (n *Notify) ConfirmRegister(ctx context.Context, code *pb.EmailWithCode) (*empty.Empty, error) {
	codeDto := dto.NewCode(code.GetEmail(), code.GetCode(), code.GetLang())
	if err := n.authSrv.ConfirmRegister(ctx, codeDto); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (n *Notify) DisableGa(_ context.Context, code *pb.EmailWithCode) (*empty.Empty, error) {
	codeDto := dto.NewCode(code.GetEmail(), code.GetCode(), code.GetLang())
	subject := n.i18n.T("disableGaSubject", nil, code.GetLang())
	templateData := model.NewCodeLocale(codeDto.Code,
		subject,
		n.i18n.T("disableGaDescription", nil, code.GetLang()))

	confirmTemplate, err := n.template.Parse("mail_template", templateData)
	if err != nil {
		n.logger.ErrorLog.Println("err disable ga template: ", err)

		return nil, err
	}

	if err = n.sender.Send(subject,
		confirmTemplate,
		codeDto.Email); err != nil {
		n.logger.ErrorLog.Println("err disable ga send: ", err)

		return nil, err
	}

	return &empty.Empty{}, nil
}

func (n *Notify) RecoveryPassword(_ context.Context, code *pb.EmailWithCode) (*empty.Empty, error) {
	codeDto := dto.NewCode(code.GetEmail(), code.GetCode(), code.GetLang())
	subject := n.i18n.T("recoveryPasswordSubject", nil, code.GetLang())
	templateData := model.NewCodeLocale(codeDto.Code,
		subject,
		n.i18n.T("recoveryPasswordDescription", nil, code.GetLang()))
	confirmTemplate, err := n.template.Parse("mail_template", templateData)
	if err != nil {
		n.logger.ErrorLog.Println("err recovery password template: ", err)

		return nil, err
	}

	if err = n.sender.Send(subject,
		confirmTemplate,
		codeDto.Email); err != nil {
		n.logger.ErrorLog.Println("err recovery password send: ", err)

		return nil, err
	}

	return &empty.Empty{}, nil
}
