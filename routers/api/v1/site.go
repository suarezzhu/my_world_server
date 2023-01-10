package v1

import (
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Site struct {
	id   int    `form:"id" valid:"Required;Min(1)"`
	name string `form:"name" valid:"Required;MaxSize(100)"`
	url  string `form:"url" valid:"Required;MaxSize(100)"`
	note string `form:"note" valid:"Required;MaxSize(255)"`
	tag  string `form:"tag" valid:"Required;MaxSize(255)"`
}

func AddSite(c *gin.Context) {
	appG := app.Gin{C: c}

	name := c.Query("name")
	url := c.Query("url")
	note := c.Query("note")
	tag := c.Query("tag")
	type_id := c.Query("type_id")

	err := models.AddSite(name, url, note, tag, type_id)
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
func GetAllSites(c *gin.Context) {

	code := e.SUCCESS
	name := "%" + c.Query("name") + "%"
	url := "%" + c.Query("url") + "%"
	note := "%" + c.Query("note") + "%"
	tag := "%" + c.Query("tag") + "%"
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	data := models.GetAllSites2(name, url, note, tag, pageNum, pageSize)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
