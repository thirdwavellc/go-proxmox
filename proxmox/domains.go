package proxmox

import (
	"encoding/json"
	"net/url"
)

type DomainList struct {
	Data []Domain
}

type Domain struct {
	Comment string `json:"comment"`
	Realm   string `json:"realm"`
	Type    string `json:"type"`
}

func (p ProxmoxClient) GetDomains() ([]Domain, error) {
	endpoint_url := "/api2/json/access/domains"
	payload := url.Values{}
	body, err := p.GetContent(endpoint_url, payload)

	if err != nil {
		return nil, err
	}

	var domains DomainList
	json.Unmarshal(body, &domains)
	return domains.Data, nil
}

type RealmConfigList struct {
	Data RealmConfig
}

type RealmConfig struct {
	Comment string `json:"comment"`
	Digest  string `json:"digest"`
	Plugin  string `json:"plugin"`
	Type    string `json:"type"`
}

type RealmConfigRequest struct {
	Realm string `json:"realm"`
}

func (p ProxmoxClient) GetRealmConfig(req *RealmConfigRequest) (RealmConfig, error) {
	endpoint_url := "/api2/json/access/domains/" + req.Realm
	payload := url.Values{}
	body, err := p.GetContent(endpoint_url, payload)

	if err != nil {
		return RealmConfig{}, err
	}

	var realmConfig RealmConfigList
	json.Unmarshal(body, &realmConfig)
	return realmConfig.Data, nil
}
