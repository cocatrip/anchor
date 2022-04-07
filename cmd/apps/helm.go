package apps

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/cocatrip/anchor/pkg/common"
	"github.com/cocatrip/anchor/pkg/files"
)

// Create helm directory and put contents into it.
// Pass --no-secrets to not generate secret.yaml
func InitHelm(c Config) {
	appName := fmt.Sprintf("%v", c.Global["APPLICATION_NAME"])

	helmDir := "helm"
	chartDir := fmt.Sprintf("%s/%s", helmDir, appName)
	templateDir := fmt.Sprintf("%s/templates", chartDir)

	isNoSecretStr := fmt.Sprintf("%t", c.Helm["isNoSecret"])
	isNoSecret, err := strconv.ParseBool(isNoSecretStr)
	if err != nil {
		panic(err)
	}

	_, err = os.Stat(helmDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(helmDir, 0755)
		if err != nil {
			panic(err)
		}
	}

	helmCreate := exec.Command("helm", "create", appName)
	helmCreate.Dir = helmDir

	if err := helmCreate.Run(); err != nil {
		panic(err)
	}

	removeList := []string{
		templateDir + "/deployment.yaml",
		templateDir + "/service.yaml",
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
