package routers

import (
	"fmt"
	"gin-blog/pkg/setting"
	"gin-blog/routers/api"
	v1 "gin-blog/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth) //jwt get token

	apiv1 := r.Group("/api/v1")
	//apiv1.Use(jwt.JWT()) //对该组api做认证授权 check token

	r.Use(func(c *gin.Context) {
		fmt.Println("r->middleware test")
		c.Next()
	})
	apiv1.Use(func(c *gin.Context) {
		fmt.Println("apiv1->middleware test")
		c.Next()
	})

	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)

		apiv1.GET("/writers/:id", v1.GetWriter)
		apiv1.GET("/writers", v1.GetWriters)
		apiv1.POST("/writer", v1.AddWriter)

	}

	return r
}
