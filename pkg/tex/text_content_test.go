package tex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var fakeChapterText = `Essa é a introdução do nosso trabalho.

Essa é uma outra parte da introdução.`

var fakeSectionText = `De acordo com \citeonline{Doe2019}, esse é um exemplo de um arquivo yaml.

\begin{citacao}
John Doe is not my name. This is a very common misconception. I could be called José da Silva, as well.
\cite{Doe2019}
\end{citacao}`

func fakeContent() *Content {
	return &Content{
		Contents: []*TextContent{
			fakeChapter(),
			fakeSection(),
		},
	}
}

func fakeReferencedContent() *ReferencedContent {
	return &ReferencedContent{
		ID:       "Doe2019",
		Title:    "Testando Referências",
		Subtitle: "Uma forma eficaz de testar referências",
		Authors: []string{
			"José da Silva",
			"John Doe",
		},
		Publisher:   "Publicações Tabajara",
		ReleaseYear: "2019",
		ReleaseCity: "São Paulo",
		Kind:        "book",
	}
}

func fakeChapter() *TextContent {
	return &TextContent{
		Title: "Introdução",
		ID:    "introducao",
		Kind:  "chapter",
		Text:  fakeChapterText,
	}
}

func fakeSection() *TextContent {
	return &TextContent{
		Title:     "Sobre o que é esse artigo?",
		ID:        "sobreoque",
		Kind:      "section",
		Text:      fakeSectionText,
		Citations: []*ReferencedContent{fakeReferencedContent()},
	}
}

var expectedChapter = `\chapter{Introdução}
\label{chap:introducao}

Essa é a introdução do nosso trabalho.

Essa é uma outra parte da introdução.`

var expectedSection = `\section{Sobre o que é esse artigo?}
\label{chap:sobreoque}

De acordo com \citeonline{Doe2019}, esse é um exemplo de um arquivo yaml.
      
\begin{citacao}
John Doe is not my name. This is a very common misconception. I could be called José da Silva, as well.
\end{citacao}`

var expectedContentString = `\chapter{Introdução}
\label{chap:introducao}

Essa é a introdução do nosso trabalho.

Essa é uma outra parte da introdução.
section{Sobre o que é esse artigo?}
\label{chap:sobreoque}

De acordo com \citeonline{Doe2019}, esse é um exemplo de um arquivo yaml.
      
\begin{citacao}
John Doe is not my name. This is a very common misconception. I could be called José da Silva, as well.
\end{citacao}`

func fakeTextContent() *TextContent {
	return &TextContent{}
}

func TestGenerateChapter(t *testing.T) {
	generatedTextContent := fakeTextContent().GenerateTextContent()
	assert.Equal(t, expectedChapter, generatedTextContent)
}

func TestGenerateSection(t *testing.T) {
	generatedTextContent := fakeTextContent().GenerateTextContent()
	assert.Equal(t, expectedSection, generatedTextContent)
}

func TestContentWithCitationShouldUpdateReferences(t *testing.T) {}

func TestGenerateContentString(t *testing.T) {
	generatedContentString := fakeContent().GenerateContentString()
	assert.Equal(t, expectedContentString, generatedContentString)
}

func TestAddImage(t *testing.T) {}
