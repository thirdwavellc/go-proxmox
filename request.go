package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func (p Proxmox) Url() string {
	return p.host + p.api_endpoint
}

func (p Proxmox) GetContent() []byte {
	fmt.Println("Fetching:", p.Url())

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
