package main

import (
	"encoding/json"
)

type NodeDatastoreList struct {
	Data []NodeDatastore
}

type NodeDatastore struct {
	Active  int
	Avail   int
	Content string
	Enabled int
	Shared  int
	Storage string
	Total   int
	Type    string
	Used    int
}

type NodeDatastoreContentList struct {
	Data []NodeDatastoreContent
}

type NodeDatastoreContent struct {
	Content string
	Format  string
	Size    int
	Volid   string
}

func (p Proxmox) GetNodeDatastores(node string) ([]NodeDatastore, error) {
	endpoint_url := "/api2/json/nodes/" + node + "/storage"
	body, err := p.GetContent(endpoint_url)

	if err != nil {
		return nil, err
	}

	var nodeDatastores NodeDatastoreList
	json.Unmarshal(body, &nodeDatastores)
	return nodeDatastores.Data, nil
}

func (p Proxmox) GetNodeDatastoreContent(node string, datastore string) ([]NodeDatastoreContent, error) {
	endpoint_url := "/api2/json/nodes/" + node + "/storage/" + datastore + "/content"
	body, err := p.GetContent(endpoint_url)

	if err != nil {
		return nil, err
	}

	var nodeDatastoreContent NodeDatastoreContentList
	json.Unmarshal(body, &nodeDatastoreContent)
	return nodeDatastoreContent.Data, nil
}
