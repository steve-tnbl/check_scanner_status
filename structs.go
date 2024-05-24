package main

type IOScannerList struct {
	Scanners []struct {
		CreationDate           int      `json:"creation_date"`
		Distro                 string   `json:"distro,omitempty"`
		EngineVersion          string   `json:"engine_version,omitempty"`
		Group                  bool     `json:"group"`
		Hostname               string   `json:"hostname,omitempty"`
		ID                     int      `json:"id"`
		IPAddresses            []string `json:"ip_addresses,omitempty"`
		Key                    string   `json:"key"`
		LastConnect            int      `json:"last_connect"`
		LastModificationDate   int      `json:"last_modification_date"`
		Linked                 int      `json:"linked"`
		LoadedPluginSet        string   `json:"loaded_plugin_set,omitempty"`
		Name                   string   `json:"name"`
		NetworkName            string   `json:"network_name"`
		NumScans               int      `json:"num_scans"`
		Owner                  string   `json:"owner"`
		OwnerID                int      `json:"owner_id"`
		OwnerName              string   `json:"owner_name"`
		OwnerUUID              string   `json:"owner_uuid"`
		Platform               string   `json:"platform,omitempty"`
		Pool                   bool     `json:"pool"`
		ScanCount              int      `json:"scan_count"`
		Shared                 int      `json:"shared"`
		Source                 string   `json:"source"`
		Status                 string   `json:"status"`
		Timestamp              int      `json:"timestamp"`
		Type                   string   `json:"type"`
		UIBuild                string   `json:"ui_build,omitempty"`
		UIVersion              string   `json:"ui_version,omitempty"`
		UserPermissions        int      `json:"user_permissions"`
		UUID                   string   `json:"uuid"`
		RemoteUUID             string   `json:"remote_uuid,omitempty"`
		SupportsRemoteLogs     bool     `json:"supports_remote_logs"`
		SupportsWebapp         bool     `json:"supports_webapp"`
		SupportsRemoteSettings bool     `json:"supports_remote_settings"`
		License                struct {
			Type            string `json:"type"`
			Agents          int    `json:"agents"`
			Ips             int    `json:"ips"`
			Scanners        int    `json:"scanners"`
			Users           int    `json:"users"`
			EnterprisePause bool   `json:"enterprise_pause"`
			ExpirationDate  int    `json:"expiration_date"`
			Evaluation      bool   `json:"evaluation"`
			Apps            struct {
				Pci struct {
					Mode string `json:"mode"`
				} `json:"pci"`
				Consec struct {
					Mode           string `json:"mode"`
					ExpirationDate int    `json:"expiration_date"`
				} `json:"consec"`
				Was struct {
					Mode           string `json:"mode"`
					ExpirationDate int    `json:"expiration_date"`
				} `json:"was"`
				Lumin struct {
					Mode           string `json:"mode"`
					ExpirationDate int    `json:"expiration_date"`
				} `json:"lumin"`
				Cns struct {
					Mode           string `json:"mode"`
					ExpirationDate int    `json:"expiration_date"`
				} `json:"cns"`
				Asm struct {
					Mode           string `json:"mode"`
					ExpirationDate int    `json:"expiration_date"`
					Domains        int    `json:"domains"`
					ScanFrequency  string `json:"scan_frequency"`
				} `json:"asm"`
			} `json:"apps"`
			ScannersUsed int `json:"scanners_used"`
			AgentsUsed   int `json:"agents_used"`
		} `json:"license,omitempty"`
	} `json:"scanners"`
}

type IOScannerDetails struct {
	CreationDate           int      `json:"creation_date"`
	Distro                 string   `json:"distro"`
	EngineVersion          string   `json:"engine_version"`
	Group                  bool     `json:"group"`
	Hostname               string   `json:"hostname"`
	ID                     int      `json:"id"`
	IPAddresses            []string `json:"ip_addresses"`
	Key                    string   `json:"key"`
	LastConnect            int      `json:"last_connect"`
	LastModificationDate   int      `json:"last_modification_date"`
	Linked                 int      `json:"linked"`
	LoadedPluginSet        string   `json:"loaded_plugin_set"`
	Name                   string   `json:"name"`
	NetworkName            string   `json:"network_name"`
	NumScans               int      `json:"num_scans"`
	Owner                  string   `json:"owner"`
	OwnerID                int      `json:"owner_id"`
	OwnerName              string   `json:"owner_name"`
	OwnerUUID              string   `json:"owner_uuid"`
	Platform               string   `json:"platform"`
	Pool                   bool     `json:"pool"`
	ScanCount              int      `json:"scan_count"`
	Shared                 int      `json:"shared"`
	Source                 string   `json:"source"`
	Status                 string   `json:"status"`
	Timestamp              int      `json:"timestamp"`
	Type                   string   `json:"type"`
	UIBuild                string   `json:"ui_build"`
	UIVersion              string   `json:"ui_version"`
	UserPermissions        int      `json:"user_permissions"`
	UUID                   string   `json:"uuid"`
	RemoteUUID             string   `json:"remote_uuid"`
	SupportsRemoteLogs     bool     `json:"supports_remote_logs"`
	SupportsWebapp         bool     `json:"supports_webapp"`
	SupportsRemoteSettings bool     `json:"supports_remote_settings"`
	RemoteSettings         []struct {
		Name            string `json:"name"`
		Setting         string `json:"setting"`
		Type            string `json:"type"`
		Description     string `json:"description"`
		BackendReload   bool   `json:"backend_reload"`
		Status          string `json:"status"`
		Value           string `json:"value"`
		AllowableValues []struct {
			Value string `json:"value"`
		} `json:"allowable_values,omitempty"`
		Default string `json:"default"`
		Min     int    `json:"min,omitempty"`
	} `json:"remote_settings"`
}

type IOScannerGroups struct {
	ScannerPools []struct {
		CreationDate         int    `json:"creation_date"`
		LastModificationDate int    `json:"last_modification_date"`
		OwnerID              int    `json:"owner_id"`
		Owner                string `json:"owner"`
		OwnerUUID            string `json:"owner_uuid"`
		DefaultPermissions   int    `json:"default_permissions"`
		UserPermissions      int    `json:"user_permissions"`
		Shared               int    `json:"shared"`
		ScanCount            int    `json:"scan_count"`
		ScannerCount         int    `json:"scanner_count"`
		UUID                 string `json:"uuid"`
		Type                 string `json:"type"`
		Name                 string `json:"name"`
		NetworkName          string `json:"network_name"`
		ID                   int    `json:"id"`
		SupportsWebapp       bool   `json:"supports_webapp"`
		ScannerID            int    `json:"scanner_id"`
		ScannerUUID          string `json:"scanner_uuid"`
		OwnerName            string `json:"owner_name"`
	} `json:"scanner_pools"`
}

type IOGroupDetails struct {
	CreationDate         int    `json:"creation_date"`
	LastModificationDate int    `json:"last_modification_date"`
	OwnerID              int    `json:"owner_id"`
	Owner                string `json:"owner"`
	OwnerUUID            string `json:"owner_uuid"`
	DefaultPermissions   int    `json:"default_permissions"`
	UserPermissions      int    `json:"user_permissions"`
	Shared               int    `json:"shared"`
	ScanCount            int    `json:"scan_count"`
	UUID                 string `json:"uuid"`
	Type                 string `json:"type"`
	Name                 string `json:"name"`
	NetworkName          string `json:"network_name"`
	ID                   int    `json:"id"`
	SupportsWebapp       bool   `json:"supports_webapp"`
	ScannerID            int    `json:"scanner_id"`
	ScannerUUID          string `json:"scanner_uuid"`
	OwnerName            string `json:"owner_name"`
}

type IOScannersInGroup struct {
	Scanners []struct {
		CreationDate           int      `json:"creation_date"`
		Distro                 string   `json:"distro"`
		EngineVersion          string   `json:"engine_version"`
		Group                  bool     `json:"group"`
		Hostname               string   `json:"hostname"`
		ID                     int      `json:"id"`
		IPAddresses            []string `json:"ip_addresses"`
		Key                    string   `json:"key"`
		LastConnect            int      `json:"last_connect"`
		LastModificationDate   int      `json:"last_modification_date"`
		Linked                 int      `json:"linked"`
		LoadedPluginSet        string   `json:"loaded_plugin_set"`
		Name                   string   `json:"name"`
		NumScans               int      `json:"num_scans"`
		Owner                  string   `json:"owner"`
		OwnerID                int      `json:"owner_id"`
		OwnerName              string   `json:"owner_name"`
		OwnerUUID              string   `json:"owner_uuid"`
		Platform               string   `json:"platform"`
		Pool                   bool     `json:"pool"`
		ScanCount              int      `json:"scan_count"`
		Source                 string   `json:"source"`
		Status                 string   `json:"status"`
		Timestamp              int      `json:"timestamp"`
		Type                   string   `json:"type"`
		UIBuild                string   `json:"ui_build"`
		UIVersion              string   `json:"ui_version"`
		UUID                   string   `json:"uuid"`
		RemoteUUID             string   `json:"remote_uuid"`
		SupportsRemoteLogs     bool     `json:"supports_remote_logs"`
		SupportsRemoteSettings bool     `json:"supports_remote_settings"`
	} `json:"scanners"`
}

type SCScanZones struct {
	Type     string `json:"type"`
	Response []struct {
		ID             string        `json:"id"`
		Name           string        `json:"name"`
		Description    string        `json:"description"`
		IPList         string        `json:"ipList"`
		CreatedTime    string        `json:"createdTime"`
		ModifiedTime   string        `json:"modifiedTime"`
		Scanners       []interface{} `json:"scanners"`
		Organizations  []interface{} `json:"organizations"`
		ActiveScanners int           `json:"activeScanners"`
		TotalScanners  int           `json:"totalScanners"`
		UUID           string        `json:"uuid"`
		Sci            struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"SCI"`
	} `json:"response"`
	ErrorCode int           `json:"error_code"`
	ErrorMsg  string        `json:"error_msg"`
	Warnings  []interface{} `json:"warnings"`
	Timestamp int           `json:"timestamp"`
}

type SCScannerList struct {
	Type     string `json:"type"`
	Response []struct {
		ID            string      `json:"id"`
		Name          string      `json:"name"`
		Description   string      `json:"description"`
		AgentCapable  string      `json:"agentCapable"`
		Status        string      `json:"status"`
		StatusMessage interface{} `json:"statusMessage"`
		Sci           struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"SCI"`
	} `json:"response"`
	ErrorCode int           `json:"error_code"`
	ErrorMsg  string        `json:"error_msg"`
	Warnings  []interface{} `json:"warnings"`
	Timestamp int           `json:"timestamp"`
}
