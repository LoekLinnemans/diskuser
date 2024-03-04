package main

import (
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

func scanFiles(dir string) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println("error finding directory")
			return err
		}
		log.Printf("Path: %v, Name: %v, Size: %v bytes\n", path, info.Name(), info.Size())
		return nil
	})
}

func main() {
	directory := "C:/Users/username/Downloads/"
	scanFiles(directory)
}
