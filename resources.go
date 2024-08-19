package resources

import (
	"embed"
)

//go:embed public/*
//go:embed app/views/*
var files embed.FS

func LoadFS() embed.FS {
	return files
}
