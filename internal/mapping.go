package internal

import (
	"github.com/spf13/afero"
	"net/http"
)

type Mapping struct {
	fs afero.Fs
}

func (m Mapping) Open(name string) (http.File, error) {
	return m.fs.Open(name)
}

func NewFileSystemMapping(fs *afero.Fs) http.FileSystem {
	return &Mapping{*fs}
}
