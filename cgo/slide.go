package cgo

/*
#include "C:/BioImgServer/BioImgCore/BioImgCore/src/GoInterfaces.h"

#cgo windows LDFLAGS: -LC:/BioImgServer/BioImgCore/x64/Release -lBioImgCore
*/
import "C"
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"unsafe"
)

// 获取玻片的切片
func SlideTile(path string, level, x, y int) ([]byte, error) {
	r := append([]rune(path), 0)
	c := (*C.int)(unsafe.Pointer(&r[0]))

	bsize := 256 * 256 * 4
	buf := make([]byte, bsize)
	bs := C.ReadSlideTile(c, C.int(level), C.int(x), C.int(y), 256, (*C.char)(unsafe.Pointer(&buf[0])), C.int(bsize))
	if bs < 0 {
		return nil, fmt.Errorf("read slide tile falied.")
	}
	buf = buf[:bs]

	return buf, nil
}

func SlideWidthHeight(path string) (int, int, error) {
	r := append([]rune(path), 0)
	c := (*C.int)(unsafe.Pointer(&r[0]))

	type Info struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	}

	b := make([]byte, 1024*1024)
	cinfo := (*C.char)(unsafe.Pointer(&b[0]))
	sz := C.ReadSlideInfo(c, cinfo)

	ginfo := C.GoString(cinfo)

	if sz < 0 || ginfo == "" {
		return 0, 0, fmt.Errorf("read slide info failed.")
	}

	info := Info{}
	if err := json.Unmarshal([]byte(ginfo), &info); err != nil {
		return 0, 0, fmt.Errorf("unmarshal slideinfo failed.")
	}

	fmt.Println(info.Width, info.Height)

	return info.Width, info.Height, nil
}

// 获取缩略图
func SlideNail(path string) ([]byte, error) {

	nailPath := filepath.Join(path, "smallPic.jpg")

	fmt.Println(nailPath)

	pn, err := os.Open(nailPath)
	if err != nil {
		return nil, err
	}
	defer pn.Close()

	buf, err := ioutil.ReadAll(pn)
	if err != nil {
		return nil, err
	}

	return buf, nil

}

// 释放内存
func SlideClose(path string) error {
	r := append([]rune(path), 0)
	c := (*C.int)(unsafe.Pointer(&r[0]))

	C.CloseSlide(c)

	return nil
}
