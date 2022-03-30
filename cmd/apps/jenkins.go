package apps

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"strings"
	"unicode"
)

type Jenkins struct {
	File            string `yaml:"file"`
	ImageName       string `yaml:"imageName"`
	ImageTag        string `yaml:"imageTag"`
  BuildContainer  string `yaml:"buildContainer"`
  BuildStep       string `yaml:"buildStep"`
}

func (j *Jenkins) GetValues() ([]string, []string) {
  value := reflect.ValueOf(j).Elem()
  field := value.Type()
  
  var v, f []string
  for i:=0; i<value.NumField(); i++ {
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
  jenkinsByte, err := ioutil.ReadFile(j.File)
  if err != nil {
    panic(err)
  }; jenkinsFile := string(jenkinsByte)

  re, err := regexp.Compile(`%\{(\w+)\}`)
  if err != nil {
    fmt.Println(err)
  }

  values := re.FindAllStringSubmatch(jenkinsFile, -1)

  return values
}

func format(s string) string {
  a := []rune(s)
  a[0] = unicode.ToLower(a[0])
  s = string(a)
  return s
}

func (j *Jenkins) TemplateJenkins() string {
  jenkinsByte, err := ioutil.ReadFile(j.File)
  if err != nil {
    panic(err)
  }; jenkinsFile := string(jenkinsByte)
  
  configValue, configField := j.GetValues()
  jenkinsValue := j.ReadJenkins()

  for i := 0; i < j.GetValuesLength(); i++ {
    for k := 0; k < len(j.ReadJenkins()); k++ {
      if format(configField[i]) == "buildStep" {
        if configValue[i] == "" {
          if j.BuildContainer == "maven" {
            configValue[i] += "mvn clean package"
          }

          if j.BuildContainer == "node" {
            configValue[i] += `npm install
            npm run build`
          }
        }

        fmt.Println(configValue[i])
      }
      if format(configField[i]) == jenkinsValue[k][1] {
        jenkinsFile = strings.ReplaceAll(jenkinsFile, jenkinsValue[k][0], configValue[i])
      }
    }
  }

  fmt.Println(jenkinsFile)
  return jenkinsFile
}
