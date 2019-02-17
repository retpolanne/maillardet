package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// FileProperties is a struct that contains properties about a file, such as if it should be created if it doesn't exist,
// if it should be opened as read only and the private fd that can be accessed by the functions in this package only.
type FileProperties struct {
	FilePath   string
	FileName   string
	ReadOnly   bool
	CreateFile bool
	Append     bool
	flag       int         // File flags (as described in https://golang.org/pkg/os/#OpenFile)
	permission os.FileMode // File permissions (as described in https://golang.org/pkg/os/#OpenFile)
	fd         *os.File    // File descriptor (return value of https://golang.org/pkg/os/#OpenFile)
}

// handleFlags will mask the flags that describe a file as Read Only, Read-Write, Append, etc.
// For now, we hardcode the permissions as 0644, as we still haven't researched so well about https://golang.org/pkg/os/#FileMode
func (fp *FileProperties) handleFlags() {
	flag := 0
	fp.permission = 0644

	if fp.ReadOnly {
		flag = os.O_RDONLY
	} else if fp.CreateFile {
		flag = os.O_CREATE | os.O_RDWR
		if fp.Append {
			flag = flag | os.O_APPEND
		}
	} else {
		flag = os.O_RDWR
		if fp.Append {
			flag = flag | os.O_APPEND
		}
	}
	fp.flag = flag
}

// openFile will call getFullPath to get the full path for a file (expanding env vars if needed), handle its flags and opens it,
// keeping the file pointer in the fd property.
// TODO - maybe this file won't be used because of ioutil.ReadFile.
func (fp *FileProperties) openFile() error {
	fullPath := filepath.Join(os.ExpandEnv(fp.FilePath), fp.FileName)
	fp.handleFlags()
	fd, err := os.OpenFile(fullPath, fp.flag, fp.permission)
	if err != nil {
		return err
	}
	fp.fd = fd
	return err
}

//ParseYaml parses a yaml to the desired struct
func ParseYaml(fp *FileProperties, yamlStruct interface{}) error {
	fullPath := filepath.Join(os.ExpandEnv(fp.FilePath), fp.FileName)
	fileContent, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal([]byte(fileContent), yamlStruct)
	return err
}
