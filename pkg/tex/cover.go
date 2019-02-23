package tex

import (
	"bytes"

	"github.com/vinicyusmacedo/maillardet/pkg/utils"
)

// Cover is the struct that represents the document cover
type Cover struct {
	Title        string   `yaml:"title"`
	Institution  string   `yaml:"institution"`
	CourseArea   string   `yaml:"courseArea"`
	City         string   `yaml:"city"`
	Period       string   `yaml:"period"`
	Authors      []string `yaml:"authors"`
	Orientation  string   `yaml:"orientation"`
	TemplateInfo *utils.TemplateInfo
}

// GenerateCover will generate the cover section from a template to a string
func (cover *Cover) GenerateCover() (string, error) {
	var bWriter bytes.Buffer
	err := cover.TemplateInfo.RenderTemplateToWriter(&bWriter, cover)
	if err != nil {
		return "", err
	}
	return bWriter.String(), nil
}
