package main

import (
	"encoding/json"
)

type GroupList struct {
	Data []Group
}

type Group struct {
	GroupId string
}

func (p Proxmox) GetGroups() []Group {
	p.api_endpoint = "/api2/json/access/groups"
	body := p.GetContent()
	var groups GroupList
	json.Unmarshal(body, &groups)
	return groups.Data
}

type GroupConfigList struct {
	Data GroupConfig
}

type GroupConfig struct {
	Members []string
}

func (p Proxmox) GetGroupConfig(group Group) GroupConfig {
	p.api_endpoint = "/api2/json/access/groups/" + group.GroupId
	body := p.GetContent()
	var groupConfig GroupConfigList
	json.Unmarshal(body, &groupConfig)
	return groupConfig.Data
}
