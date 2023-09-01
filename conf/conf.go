package conf

import (
	_ "embed"
	"encoding/json"
	"github.com/MenciusCheng/bookmark-manager/util/log"
	"os"
)

func Init(configPath string) {
	if configPath != "" {
		loadConfig, err := os.ReadFile(configPath)
		if err != nil {
			log.Error("ReadFile", err)
			panic(err)
		}
		devConfig = loadConfig
	}
	err := json.Unmarshal(devConfig, &Conf)
	if err != nil {
		log.Error("Unmarshal", err)
		panic(err)
	}

	InitMySQL(Conf.Databases)
}

//go:embed config/dev/config.json
var devConfig []byte

var Conf struct {
	ServiceName  string         `json:"serviceName"`
	Port         int            `json:"port"`
	GinMode      string         `json:"ginMode"`
	Databases    []DatabaseConf `json:"databases"`
	RedisConfigs []RedisConf    `json:"redisConfigs"`
}

type DatabaseConf struct {
	Name   string `json:"name"`
	Master string `json:"master"`
}

type RedisConf struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Database int    `json:"database"`
}
