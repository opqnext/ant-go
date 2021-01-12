package main

import (
	"ant-go/ant"
	"ant-go/config"
	"ant-go/logger"
	"flag"
	"fmt"
	"os"
)

func main() {
	var configPath string

	//获取配置文件
	flag.StringVar(&configPath, "config", "", "配置文件路径")
	flag.Parse()

	if configPath == "" {
		fmt.Printf("Config Path must be assigned.")
		os.Exit(1)
	}

	var err error

	//校验配置文件格式
	err = config.InitConfig(configPath)
	if err != nil {
		fmt.Printf("Init config failed. Error is %v", err)
	}

	logConfig := config.GetConfig().LogConfig

	err = logger.InitLogger(logConfig.LogPath, logConfig.LogLevel)
	if err != nil {
		fmt.Printf("Init logger failed. Error is %v", err)
	}

	logger.GetLogger().Info("ant-go init success")

	app := ant.New()
	app.Get("/hello", helloHandlerFunc)
	//app.Get("/hi", hiHandlerFunc)
	app.Run()
}

func helloHandlerFunc(c *ant.Context) {
	c.Write.Write([]byte("Hello ant-go"))
}
