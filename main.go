//package go_bioserver
package main

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
	//"fmt"
	//"io/ioutil"
	"net/http"
	//"os"

	"unsafe"

	"github.com/gin-gonic/gin"
)

func main() {
	C.test()
	r := gin.Default()
	r.GET("/ping", handlePing)
	r.GET("/image", handleImage)
	r.Run(":9999")
}

func handlePing(c *gin.Context) {
	c.String(http.StatusOK, "pong!")
}

func handleImage(c *gin.Context) {
	str1 := "D:\\256.jpg"
	cstr1 := C.CString(str1)
	defer C.free(unsafe.Pointer(cstr1))
	b := make([]byte, 1024*1024) // 1MB
	imgSize := C.readfile(cstr1, (*C.char)(unsafe.Pointer(&b[0])), C.int32_t(len(b)))
	fmt.Println(imgSize)

	b = b[:imgSize]

	c.Data(200, "image/jpeg", b)

	/*file, _ := os.Open(str1)
	b, _ := ioutil.ReadAll(file)
	c.Data(200, "image/jpeg", b)*/
}
