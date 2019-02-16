package utils

import (
	"html/template"
	"io"
	"path/filepath"
)

// RenderTemplateToWriter will render a template to any io.Writer
func RenderTemplateToWriter(
	writer io.Writer, templatePath string, templateName string,
	data interface{}, delims []string,
) error {
	fp := filepath.Join(templatePath, templateName)
	tpl, err := template.New(templateName).Delims(delims[0], delims[1]).ParseFiles(fp)
	if err != nil {
		return err
	}
	return tpl.ExecuteTemplate(writer, templateName, data)
}
