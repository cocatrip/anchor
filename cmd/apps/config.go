package apps

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/cocatrip/anchor/pkg/common"
)

type Config struct {
	// Global values
	Global map[string]interface{} `yaml:",inline"`
	// Jenkins values
	Jenkins map[string]interface{} `yaml:"jenkins,omitempty"`
	// Docker values
	Docker map[string]interface{} `yaml:"docker,omitempty"`
	// Helm values
	Helm map[string]interface{} `yaml:"helm,omitempty"`
}

func (c *Config) Template(templateFileName string, resultFileName string) error {
	// resultFileName := fmt.Sprintf("%s-%s", templateFileName, c.Global["TESTING_TAG"])
	t := template.New(filepath.Base(templateFileName))
	t = t.Delims("[[", "]]")

	// parse pertama
	t, err := t.ParseFiles(templateFileName)
	if err != nil {
		return err
	}

	// buat file untuk hasil parse
	resultFile, err := os.Create(resultFileName)
	if err != nil {
		return err
	}
	defer resultFile.Close()

	// save parse ke file
	err = t.Execute(resultFile, c)
	if err != nil {
		return err
	}

	contentByte, err := ioutil.ReadFile(resultFileName)
	if err != nil {
		return err
	}
	content := string(contentByte)

	fmt.Println(resultFileName)

	seperator := strings.Repeat("-", len(resultFileName))
	fmt.Println(seperator)

	fmt.Println(content)

	if strings.Contains(content, "<no value>") {
		return errorNoValues(resultFileName)
	} else {
		common.Success("SUCCESS!")
		return nil
	}
}

func errorNoValues(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	var errorMessage []string

	errorMessage = append(errorMessage, fmt.Sprintf("<no value> found in %s:", fileName))
	for n, line := range text {
		if strings.Contains(line, "<no value>") {
			errorMessage = append(errorMessage, fmt.Sprintf("\t%d: %s", n, line))
		}
	}
	error := strings.Join(errorMessage, "\n")
	return errors.New(error)
}
