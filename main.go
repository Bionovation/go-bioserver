//package go_bioserver
//-ldflags="-H windowsgui"
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	//"github.com/sirupsen/logrus"
)

// 配置文件
const cfile = "./config.toml"

// 日志输出
// var log = logrus.New()

func stdToFile() {
	f, _ := os.OpenFile("./go-bioserver.log", os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0755)
	os.Stdout = f
	os.Stderr = f
	log.SetOutput(f)
}

func main() {
	exPath, _ := os.Executable()
	exFolder := filepath.Dir(exPath)
	confFile := filepath.Join(exFolder, cfile)

	checkConfigFile(confFile)

	bioConfig.readConfig(confFile) // 读取配置文件

	// os.Args = append(os.Args, "F:\\BioSlides\\北京2019-11-18-09-52-29\\data.bimg")

	if len(os.Args) == 2 {
		// 如果是传入一个参数，则可能是打开本地文件
		viewLocalFile(os.Args[1])
	} else {
		regesterBimg()
		runServ()
	}

}

func viewLocalFile(path string) {

	_, err := os.Stat(os.Args[1])
	if os.IsNotExist(err) {
		panic("error calling")
	}

	if filepath.Base(os.Args[1]) != "data.bimg" && filepath.Base(os.Args[1]) != "downlayer.bimg" {
		panic("error ext")
	}

	folder := filepath.Dir(os.Args[1])
	folder = strings.ReplaceAll(folder, "\\", "/")

	// 拼url
	// http://localhost:8890/html/sample.html?samplePath=D%3A%5CBioScan%5C2019-09-12-13-33-16
	urlStr := fmt.Sprintf("http://localhost:%v/html/sample.html?samplePath=%v", bioConfig.Common.ListenPort, folder)
	/*l, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}

	urlStr = l.Scheme + "://" + l.Host + "?" + l.Query().Encode()
	*/
	//fmt.Println(urlStr)

	// 检查服务是否已经运行
	n := servIsRunning()
	if n == false {
		// 如果只有这一个实例，则需要先启动看图服务再打开
		go func() {
			time.Sleep(time.Second * 2)
			exec.Command(`cmd`, `/c`, `start`, urlStr).Start()
		}()

		// 	启动服务
		runServ()

		// 改成启动进程
		/*exPath, _ := os.Executable()
		exec.Command(`cmd`, `/c`, `start`, exPath).Start()
		time.Sleep(time.Second * 2)*/

	}
	exec.Command(`cmd`, `/c`, `start`, urlStr).Start()

}

func runServ() {
	log.Println("bioserver startting...")

	// 重定向标准输出到文件
	//stdToFile()

	// 登录frp代理服务
	if bioConfig.Proxy.ServerPort != 0 {
		go frpLogin()
	}

	go clearRoutine(nil) // 运行内存清理线程

	log.Println("bioserver is running.")
	log.Printf("data folder is  : %v", bioConfig.Common.DataFolder)
	log.Printf("please visit url: http://localhost:%v\n", bioConfig.Common.ListenPort)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	exPath, _ := os.Executable()
	folder := filepath.Dir(exPath)

	r.StaticFS("/html", http.Dir(filepath.Join(folder, "./html")))
	r.StaticFS("/openseadragon", http.Dir(filepath.Join(folder, "./html/openseadragon")))

	r.GET("/", handleIndex)
	r.GET("/host", handleHost)
	r.GET("/ping", handlePing)
	r.GET("/image", handleImage)
	r.GET("/slides", handleSlideList)
	r.GET("/slideinfo", handleSlideInfo)
	r.GET("/slidetile", handleSlideTile)
	r.GET("/slidenail", handleSlideNail)

	r.GET("/slidedelete", handleSlideDel)

	r.GET("/test", handleTest)

	r.Run(fmt.Sprintf(":%v", bioConfig.Common.ListenPort))
}

func servIsRunning() bool {
	pingUrl := fmt.Sprintf("http://localhost:%v/ping", bioConfig.Common.ListenPort)

	//fmt.Println(pingUrl)

	client := &http.Client{}
	resp, err := client.Get(pingUrl)

	if err != nil {
		return false
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	strBody := string(body)

	//fmt.Println(strBody)
	return strings.Contains(strBody, "pong")
}

func regesterBimg() {
	exepath, _ := os.Executable()
	cmd := exec.Command("cmd", "/c", "assoc .bimg=BioData")
	cmd.Output()
	//output, _ := cmd.Output()
	//fmt.Println(string(output))

	ftype := fmt.Sprintf("ftype BioData=%v %%1", exepath)

	cmd = exec.Command("cmd", "/c", ftype)
	cmd.Output()
	//output, _ = cmd.Output()
	//fmt.Println(string(output))
}

func checkConfigFile(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) == false {
		return
	}

	defaultConfig := `[common]
DataFolder = "D:/BioScan"
ListenPort = 8890`

	ioutil.WriteFile(path, []byte(defaultConfig), 0644)
}
