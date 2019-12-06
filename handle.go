package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Bionovation/go-bioserver/cgo"
	"github.com/gin-gonic/gin"
	"github.com/yuin/gopher-lua"
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
/*func handleSlideList(c *gin.Context) {
	sl, err := SlideList(bioConfig.Common.DataFolder)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, sl)
	}

}*/

func handleSlideList(c *gin.Context) {
	res := NewRes()
	sl, err := SlideList(bioConfig.Common.DataFolder)
	if err != nil {
		res.FailErr(c, err)
		return
	}

	// 分页显示
	type Page struct {
		Index  int      `json:"index"`
		Count  int      `json:"count"`
		Total  int      `json:"total"`
		Slides []string `json:"slides"`
	}

	page := &Page{}

	page.Total = len(sl)

	if page.Index, err = strconv.Atoi(c.Query("index")); err != nil {
		page.Index = 0
	}

	if page.Count, err = strconv.Atoi(c.Query("count")); err != nil {
		page.Count = 30
	}

	if page.Index >= page.Total {
		page.Count = 0
	} else if page.Index+page.Count >= page.Total {
		page.Count = page.Total - page.Index
	}

	page.Slides = sl[page.Index : page.Index+page.Count]

	//log.Println(page)

	res.DoneData(c, page)
}

func handleSlideInfo(c *gin.Context) {
	path := c.Query("path")
	fmt.Println(path)

	res := NewRes()
	infoPath := filepath.Join(path, "slideinfo.bic")

	fp, err := os.Open(infoPath)
	if err != nil {
		fmt.Println(err)
		res.FailErr(c, err)
		return
	}
	defer fp.Close()

	var s []byte
	if s, err = ioutil.ReadAll(fp); err != nil {
		fmt.Println(err)
		res.FailErr(c, err)
		return
	}

	/*type SlideInfo struct {
		Levels     int     `json:"Levels"`
		Width      int     `json:"PhysicalWidth"`
		Height     int     `json:"PhysicalHeight"`
		Mag        float32 `json:"SourceLens"`
		TimeUse    float32 `json:"ScanTime"`
		Lens       string  `json:"SourceLensStr"`
		FileSize   string
		CellCount  int
		CreateTime time.Time
	}

	tinfo := SlideInfo{}
	err = json.Unmarshal(s, &tinfo)
	if err != nil {
		res.FailErr(c, err)
		return
	}

	res.DoneData(c, tinfo)*/

	c.Data(http.StatusOK, "application/json", s)

	/*var d map[string]interface{}
	if err := json.Unmarshal(s, &d); err != nil {
		res.FailErr(c, err)
		return
	}

	res.DoneData(c, d)*/

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

func handleTest(c *gin.Context) {
	res := NewRes()
	str := `{"name":"gray"}`
	var d map[string]interface{}
	if err := json.Unmarshal([]byte(str), &d); err != nil {
		res.FailErr(c, err)
		return
	}

	res.DoneData(c, d)
}
