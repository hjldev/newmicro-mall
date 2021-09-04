package main

import (
	"context"
	"github.com/hjldev/newmicro-mall/gateway/dao"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	dao.ServiceManagerHandler.LoadOnce()

	var httpSrvHandler *http.Server

	go func() {
		r := initRouter()
		httpSrvHandler = &http.Server{
			Addr:           ":8880",
			Handler:        r,
			ReadTimeout:    time.Duration(10) * time.Second,
			WriteTimeout:   time.Duration(10) * time.Second,
			MaxHeaderBytes: 1 << uint(20),
		}
		log.Printf(" [INFO] http_proxy_run %s\n", "8880")
		if err := httpSrvHandler.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf(" [ERROR] http_proxy_run %s err:%v\n", "8880", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := httpSrvHandler.Shutdown(ctx); err != nil {
		log.Printf(" [ERROR] http_proxy_stop err:%v\n", err)
	}
	log.Printf(" [INFO] http_proxy_stop %v stopped\n", "8880")
}
