package apps

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"strings"

	"github.com/cocatrip/anchor/pkg/common"
	"github.com/cocatrip/anchor/pkg/files"
)

type Helm struct {
	APPLICATION_NAME   string
	BUSINESS_NAME      string
	TESTING_TAG        string
	SERVER_NAME        string
	FILE               string
	Version_Major      string `yaml:"Version_Major,omitempty"`
	Version_Minor      string `yaml:"Version_Minor,omitempty"`
	Version_Patch      string `yaml:"Version_Patch,omitempty"`
	BUILD_TIMESTAMP    string `yaml:"BUILD_TIMESTAMP,omitempty"`
	BUILD_NUMBER       string `yaml:"BUILD_NUMBER,omitempty"`
	SERVICE_TYPE       string `yaml:"SERVICE_TYPE,omitempty"`
	SERVICE_PORT       string `yaml:"SERVICE_PORT,omitempty"`
	SERVICE_TARGETPORT string `yaml:"SERVICE_TARGETPORT,omitempty"`
	CPU_LIMIT          string `yaml:"CPU_LIMIT,omitempty"`
	MEM_LIMIT          string `yaml:"MEM_LIMIT,omitempty"`
	CPU_REQUEST        string `yaml:"CPU_REQUEST,omitempty"`
	MEM_REQUEST        string `yaml:"MEM_REQUEST,omitempty"`
}

func (h *Helm) New(c Config) {
	h.APPLICATION_NAME = c.APPLICATION_NAME
	h.BUSINESS_NAME = c.BUSINESS_NAME
	h.TESTING_TAG = c.TESTING_TAG
	h.SERVER_NAME = c.SERVER_NAME
	h.FILE = fmt.Sprintf("values-%s.yaml", h.TESTING_TAG)
}

func (h *Helm) GetValues() ([]string, []string) {
	value := reflect.ValueOf(h).Elem()
	field := value.Type()

	var v, f []string
	for i := 0; i < value.NumField(); i++ {
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
	helmByte, err := ioutil.ReadFile("helm/" + h.APPLICATION_NAME + "/values.yaml")
	if err != nil {
		panic(err)
	}
	helmFile := string(helmByte)

	re, err := regexp.Compile(`%\{(\w+)\}`)
	if err != nil {
		fmt.Println(err)
	}

	values := re.FindAllStringSubmatch(helmFile, -1)

	return values
}

func (h *Helm) TemplateHelm() string {
	helmByte, err := ioutil.ReadFile("helm/" + h.APPLICATION_NAME + "/values.yaml")
	if err != nil {
		panic(err)
	}
	helmFile := string(helmByte)

	configValue, configField := h.GetValues()
	helmValue := h.ReadHelm()

	for i := 0; i < h.GetValuesLength(); i++ {
		for k := 0; k < len(h.ReadHelm()); k++ {
			if configField[i] == helmValue[k][1] {
				helmFile = strings.ReplaceAll(helmFile, helmValue[k][0], configValue[i])
			}
		}
	}

	return helmFile
}

func (h *Helm) InitHelm() {
	helmDir := "helm"
	chartDir := fmt.Sprintf("%s/%s", helmDir, h.APPLICATION_NAME)
	templateDir := fmt.Sprintf("%s/templates", chartDir)

	// create directory
	_, err := os.Stat(helmDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(helmDir, 0755)
		if err != nil {
			panic(err)
		}
	}

	// helm create inside helm dir
	helmCreate := exec.Command("helm", "create", h.APPLICATION_NAME)
	helmCreate.Dir = helmDir

	if err := helmCreate.Run(); err != nil {
		panic(err)
	}

	removeList := []string{
		templateDir + "/deployment.yaml",
		templateDir + "/service.yaml",
		templateDir + "/tests",
		chartDir + "/values.yaml",
		chartDir + "/charts",
	}

	for _, i := range removeList {
		if err := os.RemoveAll(i); err != nil {
			panic(err)
		}
	}

	files.Deployment = strings.ReplaceAll(files.Deployment, "%{APPLICATION_NAME}", h.APPLICATION_NAME)
	common.SaveFile(templateDir+"/deployment.yaml", files.Deployment)

	files.Service = strings.ReplaceAll(files.Service, "%{APPLICATION_NAME}", h.APPLICATION_NAME)
	common.SaveFile(templateDir+"/service.yaml", files.Service)

	common.SaveFile(chartDir+"/values.yaml", files.Values)
}
