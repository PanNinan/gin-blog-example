package main

import (
	"example.com/example/models"
	"example.com/example/pkg/gredis"
	"example.com/example/pkg/logging"
	"example.com/example/pkg/setting"
	"example.com/example/pkg/task"
	"example.com/example/routers"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	setting.SetUp()
	models.Setup()
	logging.SetUp()
	err := gredis.Setup()
	if err != nil {
		log.Fatalf("gredis setup failed:%v", err)
	}
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	if setting.AppSetting.EnableCron {
		task.Start()
	}
	log.Printf("Server started on %s", s.Addr)
	log.Printf("pid: %d", os.Getpid())
	err = s.ListenAndServe()
	if err != nil {
		log.Fatalf("server run failed:%v", err)
	}

	//go func() {
	//	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
	//		log.Fatalf("Server error: %v", err)
	//	}
	//}()
}
