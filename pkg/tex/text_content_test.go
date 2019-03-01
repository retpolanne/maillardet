package tex

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/vinicyusmacedo/maillardet/pkg/utils"
)

var fakeChapterText = `Essa é a introdução do nosso trabalho.

Essa é uma outra parte da introdução.`

var fakeSectionText = `De acordo com \citeonline{Doe2019}, esse é um exemplo de um arquivo yaml.

\begin{citacao}
John Doe is not my name. This is a very common misconception. I could be called José da Silva, as well.
\cite{Doe2019}
\end{citacao}`

func fakeContents() *Contents {
	return &Contents{
		fakeChapter(),
		fakeSection(),
	}
}

func fakeChapterTemplateInfo() *utils.TemplateInfo {
	return &utils.TemplateInfo{
		TemplatePath: "$GOPATH/src/github.com/vinicyusmacedo/maillardet/fateczl-abntex2-templates",
		Delims:       []string{"[[", "]]"},
		Path:         "textuais",
		FileName:     "capitulos.tex",
	}
}

func fakeSectionTemplateInfo() *utils.TemplateInfo {
	return &utils.TemplateInfo{
		TemplatePath: "$GOPATH/src/github.com/vinicyusmacedo/maillardet/fateczl-abntex2-templates",
		Delims:       []string{"[[", "]]"},
		Path:         "textuais",
		FileName:     "secoes.tex",
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
		Title:        "Introdução",
		ID:           "introducao",
		Kind:         "chapter",
		Text:         fakeChapterText,
		TemplateInfo: fakeChapterTemplateInfo(),
	}
}

func fakeSection() *TextContent {
	return &TextContent{
		Title:        "Sobre o que é esse artigo?",
		ID:           "sobreoque",
		Kind:         "section",
		Text:         fakeSectionText,
		Citations:    []*ReferencedContent{fakeReferencedContent()},
		TemplateInfo: fakeSectionTemplateInfo(),
	}
}

var expectedChapter = `\chapter{Introdução}
\label{chap:introducao}

Essa é a introdução do nosso trabalho.

Essa é uma outra parte da introdução.`

var expectedSection = `\section{Sobre o que é esse artigo?}
\label{sec:sobreoque}

De acordo com \citeonline{Doe2019}, esse é um exemplo de um arquivo yaml.

\begin{citacao}
John Doe is not my name. This is a very common misconception. I could be called José da Silva, as well.
\cite{Doe2019}
\end{citacao}`

var expectedContentString = `\chapter{Introdução}
\label{chap:introducao}

Essa é a introdução do nosso trabalho.

Essa é uma outra parte da introdução.
\section{Sobre o que é esse artigo?}
\label{sec:sobreoque}

De acordo com \citeonline{Doe2019}, esse é um exemplo de um arquivo yaml.

\begin{citacao}
John Doe is not my name. This is a very common misconception. I could be called José da Silva, as well.
\cite{Doe2019}
\end{citacao}
`

func fakeTextContent() *TextContent {
	return &TextContent{}
}

func TestGenerateTextContentError(t *testing.T) {
	textContent := fakeChapter()
	templateInfo := fakeChapterTemplateInfo()
	templateInfo.FileName = "inexistent.tex"
	textContent.TemplateInfo = templateInfo
	_, err := textContent.GenerateTextContent()
	assert.Error(t, err)
}
func TestGenerateChapter(t *testing.T) {
	generatedTextContent, err := fakeChapter().GenerateTextContent()
	assert.Equal(t, expectedChapter, generatedTextContent)
	assert.NoError(t, err)
}

func TestGenerateSection(t *testing.T) {
	generatedTextContent, err := fakeSection().GenerateTextContent()
	assert.Equal(t, expectedSection, generatedTextContent)
	assert.NoError(t, err)
}

func TestGenerateContentString(t *testing.T) {
	generatedContentString, err := fakeContents().GenerateContentString()
	assert.Equal(t, expectedContentString, generatedContentString)
	assert.NoError(t, err)
}

func TestGenerateContentErrorShouldReturnError(t *testing.T) {
	textContent := fakeChapter()
	templateInfo := fakeChapterTemplateInfo()
	templateInfo.FileName = "inexistent.tex"
	textContent.TemplateInfo = templateInfo
	fakeContents := fakeContents()
	newContent := append(*fakeContents, textContent)
	_, err := newContent.GenerateContentString()
	assert.Error(t, err)
}
