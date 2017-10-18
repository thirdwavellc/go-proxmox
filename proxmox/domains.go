package proxmox

import (
	"encoding/json"
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
	body, err := p.GetContent(endpoint_url)

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

func (p ProxmoxClient) GetRealmConfig(domain Domain) (RealmConfig, error) {
	endpoint_url := "/api2/json/access/domains/" + domain.Realm
	body, err := p.GetContent(endpoint_url)

	if err != nil {
		return RealmConfig{}, err
	}

	var realmConfig RealmConfigList
	json.Unmarshal(body, &realmConfig)
	return realmConfig.Data, nil
}
