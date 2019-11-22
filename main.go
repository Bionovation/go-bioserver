//package go_bioserver
package main

import (
	"github.com/Bionovation/go-bioserver/handle"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", handle.Ping)
	r.GET("/image", handle.Image)
	r.GET("/slides", handle.SlideList)
	r.GET("/slideinfo", handle.SlideInfo)
	r.GET("/slidetile", handle.SlideTile)

	r.Run(":9999")
}
