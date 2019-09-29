package main

import (
	"fmt"
	"gin-blog/pkg/setting"
	"gin-blog/routers"
	"net/http"
)

func main() {
	//router := gin.Default()
	//router.GET("/test", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "test",
	//	})
	//})
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

	//gin.SetMode(setting.ServerSetting.RunMode)
	//
	//routersInit := routers.InitRouter()
	//readTimeout := setting.ServerSetting.ReadTimeout
	//writeTimeout := setting.ServerSetting.WriteTimeout
	//endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	//maxHeaderBytes := 1 << 20
	//
	//server := &http.Server{
	//	Addr:           endPoint,
	//	Handler:        routersInit,
	//	ReadTimeout:    readTimeout,
	//	WriteTimeout:   writeTimeout,
	//	MaxHeaderBytes: maxHeaderBytes,
	//}
	//
	//log.Printf("[info] start http server listening %s", endPoint)
	//
	//server.ListenAndServe()

}
