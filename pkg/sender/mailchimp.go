package sender

import (
	"fmt"
	"github.com/mattbaird/gochimp"
)

type Mailchimp struct {
	mandrill *gochimp.MandrillAPI
}

func NewMailchimp(apiKey string) *Mailchimp {
	mandrillApi, err := gochimp.NewMandrill(apiKey)

	if err != nil {
		panic("Error instantiating client: " + err.Error())
	}

	return &Mailchimp{
		mandrill: mandrillApi,
	}
}

func (m *Mailchimp) Send(subject, body string, recipients ...string) error {
	to := m.getRecipients(recipients)
	message := gochimp.Message{
		Html:      body,
		Subject:   subject,
		FromEmail: "person@place.com",
		FromName:  "Boss Man",
		To:        to,
	}

	_, err := m.mandrill.MessageSend(message, false)

	if err != nil {
		return err
	}

	return nil
}

func (m *Mailchimp) getTemplate() string {
	templateName := "welcome email"
	contentVar := gochimp.Var{"main", "<h1>Welcome aboard!</h1>"}
	content := []gochimp.Var{contentVar}
	_, err := m.mandrill.TemplateAdd(templateName, fmt.Sprintf("%s", contentVar.Content), true)
	if err != nil {
		fmt.Println("Error adding template: %v", err)
		return ""
	}
	defer m.mandrill.TemplateDelete(templateName)
	renderedTemplate, err := m.mandrill.TemplateRender(templateName, content, nil)

	if err != nil {
		fmt.Println("Error rendering template: %v", err)
		return ""
	}

	return renderedTemplate

}

func (m *Mailchimp) getRecipients(recipients []string) []gochimp.Recipient {
	result := make([]gochimp.Recipient, len(recipients))
	for i := 0; i < len(recipients); i++ {
		recipient := gochimp.Recipient{Email: recipients[i]}
		result[i] = recipient
	}

	return result
}
