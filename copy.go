package main

import (
	"fmt"
	"github.com/spf13/afero"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func CopyFile(src, dst string, destFs *afero.Fs) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := (*destFs).Create(dst)
	if err != nil {
		return
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		return
	}

	err = out.Sync()
	if err != nil {
		return
	}

	si, err := os.Stat(src)
	if err != nil {
		return
	}
	err = (*destFs).Chmod(dst, si.Mode())
	if err != nil {
		return
	}

	return
}

func CopyDir(src string, dst string, destFs *afero.Fs) (count int, err error) {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	si, err := os.Stat(src)
	if err != nil {
		return count, err
	}
	if !si.IsDir() {
		return count, fmt.Errorf("source is not a directory")
	}

	err = (*destFs).MkdirAll(dst, si.Mode())
	if err != nil {
		return
	}

	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return
	}


	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			var subdirCount, subDirErr = CopyDir(srcPath, dstPath, destFs)
			if subDirErr != nil {
				return
			}
			count += subdirCount
		} else {
			// Skip symlinks.
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}

			err = CopyFile(srcPath, dstPath, destFs)
			if err != nil {
				return
			}
			count += 1
		}
	}

	return
}
