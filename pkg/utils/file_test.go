package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeYaml struct {
	Test string `yaml:"test"`
}

func fakeYamlFileProperties() *FileProperties {
	return &FileProperties{
		FilePath: "$GOPATH/src/github.com/vinicyusmacedo/maillardet/fixtures",
		FileName: "fake.yaml",
		ReadOnly: true,
	}
}

func fakeYamlInvalidFileProperties() *FileProperties {
	return &FileProperties{
		FilePath: "$GOPATH/src/github.com/vinicyusmacedo/maillardet/fixtures",
		FileName: "fakeInvalid.yaml",
		ReadOnly: true,
	}
}

func fakeFileProperties() *FileProperties {
	return &FileProperties{
		FilePath: "$GOPATH/src/github.com/vinicyusmacedo/maillardet/fateczl-abntex2-templates/textuais",
		FileName: "capitulos.tex",
	}
}

func fakeInexistentFileProperties() *FileProperties {
	return &FileProperties{
		FilePath: "$WHATEVER/fixtures",
		FileName: "filedonotexist.tex",
	}
}

func fakeFilePropertiesFlags() *FileProperties {
	return &FileProperties{
		ReadOnly:   true,
		CreateFile: true,
		Append:     true,
	}
}

func fakeFilePropertiesMissingFlags() *FileProperties {
	return &FileProperties{}
}

func TestHandleFlagsReadOnly(t *testing.T) {
	fp := fakeFilePropertiesFlags()
	fp.ReadOnly = true
	fp.handleFlags()
	assert.Equal(t, os.O_RDONLY, fp.flag)
}

func TestHandleFlagsCreateFile(t *testing.T) {
	fp := fakeFilePropertiesFlags()
	fp.ReadOnly = false
	fp.CreateFile = true
	fp.Append = false
	fp.handleFlags()
	assert.Equal(t, os.O_CREATE|os.O_RDWR, fp.flag)
}

func TestHandleFlagsNoCreate(t *testing.T) {
	fp := fakeFilePropertiesFlags()
	fp.ReadOnly = false
	fp.CreateFile = false
	fp.Append = false
	fp.handleFlags()
	assert.Equal(t, os.O_RDWR, fp.flag)
}

func TestHandleFlagsAppend(t *testing.T) {
	fp := fakeFilePropertiesFlags()
	fp.ReadOnly = false
	fp.CreateFile = false
	fp.Append = true
	fp.handleFlags()
	assert.Equal(t, os.O_RDWR|os.O_APPEND, fp.flag)
	fp.CreateFile = true
	fp.handleFlags()
	assert.Equal(t, os.O_CREATE|os.O_RDWR|os.O_APPEND, fp.flag)
}

func TestHandleFlagsEmptyProperties(t *testing.T) {
	fp := fakeFilePropertiesMissingFlags()
	fp.handleFlags()
	assert.Equal(t, os.O_RDWR, fp.flag)
}

func TestHandleFlagsConflictingProperties(t *testing.T) {
	fp := fakeFilePropertiesFlags()
	fp.handleFlags()
	assert.Equal(t, os.O_RDONLY, fp.flag)
}

func TestOpenFile(t *testing.T) {
	fp := fakeFileProperties()
	err := fp.openFile()
	defer fp.fd.Close()
	assert.NoError(t, err)
	assert.Contains(t, fp.fd.Name(), "github.com")
}

func TestOpenFileNotFoundError(t *testing.T) {
	fp := fakeInexistentFileProperties()
	err := fp.openFile()
	defer fp.fd.Close()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no such file or directory")
}

func TestOpenFileError(t *testing.T) {
	fp := fakeFilePropertiesMissingFlags()
	err := fp.openFile()
	defer fp.fd.Close()
	assert.Error(t, err)
}

func TestParseYaml(t *testing.T) {
	fp := fakeYamlFileProperties()
	yaml := &fakeYaml{}
	err := ParseYaml(fp, yaml)
	assert.NoError(t, err)
	assert.Contains(t, yaml.Test, "test")
}

func TestParseYamlInvalid(t *testing.T) {
	fp := fakeYamlInvalidFileProperties()
	yaml := &fakeYaml{}
	err := ParseYaml(fp, yaml)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot unmarshal")
}

func TestParseYamlInexistent(t *testing.T) {
	fp := fakeInexistentFileProperties()
	yaml := &fakeYaml{}
	err := ParseYaml(fp, yaml)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no such file or directory")
}
