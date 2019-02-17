package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// ParseYaml parses a yaml to the desired struct.
// It extends any environment var that is found on the filePath string.
func ParseYaml(filePath string, fileName string, yamlStruct interface{}) error {
	fullPath := filepath.Join(os.ExpandEnv(filePath), fileName)
	fileContent, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal([]byte(fileContent), yamlStruct)
	return err
}
