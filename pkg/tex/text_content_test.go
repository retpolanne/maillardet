package tex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func fakeChapter() *Chapter {
	return &Chapter{
		Title:   "Some chapter",
		ID:      "somechapter",
		Content: "That's a line of chapter\nThats another line of chapter.",
	}
}

func fakeSection() *Chapter {
	return &Chapter{
		Title:   "Some section",
		ID:      "somesection",
		Content: "That's a line of section\nThats another line of section.",
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

func TestAddImage(t *testing.T) {}
