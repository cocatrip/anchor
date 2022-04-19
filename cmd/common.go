// commonly used function in cmd package
package cmd

import (
	"io/ioutil"

	"github.com/cocatrip/anchor/cmd/apps"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// function to read config to config struct
func readConfig(c *apps.Config) error {
	// read config files used
	f, err := ioutil.ReadFile(viper.ConfigFileUsed())
	if err != nil {
		return err
	}

	// unmarshal it to a variables
	if err := yaml.Unmarshal(f, c); err != nil {
		return err
	}

	return nil
}

