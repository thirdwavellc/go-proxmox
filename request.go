package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func GetContent(url string, auth AuthInfo) []byte {
	fmt.Println("Fetching:", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		PrintError(err)
	}

	cookie := http.Cookie{Name: "PVEAuthCookie", Value: auth.Ticket, Expires: time.Now().Add(356 * 24 * time.Hour), HttpOnly: true}
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
