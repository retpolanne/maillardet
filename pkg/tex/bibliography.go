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
	Title         string
	Author        []Author
	AccessDate    string
	ReleaseYear   int
	ReleaseCity   string
	ReleaseEditor string
	URL           string
	Kind          string
}

// replaceAccents will find and replace accents to escaped accents, so they won't break on
// bibtex when converted to uppercase
func (content *ReferencedContent) replaceAccents() {

}

// AddReferencedContent adds ReferencedContent to Bibliography
func (biblio *Bibliography) AddReferencedContent(content *ReferencedContent) {

}

// GenerateBibliography renders the bibliography to a string
func (biblio *Bibliography) GenerateBibliography() string {
	return ""
}
