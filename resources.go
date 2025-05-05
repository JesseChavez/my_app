package resources

import (
	"embed"
)

//go:embed config/database.yml
//go:embed public/*
//go:embed app/templates/*
var files embed.FS

func LoadFS() embed.FS {
	return files
}
