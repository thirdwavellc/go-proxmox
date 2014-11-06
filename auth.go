package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type AuthRequest struct {
	Data AuthInfo
}

type AuthInfo struct {
	CSRFPreventionToken string
	Ticket              string
	Username            string
}

func (p Proxmox) GetAuth() AuthInfo {
	p.api_endpoint = "/api2/json/access/ticket"
	values := make(url.Values)
	values.Set("username", p.user)
	values.Set("password", p.password)
	resp, err := http.PostForm(p.Url(), values)
	if err != nil {
		PrintError(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		PrintError(err)
	}

	var auth AuthRequest
	json.Unmarshal(body, &auth)

	return auth.Data
}
