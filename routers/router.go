package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hucongyang/city-taste-service/pkg/setting"
	v1 "github.com/hucongyang/city-taste-service/routers/api/v1"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test1",
		})
	})

	apiv1 := router.Group("/api/v1")
	{
		// 获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		// 获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)

		// 获取up列表
		apiv1.GET("/ups", v1.GetUps)
		// 获取指定up
		apiv1.GET("/ups/:id", v1.GetUp)

	}

	return router
}