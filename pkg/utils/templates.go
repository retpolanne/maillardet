package utils

import (
	"html/template"
	"io"
	"os"
	"path/filepath"
)

// RenderTemplateToWriter will render a template to any io.Writer
// TODO - create an interface so that all data will contain path, filename and delims
// TODO - also, should keep delims within the utils package
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
