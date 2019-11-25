//package go_bioserver
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const BioFolder = "F:/BioSlides"

func main() {
	r := gin.Default()

	r.StaticFS("/html", http.Dir("./html"))
	r.StaticFS("/openseadragon", http.Dir("./html/openseadragon"))

	r.GET("/ping", handlePing)
	r.GET("/image", handleImage)
	r.GET("/slides", handleSlideList)
	r.GET("/slideinfo", handleSlideInfo)
	r.GET("/slidetile", handleSlideTile)
	r.GET("/slidenail", handleSlideNail)

	r.GET("/lua", handleLua)

	r.Run(":9999")
}
