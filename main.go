package main

import (
	"example.com/example/pkg/setting"
	"example.com/example/routers"
	"fmt"
	//_ "gorm.io/driver/mysql"
	"log"
	"net/http"
	"os"
)

func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("Server started on %s", s.Addr)
	log.Printf("pid: %d", os.Getpid())
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("server run failed:%v", err)
	}

	//go func() {
	//	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
	//		log.Fatalf("Server error: %v", err)
	//	}
	//}()
}
