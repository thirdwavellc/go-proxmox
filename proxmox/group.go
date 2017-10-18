package proxmox

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
)

type GroupList struct {
	Data []Group
}

type Group struct {
	GroupId string `json:"groupid" url:"groupid"`
	Comment string `json:"comment" url:"comment,omitempty"`
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

type GroupConfig struct {
	Comment string   `json:"comment"`
	Members []string `json:"members"`
}

func (p ProxmoxClient) GetGroupConfig(group Group) (GroupConfig, error) {
	endpoint_url := "/api2/json/access/groups/" + group.GroupId
	body, err := p.GetContent(endpoint_url)

	if err != nil {
		return GroupConfig{}, err
	}

	var groupConfig GroupConfig
	json.Unmarshal(body, &groupConfig)
	return groupConfig, nil
}

func (p ProxmoxClient) CreateGroup(group Group) ([]byte, error) {
	endpoint_url := "/api2/json/access/groups"
	payload, _ := query.Values(group)
	body, err := p.PostContent(endpoint_url, payload)

	if err != nil {
		return nil, err
	}

	return body, nil
}
