package handle

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SlideList(c *gin.Context) {
	c.String(http.StatusOK, "[]")
}

func SlideInfo(c *gin.Context) {
	c.String(http.StatusOK, "{info}")
}

func SlideTile(c *gin.Context) {
	path := c.Query("path")
	z := c.Query("level")
	x := c.Query("x")
	y := c.Query("y")
	rtn := fmt.Sprintf("get tile %v/%v/%v/%v", path, z, x, y)
	c.String(http.StatusOK, rtn)
}
