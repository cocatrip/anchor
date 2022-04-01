package apps

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"strings"
)

type Jenkins struct {
	APPLICATION_NAME string
	BUSINESS_NAME    string
	TESTING_TAG      string
	FILE             string
	SERVER_NAME      string `yaml:"SERVER_NAME,omitempty"`
	LINK_SONARQUBE   string `yaml:"LINK_SONARQUBE,omitempty"`
	SONARQUBE_URL    string `yaml:"SONARQUBE_URL,omitempty"`
	SONARQUBE_KEY    string `yaml:"SONARQUBE_KEY,omitempty"`
	JAR_APP_NAME     string `yaml:"JAR_APP_NAME,omitempty"`
	CREDENTIAL_HELM  string `yaml:"CREDENTIAL_HELM,omitempty"`
}

func (j *Jenkins) New(app string, tag string, businessName string) {
	j.APPLICATION_NAME = app
	j.BUSINESS_NAME = businessName
	j.TESTING_TAG = tag
	j.FILE = fmt.Sprintf("Jenkinsfile-%s", j.TESTING_TAG)
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
