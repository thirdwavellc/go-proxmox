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

func (p Proxmox) CreateGroup(group Group) []byte {
	p.api_endpoint = "/api2/json/access/groups"

	payload := url.Values{}
	payload.Add("groupid", group.GroupId)

	body := p.PostContent(payload)
	return body
}
