package main

import (
	"encoding/json"
	"net/url"
)

type GroupList struct {
	Data []Group
}

type Group struct {
	GroupId string `json:"groupid"`
}

func (p Proxmox) GetGroups() []Group {
	endpoint_url := "/api2/json/access/groups"
	body := p.GetContent(endpoint_url)
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
	endpoint_url := "/api2/json/access/groups/" + group.GroupId
	body := p.GetContent(endpoint_url)
	var groupConfig GroupConfigList
	json.Unmarshal(body, &groupConfig)
	return groupConfig.Data
}

func (p Proxmox) CreateGroup(group Group) []byte {
	endpoint_url := "/api2/json/access/groups"
	payload := url.Values{}
	payload.Add("groupid", group.GroupId)
	body := p.PostContent(endpoint_url, payload)
	return body
}
