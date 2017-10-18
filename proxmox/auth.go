package proxmox

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
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

type TicketRequest struct {
	Password string `url:"password"`
	Username string `url:"username"`
	OTP      string `url:"otp,omitempty"`
	Path     string `url:"path,omitempty"`
	Privs    string `url:"privs,omitempty"`
	Realm    string `url:"realm,omitempty"`
}

// GetAuth requests a new auth ticket, storing the information in the
// corresponding Proxmox struct.
func (p ProxmoxClient) GetAuth(req *TicketRequest) (AuthInfo, error) {
	endpoint_url := "/api2/json/access/ticket"
	payload, _ := query.Values(req)
	body, err := p.PostContent(endpoint_url, payload)

	if err != nil {
		return AuthInfo{}, err
	}

	var auth AuthResponse
	json.Unmarshal(body, &auth)

	return auth.Data, nil
}
