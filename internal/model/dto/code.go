package dto

type Code struct {
	Email string
	Code  string
	Lang  string
}

func NewCode(email, code, lang string) *Code {
	return &Code{
		Email: email,
		Code:  code,
		Lang:  lang,
	}
}
