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

func GetAuth(host string, username string, password string) AuthInfo {
	requestUrl := host + "/api2/json/access/ticket"
	values := make(url.Values)
	values.Set("username", username)
	values.Set("password", password)
	resp, err := http.PostForm(requestUrl, values)
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
