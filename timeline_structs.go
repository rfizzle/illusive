package illusive

import (
	"encoding/json"
)

type TimelineItem struct {
	Details json.RawMessage `json:"details"`
	EventID *int            `json:"eventId,omitempty"`
	ID      string          `json:"id"`
	Source  string          `json:"source"`
	Starred bool            `json:"starred"`
	Time    int64           `json:"time"`
	Title   string          `json:"title"`
	Type    string          `json:"type"`
}

type TimelineDetailsMachineProcess struct {
	AdministratorPrivileges string `json:"administratorPrivileges"`
	CmdHistory              struct {
		ProcessConsoleOutput []string `json:"processConsoleOutput"`
	} `json:"cmdHistory"`
	Commandline string `json:"commandline"`
	Connections []struct {
		CreationTimestamp string `json:"creationTimestamp"`
		LocalIP           string `json:"localIp"`
		LocalPort         int    `json:"localPort"`
		ModuleName        string `json:"moduleName"`
		Protocol          string `json:"protocol"`
		RemoteIP          string `json:"remoteIp"`
		RemotePort        int    `json:"remotePort"`
		State             string `json:"state"`
	} `json:"connections"`
	CreationTime string `json:"creationTime"`
	Handles      struct {
		Files    []string `json:"files"`
		Registry []string `json:"registry"`
	} `json:"handles"`
	Md5               string `json:"md5"`
	ParentProcessInfo struct {
		Md5         string `json:"md5"`
		ProcessID   int    `json:"processId"`
		ProcessName string `json:"processName"`
		ProcessPath string `json:"processPath"`
		Sha256      string `json:"sha256"`
	} `json:"parentProcessInfo"`
	ProcessID   int    `json:"processId"`
	ProcessName string `json:"processName"`
	ProcessPath string `json:"processPath"`
	SessionID   int    `json:"sessionId"`
	SessionName string `json:"sessionName"`
	Sha256      string `json:"sha256"`
	User        string `json:"user"`
}

type TimelineDetailsManagementEvent struct {
	ID                   int      `json:"id"`
	Date                 string   `json:"date"`
	Type                 string   `json:"type"`
	SourceIP             string   `json:"sourceIP"`
	Trap                 string   `json:"trap"`
	DestinationIPAddress string   `json:"destinationIpAddress"`
	ServiceType          string   `json:"serviceType"`
	Data                 []string `json:"data,omitempty"`
	HasForensics         string   `json:"hasForensics"`
	Title                string   `json:"title"`
}

type TimelineDetailsMachineEventLog struct {
	Channel    string            `json:"channel"`
	DateTime   string            `json:"dateTime"`
	EventData  map[string]string `json:"eventData"`
	EventLogID int               `json:"eventLogId"`
	Level      string            `json:"level"`
	Message    string            `json:"message"`
	OpCode     string            `json:"opCode"`
	Provider   string            `json:"provider"`
	Source     string            `json:"source"`
	Task       string            `json:"task"`
	XMLData    struct {
		Event json.RawMessage `json:"Event"`
	} `json:"xmlData"`
}

type TimelineDetailsMachineNetstat struct {
	OwnerProcessID    int    `json:"ownerProcessId"`
	Protocol          string `json:"protocol"`
	RemoteIP          string `json:"remoteIp"`
	LocalPort         int    `json:"localPort"`
	ModuleName        string `json:"moduleName"`
	RemotePort        int    `json:"remotePort"`
	CreationTimestamp string `json:"creationTimestamp"`
	LocalIP           string `json:"localIp"`
	State             string `json:"state"`
}

type TimelineDetailsMachineDns struct {
	Data         string `json:"data"`
	CreationTime string `json:"creationTime"`
	DataLength   int    `json:"dataLength"`
	Name         string `json:"name"`
	Section      string `json:"section"`
	Type         string `json:"type"`
	TTL          int    `json:"ttl"`
}

type TimelineDetailsMachineKlist struct {
	BranchID             int    `json:"branchId"`
	CreationTime         string `json:"creationTime"`
	ClientName           string `json:"clientName"`
	ServerRealm          string `json:"serverRealm"`
	ServerName           string `json:"serverName"`
	TicketType           string `json:"ticketType"`
	UserName             string `json:"userName"`
	RenewTime            string `json:"renewTime"`
	TicketFlags          int    `json:"ticketFlags"`
	UserSid              string `json:"userSid"`
	EndTime              string `json:"endTime"`
	TicketEncryptionType int    `json:"ticketEncryptionType"`
}

type TimelineDetailsMachineSession struct {
	LogonTime      string `json:"logonTime"`
	SessionState   string `json:"sessionState"`
	DisconnectTime string `json:"disconnectTime"`
	SessionName    string `json:"sessionName"`
	SessionID      int    `json:"sessionId"`
	UserName       string `json:"userName"`
	ConnectionTime string `json:"connectionTime"`
}

type TimelineDetailsMachineUserAssist struct {
	ExecutionTime string `json:"executionTime"`
	FileName      string `json:"fileName"`
	Sha256        string `json:"sha256"`
	FilePath      string `json:"filePath"`
	Count         int    `json:"count"`
	FileExists    string `json:"fileExists"`
	UserName      string `json:"userName"`
	Md5           string `json:"md5"`
	FileDates     struct {
		AccessTime       string `json:"accessTime"`
		CreationTime     string `json:"creationTime"`
		ModificationTime string `json:"modificationTime"`
	} `json:"fileDates"`
}

type TimelineDetailsMachineShim struct {
	ImageName            string `json:"imageName"`
	Sha256               string `json:"sha256"`
	FileExist            string `json:"fileExist"`
	LastModificationTime string `json:"lastModificationTime"`
	ImagePath            string `json:"imagePath"`
	Executed             string `json:"executed"`
	Version              string `json:"version"`
	Md5                  string `json:"md5"`
}

type TimelineDetailsMachineInstalled struct {
	InstallationTime string `json:"installationTime"`
	Sha256           string `json:"sha256"`
	DisplayName      string `json:"displayName"`
	FilePath         string `json:"filePath"`
	FileExists       string `json:"fileExists"`
	Version          string `json:"version"`
	Md5              string `json:"md5"`
}

type TimelineDetailsManagementLastDeployment struct {
	LastDeployment string `json:"lastDeployment"`
}

type TimelineDetailsMachineMft struct {
	StdAccessTime        string `json:"stdAccessTime"`
	FilePath             string `json:"filePath"`
	StdCreatedTime       string `json:"stdCreatedTime"`
	StdMftTime           string `json:"stdMftTime"`
	FileTimesEqual       string `json:"fileTimesEqual"`
	FilenameMftTime      string `json:"filenameMftTime"`
	FilenameAccessTime   string `json:"filenameAccessTime"`
	Deleted              string `json:"deleted"`
	FileSize             int    `json:"fileSize"`
	StdModifiedTime      string `json:"stdModifiedTime"`
	FilenameCreatedTime  string `json:"filenameCreatedTime"`
	FilenameModifiedTime string `json:"filenameModifiedTime"`
}
