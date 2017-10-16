package proxmox

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

func (p ProxmoxClient) GetGroups() ([]Group, error) {
	endpoint_url := "/api2/json/access/groups"
	body, err := p.GetContent(endpoint_url)

	if err != nil {
		return nil, err
	}

	var groups GroupList
	json.Unmarshal(body, &groups)
	return groups.Data, nil
}

type GroupConfigList struct {
	Data GroupConfig
}

type GroupConfig struct {
	Members []string
}

func (p ProxmoxClient) GetGroupConfig(group Group) (GroupConfig, error) {
	endpoint_url := "/api2/json/access/groups/" + group.GroupId
	body, err := p.GetContent(endpoint_url)

	if err != nil {
		return GroupConfig{}, err
	}

	var groupConfig GroupConfigList
	json.Unmarshal(body, &groupConfig)
	return groupConfig.Data, nil
}

func (p ProxmoxClient) CreateGroup(group Group) ([]byte, error) {
	endpoint_url := "/api2/json/access/groups"
	payload := url.Values{}
	payload.Add("groupid", group.GroupId)
	body, err := p.PostContent(endpoint_url, payload)

	if err != nil {
		return nil, err
	}

	return body, nil
}
