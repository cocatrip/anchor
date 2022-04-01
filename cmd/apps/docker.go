package apps

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"strings"
)

type Docker struct {
	APPLICATION_NAME string
	BUSINESS_NAME    string
	TESTING_TAG      string
	FILE             string
	JAR_SNAPSHOT     string `yaml:"JAR_SNAPSHOT,omitempty"`
	JAR_ELASTIC      string `yaml:"JAR_ELASTIC,omitempty"`
	ELASTIC_URL      string `yaml:"ELASTIC_URL,omitempty"`
	ELASTIC_PORT     string `yaml:"ELASTIC_PORT,omitempty"`
}

func (d *Docker) New(app string, tag string, businessName string) {
	d.APPLICATION_NAME = app
	d.BUSINESS_NAME = businessName
	d.TESTING_TAG = tag
	d.FILE = fmt.Sprintf("Dockerfile-%s", d.TESTING_TAG)
}

func (d *Docker) GetValues() ([]string, []string) {
	value := reflect.ValueOf(d).Elem()
	field := value.Type()

	var v, f []string
	for i := 0; i < value.NumField(); i++ {
		v = append(v, value.Field(i).String())
		f = append(f, field.Field(i).Name)
	}
	return v, f
}

func (d *Docker) GetValuesLength() int {
	value := reflect.ValueOf(d).Elem()
	return value.NumField()
}

func (d *Docker) ReadDocker() [][]string {
	dockerByte, err := ioutil.ReadFile("Dockerfile")
	if err != nil {
		panic(err)
	}
	dockerFile := string(dockerByte)

	re, err := regexp.Compile(`%\{(\w+)\}`)
	if err != nil {
		fmt.Println(err)
	}

	values := re.FindAllStringSubmatch(dockerFile, -1)

	return values
}

func (d *Docker) TemplateDocker() string {
	dockerByte, err := ioutil.ReadFile("Dockerfile")
	if err != nil {
		panic(err)
	}
	dockerFile := string(dockerByte)

	configValue, configField := d.GetValues()
	dockerValue := d.ReadDocker()

	for i := 0; i < d.GetValuesLength(); i++ {
		for k := 0; k < len(d.ReadDocker()); k++ {
			if configField[i] == dockerValue[k][1] {
				dockerFile = strings.ReplaceAll(dockerFile, dockerValue[k][0], configValue[i])
			}
		}
	}

	return dockerFile
}
