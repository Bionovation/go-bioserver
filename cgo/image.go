package cgo

/*
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>

void test(){
    printf("hello there!\n");
}

int readfile(char* path, char* buf, int32_t len){
	FILE* pf = fopen(path, "rb");
    if(pf == NULL) return 0;
    fseek(pf, 0, SEEK_END);
    long lSize = ftell(pf);
    if(lSize > len) return 0;
    rewind(pf);
    fread(buf, sizeof(char), lSize, pf);
    fclose(pf);
    return lSize;
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func Read(imgPath string) ([]byte, error)  {
	cstr1 := C.CString(imgPath)
	defer C.free(unsafe.Pointer(cstr1))
	b := make([]byte, 1024*1024) // 1MB
	imgSize := C.readfile(cstr1, (*C.char)(unsafe.Pointer(&b[0])), C.int32_t(len(b)))
	fmt.Println(imgSize)

	b = b[:imgSize]
	return b, nil
}

