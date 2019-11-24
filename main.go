//package go_bioserver
package main

import (
	"github.com/gin-gonic/gin"
)

const BioFolder  = "D:/BioScan"

func main() {
	r := gin.Default()

	//r.Static("/","C:/BioImgServer/Html/")
	//r.StaticFile("/","C:/BioImgServer/Html/")

	r.GET("/ping", handlePing)
	r.GET("/image", handleImage)
	r.GET("/slides", handleSlideList)
	r.GET("/slideinfo", handleSlideInfo)
	r.GET("/slidetile", handleSlideTile)

	r.GET("/lua", handleLua)


	r.Run(":9999")
}
