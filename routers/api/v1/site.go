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

	err := models.AddSite(name, url, notes, tags)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

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
