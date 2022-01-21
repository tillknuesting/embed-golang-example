package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed static
var embeddedFiles embed.FS

func getFileSystem() http.FileSystem {
	log.Println("using embed mode")
	fileSys, err := fs.Sub(embeddedFiles, "static")
	if err != nil {
		log.Fatal(err)
	}

	return http.FS(fileSys)
}

func main() {
	http.Handle("/", http.FileServer(getFileSystem()))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err)
	}
}
