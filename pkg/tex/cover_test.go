package tex

import "testing"

var providedCover = &Cover{
	Title:       "Maillardet: ferramenta de renderização de templates ABNT",
	Institution: "Faculdade de Tecnologia da Zona Leste",
	CourseArea:  "Curso de Análise e Desenvolvimento de Sistemas",
	City:        "São Paulo, SP",
	Period:      "Janeiro, 2019",
	Orientation: "Prof. Fulano de Tal",
	Authors: []string{
		"José Da Silva",
		"John Doe",
		"Jane Doe",
	},
}

var expectedCover = `\titulo{Maillardet: ferramenta de renderização de templates ABNT}
\autor{José Da Silva}
\autor{John Doe}
\autor{Jane Doe}
\local{São Paulo, SP}
\data{2019}
\projeto{Trabalho de Conclusão de Curso}
\tituloAcademico{Tecnólogo}
\areaconcentracao{}
\linhapesquisa{}
\instituicao{Faculdade de Tecnologia da Zona Leste}
\departamento{}
\programa{Curso de Análise e Desenvolvimento de Sistemas}
\logoinstituicao{0.2}{figuras/logo-instituicao.png}
\orientador{Prof. Fulano de Tal}`

func TestShouldGenerateCover(t *testing.T) {
	coverString := providedCover.GenerateCover()
	if coverString != expectedCover {
		t.Errorf("The generated covers differ: expected %s, got %s\n", expectedCover, coverString)
	}
}
