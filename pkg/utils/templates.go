package utils

import (
	"html/template"
	"io"
	"os"
	"path/filepath"
)

// RenderTemplateToWriter will render a template to any io.Writer
func RenderTemplateToWriter(
	writer io.Writer, fp *FileProperties,
	data interface{}, delims []string,
) error {
	fullPath := filepath.Join(os.ExpandEnv(fp.FilePath), fp.FileName)
	tpl, err := template.New(fp.FileName).Delims(delims[0], delims[1]).ParseFiles(fullPath)
	if err != nil {
		return err
	}
	return tpl.ExecuteTemplate(writer, fp.FileName, data)
}
