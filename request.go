package main

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func (p Proxmox) BuildUrl(endpoint_url string) string {
	return p.host + endpoint_url
}

func (p Proxmox) GetContent(endpoint_url string) ([]byte, error) {
	req, err := http.NewRequest("GET", p.BuildUrl(endpoint_url), nil)
	if err != nil {
		return nil, err
	}

	cookie := http.Cookie{Name: "PVEAuthCookie", Value: p.auth.Ticket, Expires: time.Now().Add(356 * 24 * time.Hour), HttpOnly: true}
	req.AddCookie(&cookie)

	// TODO: remove me or refactor to be optional
	// This is only needed for testing local instance with self-signed cert
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err := errors.New(resp.Status)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (p Proxmox) PostContent(endpoint_url string, payload url.Values) ([]byte, error) {
	body, err := p.SendContent("POST", endpoint_url, payload)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (p Proxmox) PutContent(endpoint_url string, payload url.Values) ([]byte, error) {
	body, err := p.SendContent("PUT", endpoint_url, payload)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (p Proxmox) DeleteContent(endpoint_url string, payload url.Values) ([]byte, error) {
	body, err := p.SendContent("DELETE", endpoint_url, payload)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (p Proxmox) SendContent(method string, endpoint_url string, payload url.Values) ([]byte, error) {
	fmt.Println("SendContent Request:")
	fmt.Printf("Method: %s\n", method)
	fmt.Printf("URL: %s\n", endpoint_url)
	fmt.Printf("Payload: %+v\n\n", payload)

	request, err := http.NewRequest(method, p.BuildUrl(endpoint_url), bytes.NewBufferString(payload.Encode()))
	if err != nil {
		return nil, err
	}

	if p.auth.Ticket != "" {
		request.Header.Add("CSRFPreventionToken", p.auth.CSRFPreventionToken)

		cookie := http.Cookie{Name: "PVEAuthCookie", Value: p.auth.Ticket,
			Expires: time.Now().Add(356 * 24 * time.Hour), HttpOnly: true}
		request.AddCookie(&cookie)
	}

	// TODO: remove me or refactor to be optional
	// This is only needed for testing local instance with self-signed cert
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		err := errors.New(resp.Status)
		return nil, err
	}

	fmt.Printf("SendContent Response: %+v\n", resp)
	fmt.Println("----------------------------------------------------------------------------")

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		PrintError(err)
		return body, err
	}

	return body, nil
}
