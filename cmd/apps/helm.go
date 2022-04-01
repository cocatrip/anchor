package apps

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"strings"
)

type Helm struct {
  File            string `yaml:"file"`
	ImageRepository string `yaml:"imageRepository"`
	ImageTag        string `yaml:"imageTag"`
}

func (h *Helm) GetValues() ([]string, []string) {
  value := reflect.ValueOf(h).Elem()
  field := value.Type()
  
  var v, f []string
  for i:=0; i<value.NumField(); i++ {
    v = append(v, value.Field(i).String())
    f = append(f, field.Field(i).Name)
  }
  return v, f
}

func (h *Helm) GetValuesLength() int {
  value := reflect.ValueOf(h).Elem()
  return value.NumField()
}

func (h *Helm) ReadHelm() [][]string {
  helmByte, err := ioutil.ReadFile(h.File)
  if err != nil {
    panic(err)
  }; helmFile := string(helmByte)

  re, err := regexp.Compile(`%\{(\w+)\}`)
  if err != nil {
    fmt.Println(err)
  }

  values := re.FindAllStringSubmatch(helmFile, -1)

  return values
}

// func format(s string) string {
//   a := []rune(s)
//   a[0] = unicode.ToLower(a[0])
//   s = string(a)
//   return s
// }

func (h *Helm) TemplateHelm() string {
  helmByte, err := ioutil.ReadFile(h.File)
  if err != nil {
    panic(err)
  }; helmFile := string(helmByte)
  
  configValue, configField := h.GetValues()
  helmValue := h.ReadHelm()

  for i := 0; i < h.GetValuesLength(); i++ {
    for j := 0; j < len(h.ReadHelm()); j++ {
      if format(configField[i]) == helmValue[j][1] {
        helmFile = strings.ReplaceAll(helmFile, helmValue[j][0], configValue[i])
      }
    }
  }

  return helmFile
}
