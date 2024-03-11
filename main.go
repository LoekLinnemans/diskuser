package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func init() {
	logFile, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println("couldn't create logfile", err)
		os.Exit(1)
	}

	writer := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(writer)

}

func UserInput() string {

	var input string

	fmt.Println("enter the directory to scan:")

	fmt.Scanln(&input)
	return input
}

func ScanFiles(dir string) {
	var totalSize int64

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println("Error finding directory or file: ", err)
			os.Exit(2)
			return err
		}
		if info.IsDir() {
			totalSize += info.Size()
		}
		return nil
	})
	WriteToFile(fmt.Sprintf("Memory Usage of %s: %v bytes\n", dir, totalSize))
}

func WriteToFile(dir string) {
	file, err := os.OpenFile("result.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println("Error creating result.txt:", err)
		os.Exit(1)
		return
	}
	_, err = file.WriteString(dir)
	if err != nil {
		log.Println("Error writing to file:", err)
		os.Exit(3)
		return
	}
	defer file.Close()
}

func main() {
	directory := UserInput()
	ScanFiles(directory)
}
