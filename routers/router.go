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
		
	}

	return router
}