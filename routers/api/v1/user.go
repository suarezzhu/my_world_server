package v1

import (
	"encoding/json"
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID           int    `gorm:"primary_key" json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	roles        string `json:"roles"`
	avatar       string `json:"avatar"`
	introduction string `json:"introduction"`
	email        string `json:"email"`
}

//
//func AddUser(c *gin.Context) {
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

func GetUserInfo(c *gin.Context) {
	data, _ := c.GetRawData()
	var body map[string]int
	_ = json.Unmarshal(data, &body)

	userId := body["id"]
	//var userId int =
	//rsData := make(map[string]interface{})
	//if userId != "" {
	//	maps["id"] = userId
	//}

	code := e.SUCCESS

	user, _ := models.GetUserInfoId(userId)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": user,
	})
}
