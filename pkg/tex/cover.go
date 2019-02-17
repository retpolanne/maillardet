package tex

import (
	"bytes"

	"github.com/vinicyusmacedo/maillardet/pkg/utils"
)

// Cover is the struct that represents the document cover
type Cover struct {
	Title            string
	Institution      string
	CourseArea       string
	City             string
	Period           string
	Authors          []string
	Orientation      string
	templatePath     string
	templateFilename string
	delims           []string
}

// GenerateCover will generate the cover section from a template to a string
func (cover *Cover) GenerateCover() (string, error) {
	var bWriter bytes.Buffer
	err := utils.RenderTemplateToWriter(&bWriter, cover.templatePath, cover.templateFilename, cover, cover.delims)
	if err != nil {
		return "", err
	}
	return bWriter.String(), nil
}
