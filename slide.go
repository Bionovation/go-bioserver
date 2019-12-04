package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type Slide struct {
	path string
	size int
	time time.Time
}

func SlideList(folder string) ([]string, error) {
	dir, err := ioutil.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)

	for _, fi := range dir {
		if !fi.IsDir() {
			continue
		}
		fp := filepath.Join(folder, fi.Name())
		fn := filepath.Join(fp, "data.bimg")
		fn2 := filepath.Join(fp, "downlayer.bimg")
		if finfo, err := os.Stat(fn); !os.IsNotExist(err) {
			if _, err = os.Stat(fn2); !os.IsNotExist(err) {
				//sl = append(sl, fp)
				m[finfo.ModTime().String()] = fp
			}
		}
	}

	// sort
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}

	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	sl := make([]string, 0)
	for _, k := range keys {
		sl = append(sl, m[k])
	}

	return sl, nil
}
