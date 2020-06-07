package main

import (
	"fmt"
	"github.com/spf13/afero"
	"log"
	"micro-frontend-go-server/internal"
	"net/http"
	"time"
)

const HttpPort = ":8080";

func main() {
	var memFs = afero.NewMemMapFs()
	var httpFs = internal.NewFileSystemMapping(&memFs)

	start := time.Now()
	var filesCopiedCount, _ = internal.CopyDir("./web", "/", &memFs)
	fmt.Println("Copied", filesCopiedCount, "files to memory in", time.Since(start))

	fmt.Println("Server listening on port", HttpPort)
	log.Fatal(http.ListenAndServe(HttpPort, http.FileServer(httpFs)))
}
