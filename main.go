//package go_bioserver
package main

/*
#include "ccode.h"
*/
import "C"
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"unsafe"
)

func main()  {
	C.test();
	r := gin.Default()
	r.GET("/ping",handlePing)
	r.GET("/image",handleImage)
	r.Run(":9999")
}

func handlePing(c *gin.Context)  {
	c.String(http.StatusOK,"pong!")
}

func handleImage(c *gin.Context)  {
	str1 := "D:/test.jpg"
	cstr1 := C.CString(str1)
	defer C.free(unsafe.Pointer(cstr1))
	b := make([]byte,1024*1024) // 1MB
	C.readfile(cstr1, (*C.char)(unsafe.Pointer(&b[0])), C.int32_t(len(b)))
	c.String(http.StatusOK,"ing!")
}