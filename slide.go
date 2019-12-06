package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type Slide struct {
	Levels  int     `json:"levels"`
	Width   int     `json:"width"`
	Height  int     `json:"height"`
	Mag     int     `json:"mag"`
	TimeUse float32 `json:"timeuse"`

	FileSize   string
	Version    string
	ScanFolder string
	CellCount  int
	CreateTime time.Time
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
		fn3 := filepath.Join(fp, "slideinfo.bic")
		fn4 := filepath.Join(fp, fmt.Sprintf("%v.bic", fi.Name()))
		if finfo, err := os.Stat(fn); !os.IsNotExist(err) {
			if _, err = os.Stat(fn2); !os.IsNotExist(err) {
				if _, err = os.Stat(fn3); !os.IsNotExist(err) {
					m[finfo.ModTime().String()] = fp
				} else if _, err = os.Stat(fn4); !os.IsNotExist(err) {
					m[finfo.ModTime().String()] = fp
				}
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
