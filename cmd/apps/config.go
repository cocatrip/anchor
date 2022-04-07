package apps

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/cocatrip/anchor/pkg/common"
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

	fmt.Println(content)

	if strings.Contains(content, "<no value>") {
		printNoValues(resultFileName)
	} else {
		common.Success.Println("tidak ada yang no value")
	}
}

func printNoValues(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	common.Error.Printf("Err: <no value> found in %s:\n", fileName)
	for n, line := range text {
		if strings.Contains(line, "<no value>") {
			fmt.Printf("\t%d: %s\n", n, line)
		}
	}
}
