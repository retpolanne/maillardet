package tex

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var fakeCoverTemplatePath = "$GOPATH/src/github.com/vinicyusmacedo/maillardet/fateczl-abntex2-templates"

func fakeCover() *Cover {
	return &Cover{
		Title:       "Maillardet: ferramenta de renderização de templates ABNT",
		Institution: "Faculdade de Tecnologia da Zona Leste",
		CourseArea:  "Curso de Análise e Desenvolvimento de Sistemas",
		City:        "São Paulo, SP",
		Period:      "2019",
		Orientation: "Prof. Fulano de Tal",
		Authors: []string{
			"José Da Silva",
			"John Doe",
			"Jane Doe",
		},
		templatePath:     filepath.Join(fakeCoverTemplatePath, "pre-textuais"),
		templateFilename: "capa.tex",
		delims:           []string{"[[", "]]"},
	}
}

var expectedCover = `\titulo{Maillardet: ferramenta de renderização de templates ABNT}
\autor{José Da Silva}
\autor{John Doe}
\autor{Jane Doe}
\local{São Paulo, SP}
\data{2019}
\projeto{Trabalho de Conclusão de Curso}
\tituloAcademico{Tecnólogo}
\instituicao{Faculdade de Tecnologia da Zona Leste}
\programa{Curso de Análise e Desenvolvimento de Sistemas}
\logoinstituicao{0.2}{figuras/logo-instituicao.png}
\orientador{Prof. Fulano de Tal}`

func TestShouldGenerateCover(t *testing.T) {
	coverString, err := fakeCover().GenerateCover()
	assert.NoError(t, err)
	assert.Equal(t, expectedCover, coverString)
}

func TestGenerateCoverError(t *testing.T) {
	cover := fakeCover()
	cover.templateFilename = "invalid.tex"
	_, err := cover.GenerateCover()
	assert.Error(t, err)
}
