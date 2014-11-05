package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Cluster struct {
	Data []ClusterNode
}

type ClusterNode struct {
	Id     string
	Level  string
	Local  int
	Name   string
	NodeId int
	Pmxcfs int
	State  int
	Type   string
}

func GetClusterStatus(host string, auth AuthInfo) []ClusterNode {
	url := host + "/api2/json/cluster/status"
	body := GetContent(url, auth)
	var cluster Cluster
	json.Unmarshal(body, &cluster)
	return cluster.Data
}

func printClusterStatus(cluster []ClusterNode) {
	for _, clusterNode := range cluster {
		printClusterNode(clusterNode)
	}
}

func printClusterNode(clusterNode ClusterNode) {
	s := reflect.ValueOf(clusterNode)
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%s: %v\n", typeOfT.Field(i).Name, f.Interface())
	}
	fmt.Printf("\n")
}
