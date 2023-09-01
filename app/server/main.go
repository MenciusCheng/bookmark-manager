package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/MenciusCheng/bookmark-manager/conf"
	"github.com/MenciusCheng/bookmark-manager/server"
	"github.com/MenciusCheng/bookmark-manager/util/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 读取配置
	var configPath string
	flag.StringVar(&configPath, "c", "", "Configuration file")
	conf.Init(configPath)

	// 初始化路由
	gin.SetMode(conf.Conf.GinMode)
	router := server.NewRouter()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Conf.Port),
		Handler: router,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("ListenAndServe", err)
		}
	}()

	log.Info("Server Start", conf.Conf.ServiceName)
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Server Shutdown", conf.Conf.ServiceName)
}
