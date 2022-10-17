package helper

import "dysn/notify/internal/model/consts"

var LangList = []string{
	consts.LangRu,
	consts.LangEn,
}

func GetDefaultLang() string {
	return consts.LangEn
}
