package proxmox

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"net/url"
)

type UserList struct {
	Data []User
}

type User struct {
	UserId    string `json:"userid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Comment   string `json:"comment"`
	Expire    int    `json:"expire"`
	Keys      string `json:"keys"`
	Enable    int    `json:"enable"`
}

func (p ProxmoxClient) GetUsers() ([]User, error) {
	endpoint_url := "/api2/json/access/users"
	payload := url.Values{}
	body, err := p.GetContent(endpoint_url, payload)

	if err != nil {
		return nil, err
	}

	var users UserList
	json.Unmarshal(body, &users)
	return users.Data, nil
}

type UserConfigWrapper struct {
	Data UserConfig
}

type UserConfig struct {
	FirstName string   `json:"firstname"`
	LastName  string   `json:"lastname"`
	Email     string   `json:"email"`
	Groups    []string `json:"groups"`
	Comment   string   `json:"comment"`
	Expire    int      `json:"expire"`
	Keys      string   `json:"keys"`
	Enable    int      `json:"enable"`
}

type UserConfigRequest struct {
	UserId string `url:"userid"`
}

func (p ProxmoxClient) GetUserConfig(req *UserConfigRequest) (UserConfig, error) {
	endpoint_url := "/api2/json/access/users/" + req.UserId
	payload := url.Values{}
	body, err := p.GetContent(endpoint_url, payload)

	if err != nil {
		return UserConfig{}, err
	}

	var userConfig UserConfigWrapper
	json.Unmarshal(body, &userConfig)
	return userConfig.Data, nil
}

type NewUserRequest struct {
	UserId    string   `url:"userid"`
	Comment   string   `url:"comment,omitempty"`
	Email     string   `url:"email,omitempty"`
	Enable    int      `url:"enable,omitempty"`
	Expire    int      `url:"expire,omitempty"`
	FirstName string   `url:"firstname,omitempty"`
	Groups    []string `url:"groups,omitempty"`
	Keys      string   `url:"keys,omitempty"`
	LastName  string   `url:"lastname,omitempty"`
	Password  string   `url:"password,omitempty"`
}

func (p ProxmoxClient) CreateUser(req *NewUserRequest) ([]byte, error) {
	endpoint_url := "/api2/json/access/users"
	payload, _ := query.Values(req)
	body, err := p.PostContent(endpoint_url, payload)

	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserResponse struct {
	Data string
}

type ExistingUserRequest struct {
	UserId    string   `url:"-"`
	Comment   string   `url:"comment,omitempty"`
	Email     string   `url:"email,omitempty"`
	Enable    int      `url:"enable,omitempty"`
	Expire    int      `url:"expire,omitempty"`
	FirstName string   `url:"firstname,omitempty"`
	Groups    []string `url:"groups,omitempty"`
	Keys      string   `url:"keys,omitempty"`
	LastName  string   `url:"lastname,omitempty"`
	Password  string   `url:"password,omitempty"`
}

func (p ProxmoxClient) UpdateUser(req *ExistingUserRequest) (string, error) {
	endpoint_url := "/api2/json/access/users/" + req.UserId
	payload, _ := query.Values(req)
	body, err := p.PutContent(endpoint_url, payload)

	if err != nil {
		return "", err
	}

	var response UserResponse
	json.Unmarshal(body, &response)
	return response.Data, nil
}

func (p ProxmoxClient) DeleteUser(req *ExistingUserRequest) (string, error) {
	endpoint_url := "/api2/json/access/users/" + req.UserId
	payload := url.Values{}
	body, err := p.DeleteContent(endpoint_url, payload)

	if err != nil {
		return "", err
	}

	var response UserResponse
	json.Unmarshal(body, &response)
	return response.Data, nil
}
