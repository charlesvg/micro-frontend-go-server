package internal

import (
	"fmt"
	"github.com/spf13/afero"
	"net/http"
)

type Mapping struct {
	fs afero.Fs
}

func (m Mapping) Open(name string) (http.File, error) {
	fmt.Println("open", name)
	return m.fs.Open(name)
}

func NewFileSystemMapping(fs *afero.Fs) http.FileSystem {
	return &Mapping{*fs}
}