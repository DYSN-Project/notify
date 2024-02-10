package service

import (
	"context"
	"dysn/notify/internal/model"
	"dysn/notify/internal/model/dto"
	"dysn/notify/pkg/i18n"
	"dysn/notify/pkg/log"
	"dysn/notify/pkg/sender"
	"dysn/notify/pkg/template"
)

type AuthService struct {
	i18n     *i18n.I18n
	template *template.Template
	logger   *log.Logger
	sender   sender.Sender
}

func NewAuthService(i18n *i18n.I18n,
	template *template.Template,
	logger *log.Logger,
	sender sender.Sender) *AuthService {
	return &AuthService{i18n: i18n,
		template: template,
		logger:   logger,
		sender:   sender,
	}
}

func (a *AuthService) ConfirmRegister(ctx context.Context, code dto.Code) error {
	subject := a.i18n.T("registerSubject", nil, code.Lang)
	templateData := model.NewCodeLocale(code.Code,
		subject,
		a.i18n.T("registerDescription", nil, code.Lang))

	confirmTemplate, err := a.template.Parse("mail_template", templateData)
	if err != nil {
		a.logger.ErrorLog.Println("err confirm register template: ", err)

		return err
	}

	if err = a.sender.Send(subject,
		confirmTemplate,
		code.Email); err != nil {
		a.logger.ErrorLog.Println("err confirm register send: ", err)

		return err
	}

	return nil
}
