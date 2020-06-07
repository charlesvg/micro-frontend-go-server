package main

import (
	"fmt"
	"github.com/spf13/afero"
	"io"
	"log"
	"net/http"
	"os"
)


func main() {

	fmt.Println("started")
	var AppFs = afero.NewMemMapFs()
	var appMapping = NewMapping(&AppFs)
	Copy("test.html", "/test.html", &AppFs)
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(appMapping)))


}

// Copy the src file to dst. Any existing file will be overwritten and will not
// copy file attributes.
func Copy(src string, dst string, destFs *afero.Fs) error {

	fmt.Println("copy start")
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := (*destFs).Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	fmt.Println("copied")
	return out.Close()
}