package tex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var expectedIntro = ``

var expectedChapter = ``

var expectedChapterWithSection = ``

var expectedCitation = ``

var expectedImage = ``

func fakeTextContent() *TextContent {
	return &TextContent{}
}

func TestGenerateIntro(t *testing.T) {
	generatedTextContent := fakeTextContent().GenerateTextContent()
	assert.Equal(t, expectedIntro, generatedTextContent)
}

func TestGenerateChapter(t *testing.T) {
	generatedTextContent := fakeTextContent().GenerateTextContent()
	assert.Equal(t, expectedChapter, generatedTextContent)
}

func TestGenerateChapterWithSection(t *testing.T) {
	generatedTextContent := fakeTextContent().GenerateTextContent()
	assert.Equal(t, expectedChapterWithSection, generatedTextContent)
}

func TestAddCitation(t *testing.T) {}

func TestAddImage(t *testing.T) {}
