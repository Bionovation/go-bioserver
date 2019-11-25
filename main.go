//package go_bioserver
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const BioFolder = "D:/BioScan"

func main() {
	r := gin.Default()

	r.StaticFS("/", http.Dir("./html"))

	r.GET("/ping", handlePing)
	r.GET("/image", handleImage)
	r.GET("/slides", handleSlideList)
	r.GET("/slideinfo", handleSlideInfo)
	r.GET("/slidetile", handleSlideTile)

	r.GET("/lua", handleLua)

	r.Run(":9999")
}
