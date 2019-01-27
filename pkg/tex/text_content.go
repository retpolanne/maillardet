package tex

// TextContent is a struct to represent any content such as chapter or section
// Chapters are what separate a whole article in parts
// Sections are what separate a chapter in parts
type TextContent struct {
	Title     string
	ID        string
	Type      string
	Text      string
	Citations []ReferencedContent
}

// addImage will render an image to a string and render it to the Content string
func (txtContent *TextContent) addImage(image *ReferencedContent) {
}

// GenerateTextContent will render a TextContent to a string
func (txtContent *TextContent) GenerateTextContent() string {
	return ""
}
