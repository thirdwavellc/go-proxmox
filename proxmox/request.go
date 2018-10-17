package proxmox

import (
	"bytes"
	"crypto/tls"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func (p ProxmoxClient) BuildUrl(endpoint_url string) string {
	return p.Host + endpoint_url
}

func (p ProxmoxClient) GetContent(endpoint_url string, payload url.Values) ([]byte, error) {
	body, err := p.apiRequest("GET", endpoint_url, payload)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (p ProxmoxClient) PostContent(endpoint_url string, payload url.Values) ([]byte, error) {
	body, err := p.apiRequest("POST", endpoint_url, payload)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (p ProxmoxClient) PutContent(endpoint_url string, payload url.Values) ([]byte, error) {
	body, err := p.apiRequest("PUT", endpoint_url, payload)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (p ProxmoxClient) DeleteContent(endpoint_url string, payload url.Values) ([]byte, error) {
	body, err := p.apiRequest("DELETE", endpoint_url, payload)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (p ProxmoxClient) apiRequest(method string, endpoint_url string, payload url.Values) ([]byte, error) {
	log.Println("API Request:")
	log.Printf("Method: %s\n", method)
	log.Printf("URL: %s\n", endpoint_url)
	log.Printf("Payload: %+v\n\n", payload)

	request, err := http.NewRequest(method, p.BuildUrl(endpoint_url), bytes.NewBufferString(payload.Encode()))
	if err != nil {
		return nil, err
	}

	if p.Auth.Ticket != "" {
		request.Header.Add("CSRFPreventionToken", p.Auth.CSRFPreventionToken)

		cookie := http.Cookie{Name: "PVEAuthCookie", Value: p.Auth.Ticket,
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

	log.Printf("API Response: %+v\n", resp)
	log.Println("----------------------------------------------------------------------------")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		PrintError(err)
		return body, err
	}

	log.Println("Body:")
	log.Printf("%s", body)
	log.Println("----------------------------------------------------------------------------")

	if resp.StatusCode != 200 {
		err := errors.New(resp.Status)
		return nil, err
	}

	return body, nil
}
