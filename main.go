//package go_bioserver
package main

/*
#include "ccode.h"
*/
import "C"
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	C.test();
	r := gin.Default()
	r.GET("/ping",handlePing)
	r.Run(":9999")
}

func handlePing(c *gin.Context)  {
	c.String(http.StatusOK,"pong!")
}