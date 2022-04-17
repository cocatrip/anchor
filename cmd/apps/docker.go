package apps

import (
	"os"

	"github.com/cocatrip/anchor/pkg/common"
	"github.com/cocatrip/anchor/pkg/files"
)

// check if Dockerfile exists, if it doesn't create it
func InitDocker() {
	_, err := os.Stat("Dockerfile")
	if os.IsNotExist(err) {
		common.SaveFile("Dockerfile", files.Dockerfile)
	}
}
