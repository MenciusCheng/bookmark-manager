package main

import (
	"flag"
	"github.com/MenciusCheng/bookmark-manager/bookmark-service/conf"
	"github.com/MenciusCheng/bookmark-manager/bookmark-service/server/http"
	"github.com/MenciusCheng/bookmark-manager/bookmark-service/service"
	"github.com/MenciusCheng/bookmark-manager/bookmark-service/util/logging"
	"github.com/MenciusCheng/superman/util/dragons"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	configS := flag.String("config", "config/config.toml", "Configuration file")
	flag.Parse()

	dragons.Init(
		dragons.ConfigPath(*configS),
	)
}

// @title 书签服务
// @version 1.0
// @description 书签服务相关接口
// @host 127.0.0.1:8080
// @BasePath /
// @accept json
// @produce json
func main() {

	defer dragons.Shutdown()

	// init local config
	cfg, err := conf.Init()
	if err != nil {
		logging.Fatalf("service config init error %s", err)
	}

	// create a service instance
	srv := service.New(cfg)

	// init and start http server
	http.Init(srv, cfg)

	defer http.Shutdown()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-sigChan
		log.Printf("get a signal %s\n", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("cat.demo server exit now...")
			return
		case syscall.SIGHUP:
		default:
		}
	}
}
