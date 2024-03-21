package v1

// 文章相关路由

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hucongyang/city-taste-service/models"
	"github.com/hucongyang/city-taste-service/pkg/e"
	"github.com/hucongyang/city-taste-service/pkg/setting"
	"github.com/hucongyang/city-taste-service/pkg/util"
	"github.com/unknwon/com"
)

// 获取多个文章
func GetArticles(c *gin.Context) {
	title := c.Query("title")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if title != "" {
		maps["title"] = title
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["list"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetArticleTotal(maps)

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