package helper

import (
	"encoding/csv"
	"log"
	"os"
)

const ROOT_FOLDER = "./seeds/files/"

func Write(filename string, data [][]string) {

	var file *os.File
	var err error
	if fileExists(ROOT_FOLDER + filename) {
		file, err = os.OpenFile(ROOT_FOLDER+filename,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	} else {
		file, err = os.Create(ROOT_FOLDER + filename)
	}
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
