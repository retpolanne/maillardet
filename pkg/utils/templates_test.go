package utils

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vinicyusmacedo/maillardet/pkg/tex"
)

var fakeChapterText = `Essa é a introdução do nosso trabalho.

Essa é uma outra parte da introdução.`
var fakeDelims = []string{"[[", "]]"}

var expectedRenderedTemplate = `\chapter{Introdução}
\label{chap:introducao}

Essa é a introdução do nosso trabalho.

Essa é uma outra parte da introdução.`

func fakeChapter() *tex.TextContent {
	return &tex.TextContent{
		Title: "Introdução",
		ID:    "introducao",
		Kind:  "chapter",
		Text:  fakeChapterText,
	}
}

func fakeTexFileProperties() *FileProperties {
	return &FileProperties{
		FilePath: "../../fateczl-abntex2-templates/textuais",
		FileName: "capitulos.tex",
		ReadOnly: true,
	}
}

func fakeInvalidTexFileProperties() *FileProperties {
	return &FileProperties{
		FilePath: "../../fixtures",
		FileName: "invalid.tex",
		ReadOnly: true,
	}
}

func fakeInexistentTexFileProperties() *FileProperties {
	return &FileProperties{
		FilePath: "../../fixtures",
		FileName: "filedonotexist.tex",
		ReadOnly: true,
	}
}

func TestRenderTemplateToWriter(t *testing.T) {
	var bWriter bytes.Buffer
	err := RenderTemplateToWriter(&bWriter, fakeTexFileProperties(), fakeChapter(), fakeDelims)
	assert.Equal(t, expectedRenderedTemplate, bWriter.String())
	assert.NoError(t, err)
}

func TestRenderTemplateError(t *testing.T) {
	var bWriter bytes.Buffer
	err := RenderTemplateToWriter(&bWriter, fakeInvalidTexFileProperties(), fakeChapter(), fakeDelims)
	assert.Empty(t, bWriter)
	assert.Error(t, err)
}

func TestRenderTemplateFileNotFoundError(t *testing.T) {
	var bWriter bytes.Buffer
	err := RenderTemplateToWriter(&bWriter, fakeInexistentTexFileProperties(), fakeChapter(), fakeDelims)
	assert.Empty(t, bWriter)
	assert.Contains(t, err.Error(), "no such file or directory")
	assert.Error(t, err)
}

func TestRenderTemplateInvalidError(t *testing.T) {
	var bWriter bytes.Buffer
	err := RenderTemplateToWriter(&bWriter, fakeInvalidTexFileProperties(), fakeChapter(), fakeDelims)
	assert.Empty(t, bWriter)
	assert.Contains(t, err.Error(), "undefined variable")
	assert.Error(t, err)
}
