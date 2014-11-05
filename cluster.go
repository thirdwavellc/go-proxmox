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

type ClusterTaskList struct {
	Data []ClusterTask
}

type ClusterTask struct {
	EndTime   int
	Id        string
	Node      string
	Saved     string
	StartTime string
	Status    string
	Type      string
	UPId      string
	User      string
}

func GetClusterTasks(host string, auth AuthInfo) []ClusterTask {
	url := host + "/api2/json/cluster/tasks"
	body := GetContent(url, auth)
	var clusterTasks ClusterTaskList
	json.Unmarshal(body, &clusterTasks)
	return clusterTasks.Data
}

func printClusterTasks(clusterTasks []ClusterTask) {
	for _, clusterTask := range clusterTasks {
		printClusterTask(clusterTask)
	}
}

func printClusterTask(clusterTask ClusterTask) {
	s := reflect.ValueOf(clusterTask)
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%s: %v\n", typeOfT.Field(i).Name, f.Interface())
	}
	fmt.Printf("\n")
}

type ClusterBackupSchedule struct {
	Data []ClusterBackupScheduleItem
}

type ClusterBackupScheduleItem struct {
	All       int
	Compress  string
	DOW       string
	Exclude   string
	Id        string
	MailTo    string
	Mode      string
	Quiet     int
	StartTime string
	Storage   string
}

func GetClusterBackupSchedule(host string, auth AuthInfo) []ClusterBackupScheduleItem {
	url := host + "/api2/json/cluster/backup"
	body := GetContent(url, auth)
	var clusterBackupSchedule ClusterBackupSchedule
	json.Unmarshal(body, &clusterBackupSchedule)
	return clusterBackupSchedule.Data
}

func printClusterBackupSchedule(clusterBackupSchedule []ClusterBackupScheduleItem) {
	for _, clusterBackupScheduleItem := range clusterBackupSchedule {
		printClusterBackupScheduleItem(clusterBackupScheduleItem)
	}
}

func printClusterBackupScheduleItem(clusterBackupScheduleItem ClusterBackupScheduleItem) {
	s := reflect.ValueOf(clusterBackupScheduleItem)
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%s: %v\n", typeOfT.Field(i).Name, f.Interface())
	}
	fmt.Printf("\n")
}
