package main

import (
	"fmt"
	"github.com/spf13/afero"
	"log"
	"net/http"
)


func main() {

	fmt.Println("app started")
	var aferoMemFs = afero.NewMemMapFs()
	var httpFs = NewFileSystemMapping(&aferoMemFs)

	var filesCopiedCount, _ = CopyDir("./files", "/", &aferoMemFs)
	fmt.Println("copied", filesCopiedCount, "files to memory" )

	fmt.Println("serving files")
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(httpFs)))


}
