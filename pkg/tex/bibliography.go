package tex

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

// AddReferencedContent adds ReferencedContent to Bibliography
func AddReferencedContent(content *ReferencedContent) {

}

// GenerateBibliography renders the bibliography to a string
func (biblio *Bibliography) GenerateBibliography() string {
	return ""
}
