package main

import (
	"fmt"
	"os"

	"github.com/kardianos/service"
)

func main() {
	fmt.Println("EarlySleepService v1.0, code by shrek@2022-08-28")
	if len(os.Args) == 1 {
		fmt.Println("Usage: autosleep [command]")
		fmt.Println("                 install -- 安装服务 install service")
		fmt.Println("                 remove  -- 卸载服务 uninstall service")
		fmt.Println("                 start   -- 启动服务 start service")
		fmt.Println("                 stop    -- 停止服务 stop service")
	}

	svcConfig := &service.Config{
		Name:        "EarlySleepService",                          //服务显示名称
		DisplayName: "EarlySleep",                                 //服务名称
		Description: "If u are a student,u should sleep earlier.", //服务描述
		Option: service.KeyValue{
			"StartType": "automatic",
		},
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			s.Install()
			fmt.Println("服务安装成功")
			return
		} else if os.Args[1] == "remove" {
			s.Uninstall()
			fmt.Println("服务卸载成功")
			return
		} else if os.Args[1] == "stop" {
			s.Stop()
			fmt.Println("服务停止成功")
			return
		} else if os.Args[1] == "start" {
			s.Start()
			fmt.Println("服务启动成功")
			return
		} else if os.Args[1] == "status" {
			s.Status()
			return
		}
	}

	err = s.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
}
