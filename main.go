//package go_bioserver
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 配置文件
const cfile = "./config.toml"

func main() {
	bioConfig.readConfig(cfile) // 读取配置文件
	go frpLogin()               // 登录frp代理服务
	go clearRoutine(nil)        // 运行内存清理线程

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.StaticFS("/html", http.Dir("./html"))
	r.StaticFS("/openseadragon", http.Dir("./html/openseadragon"))

	r.GET("/", handleIndex)
	r.GET("/host", handleHost)
	r.GET("/ping", handlePing)
	r.GET("/image", handleImage)
	r.GET("/slides", handleSlideList)
	r.GET("/slideinfo", handleSlideInfo)
	r.GET("/slidetile", handleSlideTile)
	r.GET("/slidenail", handleSlideNail)

	r.GET("/lua", handleLua)

	r.Run(fmt.Sprintf(":%v", bioConfig.Common.ListenPort))
}
