package main

import (
	"embed"

	helloworld "github.com/maldan/gam-app-passman/internal/app/passman"
)

//go:embed frontend/build/*
var frontFs embed.FS

func main() {
	helloworld.Start(frontFs) //
}
