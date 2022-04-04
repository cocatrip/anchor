package apps

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Config struct {
	// global values
	Global map[string]interface{} `yaml:",inline"`
	// jenkins values
	Jenkins map[string]interface{} `yaml:"jenkins,omitempty"`
	// docker values
	Docker map[string]interface{} `yaml:"docker,omitempty"`
	// helm values
	Helm map[string]interface{} `yaml:"helm,omitempty"`
}

func (c *Config) Template(templateFileName string, resultFileName string) {
	// resultFileName := fmt.Sprintf("%s-%s", templateFileName, c.Global["TESTING_TAG"])

	t := template.New(filepath.Base(templateFileName))
	t = t.Delims("[[", "]]")

	// parse pertama
	t, err := t.ParseFiles(templateFileName)
	if err != nil {
		panic(err)
	}

	// buat file untuk hasil parse
	resultFile, err := os.Create(resultFileName)
	if err != nil {
		panic(err)
	}
	defer resultFile.Close()

	// save parse ke file
	err = t.Execute(resultFile, c)
	if err != nil {
		panic(err)
	}

	contentByte, err := ioutil.ReadFile(resultFileName)
	if err != nil {
		panic(err)
	}
	content := string(contentByte)

	if strings.Contains(content, "<no value>") {
		fmt.Println("ada yang no value")
	} else {
		fmt.Println("tidak ada yang no value")
	}
}
