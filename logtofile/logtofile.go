// Package logtofile is used to log to a file.
package logtofile

import (
	"log"
	"os"
)

var fileToWriteTo string

// FileExists check that a file exists.
func FileExists(fileName string) bool {
	fileToWriteTo = fileName
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// CreateFile creates the file at the specified location.
func CreateFile() error {
	fo, err := os.Create(fileToWriteTo)
	if err != nil {
		return err
	}
	defer func() {
		fo.Close()
	}()
	return nil
}

// WriteMessage writes the message.
func WriteMessage(message string) {
	f, err := os.OpenFile(fileToWriteTo, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(message)
}
