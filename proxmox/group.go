package proxmox

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"net/url"
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
	payload := url.Values{}
	body, err := p.GetContent(endpoint_url, payload)

	if err != nil {
		return nil, err
	}

	var groups GroupList
	json.Unmarshal(body, &groups)
	return groups.Data, nil
}

type GroupConfigWrapper struct {
	Data GroupConfig
}

type GroupConfig struct {
	Comment string   `json:"comment"`
	Members []string `json:"members"`
}

type GroupConfigRequest struct {
	GroupId string `url:"groupid"`
}

func (p ProxmoxClient) GetGroupConfig(req *GroupConfigRequest) (GroupConfig, error) {
	endpoint_url := "/api2/json/access/groups/" + req.GroupId
	payload := url.Values{}
	body, err := p.GetContent(endpoint_url, payload)

	if err != nil {
		return GroupConfig{}, err
	}

	var groupConfig GroupConfigWrapper
	json.Unmarshal(body, &groupConfig)
	return groupConfig.Data, nil
}

type NewGroupRequest struct {
	GroupId string `url:"groupid"`
	Comment string `url:"comment,omitempty"`
}

func (p ProxmoxClient) CreateGroup(req *NewGroupRequest) ([]byte, error) {
	endpoint_url := "/api2/json/access/groups"
	payload, _ := query.Values(req)
	body, err := p.PostContent(endpoint_url, payload)

	if err != nil {
		return nil, err
	}

	return body, nil
}
