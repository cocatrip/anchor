// commonly used function in apps package
package apps

import (
	"os"

	"github.com/fatih/color"
)

// Success color
func success(s string) {
	color := color.New(color.FgGreen, color.Bold)
	color.Fprintf(os.Stdout, "%s\n", s)
}

// Save a string to a file
func saveFile(fileName string, content string) error {
	// check if file exist
	if _, err := os.Stat(fileName); err != nil {
		// if not exist then create file
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer file.Close()

		// save the file
		_, err = file.WriteAt([]byte(content), 0)
		if err != nil {
			return err
		}
	} else {
		// if file exist then open file
		file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
		if err != nil {
			return err
		}
		defer file.Close()

		// save the file
		_, err = file.WriteAt([]byte(content), 0)
		if err != nil {
			return err
		}
	}

	return nil
}
