package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type Slide struct {
	path string
	size int
}

func SlideList(folder string) ([]string, error) {
	dir, err := ioutil.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	sl := make([]string, 0)
	for _, fi := range dir {
		if !fi.IsDir() {
			continue
		}
		fp := filepath.Join(folder, fi.Name())
		fn := filepath.Join(fp, "data.bimg")
		fn2 := filepath.Join(fp, "downlayer.bimg")
		if _, err = os.Stat(fn); !os.IsNotExist(err) {
			if _, err = os.Stat(fn2); !os.IsNotExist(err) {
				sl = append(sl, fp)
			}
		}
	}
	return sl, nil
}
