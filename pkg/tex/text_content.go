package tex

// Content keeps a list of TextContents.
// This list should be ordered the same way the contents will appear on the final work
type Content struct {
	Contents []*TextContent
}

// TextContent is a struct to represent any content such as chapter or section
// Chapters are what separate a whole article in parts
// Sections are what separate a chapter in parts
type TextContent struct {
	Title     string
	ID        string
	Kind      string
	Text      string
	Citations []*ReferencedContent
	Images    []*ReferencedContent
}

// AddImage will render an image to a string and render it to the Text string
func (txtContent *TextContent) addImage() {
}

// GenerateTextContent will render a TextContent to a string
func (txtContent *TextContent) GenerateTextContent() string {
	return ""
}

// GenerateContentString will render all the TextContents into a single string
func (content *Content) GenerateContentString() string {
	return ""
}
