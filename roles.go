package main

import (
	"encoding/json"
)

type RoleList struct {
	Data []Role
}

type Role struct {
	Privs  string
	RoleId string
}

func (p Proxmox) GetRoles() []Role {
	endpoint_url := "/api2/json/access/roles"
	body := p.GetContent(endpoint_url)
	var roles RoleList
	json.Unmarshal(body, &roles)
	return roles.Data
}

type RoleConfigList struct {
	Data RoleConfig
}

type RoleConfig struct {
	GroupAllocate     int `json:"Group.Allocate"`
	PermissionsModify int `json:"Permissions.Modify"`
	PoolAllocate      int `json:"Pool.Allocate"`
	RealmAllocateUser int `json:"Realm.AllocateUser"`
	SysAudit          int `json:"Sys.Audit"`
	SysConsole        int `json:"Sys.Console"`
	SysSyslog         int `json:"Sys.Syslog"`
	UserModify        int `json:"User.Modify"`
	VMAllocate        int `json:"VM.Allocate"`
	VMAudit           int `json:"VM.Audit"`
	VMBackup          int `json:"VM.Backup"`
	VMConfigCDROM     int `json:"VM.Config.CDROM"`
	VMConfigCPU       int `json:"VM.Config.CPU"`
	VMConfigDisk      int `json:"VM.Config.Disk"`
	VMConfigHWType    int `json:"VM.Config.HWType"`
	VMConfigMemory    int `json:"VM.Config.Memory"`
	VMConfigNetwork   int `json:"VM.Config.Network"`
	VMConfigOptions   int `json:"VM.Config.Options"`
	VMConsole         int `json:"VM.Console"`
	VMMigrate         int `json:"VM.Migrate"`
	VMMonitor         int `json:"VM.Monitor"`
	VMPowerMgmt       int `json:"VM.PowerMgmt"`
	VMSnapshot        int `json:"VM.Snapshot"`
}

func (p Proxmox) GetRoleConfig(role Role) RoleConfig {
	endpoint_url := "/api2/json/access/roles/" + role.RoleId
	body := p.GetContent(endpoint_url)
	var roleConfig RoleConfigList
	json.Unmarshal(body, &roleConfig)
	return roleConfig.Data
}
