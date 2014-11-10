package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func (p Proxmox) Url() string {
	return p.host + p.api_endpoint
}

func (p Proxmox) GetContent() []byte {
	req, err := http.NewRequest("GET", p.Url(), nil)
	if err != nil {
		PrintError(err)
	}

	cookie := http.Cookie{Name: "PVEAuthCookie", Value: p.auth.Ticket, Expires: time.Now().Add(356 * 24 * time.Hour), HttpOnly: true}
	req.AddCookie(&cookie)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		PrintError(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		PrintError(err)
	}

	return body
}

func (p Proxmox) PostContent(payload url.Values) []byte {
	request, err := http.NewRequest("POST", p.Url(), bytes.NewBufferString(payload.Encode()))
	if err != nil {
		PrintError(err)
	}

	request.Header.Add("CSRFPreventionToken", p.auth.CSRFPreventionToken)

	cookie := http.Cookie{Name: "PVEAuthCookie", Value: p.auth.Ticket,
		Expires: time.Now().Add(356 * 24 * time.Hour), HttpOnly: true}
	request.AddCookie(&cookie)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		PrintError(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		PrintError(err)
	}

	return body
}
