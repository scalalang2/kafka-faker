package main

import (
	"bytes"
	"fmt"
	"html/template"
)

type Template string

func ParseTemplate(tpl Template) (string, error) {
	tmpl, err := template.New("tpl").Funcs(TemplateFunctions).Parse(string(tpl))
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %s, err: %w", tpl, err)
	}

	// create bytes buffer
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, nil)
	if err != nil {
		return "", fmt.Errorf("error occurred while executing template. err: %w", err)
	}

	return buf.String(), nil
}
