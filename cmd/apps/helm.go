package apps

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/cocatrip/anchor/pkg/common"
	"github.com/cocatrip/anchor/pkg/files"
)

func InitHelm(c Config, isNoSecret bool) {
	// testingTag := fmt.Sprintf("%v", c.Global["TESTING_TAG"])
	appName := fmt.Sprintf("%v", c.Global["APPLICATION_NAME"])
	helmDir := "helm"
	chartDir := fmt.Sprintf("%s/%s", helmDir, appName)
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
	helmCreate := exec.Command("helm", "create", appName)
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

	common.SaveFile(templateDir+"/deployment.yaml", files.Deployment)
	c.Template(templateDir+"/deployment.yaml", templateDir+"/deployment.yaml")

	common.SaveFile(templateDir+"/service.yaml", files.Service)
	c.Template(templateDir+"/service.yaml", templateDir+"/service.yaml")

	common.SaveFile(templateDir+"/configmap.yaml", files.ConfigMap)
	c.Template(templateDir+"/configmap.yaml", templateDir+"/configmap.yaml")

	if !isNoSecret {
		common.SaveFile(templateDir+"/secret.yaml", files.Secret)
	}

	common.SaveFile(chartDir+"/values.yaml", files.Values)
}
