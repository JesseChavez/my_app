package resources

import (
	"embed"
	"log"
	"os"
)

//go:embed config/database.yml
//go:embed public/*
//go:embed app/templates/*
var files embed.FS

func LoadFS() embed.FS {
	return files
}

func BaseDir() string {
	path, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	return path
}
