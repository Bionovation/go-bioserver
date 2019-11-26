//package go_bioserver
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
)

// 配置文件
const cfile = "./config.toml"

// 参数
type BioConfig struct {
	DataFolder string
	ListenPort int
}

var bioConfig = BioConfig{
	DataFolder: "D:/BioScan",
	ListenPort: 9999,
}

func readConfig(conFile string) {
	newConfig := BioConfig{}
	if _, err := toml.DecodeFile(conFile, &newConfig); err != nil {
		log.Println("read config failed:", err)
		return
	}

	bioConfig.DataFolder = newConfig.DataFolder
	bioConfig.ListenPort = newConfig.ListenPort
	log.Printf("read config success, config: %v", bioConfig)
}

func main() {
	//frpLogin()
	readConfig(cfile)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.StaticFS("/html", http.Dir("./html"))
	r.StaticFS("/openseadragon", http.Dir("./html/openseadragon"))

	r.GET("/host", handleHost)
	r.GET("/ping", handlePing)
	r.GET("/image", handleImage)
	r.GET("/slides", handleSlideList)
	r.GET("/slideinfo", handleSlideInfo)
	r.GET("/slidetile", handleSlideTile)
	r.GET("/slidenail", handleSlideNail)

	r.GET("/lua", handleLua)

	r.Run(fmt.Sprintf(":%v", bioConfig.ListenPort))
}
