package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"micro-frontend-go-server/internal"
	"net/http"
	"os"
	"time"
)

const HttpPort = ":8080"
const contextPath = "/test/"

func initlog() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{ ForceColors: true })
}

func main() {

	initlog()

	var memFs = afero.NewMemMapFs()
	var httpFs = internal.NewFileSystemMapping(&memFs)

	start := time.Now()
	filesCopiedCount, _ := internal.CopyDir("./web", "/", &memFs)
	log.Println("Copied", filesCopiedCount, "files (",fmt.Sprintf("%.2f", DirSizeMB("/", &memFs)), "mb ) to memory in", time.Since(start))

	log.Println("Server listening on port", HttpPort, "and context path", contextPath)
	http.Handle(contextPath,  http.StripPrefix(contextPath, customHeaders(http.FileServer(httpFs))))
	http.ListenAndServe(HttpPort, nil)
}

func customHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer duration(track("Serving ", r.URL))
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

func track(msg ...interface{}) ( time.Time, string) {
	return time.Now(), fmt.Sprint(msg...)
}

func duration(start time.Time, msg string ) {
	log.Println(msg,"took", time.Since(start))
}