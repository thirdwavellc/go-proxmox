package proxmox

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"net/url"
)

type BackupList struct {
	Data []Backup
}

type Backup struct {
	Id               string `json:"id"`
	StartTime        string `json:"starttime"`
	All              int    `json:"all"`
	BwLimit          int    `json:"bwlimit"`
	Compress         string `json:"compress"`
	DOW              string `json:"dow"`
	DumpDir          string `json:"dumpdir"`
	Enabled          int    `json:"enabled"`
	Exclude          string `json:"exclude"`
	ExcludePath      string `json:"exlude-path"`
	IoNice           int    `json:"ionice"`
	LockWait         int    `json:"lockwait"`
	MailNotification string `json:"mailnotification"`
	MailTo           string `json:"mailto"`
	MaxFiles         int    `json:"maxfiles"`
	Mode             string `json:"mode"`
	Node             string `json:"node"`
	Pigz             int    `json:"pigz"`
	Quiet            int    `json:"quiet"`
	Remove           int    `json:"remove"`
	Script           string `json:"script"`
	StdExcludes      int    `json:"stdexcludes"`
	Stop             int    `json:"stop"`
	StopWait         int    `json:"stopwait"`
	Storage          string `json:"storage"`
	TmpDir           string `json:"tmpdir"`
	VMID             string `json:"vmid"`
}

func (p ProxmoxClient) GetBackups() ([]Backup, error) {
	endpoint_url := "/api2/json/cluster/backup"
	payload := url.Values{}
	body, err := p.GetContent(endpoint_url, payload)

	if err != nil {
		return nil, err
	}

	var backups BackupList
	json.Unmarshal(body, &backups)
	return backups.Data, nil
}

type ExistingBackupRequest struct {
	Id               string   `url:"-"`
	StartTime        string   `url:"starttime,omitempty"`
	All              int      `url:"all,omitempty"`
	BwLimit          int      `url:"bwlimit,omitempty"`
	Compress         string   `url:"compress,omitempty"`
	DOW              []string `url:"dow,omitempty"`
	DumpDir          string   `url:"dumpdir,omitempty"`
	Enabled          int      `url:"enabled,omitempty"`
	Exclude          string   `url:"exclude,omitempty"`
	ExcludePath      string   `url:"exclude-path,omitempty"`
	IoNice           int      `url:"ionice,omitempty"`
	LockWait         int      `url:"lockwait,omitempty"`
	MailNotification string   `url:"mailnotification,omitempty"`
	MailTo           string   `url:"mailto,omitempty"`
	MaxFiles         int      `url:"maxfiles,omitempty"`
	Mode             string   `url:"mode,omitempty"`
	Node             string   `url:"node,omitempty"`
	Pigz             int      `url:"pigz,omitempty"`
	Quiet            int      `url:"quiet,omitempty"`
	Remove           int      `url:"remove,omitempty"`
	Script           string   `url:"script,omitempty"`
	StdExcludes      int      `url:"stdexcludes,omitempty"`
	Stop             int      `url:"stop,omitempty"`
	StopWait         int      `url:"stopwait,omitempty"`
	Storage          string   `url:"storage,omitempty"`
	TmpDir           string   `url:"tmpdir,omitempty"`
	VMID             []string `url:"vmid,omitempty"`
}

type BackupConfigWrapper struct {
	Data Backup
}

func (p ProxmoxClient) GetBackupConfig(req *ExistingBackupRequest) (Backup, error) {
	endpoint_url := "/api2/json/cluster/backup/" + req.Id
	payload := url.Values{}
	body, err := p.GetContent(endpoint_url, payload)

	if err != nil {
		return Backup{}, err
	}

	var backup BackupConfigWrapper
	json.Unmarshal(body, &backup)
	return backup.Data, nil
}

type NewBackupRequest struct {
	StartTime        string   `url:"starttime"`
	All              int      `url:"all,omitempty"`
	BwLimit          int      `url:"bwlimit,omitempty"`
	Compress         string   `url:"compress,omitempty"`
	DOW              []string `url:"dow,omitempty"`
	DumpDir          string   `url:"dumpdir,omitempty"`
	Enabled          int      `url:"enabled,omitempty"`
	Exclude          string   `url:"exclude,omitempty"`
	ExcludePath      string   `url:"exclude-path,omitempty"`
	IoNice           int      `url:"ionice,omitempty"`
	LockWait         int      `url:"lockwait,omitempty"`
	MailNotification string   `url:"mailnotification,omitempty"`
	MailTo           string   `url:"mailto,omitempty"`
	MaxFiles         int      `url:"maxfiles,omitempty"`
	Mode             string   `url:"mode,omitempty"`
	Node             string   `url:"node,omitempty"`
	Pigz             int      `url:"pigz,omitempty"`
	Quiet            int      `url:"quiet,omitempty"`
	Remove           int      `url:"remove,omitempty"`
	Script           string   `url:"script,omitempty"`
	StdExcludes      int      `url:"stdexcludes,omitempty"`
	Stop             int      `url:"stop,omitempty"`
	StopWait         int      `url:"stopwait,omitempty"`
	Storage          string   `url:"storage,omitempty"`
	TmpDir           string   `url:"tmpdir,omitempty"`
	VMID             []string `url:"vmid,omitempty"`
}

func (p ProxmoxClient) CreateBackup(req *NewBackupRequest) ([]byte, error) {
	endpoint_url := "/api2/json/cluster/backup"
	payload, _ := query.Values(req)
	body, err := p.PostContent(endpoint_url, payload)

	if err != nil {
		return nil, err
	}

	return body, nil
}

type BackupResponse struct {
	Data string
}

func (p ProxmoxClient) UpdateBackup(req *ExistingBackupRequest) (string, error) {
	endpoint_url := "/api2/json/cluster/backup/" + req.Id
	payload, _ := query.Values(req)
	body, err := p.PutContent(endpoint_url, payload)

	if err != nil {
		return "", err
	}

	var response BackupResponse
	json.Unmarshal(body, &response)
	return response.Data, nil
}

func (p ProxmoxClient) DeleteBackup(req *ExistingBackupRequest) (string, error) {
	endpoint_url := "/api2/json/cluster/backup/" + req.Id
	payload, _ := query.Values(req)
	body, err := p.DeleteContent(endpoint_url, payload)

	if err != nil {
		return "", err
	}

	var response BackupResponse
	json.Unmarshal(body, &response)
	return response.Data, nil
}
