package rdb_config

type AddressInfo struct {
	Host string // localhost,
	Port string
}

type RdbCommonConfig struct {
	Timeout            string
	MaxIdleConnections string
	MaxOpenConnections string
	Charset            string
	Timezone           string
}

type RdbServerMaster struct {
	RdbType       string
	User          string
	Password      string
	Database      string
	ConnectionStr string
	DbType        string
	MaxSetting    RdbCommonConfig
	Address       AddressInfo
	DebugLevel    string
}

type RdbServerLog struct {
	RdbType       string
	User          string
	Password      string
	Database      string
	ConnectionStr string
	DbType        string
	MaxSetting    RdbCommonConfig
	Address       AddressInfo
	DebugLevel    string
}

type RdbServerStatistics struct {
	RdbType       string
	User          string
	Password      string
	Database      string
	ConnectionStr string
	DbType        string
	MaxSetting    RdbCommonConfig
	Address       AddressInfo
	DebugLevel    string
}
