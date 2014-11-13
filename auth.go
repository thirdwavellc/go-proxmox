package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// AuthResponse is the wrapper struct for the auth JSON response.
type AuthResponse struct {
	Data AuthInfo
}

// AuthInfo maps to the JSON response for creating an auth ticket.
type AuthInfo struct {
	CSRFPreventionToken string
	Ticket              string
	Username            string
}

// GetAuth requests a new auth ticket, storing the information in the
// corresponding Proxmox struct.
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

	var auth AuthResponse
	json.Unmarshal(body, &auth)

	return auth.Data
}
