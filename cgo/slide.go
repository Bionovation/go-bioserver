package cgo

/*
#include "C:/BioImgServer/BioImgCore/BioImgCore/src/GoInterfaces.h"

#cgo windows LDFLAGS: -LC:/BioImgServer/BioImgCore/x64/Release -lBioImgCore
*/
import "C"
import (
	"fmt"
	"os"
	"unsafe"
)

// 获取玻片的切片
func SlideTile(path string, level, x, y int) ([]byte, error) {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	bsize := 256 * 256 * 4
	buf := make([]byte, bsize)
	bs := C.ReadSlideTile(cpath, C.int(level), C.int(x), C.int(y), 256, (*C.char)(unsafe.Pointer(&buf[0])), C.int(bsize))
	if bs < 0 {
		return nil, fmt.Errorf("read slide tile falied.")
	}
	buf = buf[:bs]

	return buf, nil
}

// 玻片信息
func SlideInfo(path string) (string, error) {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	b := make([]byte, 1024*1024)
	cinfo := (*C.char)(unsafe.Pointer(&b[0]))
	sz := C.ReadSlideInfo(cpath, cinfo)

	ginfo := C.GoString(cinfo)
	if sz < 0 || ginfo == "" {
		return "", fmt.Errorf("ReadSlideInfo falied.")
	}
	return ginfo, nil

	// slideinfopath := fmt.Sprintf("%s/slideinfo.json", path)
	// inputFile, inputError := os.Open(slideinfopath)
	// if inputError != nil {
	// 	return nil, nil
	// } else {
	// 	return ioutil.ReadAll(f), nil
	// }
	// defer inputFile.Close()

}

// 获取缩略图
func SlideNail(path string) ([]byte, error) {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	bsize := 1024 * 1024 / 2
	buf := make([]byte, bsize)
	bs := C.ReadSlideNail(cpath, (*C.char)(unsafe.Pointer(&buf[0])), C.int(bsize))
	if bs < 0 {
		return nil, fmt.Errorf("read slide tile falied.")
	}
	buf = buf[:bs]

	return buf, nil
}

// 释放内存
func SlideClose(path string) error {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	C.CloseSlide(cpath)
	return nil
}
