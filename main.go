package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/JanFant/TLServer/logger"
	"os"
	"rura/bb-server/extcon"
	"rura/downLoad/client"
	"rura/downLoad/server"
	"rura/downLoad/setup"
	"strings"
	"time"
)

var err error
var isserver bool

//Секция инициализации программы
func init() {
	isserver = false
	if len(os.Args) > 1 {
		if strings.Contains(os.Args[1], "server") {
			isserver = true
		}
	}
	setup.SetServer = new(setup.Server)
	setup.SetClient = new(setup.Client)
	if isserver {
		if _, err := toml.DecodeFile("configServer.toml", &setup.SetServer); err != nil {
			fmt.Println("Can't load config file - ", err.Error())
			os.Exit(1)
		}
		err = logger.Init(setup.SetServer.LogPath)
	} else {
		if _, err := toml.DecodeFile("configClient.toml", &setup.SetClient); err != nil {
			fmt.Println("Can't load config file - ", err.Error())
			os.Exit(1)
		}
		err = logger.Init(setup.SetClient.LogPath)
	}
	if err != nil {
		fmt.Println("Error opening logger subsystem ", err.Error())
		return
	}
	if isserver {
		logger.Info.Println("Starting as server")
		fmt.Println("Starting as server...")
	} else {
		logger.Info.Println("Starting as client")
		fmt.Println("Starting as client...")
	}

}

func main() {
	extcon.BackgroundInit()
	m, _ := extcon.NewContext("main")
	stop := make(chan int)
	if isserver {
		//Подготовим свои перекодировки
		for _, s := range setup.SetServer.Regions {
			if s.IDOut == 0 {
				s.IDOut = s.IDIn
			}
		}
		go server.Server(m, stop)
	} else {
		go client.Client(m, stop)
	}
	extcon.BackgroundWork(time.Duration(1*time.Second), stop)
	fmt.Println("Exit work")
	logger.Info.Println("Exit work")
}
