package tex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var providedAccentuation = `áàâãéêèîíôõóûúç`
var expectedAccentuation = `{\'a}{\` + "`" + `a}{\^a}{\~a}{\'e}{\^e}{\` + "`" + `e}{\^\i}{\'\i}{\~o}{\^o}{\'o}{\^u}{\'u}{\c c}`

var providedAuthors = []string{
	"José da Silva",
	"John Doe",
}

var expectedAuthors = "José da Silva and John Doe"

var providedBookReference = &ReferencedContent{
	ID:          "Doe2019",
	Title:       "Testando Referências",
	Subtitle:    "Uma forma eficaz de testar referências",
	Authors:     providedAuthors,
	Publisher:   "Publicações Tabajara",
	ReleaseYear: "2019",
	ReleaseCity: "São Paulo",
	Kind:        "Book",
}

var expectedBookReference = `@Book{Doe2019,
	Title                    = {Testando Refer{\^e}ncias},
	Author                   = {Jos{\'e} da Silva and John Doe},
	Publisher                = {Publicações Tabajara},
	Year                     = {2019},
	Address                  = {S{\~a}o Paulo},
	Subtitle                 = {Uma forma eficaz de testar refer{\^e}ncias},
  }`

var providedWebsiteReference = &ReferencedContent{
	ID:         "Doe2018",
	Title:      "Outra forma de testar referências",
	Authors:    providedAuthors,
	AccessDate: "19 de janeiro de 2019",
	URL:        "https://wikipedia.org",
	Kind:       "Website",
}

var expectedWebsiteReference = `@Misc{Doe2018,
	Title                    = {Outra forma de testar refer{\^encias},
	Author                   = {Jos{\'e} da Silva and John Doe},
	Year                     = {2018},
	Url                      = {https://wikipedia.org}
	Urlaccessdate            = {19 de janeiro de 2019}
  }`

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

func TestGenerateBibliography(t *testing.T) {}

func TestGenerateBookReference(t *testing.T) {
	generatedBookReference := GenerateReference(providedBookReference)
	assert.Equal(t, expectedBookReference, generatedBookReference)
}

func TestGenerateWebsiteReference(t *testing.T) {
	generatedWebsiteReference := GenerateReference(providedWebsiteReference)
	assert.Equal(t, expectedWebsiteReference, generatedWebsiteReference)
}

func TestAccentuationIsEscaped(t *testing.T) {
	accents := replaceAccents(providedAccentuation)
	assert.Equal(t, expectedAccentuation, accents)
}

func TestAccentuationEscapeShouldBeCalled(t *testing.T) {

}

func TestGenerateEtAll(t *testing.T) {
	etAll := generateEtAll(providedAuthors)
	assert.Equal(t, expectedAuthors, etAll)
}
func TestGenerateEtAllShouldBeCalled(t *testing.T) {}
