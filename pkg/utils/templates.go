package utils

import (
	"html/template"
	"io"
	"os"
	"path/filepath"
)

// Mappings is a struct for holding the templates mappings
type Mappings struct {
	TemplatePath string
	Delims       []string    `yaml:"delims"`
	BaseTexFile  string      `yaml:"baseTexFile"`
	Pretextual   *Pretextual `yaml:"pretextual"`
	Textual      *Textual    `yaml:"textual"`
	Postextual   *Postextual `yaml:"postextual"`
}

// Pretextual is a struct that holds pretextual mappings
type Pretextual struct {
	Path       string `yaml:"path"`
	AbstractEn string `yaml:"abstractEn"`
	AbstractPt string `yaml:"abstractPt"`
	Cover      string `yaml:"cover"`
	ToC        string `yaml:"toc"`
}

// Textual is a struct that holds textual mappings
type Textual struct {
	Path    string `yaml:"path"`
	Section string `yaml:"section"`
	Chapter string `yaml:"chapter"`
}

// Postextual is a struct that holds postextual mappings
type Postextual struct {
	Path         string `yaml:"path"`
	Bibliography string `yaml:"bibliography"`
}

// TemplateInfo is a struct that holds information about the template to be rendered
type TemplateInfo struct {
	Delims       []string
	TemplatePath string
	Path         string
	FileName     string
}

// RenderTemplateToWriter will render a template to any io.Writer
func (ti *TemplateInfo) RenderTemplateToWriter(
	writer io.Writer, data interface{},
) error {
	fullPath := filepath.Join(os.ExpandEnv(ti.TemplatePath), ti.Path, ti.FileName)
	tpl, err := template.New(ti.FileName).Delims(ti.Delims[0], ti.Delims[1]).ParseFiles(fullPath)
	if err != nil {
		return err
	}
	return tpl.ExecuteTemplate(writer, ti.FileName, data)
}
