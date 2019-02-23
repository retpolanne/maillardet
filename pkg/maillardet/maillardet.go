package maillardet

import (
	"github.com/vinicyusmacedo/maillardet/pkg/tex"
)

// Document is the yaml file struct
type Document struct {
	Template     string        `yaml:"template"`
	TemplatePath string        `yaml:"templatePath"`
	Cover        *tex.Cover    `yaml:"cover"`
	Contents     *tex.Contents `yaml:"contents"`
}

// Maillardet is the main facade. It will take care of:
// 1 - Interpreting the document yaml in the provided path;
// 2 - Finding the specified template and interpreting its mappings;
// 3 - Rendering the templates based on the yaml and printing to stdout.
func Maillardet(yamlFilePath string) {

}

// ValidateYaml will validate if all required fields are found on the yaml file
func ValidateYaml(yamlFilePath string) error {
	return nil
}
