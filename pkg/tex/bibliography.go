package tex

var accents = map[string]string{
	"á": `{\'a}`,
	"â": `{\^a}`,
	"ã": `{\~a}`,
	"à": `{\` + "`" + `a}`,
	"é": `{\'e}`,
	"è": `{\` + "`" + `e}`,
	"ê": `{\^e}`,
	"í": `{\'\i}`,
	"î": `{\^\i}`,
	"ó": `{\'o}`,
	"ô": `{\^o}`,
	"õ": `{\~o}`,
	"ú": `{\'u}`,
	"û": `{\^u}`,
	"ç": `{\c c}`,
}

// Bibliography is a struct that contains the list of books and websites
type Bibliography struct {
	Books    []ReferencedContent
	Websites []ReferencedContent
}

// Author is a struct that represents an author
type Author struct {
	FirstName string
	Surname   string
}

// ReferencedContent represents the metadata of an image or citation
type ReferencedContent struct {
	// The ID should be something like LastnameYear (as in Macedo2019)
	ID          string
	Title       string
	Subtitle    string
	Authors     []string
	AccessDate  string
	ReleaseYear int
	ReleaseCity string
	Publisher   string
	URL         string
	Kind        string
}

// replaceAccents will find and replace accents to escaped accents, so they won't break on
// bibtex when converted to uppercase
func replaceAccents(text string) string {
	return ""
}

// generateEtAll will take a list of authors and add them to a single string, separated by the word "and"
func generateEtAll(authors []string) string {
	return ""
}

// AddReferencedContent adds ReferencedContent to Bibliography
func (biblio *Bibliography) AddReferencedContent(content *ReferencedContent) {

}

// GenerateReference will render a reference of any type from a template
func GenerateReference(content *ReferencedContent) string {
	return ""
}

// GenerateBibliography renders the whole bibliography to a string
func (biblio *Bibliography) GenerateBibliography() string {
	return ""
}
