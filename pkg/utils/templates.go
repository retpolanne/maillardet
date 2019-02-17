package utils

import (
	"html/template"
	"io"
	"os"
	"path/filepath"
)

// RenderTemplateToWriter will render a template to any io.Writer
func RenderTemplateToWriter(
	writer io.Writer, filePath string, fileName string,
	data interface{}, delims []string,
) error {
	fullPath := filepath.Join(os.ExpandEnv(filePath), fileName)
	tpl, err := template.New(fileName).Delims(delims[0], delims[1]).ParseFiles(fullPath)
	if err != nil {
		return err
	}
	return tpl.ExecuteTemplate(writer, fileName, data)
}
