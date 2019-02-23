package tex

import (
	"bytes"

	"github.com/vinicyusmacedo/maillardet/pkg/utils"
)

// Contents keeps a list of TextContents.
// This list should be ordered the same way the contents will appear on the final work
type Contents []*TextContent

// TextContent is a struct to represent any content such as chapter or section
// Chapters are what separate a whole article in parts
// Sections are what separate a chapter in parts
type TextContent struct {
	Title        string               `yaml:"title"`
	ID           string               `yaml:"id"`
	Kind         string               `yaml:"kind"`
	Text         string               `yaml:"text"`
	Citations    []*ReferencedContent `yaml:"citations"`
	TemplateInfo *utils.TemplateInfo
}

// GenerateTextContent will render a TextContent to a string
// TODO - maybe this should not be public - use GenerateContentString as entrypoint
func (txtContent *TextContent) GenerateTextContent() (string, error) {
	var bWriter bytes.Buffer
	err := txtContent.TemplateInfo.RenderTemplateToWriter(&bWriter, txtContent)
	if err != nil {
		return "", err
	}
	return bWriter.String(), nil
}

// GenerateContentString will render all the TextContents into a single string
func (contents *Contents) GenerateContentString() (string, error) {
	var buffer bytes.Buffer
	for _, content := range *contents {
		renderedContent, err := content.GenerateTextContent()
		if err != nil {
			return "", err
		}
		buffer.WriteString(renderedContent + "\n")
	}
	return buffer.String(), nil
}
