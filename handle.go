package main

import (
	"encoding/json"
	//"log"
	"net/http"
	"strconv"

	"github.com/yuin/gopher-lua"

	"github.com/Bionovation/go-bioserver/cgo"
	"github.com/gin-gonic/gin"
)

func handleIndex(c *gin.Context) {
	c.Redirect(http.StatusFound, "/html/")
}

// 服务器状态测试
func handlePing(c *gin.Context) {
	c.String(http.StatusOK, "pong!")
}

// 返回服务器地址
func handleHost(c *gin.Context) {
	c.String(http.StatusOK, c.Request.Host)
}

// 测试图像数据读取
func handleImage(c *gin.Context) {
	b, err := cgo.Read("D:/256.jpg")
	if err != nil {
		panic(err)
	}
	c.Data(http.StatusOK, "image/jpeg", b)
}

// 获取扫描数据列表
func handleSlideList(c *gin.Context) {
	sl, err := SlideList(bioConfig.Common.DataFolder)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, sl)
	}

}

func handleSlideInfo(c *gin.Context) {
	type SlideInfo struct {
		Levels int `json:"levels"`
		Width  int `json:"width"`
		Height int `json:"height"`
	}

	path := c.Query("path")
	info, err := cgo.SlideInfo(path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	tinfo := SlideInfo{}
	err = json.Unmarshal([]byte(info), &tinfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	bioGC.Visit(path) // gc

	c.JSON(http.StatusOK, tinfo)
}

// 读取瓦片
func handleSlideTile(c *gin.Context) {
	path := c.Query("path")
	sz := c.Query("level")
	sx := c.Query("x")
	sy := c.Query("y")

	z, _ := strconv.Atoi(sz)
	x, _ := strconv.Atoi(sx)
	y, _ := strconv.Atoi(sy)
	buf, err := cgo.SlideTile(path, z, x, y)
	if err != nil {
		c.String(http.StatusInternalServerError, "读取切片失败.")
		return
	}

	bioGC.Visit(path) // gc

	c.Data(http.StatusOK, "image/jpeg", buf)
}

func handleSlideNail(c *gin.Context) {
	path := c.Query("path")
	buf, err := cgo.SlideNail(path)
	if err != nil {
		c.String(http.StatusInternalServerError, "读取缩略图失败.")
		return
	}

	c.Data(http.StatusOK, "image/jpeg", buf)
}

func handleLua(c *gin.Context) {
	l := lua.NewState()
	defer l.Close()
	l.DoFile("./test.lua")
	c.String(http.StatusOK, "ok")
}
