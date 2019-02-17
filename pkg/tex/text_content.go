package tex

import (
	"bytes"

	"github.com/vinicyusmacedo/maillardet/pkg/utils"
)

// Content keeps a list of TextContents.
// This list should be ordered the same way the contents will appear on the final work
type Content struct {
	Contents []*TextContent
}

// TextContent is a struct to represent any content such as chapter or section
// Chapters are what separate a whole article in parts
// Sections are what separate a chapter in parts
type TextContent struct {
	Title            string
	ID               string
	Kind             string
	Text             string
	Citations        []*ReferencedContent
	Images           []*ReferencedContent
	templateFilename string
	templatePath     string
	delims           []string
}

// AddImage will render an image to a string and render it to the Text string
// TODO - add image function
func (txtContent *TextContent) addImage() {
}

// GenerateTextContent will render a TextContent to a string
// TODO - maybe this should not be public - use GenerateContentString as entrypoint
func (txtContent *TextContent) GenerateTextContent() (string, error) {
	var bWriter bytes.Buffer
	err := utils.RenderTemplateToWriter(&bWriter, txtContent.templatePath, txtContent.templateFilename, txtContent, txtContent.delims)
	if err != nil {
		return "", err
	}
	return bWriter.String(), nil
}

// GenerateContentString will render all the TextContents into a single string
func (content *Content) GenerateContentString() string {
	return ""
}
