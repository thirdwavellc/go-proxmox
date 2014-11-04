package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

type AuthRequest struct {
	Data AuthInfo
}

type AuthInfo struct {
	CSRFPreventionToken string
	Ticket              string
	Username            string
}

func printError(err error) {
	fmt.Println("There was an error...")
	fmt.Printf("Error: %v", err)
	os.Exit(1)
}

func getContent(url string, auth AuthInfo) []byte {
	fmt.Println("Fetching:", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		printError(err)
	}

	cookie := http.Cookie{Name: "PVEAuthCookie", Value: auth.Ticket, Expires: time.Now().Add(356 * 24 * time.Hour), HttpOnly: true}
	req.AddCookie(&cookie)

	// Send the request via a client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		printError(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		printError(err)
	}

	return body
}

func getAuth(host string, username string, password string) AuthInfo {
	requestUrl := host + "/api2/json/access/ticket"
	values := make(url.Values)
	values.Set("username", username)
	values.Set("password", password)
	resp, err := http.PostForm(requestUrl, values)
	if err != nil {
		printError(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		printError(err)
	}

	var auth AuthRequest
	json.Unmarshal(body, &auth)

	return auth.Data
}

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

func getNodesAction(host string, auth AuthInfo) {
	nodes := getNodes(host, auth)
	printNodes(nodes)
}

func getNodes(host string, auth AuthInfo) []Node {
	url := host + "/api2/json/nodes"
	body := getContent(url, auth)
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
	fmt.Printf("Node: %s\n", node.Node)
	fmt.Printf("Disk: %d\n", node.Disk)
	fmt.Printf("CPU: %f\n", node.CPU)
	fmt.Printf("MaxDisk: %d\n", node.MaxDisk)
	fmt.Printf("MaxMem: %d\n", node.MaxMem)
	fmt.Printf("MaxCPU: %d\n", node.MaxCPU)
	fmt.Printf("Uptime: %d\n", node.Uptime)
	fmt.Printf("Id: %s\n", node.Id)
	fmt.Printf("Type: %s\n", node.Type)
	fmt.Printf("\n")
}

type ContainerList struct {
	Data []Container
}

type Container struct {
	MaxSwap   string
	Disk      string
	IP        string
	Status    string
	Netout    int
	MaxDisk   int
	MaxMem    int
	Uptime    int
	Swap      int
	VMID      string
	NProc     string
	DiskRead  int
	CPU       float64
	NetIn     int
	Name      string
	FailCnt   int
	DiskWrite int
	Mem       int
	Type      string
	CPUs      int
}

func getContainers(host string, node string, auth AuthInfo) []Container {
	url := host + "/api2/json/nodes/" + node + "/openvz"
	body := getContent(url, auth)
	fmt.Println(string(body[:]))
	var containers ContainerList
	json.Unmarshal(body, &containers)
	return containers.Data
}

func printContainers(containers []Container) {
	for _, container := range containers {
		printContainer(container)
	}
}

func printContainer(container Container) {
	fmt.Printf("Max Swap: %s\n", container.MaxSwap)
	fmt.Printf("Disk: %s\n", container.Disk)
	fmt.Printf("IP: %s\n", container.IP)
	fmt.Printf("Status: %s\n", container.Status)
	fmt.Printf("Net Out: %d\n", container.Netout)
	fmt.Printf("Max Disk: %d\n", container.MaxDisk)
	fmt.Printf("Max Mem: %d\n", container.MaxMem)
	fmt.Printf("Uptime: %d\n", container.Uptime)
	fmt.Printf("Swap: %d\n", container.Swap)
	fmt.Printf("VMID: %s\n", container.VMID)
	fmt.Printf("NProc: %s\n", container.NProc)
	fmt.Printf("Disk Read: %d\n", container.DiskRead)
	fmt.Printf("CPU: %f\n", container.CPU)
	fmt.Printf("Net In: %d\n", container.NetIn)
	fmt.Printf("Name: %s\n", container.Name)
	fmt.Printf("Fail Count: %d\n", container.FailCnt)
	fmt.Printf("Disk Write: %d\n", container.DiskWrite)
	fmt.Printf("Mem: %d\n", container.Mem)
	fmt.Printf("Type: %s\n", container.Type)
	fmt.Printf("CPUs: %d\n", container.CPUs)
	fmt.Printf("\n")
}

func main() {
	var host string
	var user string
	var password string
	var action string
	var node string
	flag.StringVar(&host, "host", "", "Proxmox host")
	flag.StringVar(&user, "user", "root@pam", "Proxmox user")
	flag.StringVar(&password, "password", "", "Proxmox user password")
	flag.StringVar(&action, "action", "", "Proxmox api action")
	flag.StringVar(&node, "node", "", "Proxmox node")

	flag.Parse()

	auth := getAuth(host, user, password)

	switch action {
	case "getNodes":
		getNodesAction(host, auth)
	case "getContainers":
		containers := getContainers(host, node, auth)
		printContainers(containers)
	}
}
