package main

import (
	"fmt"
	"ginblog/routers"
	"net/http"

	"ginblog/config"
)

func main() {
	router := routers.InitRouter()
	// router.GET("/health-dataapi", dataapi.GetDataAPIHealthHandler)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.HTTPPort),
		Handler:        router,
		ReadTimeout:    config.ReadTimeout,
		WriteTimeout:   config.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}
