package v1

// 文章相关路由

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

// 获取多个文章
func GetArticles(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	var upId int = -1
	if arg := c.Query("upId"); arg != "" {
		upId = com.StrTo(arg).MustInt()
		maps["upId"] = upId
		valid.Min(upId, 1, "upId").Message("标签ID必须大于0")
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS

		data["list"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
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

// 获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
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

// 新增文章
func AddArticle(c *gin.Context) {

}

// 修改文章
func EditArticle(c *gin.Context) {

}

// 删除文章
func DeleteTag(c *gin.Context) {

}