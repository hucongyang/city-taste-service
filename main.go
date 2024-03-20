package main

import (
	"fmt"
	"net/http"

	"github.com/hucongyang/city-taste-service/pkg/setting"
	"github.com/hucongyang/city-taste-service/routers"
)

func main() {
	router := routers.InitRouter()

	server := &http.Server{
		Addr: 			fmt.Sprintf(":%d", setting.HTTPPort),
		Handler: 		router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}