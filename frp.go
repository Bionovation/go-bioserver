package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
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

func isProcessExist(appName string) (bool, string, int) {
	appary := make(map[string]int)
	cmd := exec.Command("cmd", "/C", "tasklist")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // 隐藏窗口
	output, _ := cmd.Output()
	//fmt.Printf("fields: %v\n", output)
	n := strings.Index(string(output), "System")
	if n == -1 {
		fmt.Println("not find cmd window")
		return false, appName, -1
	}
	data := string(output)[n:]
	fields := strings.Fields(data)
	for k, v := range fields {
		if v == appName {
			appary[appName], _ = strconv.Atoi(fields[k+1])

			return true, appName, appary[appName]
		}
	}

	return false, appName, -1
}

func killProcess(appName string) error {
	c := exec.Command("cmd.exe", "/C", "taskkill", "/IM", appName)
	c.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // 隐藏窗口
	err := c.Start()
	if err != nil {
		return err
	}
	c.Wait()
	return nil
}

func frpLogin() error {

	b, _, _ := isProcessExist("frpc.exe")
	if b {
		log.Println("frpc stopping")
		killProcess("frpc.exe")
	} else {
		log.Println("frpc startting")
	}

	log.Println("frpc startting")

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

	args := []string{"-c", "./proxy/frpc.ini"}

	cmd := exec.Command("./proxy/frpc.exe", args...)

	var stdoutBuf, stderrBuf bytes.Buffer
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)

	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // 隐藏窗口
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
