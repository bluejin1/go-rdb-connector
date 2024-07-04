package rdb_config

import (
	"rdb/rdb_helper"
)

func IsUseRdbLogDatabase() bool {
	useDatabase := rdb_helper.GetEnv("RDB_USE_LOG_DB", "true")
	if useDatabase == "true" {
		return true
	}
	return false
}

func GetRdbLobClusterConnectionStr() string {
	clusterUrl := rdb_helper.GetEnv("RDB_LOG_CLUSTER_URL", "")
	return clusterUrl
}

func GetRdbLogConnectionStr() string {
	var dbConnection string = ""
	if RdbConfigLog == nil {
		return ""
	}
	if RdbConfigLog.User == "" || RdbConfigLog.Password == "" || RdbConfigLog.Address.Host == "" || RdbConfigLog.Database == "" {
		return dbConnection
	}
	dbConnection = "" + RdbConfigLog.User + ":" + RdbConfigLog.Password + "@tcp(" + RdbConfigLog.Address.Host + ":" + RdbConfigLog.Address.Port + ")/" + RdbConfigLog.Database + "?charset=" + RdbConfigLog.MaxSetting.Charset + "&parseTime=True&loc=" + RdbConfigLog.MaxSetting.Timezone
	return dbConnection
}

func InitSetConfigLog(logHost *string, logPort *string, logUser *string, logPass *string) *RdbServerLog {
	RdbLogHost = logHost
	RdbLogPort = logPort
	RdbLogUser = logUser
	RdbLogPassword = logPass
	return SetLogRdbServerConfigFromEnv()
}

func SetLogRdbServerConfigFromEnv() *RdbServerLog {
	rdbConfigLogOnce.Do(func() {

		RdbConfigLog = &RdbServerLog{
			RdbType:       rdb_helper.GetEnv("RDB_LOG_TYPE", RdbDatabasesType),
			User:          rdb_helper.GetEnv("RDB_LOG_USER", *RdbLogUser),
			Password:      rdb_helper.GetEnv("RDB_LOG_PW", *RdbLogPassword),
			Database:      rdb_helper.GetEnv("RDB_LOG_DBNAME", RdbLogDatabaseName),
			DbType:        CodeRdbConfigDatabaseTypeLog,
			ConnectionStr: GetRdbLobClusterConnectionStr(),
			Address: AddressInfo{
				Host: rdb_helper.GetEnv("RDB_LOG_HOST", *RdbLogHost),
				Port: rdb_helper.GetEnv("RDB_LOG_PORT", *RdbLogPort),
			},
			MaxSetting: RdbCommonConfig{
				Timeout:            rdb_helper.GetEnv("RDB_LOG_TIMEOUT", RdbTimeout),
				MaxIdleConnections: rdb_helper.GetEnv("RDB_LOG_MAX_IDLE_CNT", RdbMaxIdleCnt),
				MaxOpenConnections: rdb_helper.GetEnv("RDB_LOG_MAX_OPEN_CONN", RdbMaxOpenConn),
				Charset:            rdb_helper.GetEnv("RDB_LOG_CHARSET", RdbCharset),
				Timezone:           rdb_helper.GetEnv("RDB_LOG_TIMEZONE", RdbTimezone),
			},
			DebugLevel: rdb_helper.GetEnv("DEBUG_LEVEL", rdbDebugLevel),
		}
	})
	return RdbConfigLog
}
