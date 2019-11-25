package main

import (
	"github.com/fatedier/frp/client"
	"github.com/fatedier/frp/models/config"
)

func frpLogin() error {
	clientConf := config.GetDefaultClientConf()
	clientConf.ServerAddr = www.bioslide.biz
	clientConf.ServerPort = 7000

	svr, err := client.NewService(clientConf)
	if err != nil {
		return err
	}

	svr.Run()
}
