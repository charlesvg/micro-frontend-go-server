package main

import (
	"fmt"
	"github.com/spf13/afero"
	"log"
	"micro-frontend-go-server/internal"
	"net/http"
	"os"
	"time"
)

const HttpPort = ":8080"

func main() {
	var memFs = afero.NewMemMapFs()
	var httpFs = internal.NewFileSystemMapping(&memFs)

	start := time.Now()
	var filesCopiedCount, _ = internal.CopyDir("./web", "/", &memFs)
	log.Println("Copied", filesCopiedCount, "files to memory in", time.Since(start))
	log.Println("Size", fmt.Sprintf("%.2f", DirSizeMB("/", &memFs)), "mb")

	log.Println("Server listening on port", HttpPort)
	log.Fatal(http.ListenAndServe(HttpPort, customHeaders(http.FileServer(httpFs))))
}

func customHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("hello", "world")
		next.ServeHTTP(w, r)
	})
}

func DirSizeMB(path string, fs *afero.Fs) float64 {
	var dirSize int64 = 0

	readSize := func(path string, file os.FileInfo, err error) error {
		if !file.IsDir() {
			dirSize += file.Size()
		}

		return nil
	}

	afero.Walk(*fs, path, readSize)


	sizeMB := float64(dirSize) / 1024.0 / 1024.0

	return sizeMB
}
