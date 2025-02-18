package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
)

//go:embed version.txt
var version string

//go:embed disc.png
var disc []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := os.WriteFile("disc_new_1.png", disc, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dirEntries, _ := path.ReadDir("files")
	
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/"+entry.Name())
			fmt.Println(string(file))
		}
	}
}