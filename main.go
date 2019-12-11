//package go_bioserver
//-ldflags="-H windowsgui"
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	//"github.com/sirupsen/logrus"
)

// 配置文件
const cfile = "./config.toml"

// 日志输出
// var log = logrus.New()

func stdToFile() {
	f, _ := os.OpenFile("./go-bioserver.log", os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0755)
	os.Stdout = f
	os.Stderr = f
	log.SetOutput(f)
}

func main() {
	// 重定向标准输出到文件
	stdToFile()
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

	r.GET("/test", handleTest)

	r.Run(fmt.Sprintf(":%v", bioConfig.Common.ListenPort))
}
