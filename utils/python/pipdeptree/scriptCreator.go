package main

import (
	"github.com/jfrog/jfrog-client-go/utils/io/fileutils"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	pipDepTreeContentFileName     = "deptreescript.go"
	pipDepTreePythonScript        = "pipdeptree.py"
	pipDepTreeContentRelativePath = "utils/python"
	// The pip-dep-tree script version. The version should be manually incremented following changes to the pipdeptree.py source file.
	pipDepTreeVersion = "4"
)

// This main function should be executed manually following changes in pipdeptree.py. Running the function generates new 'pipDepTreeContentFileName' from 'pipDepTreePythonScript.
// Make sure to increment the value of the 'pipDepTreeVersion' constant before running this main function.
func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Check if a pip-dep-tree file of the latest version already exists
	pipDepTreeContentPath := filepath.Join(wd, pipDepTreeContentRelativePath, pipDepTreeContentFileName)
	exists, err := fileutils.IsFileExists(pipDepTreeContentPath, false)
	if err != nil {
		panic(err)
	}
	if exists {
		return
	}
	// Read the script content from the .py file
	pyFile, err := ioutil.ReadFile(path.Join(wd, pipDepTreeContentRelativePath, "pipdeptree", pipDepTreePythonScript))
	if err != nil {
		panic(err)
	}
	// Replace all backticks ( ` ) with a single quote ( ' )
	pyFileString := strings.ReplaceAll(string(pyFile), "`", "'")

	resourceString := "package python\n\n"
	// Add a const string with the script's version
	resourceString += "const pipDepTreeVersion = \"" + pipDepTreeVersion + "\"\n\n"
	// Write the script content a a byte-slice
	resourceString += "var pipDepTreeContent = []byte(`\n" + pyFileString + "`)"
	// Create .go file with the script content
	err = ioutil.WriteFile(pipDepTreeContentPath, []byte(resourceString), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
