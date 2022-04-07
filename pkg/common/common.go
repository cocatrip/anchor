package common

import (
	"os"

	"github.com/fatih/color"
)

// Success color
func Success(s string) {
	color := color.New(color.FgGreen, color.Bold)
	color.Fprintf(os.Stdout, "%s\n", s)
}

// Save a string to a file
func SaveFile(fileName string, content string) error {
	// check file ada ga
	if _, err := os.Stat(fileName); err != nil {
		// kalo gada create
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer file.Close()

		// save file
		_, err = file.WriteAt([]byte(content), 0)
		if err != nil {
			return err
		}
	} else {
		// kalo ada open
		file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
		if err != nil {
			return err
		}
		defer file.Close()

		// save file
		_, err = file.WriteAt([]byte(content), 0)
		if err != nil {
			return err
		}
	}

	return nil
}
