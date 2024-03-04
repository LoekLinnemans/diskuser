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
		log.Println("couldn't create logfile")
		os.Exit(1)
	}

	writer := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(writer)

}

func ScanFiles(dir string) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println("Error finding directory or file: ", err)
			return err
		}
		WriteToFile(fmt.Sprintf("Path: %v, Name: %v, Size: %v bytes\n", path, info.Name(), info.Size()))
		return nil
	})
}

func WriteToFile(dir string) {
	file, err := os.OpenFile("result.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println("Error creating file:", err)
		return
	}
	_, err = file.WriteString(dir)
	if err != nil {
		log.Println("Error writing to file:", err)
		return
	}
	defer file.Close()
}

func main() {
	directory := "C:\\Users\\loek\\OneDrive\\Documenten\\World Machine Documents"
	ScanFiles(directory)
}
