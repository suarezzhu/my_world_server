package v1

import (
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

//func AddSite(c *gin.Context) {
//	appG := app.Gin{C: c}
//
//	name := c.Query("name")
//	url := c.Query("url")
//	notes := c.Query("notes")
//	tags := c.Query("tags")
//
//	err := models.AddSite(name, url, notes, tags)
//	if err != nil {
//		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
//		return
//	}
//
//	appG.Response(http.StatusOK, e.SUCCESS, nil)
//}

func GetDictDataByType(c *gin.Context) {
	//data, _ := c.GetRawData()
	//var body map[string]string
	//_ = json.Unmarshal(data, &body)
	//name := body["type"]
	name := c.Query("type")
	//
	code := e.SUCCESS
	rs, _ := models.DictDataByTypeName(name)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": rs,
	})
}
