package template

import (
	"bytes"
	"html/template"
)

type TemplateInterface interface {
	Parse(fileName string, data interface{}) (string, error)
}

type Template struct {
	basePath string
}

func NewTemplate(basePath string) *Template {
	return &Template{
		basePath: basePath,
	}
}

func (t *Template) Parse(fileName string, data interface{}) (string, error) {
	path, err := t.getFilePath(fileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = path.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (t *Template) getFilePath(fileName string) (*template.Template, error) {
	basePath := t.basePath + fileName

	return template.ParseFiles(basePath + ".html")
}
