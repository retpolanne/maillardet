package utils

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vinicyusmacedo/maillardet/pkg/tex"
)

var fakeTemplatePath = "../../fateczl-abntex2-templates/textuais"
var fakeTemplateName = "capitulos.tex"
var fakeInvalidTemplatePath = "../../fixtures"
var fakeInvalidTemplateName = "invalid.tex"

var fakeChapterText = `Essa é a introdução do nosso trabalho.

Essa é uma outra parte da introdução.`
var fakeDelims = []string{"[[", "]]"}
var fakeUnexistentFile = "filedonotexist.tex"

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

func TestRenderTemplateToWriter(t *testing.T) {
	var bWriter bytes.Buffer
	err := RenderTemplateToWriter(&bWriter, fakeTemplatePath, fakeTemplateName, fakeChapter(), fakeDelims)
	assert.Equal(t, expectedRenderedTemplate, bWriter.String())
	assert.NoError(t, err)
}

func TestRenderTemplateError(t *testing.T) {
	var bWriter bytes.Buffer
	err := RenderTemplateToWriter(&bWriter, fakeTemplatePath, fakeUnexistentFile, fakeChapter(), fakeDelims)
	assert.Empty(t, bWriter)
	assert.Error(t, err)
}

func TestRenderTemplateFileNotFoundError(t *testing.T) {
	var bWriter bytes.Buffer
	err := RenderTemplateToWriter(&bWriter, fakeTemplatePath, fakeUnexistentFile, fakeChapter(), fakeDelims)
	assert.Empty(t, bWriter)
	assert.Contains(t, err.Error(), "no such file or directory")
	assert.Error(t, err)
}

func TestRenderTemplateInvalidError(t *testing.T) {
	var bWriter bytes.Buffer
	err := RenderTemplateToWriter(&bWriter, fakeInvalidTemplatePath, fakeInvalidTemplateName, fakeChapter(), fakeDelims)
	assert.Empty(t, bWriter)
	assert.Contains(t, err.Error(), "undefined variable")
	assert.Error(t, err)
}
