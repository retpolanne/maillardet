package utils

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeTestChapter struct {
	Title string
	ID    string
	Kind  string
	Text  string
}

var fakeChapterText = `Essa é a introdução do nosso trabalho.

Essa é uma outra parte da introdução.`

var expectedRenderedTemplate = `\chapter{Introdução}
\label{chap:introducao}

Essa é a introdução do nosso trabalho.

Essa é uma outra parte da introdução.`

func fakeChapter() *fakeTestChapter {
	return &fakeTestChapter{
		Title: "Introdução",
		ID:    "introducao",
		Kind:  "chapter",
		Text:  fakeChapterText,
	}
}

func fakeTemplate() *TemplateInfo {
	return &TemplateInfo{
		TemplatePath: "../../fateczl-abntex2-templates",
		Delims:       []string{"[[", "]]"},
		Path:         "textuais",
		FileName:     "capitulos.tex",
	}
}

func fakeInvalidTemplate() *TemplateInfo {
	return &TemplateInfo{
		TemplatePath: "../../fixtures",
		Delims:       []string{"[[", "]]"},
		Path:         "",
		FileName:     "invalid.tex",
	}
}

func fakeInexistentTemplate() *TemplateInfo {
	return &TemplateInfo{
		TemplatePath: "../../fixtures",
		Delims:       []string{"[[", "]]"},
		Path:         "textuais",
		FileName:     "inexistent.tex",
	}
}

func TestRenderTemplateToWriter(t *testing.T) {
	var bWriter bytes.Buffer
	err := fakeTemplate().RenderTemplateToWriter(&bWriter, fakeChapter())
	assert.Equal(t, expectedRenderedTemplate, bWriter.String())
	assert.NoError(t, err)
}

func TestRenderTemplateError(t *testing.T) {
	var bWriter bytes.Buffer
	err := fakeInvalidTemplate().RenderTemplateToWriter(&bWriter, fakeChapter())
	assert.Empty(t, bWriter)
	assert.Error(t, err)
}

func TestRenderTemplateFileNotFoundError(t *testing.T) {
	var bWriter bytes.Buffer
	err := fakeInexistentTemplate().RenderTemplateToWriter(&bWriter, fakeChapter())
	assert.Empty(t, bWriter)
	assert.Error(t, err)
}

func TestRenderTemplateInvalidError(t *testing.T) {
	var bWriter bytes.Buffer
	err := fakeInvalidTemplate().RenderTemplateToWriter(&bWriter, fakeChapter())
	assert.Empty(t, bWriter)
	assert.Contains(t, err.Error(), "undefined variable")
	assert.Error(t, err)
}
