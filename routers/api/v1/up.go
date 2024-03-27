package v1

import (
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/hucongyang/city-taste-service/models"
	"github.com/hucongyang/city-taste-service/pkg/e"
	"github.com/hucongyang/city-taste-service/pkg/setting"
	"github.com/hucongyang/city-taste-service/pkg/util"
	"github.com/unknwon/com"
)

// 获取up主列表
func GetUps(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	name := c.Query("name")
	if name != "" {
		maps["name"] = name
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		data["list"] = models.GetUps(util.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetUpTotal(maps)
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})

}

// 获取单个up
func GetUp(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistUpByID(id) {
			data = models.GetUp(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_UP
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}


// 新增up
func AddUp(c *gin.Context) {

}

// 修改up
func EditUp(c *gin.Context) {

}

// 删除up
func DeleteUp(c *gin.Context) {
	
}