package main

import (
	"fmt"
	"os"

	"github.com/fatedier/frp/client"
	"github.com/fatedier/frp/models/config"
	//"github.com/fatedier/frp/models/consts"
)

func frpLogin() error {
	clientConf := config.GetDefaultClientConf()
	clientConf.ServerAddr = "www.biopic.biz"
	clientConf.ServerPort = 7000
	err := clientConf.Check()
	if err != nil {
		fmt.Println(err)
		return err
	}

	cfg := &config.HttpProxyConf{}

	cfg.ProxyName = "web"
	cfg.ProxyType = "http"
	cfg.LocalIp = "127.0.0.1"
	cfg.LocalPort = 9999
	cfg.CustomDomains = []string{"gray.biopic.biz"}
	/*cfg.SubDomain = subDomain
	cfg.Locations = strings.Split(locations, ",")
	cfg.HttpUser = httpUser
	cfg.HttpPwd = httpPwd
	cfg.HostHeaderRewrite = hostHeaderRewrite
	cfg.UseEncryption = useEncryption
	cfg.UseCompression = useCompression*/

	err = cfg.CheckForCli()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	proxyConfs := map[string]config.ProxyConf{
		cfg.ProxyName: cfg,
	}

	svr, err := client.NewService(clientConf, proxyConfs, nil, "")
	if err != nil {
		return err
	}

	return svr.Run()
}
