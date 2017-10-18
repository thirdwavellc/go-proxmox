package proxmox

import (
	"encoding/json"
	"net/url"
)

type NodeStorageList struct {
	Data []NodeStorage
}

type NodeStorage struct {
	Active  int    `json:"active"`
	Avail   int    `json:"avail"`
	Content string `json:"content"`
	Enabled int    `json:"enabled"`
	Shared  int    `json:"shared"`
	Storage string `json:"storage"`
	Total   int    `json:"total"`
	Type    string `json:"type"`
	Used    int    `json:"used"`
}

type NodeStorageRequest struct {
	Node    string `url:"node"`
	Content string `url:"content"`
	Enabled int    `url:"enabled"`
	Storage string `url:"storage"`
	Target  string `url:"target"`
}

func (p ProxmoxClient) GetNodeStorage(req *NodeStorageRequest) ([]NodeStorage, error) {
	endpoint_url := "/api2/json/nodes/" + req.Node + "/storage"
	payload := url.Values{}
	body, err := p.GetContent(endpoint_url, payload)

	if err != nil {
		return nil, err
	}

	var nodeStorage NodeStorageList
	json.Unmarshal(body, &nodeStorage)
	return nodeStorage.Data, nil
}

type NodeStorageContentList struct {
	Data []NodeStorageContent
}

type NodeStorageContent struct {
	Content string `json:"content"`
	Format  string `json:"format"`
	Size    int    `json:"size"`
	VolId   string `json:"volid"`
}

type NodeStorageContentRequest struct {
	Node    string `url:"node"`
	Storage string `url:"storage"`
	Content string `url:"content,omitempty"`
	VMID    int    `url:"vmid,omitempty"`
}

func (p ProxmoxClient) GetNodeStorageContent(req *NodeStorageContentRequest) ([]NodeStorageContent, error) {
	endpoint_url := "/api2/json/nodes/" + req.Node + "/storage/" + req.Storage + "/content"
	payload := url.Values{}
	body, err := p.GetContent(endpoint_url, payload)

	if err != nil {
		return nil, err
	}

	var nodeStorageContent NodeStorageContentList
	json.Unmarshal(body, &nodeStorageContent)
	return nodeStorageContent.Data, nil
}
