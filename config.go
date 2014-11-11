package main

import (
	"encoding/json"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"path"
)

type ProxmoxConfig struct {
	Host        string `json:"host"`
	User        string `json:"user"`
	Password    string `json:"password"`
	DefaultNode string `json:"defaultNode"`
}

func ReadProxmoxConfig() ProxmoxConfig {
	homedir, err := homedir.Dir()
	if err != nil {
		PrintError(err)
	}

	file := path.Join(homedir, ".go-proxmox.json")
	content, err := ioutil.ReadFile(file)
	var config ProxmoxConfig
	json.Unmarshal(content, &config)
	return config
}
