package main

import (
	"encoding/json"
)

type DomainList struct {
	Data []Domain
}

type Domain struct {
	Comment string
	Realm   string
	Type    string
}

func (p Proxmox) GetDomains() []Domain {
	endpoint_url := "/api2/json/access/domains"
	body := p.GetContent(endpoint_url)
	var domains DomainList
	json.Unmarshal(body, &domains)
	return domains.Data
}

type RealmConfigList struct {
	Data RealmConfig
}

type RealmConfig struct {
	Comment string
	Digest  string
	Plugin  string
	Type    string
}

func (p Proxmox) GetRealmConfig(domain Domain) RealmConfig {
	endpoint_url := "/api2/json/access/domains/" + domain.Realm
	body := p.GetContent(endpoint_url)
	var realmConfig RealmConfigList
	json.Unmarshal(body, &realmConfig)
	return realmConfig.Data
}
