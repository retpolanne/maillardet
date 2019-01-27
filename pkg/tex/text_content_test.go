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

func fakeChapterWithCitation() *Chapter {
	return &Chapter{
		Title:   "Some chapter",
		ID:      "somechapter",
		Content: "According to [[citDoe2019]], John Doe is not his Doe's actual name.",
	}
}

func fakeSection() *Chapter {
	return &Chapter{
		Title:   "Some section",
		ID:      "somesection",
		Content: "That's a line of section\nThats another line of section.",
	}
}

var expectedChapter = `\chapter{Some chapter}
\label{chap:somechapter}

That's a line of chapter
That's another line of chapter.`

var expectedSection = `\section{Some section}
\label{chap:somesection}

That's a line of section
That's another line of section.`

var expectedDirectCitation = `\begin{citacao}
	John Doe is not my name. This is a very common misconception. I could be called José da Silva, as well \cite{Doe2019}.
\end{citacao}`

var expectedTextWithInlineCitation = `\chapter{Another chapter}
\label{chap:anotherchapter}

According to \citeonline{Doe2019}, John Doe is not his Doe's actual name.`

var expectedInlineCitation = `\citeonline{Doe2019}`

var expectedImage = ``

var expectedCompleteText = `\chapter{Some chapter}
\label{chap:somechapter}

That's a line of chapter
That's another line of chapter.

\section{Some section}
\label{chap:somesection}

That's a line of section
That's another line of section.
\chapter{Another chapter}
\label{chap:anotherchapter}

According to \citeonline{Doe2019}, John Doe is not his Doe's actual name.`

func fakeTextContent() *TextContent {
	return &TextContent{}
}

func TestGenerateTextContent(t *testing.T) {
	generatedTextContent := fakeTextContent().GenerateTextContent()
	assert.Equal(t, expectedCompleteText, generatedTextContent)
}

func TestGenerateChapter(t *testing.T) {
	generatedTextContent := fakeTextContent().GenerateTextContent()
	assert.Equal(t, expectedChapter, generatedTextContent)
}

func TestGenerateSection(t *testing.T) {
	generatedTextContent := fakeTextContent().GenerateTextContent()
	assert.Equal(t, expectedSection, generatedTextContent)
}

func TestGenerateTextWithCitation(t *testing.T) {
	generatedTextContent := fakeTextContent().GenerateTextContent()
	assert.Equal(t, expectedTextWithInlineCitation, generatedTextContent)
}

func TestAddCitation(t *testing.T) {}

func TestAddImage(t *testing.T) {}
