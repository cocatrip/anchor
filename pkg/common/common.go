package common

import "os"

func SaveFile(fileName string, content string) error {
	// check file ada ga
	if _, err := os.Stat(fileName); err != nil {
		// kalo gada create
		file, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// save file
		_, err = file.WriteAt([]byte(content), 0)
		if err != nil {
			panic(err)
		}
	} else {
		// kalo ada open
		file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// save file
		_, err = file.WriteAt([]byte(content), 0)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
