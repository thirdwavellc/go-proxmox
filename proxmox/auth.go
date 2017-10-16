package proxmox

import (
	"encoding/json"
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
func (p ProxmoxClient) GetAuth() (AuthInfo, error) {
	endpoint_url := "/api2/json/access/ticket"
	values := make(url.Values)
	values.Set("username", p.User)
	values.Set("password", p.Password)
	body, err := p.PostContent(endpoint_url, values)

	if err != nil {
		return AuthInfo{}, err
	}

	var auth AuthResponse
	json.Unmarshal(body, &auth)

	return auth.Data, nil
}
