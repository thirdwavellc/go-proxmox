package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type NodeList struct {
	Data []Node
}

type Node struct {
	Disk    int
	CPU     float64
	MaxDisk int
	MaxMem  int
	Node    string
	MaxCPU  int
	Level   string
	Uptime  int
	Id      string
	Type    string
	Mem     int
}

func GetNodes(host string, auth AuthInfo) []Node {
	url := host + "/api2/json/nodes"
	body := GetContent(url, auth)
	var nodes NodeList
	json.Unmarshal(body, &nodes)
	return nodes.Data
}

func printNodes(nodes []Node) {
	for _, node := range nodes {
		printNode(node)
	}
}

func printNode(node Node) {
	s := reflect.ValueOf(node)
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%s: %v\n", typeOfT.Field(i).Name, f.Interface())
	}
	fmt.Printf("\n")
}
