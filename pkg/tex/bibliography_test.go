package tex

import (
	"testing"

	"github.com/vinicyusmacedo/maillardet/pkg/utils"

	"github.com/stretchr/testify/assert"
)

var providedAccentuation = `áàâãéêèîíôõóûúç`
var expectedAccentuation = `{\'a}{\` + "`" + `a}{\^a}{\~a}{\'e}{\^e}{\` + "`" + `e}{\^\i}{\'\i}{\~o}{\^o}{\'o}{\^u}{\'u}{\c c}`

var providedAuthors = []string{
	"José da Silva",
	"John Doe",
}

var expectedAuthors = "José da Silva and John Doe"

func fakeBibliographyTemplateInfo() *utils.TemplateInfo {
	return &utils.TemplateInfo{
		TemplatePath: "$GOPATH/src/github.com/vinicyusmacedo/maillardet/fateczl-abntex2-templates",
		Delims:       []string{"[[", "]]"},
		Path:         "pos-textuais",
		FileName:     "base-referencias.bib",
	}
}

func fakeBookReference() *ReferencedContent {
	return &ReferencedContent{
		ID:           "Doe2019",
		Title:        "Testando Referências",
		Subtitle:     "Uma forma eficaz de testar referências",
		Authors:      providedAuthors,
		Publisher:    "Publicações Tabajara",
		ReleaseYear:  "2019",
		ReleaseCity:  "São Paulo",
		Kind:         "Book",
		TemplateInfo: fakeBibliographyTemplateInfo(),
	}
}

var expectedBookReference = `@Book{Doe2019,
	Title                    = {Testando Refer{\^e}ncias},
	Author                   = {Jos{\'e} da Silva and John Doe},
	Publisher                = {Publicações Tabajara},
	Year                     = {2019},
	Address                  = {S{\~a}o Paulo},
	Subtitle                 = {Uma forma eficaz de testar refer{\^e}ncias},
  }`

func fakeWebsiteReference() *ReferencedContent {
	return &ReferencedContent{
		ID:           "Doe2018",
		Title:        "Outra forma de testar referências",
		Authors:      providedAuthors,
		AccessDate:   "19 de janeiro de 2019",
		URL:          "https://wikipedia.org",
		Kind:         "Website",
		TemplateInfo: fakeBibliographyTemplateInfo(),
	}
}

var expectedWebsiteReference = `@Misc{Doe2018,
	Title                    = {Outra forma de testar refer{\^encias},
	Author                   = {Jos{\'e} da Silva and John Doe},
	Year                     = {2018},
	Url                      = {https://wikipedia.org}
	Urlaccessdate            = {19 de janeiro de 2019}
  }`

var fakeBibliographyTemplatePath = "$GOPATH/src/github.com/vinicyusmacedo/maillardet/fateczl-abntex2-templates"

func fakeBibliography() *Bibliography {
	return &Bibliography{
		Books:    []*ReferencedContent{fakeBookReference()},
		Websites: []*ReferencedContent{fakeWebsiteReference()},
	}
}

var expectedBibFile = `@Book{Doe2019,
	Title                    = {Testando Refer{\^e}ncias},
	Author                   = {Jos{\'e} da Silva and John Doe},
	Publisher                = {Publicações Tabajara},
	Year                     = {2019},
	Address                  = {S{\~a}o Paulo},
	Subtitle                 = {Uma forma eficaz de testar refer{\^e}ncias},
  }
  @Misc{Doe2018,
	Title                    = {Outra forma de testar refer{\^encias},
	Author                   = {Jos{\'e} da Silva and John Doe},
	Year                     = {2018},
	Url                      = {https://wikipedia.org}
	Urlaccessdate            = {19 de janeiro de 2019}
  }`

func TestGenerateBibliography(t *testing.T) {
	generatedBibliography := fakeBibliography().GenerateBibliography()
	assert.Equal(t, expectedBibFile, generatedBibliography)
}

func TestGenerateBookReference(t *testing.T) {
	generatedBookReference := GenerateReference(fakeBookReference())
	assert.Equal(t, expectedBookReference, generatedBookReference)
}

func TestGenerateWebsiteReference(t *testing.T) {
	generatedWebsiteReference := GenerateReference(fakeWebsiteReference())
	assert.Equal(t, expectedWebsiteReference, generatedWebsiteReference)
}

func TestAccentuationIsEscaped(t *testing.T) {
	accents := replaceAccents(providedAccentuation)
	assert.Equal(t, expectedAccentuation, accents)
}

func TestGenerateEtAll(t *testing.T) {
	etAll := generateEtAll(providedAuthors)
	assert.Equal(t, expectedAuthors, etAll)
}

func TestAddReferencedWebsiteContent(t *testing.T) {
	biblio := fakeBibliography()
	websiteRef := fakeWebsiteReference()
	biblio.AddReferencedContent(websiteRef)
	assert.Equal(t, biblio.Websites[0], websiteRef)
	websiteRef.Title = "Just another site"
	biblio.AddReferencedContent(websiteRef)
	assert.Equal(t, biblio.Websites[1], websiteRef)
}

func TestAddReferencedBookContent(t *testing.T) {
	biblio := fakeBibliography()
	bookRef := fakeBookReference()
	biblio.AddReferencedContent(bookRef)
	assert.Equal(t, biblio.Books[0], bookRef)
	bookRef.Title = "Just another book"
	biblio.AddReferencedContent(bookRef)
	assert.Equal(t, biblio.Websites[1], bookRef)
}
