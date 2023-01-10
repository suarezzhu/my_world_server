package v1

import (
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddSite(c *gin.Context) {
	appG := app.Gin{C: c}

	name := c.Query("name")
	url := c.Query("url")
	notes := c.Query("notes")
	tags := c.Query("tags")
	typeId := c.Query("type")

	err := models.AddSite(name, url, notes, tags, typeId)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// 获取网站
func GetSites(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}
	code := e.SUCCESS
	data["lists"] = models.GetSites(util.GetPage(c), 10, maps)
	data["total"] = models.GetSiteTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 获取全部网站地址
func getAllSites(c *gin.Context) {
	code := e.SUCCESS
	data := models.GetAllSites()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
