package apps

import (
	"os"

	"github.com/cocatrip/anchor/pkg/common"
	"github.com/cocatrip/anchor/pkg/files"
)

// check if Jenkinsfile exists, if it doesn't create it
func InitJenkins() {
	_, err := os.Stat("Jenkinsfile")
	if os.IsNotExist(err) {
		common.SaveFile("Jenkinsfile", files.Jenkinsfile)
	}
}
