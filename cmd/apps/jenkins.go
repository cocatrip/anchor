package apps

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"strings"
)

type Jenkins struct {
	FILE             string
	APPLICATION_NAME string
	BUSINESS_NAME    string
	TESTING_TAG      string
	SERVER_NAME      string
	LINK_SONARQUBE   string `yaml:"LINK_SONARQUBE,omitempty"`
	SONARQUBE_URL    string `yaml:"SONARQUBE_URL,omitempty"`
	SONARQUBE_KEY    string `yaml:"SONARQUBE_KEY,omitempty"`
	JAR_APP_NAME     string `yaml:"JAR_APP_NAME,omitempty"`
	CREDENTIAL_HELM  string `yaml:"CREDENTIAL_HELM,omitempty"`
}

func (j *Jenkins) New(c Config) {
	j.APPLICATION_NAME = c.APPLICATION_NAME
	j.BUSINESS_NAME = c.BUSINESS_NAME
	j.TESTING_TAG = c.TESTING_TAG
	j.SERVER_NAME = c.SERVER_NAME
	j.FILE = fmt.Sprintf("Jenkinsfile-%s", c.TESTING_TAG)
}

func (j *Jenkins) GetValues() ([]string, []string) {
	value := reflect.ValueOf(j).Elem()
	field := value.Type()

	var v, f []string
	for i := 0; i < value.NumField(); i++ {
		v = append(v, value.Field(i).String())
		f = append(f, field.Field(i).Name)
	}
	return v, f
}

func (j *Jenkins) GetValuesLength() int {
	value := reflect.ValueOf(j).Elem()
	return value.NumField()
}

func (j *Jenkins) ReadJenkins() [][]string {
	jenkinsByte, err := ioutil.ReadFile("Jenkinsfile")
	if err != nil {
		panic(err)
	}
	jenkinsFile := string(jenkinsByte)

	re, err := regexp.Compile(`%\{(\w+)\}`)
	if err != nil {
		fmt.Println(err)
	}

	values := re.FindAllStringSubmatch(jenkinsFile, -1)

	return values
}

func (j *Jenkins) TemplateJenkins() string {
	jenkinsByte, err := ioutil.ReadFile("Jenkinsfile")
	if err != nil {
		panic(err)
	}
	jenkinsFile := string(jenkinsByte)

	configValue, configField := j.GetValues()
	jenkinsValue := j.ReadJenkins()

	for i := 0; i < j.GetValuesLength(); i++ {
		for k := 0; k < len(j.ReadJenkins()); k++ {
			if configField[i] == jenkinsValue[k][1] {
				jenkinsFile = strings.ReplaceAll(jenkinsFile, jenkinsValue[k][0], configValue[i])
			}
		}
	}

	return jenkinsFile
}
