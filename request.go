package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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

func (p Proxmox) PostContent(payload url.Values) []byte {
	fmt.Printf("Fetching: %v\n\n", p.Url())

	request, err := http.NewRequest("POST", p.Url(), bytes.NewBufferString(payload.Encode()))
	if err != nil {
		PrintError(err)
	}

	request.Header.Add("CSRFPreventionToken", p.auth.CSRFPreventionToken)

	cookie := http.Cookie{Name: "PVEAuthCookie", Value: p.auth.Ticket,
		Expires: time.Now().Add(356 * 24 * time.Hour), HttpOnly: true}
	request.AddCookie(&cookie)

	fmt.Printf("Request Headers: %+v\n\n", request.Header)
	fmt.Printf("Request Body: %+v\n\n", request.Body)

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

	fmt.Printf("Response Status: %v\n\n", resp.Status)
	fmt.Printf("Response Headers: %v\n\n", resp.Header)
	fmt.Printf("Response Body: %v\n\n", string(body))
	return body
}
