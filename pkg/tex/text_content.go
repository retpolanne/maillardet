package tex

// TODO - organize these so we can use them on the templates

// Article holds a list of Chapters
type Article struct {
	Chapters []Chapter
}

// Chapter holds a list of TextContents
type Chapter struct {
	ID      string
	Title   string `yaml:"title, omitempty"`
	Content string
}

// TextContent is an interface from which introduction and chapter
type TextContent struct {
	Title         string
	ID            string
	Content       string
	Type          string
	Images        []map[string]ReferencedContent
	Citations     []map[string]ReferencedContent
	SubchapterIDs []string
	Subchapters   []TextContent
}

// addCitation will render a citation to a string and render it to the Content string
func (txtContent *TextContent) addCitation(citation *ReferencedContent) {
}

// addImage will render an image to a string and render it to the Content string
func (txtContent *TextContent) addImage(image *ReferencedContent) {
}

// GenerateTextContent will render a TextContent to a string
func (txtContent *TextContent) GenerateTextContent() string {
	return txtContent.Content
}
