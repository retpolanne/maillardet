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

var expectedCover = ``

var expectedCoverSheet = ``

func TestShouldGenerateCover(t *testing.T) {
	coverString := providedCover.GenerateCover()
	if coverString != expectedCover {
		t.Errorf("The generated covers differ: expected %s, got %s\n", expectedCover, coverString)
	}
}
