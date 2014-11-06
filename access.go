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
	p.api_endpoint = "/api2/json/access/domains"
	body := p.GetContent()
	var domains DomainList
	json.Unmarshal(body, &domains)
	return domains.Data
}
