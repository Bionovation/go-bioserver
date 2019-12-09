package main

import (
	"fmt"
	"math/rand"
	"time"

	_ "github.com/fatedier/frp/assets/frpc/statik"
	"github.com/fatedier/frp/client"
	"github.com/fatedier/frp/models/config"
	"github.com/fatedier/golib/crypto"
)

func frpLogin() error {
	crypto.DefaultSalt = "frp"
	rand.Seed(time.Now().UnixNano())

	cfg := config.GetDefaultClientConf()
	cfg.ServerAddr = bioConfig.Proxy.ServerAddr
	cfg.ServerPort = bioConfig.Proxy.ServerPort

	httpCfg := config.HttpProxyConf{}
	httpCfg.BaseProxyConf.ProxyName = bioConfig.Common.MachineName
	httpCfg.BaseProxyConf.ProxyType = "http"
	httpCfg.LocalIp = "127.0.0.1"
	httpCfg.LocalPort = bioConfig.Common.ListenPort

	httpCfg.CustomDomains = []string{fmt.Sprintf("%v.%v", bioConfig.Common.MachineName, bioConfig.Proxy.ServerAddr)}
	httpCfg.Locations = []string{""}

	visitorCfgs := make(map[string]config.VisitorConf)
	pxyCfgs := make(map[string]config.ProxyConf)

	pxyCfgs[bioConfig.Common.MachineName] = &httpCfg

	svr, err := client.NewService(cfg, pxyCfgs, visitorCfgs, "")
	if err != nil {
		return err
	}

	return svr.Run()
}
