package proxmox

import (
	"encoding/json"
	"net/url"
)

type RoleList struct {
	Data []Role
}

type Role struct {
	Privs  string `json:"privs"`
	RoleId string `json:"roleid"`
}

func (p ProxmoxClient) GetRoles() ([]Role, error) {
	endpoint_url := "/api2/json/access/roles"
	payload := url.Values{}
	body, err := p.GetContent(endpoint_url, payload)

	if err != nil {
		return nil, err
	}

	var roles RoleList
	json.Unmarshal(body, &roles)
	return roles.Data, nil
}

type RoleConfigList struct {
	Data RoleConfig
}

type RoleConfig struct {
	DatastoreAllocate         int `json:"Datastore.Allocate"`
	DatastoreAllocateSpace    int `json:"Datastore.AllocateSpace"`
	DatastoreAllocateTemplate int `json:"Datastore.AllocateTemplate"`
	DatastoreAudit            int `json:"Datastore.Audit"`
	GroupAllocate             int `json:"Group.Allocate"`
	PermissionsModify         int `json:"Permissions.Modify"`
	PoolAllocate              int `json:"Pool.Allocate"`
	RealmAllocate             int `json:"Realm.Allocate"`
	RealmAllocateUser         int `json:"Realm.AllocateUser"`
	SysAudit                  int `json:"Sys.Audit"`
	SysConsole                int `json:"Sys.Console"`
	SysModify                 int `json:"Sys.Modify"`
	SysPowerMgmt              int `json:"Sys.PowerMgmt"`
	SysSyslog                 int `json:"Sys.Syslog"`
	UserModify                int `json:"User.Modify"`
	VMAllocate                int `json:"VM.Allocate"`
	VMAudit                   int `json:"VM.Audit"`
	VMBackup                  int `json:"VM.Backup"`
	VMClone                   int `json:"VM.Clone"`
	VMConfigCDROM             int `json:"VM.Config.CDROM"`
	VMConfigCPU               int `json:"VM.Config.CPU"`
	VMConfigDisk              int `json:"VM.Config.Disk"`
	VMConfigHWType            int `json:"VM.Config.HWType"`
	VMConfigMemory            int `json:"VM.Config.Memory"`
	VMConfigNetwork           int `json:"VM.Config.Network"`
	VMConfigOptions           int `json:"VM.Config.Options"`
	VMConsole                 int `json:"VM.Console"`
	VMMigrate                 int `json:"VM.Migrate"`
	VMMonitor                 int `json:"VM.Monitor"`
	VMPowerMgmt               int `json:"VM.PowerMgmt"`
	VMSnapshot                int `json:"VM.Snapshot"`
}

type RoleConfigRequest struct {
	RoleId string `url:"roleid"`
}

func (p ProxmoxClient) GetRoleConfig(req *RoleConfigRequest) (RoleConfig, error) {
	endpoint_url := "/api2/json/access/roles/" + req.RoleId
	payload := url.Values{}
	body, err := p.GetContent(endpoint_url, payload)

	if err != nil {
		return RoleConfig{}, err
	}

	var roleConfig RoleConfigList
	json.Unmarshal(body, &roleConfig)
	return roleConfig.Data, nil
}
