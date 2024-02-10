package dto

type Code struct {
	Email string `json:"email"`
	Code  string `json:"code"`
	Lang  string `json:"lang"`
}

func NewCode(email, code, lang string) Code {
	return Code{
		Email: email,
		Code:  code,
		Lang:  lang,
	}
}
