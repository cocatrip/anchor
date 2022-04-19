package cmd

import (
	"fmt"
	"os"

	"github.com/cocatrip/anchor/cmd/apps"
)

// function for 'anchor clean jenkins'
func cleanJenkins() error {
	var config apps.Config

	// read from config file and put it to config struct
	if err := readConfig(&config); err != nil {
		return err
	}

	// define the resultFileName so it can be deleted
	resultFileName := fmt.Sprintf("Jenkinsfile-%s", config.Global["TESTING_TAG"])

	// list of to be removed files
	removeList := []string{
		"Jenkinsfile",
		resultFileName,
	}

	// loop through the list and delete file one by one
	for _, v := range removeList {
		// check if exist if not then don't delete so os.Remove will never return err
		_, err := os.Stat(v)
		if !os.IsNotExist(err) {
			os.Remove(v)
			fmt.Printf("Removing %s\n", v)
		}
	}

	return nil
}

// function for 'anchor clean docker'
func cleanDocker() error {
	var config apps.Config

	// read from config file and put it to config struct
	if err := readConfig(&config); err != nil {
		return err
	}

	// define the resultFileName so it can be deleted
	resultFileName := fmt.Sprintf("Dockerfile-%s", config.Global["TESTING_TAG"])

	// list of to be removed files
	removeList := []string{
		"Dockerfile",
		resultFileName,
	}

	// loop through the list and delete file one by one
	for _, v := range removeList {
		// check if exist if not then don't delete so os.Remove will never return err
		_, err := os.Stat(v)
		if !os.IsNotExist(err) {
			os.Remove(v)
			fmt.Printf("Removing %s\n", v)
		}
	}

	return nil
}

// function for 'anchor clean helm'
func cleanHelm() error {
	var config apps.Config

	// read from config file and put it to config struct
	if err := readConfig(&config); err != nil {
		return err
	}

	// simply check if helm directory exist or not
	_, err := os.Stat("helm")
	// if exist then recursively remove the directory
	if !os.IsNotExist(err) {
		os.RemoveAll("helm")
		fmt.Printf("Removing helm directory\n")
	}

	return nil
}

