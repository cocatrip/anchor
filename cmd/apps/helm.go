package apps

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/cocatrip/anchor/pkg/common"
	"github.com/cocatrip/anchor/pkg/files/maven"
)

// Create helm directory and put contents into it.
// Pass --no-secrets to not generate secret.yaml
func InitHelm(c Config) error {
	appName := fmt.Sprintf("%v", c.Global["APPLICATION_NAME"])

	helmDir := "helm"
	chartDir := fmt.Sprintf("%s/%s", helmDir, appName)
	templateDir := fmt.Sprintf("%s/templates", chartDir)

	isNoSecretStr := fmt.Sprintf("%t", c.Helm["isNoSecret"])
	isNoSecret, err := strconv.ParseBool(isNoSecretStr)
	if err != nil {
		return err
	}

	_, err = os.Stat(helmDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(helmDir, 0755)
		if err != nil {
			return err
		}
	}

	helmCreate := exec.Command("helm", "create", appName)
	helmCreate.Dir = helmDir

	if err := helmCreate.Run(); err != nil {
		return err
	}

	removeList := []string{
		templateDir + "/deployment.yaml",
		templateDir + "/service.yaml",
		chartDir + "/values.yaml",
		chartDir + "/charts",
	}

	for _, i := range removeList {
		if err := os.RemoveAll(i); err != nil {
			return err
		}
	}

	common.SaveFile(templateDir+"/deployment.yaml", maven.Deployment)
	c.Template(templateDir+"/deployment.yaml", templateDir+"/deployment.yaml")

	common.SaveFile(templateDir+"/service.yaml", maven.Service)
	c.Template(templateDir+"/service.yaml", templateDir+"/service.yaml")

	common.SaveFile(templateDir+"/configmap.yaml", maven.ConfigMap)
	c.Template(templateDir+"/configmap.yaml", templateDir+"/configmap.yaml")

	if !isNoSecret {
		common.SaveFile(templateDir+"/secret.yaml", maven.Secret)
	}

	common.SaveFile("helm/values.yaml", maven.Values)

	return nil
}
