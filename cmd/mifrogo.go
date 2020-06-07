package main

import (
	"fmt"
	"github.com/spf13/afero"
	"log"
	"micro-frontend-go-server/internal"
	"net/http"
)


func main() {

	fmt.Println("app started")
	var aferoMemFs = afero.NewMemMapFs()
	var httpFs = internal.NewFileSystemMapping(&aferoMemFs)

	var filesCopiedCount, _ = internal.CopyDir("./web", "/", &aferoMemFs)
	fmt.Println("copied", filesCopiedCount, "web to memory" )

	fmt.Println("serving web")
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(httpFs)))


}
