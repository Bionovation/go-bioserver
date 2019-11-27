package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	//"github.com/fatedier/frp/client"
	//"github.com/fatedier/frp/models/config"
	//"github.com/fatedier/frp/models/consts"
)

/*
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
*/

func frpLogin() error {
	var iniContent string
	iniContent = fmt.Sprintf(
		`[common]
server_addr=%v
server_port=%v
[%v]
type=http
local_port=%v
local_ip=127.0.0.1
custom_domains=%v.%v`,
		bioConfig.Proxy.ServerAddr,
		bioConfig.Proxy.ServerPort,
		bioConfig.Common.MachineName,
		bioConfig.Common.ListenPort,
		bioConfig.Common.MachineName,
		bioConfig.Proxy.ServerAddr,
	)

	//fmt.Println(bioConfig.Common.MachineName)

	if err := ioutil.WriteFile("./proxy/frpc.ini", []byte(iniContent), 0); err != nil {
		log.Println("write frpc.ini failed, error:", err)
		return err
	}

	/*binary, err := exec.LookPath("./proxy/frpc.exe")
	if err != nil {
		log.Println("look up frpc failed, error:", err)
		return err
	}*/

	args := []string{"-c", "./proxy/frpc.ini"}

	cmd := exec.Command("./proxy/frpc.exe", args...)

	var stdoutBuf, stderrBuf bytes.Buffer
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)

	if err := cmd.Start(); err != nil {
		log.Println("exec frpc failed, error:", err)
		return err
	}

	go func() {
		io.Copy(stdout, stdoutIn)
	}()
	go func() {
		io.Copy(stderr, stderrIn)
	}()

	log.Println("frpc start success!")

	if err := cmd.Wait(); err != nil {
		log.Println("wait frpc failed, error:", err)
		return err
	}

	return nil
}
