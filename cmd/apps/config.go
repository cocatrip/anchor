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
)

// Struct for config file
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

// function for parsing template using variable from config struct
func (c *Config) Template(templateFileName string, resultFileName string) error {
	// create new tamplate using name from templateFileName input
	t := template.New(filepath.Base(templateFileName))
	// change the delimiter so it doesn't parse helm template
	t = t.Delims("[[", "]]")

	// read the template file
	t, err := t.ParseFiles(templateFileName)
	if err != nil {
		return err
	}

	// create file for the parse result
	resultFile, err := os.Create(resultFileName)
	if err != nil {
		return err
	}
	defer resultFile.Close()

	// parse the file and save it to resultFile
	err = t.Execute(resultFile, c)
	if err != nil {
		return err
	}

	// read the result file to print the result
	contentByte, err := ioutil.ReadFile(resultFileName)
	if err != nil {
		return err
	}
	content := string(contentByte)

	// print the title (filename)
	fmt.Println(resultFileName)

	// print the sperator for title and result
	seperator := strings.Repeat("-", len(resultFileName))
	fmt.Println(seperator)

	// print the result
	fmt.Println(content)

	// print error if <no value> is found
	if strings.Contains(content, "<no value>") {
		return errorNoValues(resultFileName)
	} else {
		success("SUCCESS!")
		return nil
	}
}

// function to print error if no value is found
func errorNoValues(fileName string) error {
	// read the file to found which line contains <no value>
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	// read the file line by line and append it to array to print later
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	var errorMessage []string

	// define error message
	errorMessage = append(errorMessage, fmt.Sprintf("<no value> found in %s:", fileName))
	// loop thourg the text array and print which line contains no value
	for n, line := range text {
		if strings.Contains(line, "<no value>") {
			errorMessage = append(errorMessage, fmt.Sprintf("\t%d: %s", n, line))
		}
	}
	error := strings.Join(errorMessage, "\n")

	// return error
	return errors.New(error)
}
