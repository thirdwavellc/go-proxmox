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

func (p Proxmox) GetNodeDatastores(node string) []NodeDatastore {
	endpoint_url := "/api2/json/nodes/" + node + "/storage"
	body := p.GetContent(endpoint_url)
	var nodeDatastores NodeDatastoreList
	json.Unmarshal(body, &nodeDatastores)
	return nodeDatastores.Data
}

func (p Proxmox) GetNodeDatastoreContent(node string, datastore string) []NodeDatastoreContent {
	endpoint_url := "/api2/json/nodes/" + node + "/storage/" + datastore + "/content"
	body := p.GetContent(endpoint_url)
	var nodeDatastoreContent NodeDatastoreContentList
	json.Unmarshal(body, &nodeDatastoreContent)
	return nodeDatastoreContent.Data
}
