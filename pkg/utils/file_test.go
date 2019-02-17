package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeYaml struct {
	Test string `yaml:"test"`
}

var fakeYamlFilePath = "$GOPATH/src/github.com/vinicyusmacedo/maillardet/fixtures"
var fakeYamlFileName = "fake.yaml"
var fakeInvalidYamlName = "fakeInvalid.yaml"
var fakeInexistentYamlName = "inexistent.yaml"

func TestParseYaml(t *testing.T) {
	yaml := &fakeYaml{}
	err := ParseYaml(fakeYamlFilePath, fakeYamlFileName, yaml)
	assert.NoError(t, err)
	assert.Contains(t, yaml.Test, "test")
}

func TestParseYamlInvalid(t *testing.T) {
	yaml := &fakeYaml{}
	err := ParseYaml(fakeYamlFilePath, fakeInvalidYamlName, yaml)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot unmarshal")
}

func TestParseYamlInexistent(t *testing.T) {
	yaml := &fakeYaml{}
	err := ParseYaml(fakeYamlFilePath, fakeInexistentYamlName, yaml)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no such file or directory")
}
