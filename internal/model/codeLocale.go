package model

type CodeLocale struct {
	Code        string
	Subject     string
	Description string
}

func NewCodeLocale(code, subject, description string) *CodeLocale {
	return &CodeLocale{
		Code:        code,
		Subject:     subject,
		Description: description,
	}
}
