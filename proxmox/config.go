package proxmox

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

func ReadProxmoxConfig(file string) (ProxmoxConfig, error) {
	homedir, err := homedir.Dir()

	if err != nil {
		return ProxmoxConfig{}, err
	}

	if file == "" {
		file = path.Join(homedir, ".go-proxmox.json")
	}

	content, err := ioutil.ReadFile(file)

	if err != nil {
		return ProxmoxConfig{}, err
	}

	var config ProxmoxConfig
	json.Unmarshal(content, &config)
	return config, nil
}
