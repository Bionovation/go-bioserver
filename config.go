package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

// 机器通用参数
type ConfigCommon struct {
	MachineName string
	DataFolder  string
	ListenPort  int
}

// frp代理参数
type ConfigProxy struct {
	ServerAddr string
	ServerPort int
}

// 参数
type BioConfig struct {
	Common ConfigCommon
	Proxy  ConfigProxy
}

// 默认参数
var bioConfig = BioConfig{
	Common: ConfigCommon{
		DataFolder: "D:/BioScan",
		ListenPort: 8080,
	},
	Proxy: ConfigProxy{
		ServerAddr: "remotebionovation.com",
		ServerPort: 7000,
	},
}

func (config *BioConfig) readConfig(conFile string) {
	if config == nil {
		config = &bioConfig
	}

	if _, err := toml.DecodeFile(conFile, config); err != nil {
		log.Println("read config failed:", err)
		return
	}

	log.Printf("use config: %#v", config)
}
