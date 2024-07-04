package rdb_config

import (
	"rdb/rdb_helper"
)

func IsUseRdbMasterDatabase() bool {
	useDatabase := rdb_helper.GetEnv("RDB_USE_MASTER_DB", "true")
	if useDatabase == "true" {
		return true
	}
	return false
}

func GetRdbClusterConnectionStr() string {
	clusterUrl := rdb_helper.GetEnv("RDB_CLUSTER_URL", "")
	return clusterUrl
}

func GetRdbConnectionStr() string {
	var dbConnection string = ""
	if RdbConfigMaster == nil {
		return dbConnection
	}
	if RdbConfigMaster.User == "" || RdbConfigMaster.Password == "" || RdbConfigMaster.Address.Host == "" || RdbConfigMaster.Database == "" {
		return dbConnection
	}
	dbConnection = "" + RdbConfigMaster.User + ":" + RdbConfigMaster.Password + "@tcp(" + RdbConfigMaster.Address.Host + ":" + RdbConfigMaster.Address.Port + ")/" + RdbConfigMaster.Database + "?charset=" + RdbConfigMaster.MaxSetting.Charset + "&parseTime=True&loc=" + RdbConfigMaster.MaxSetting.Timezone
	return dbConnection
}

func InitSetConfigMaster(masterHost *string, masterPort *string, maserUser *string, masterPass *string) *RdbServerMaster {
	RdbHost = masterHost
	RdbPort = masterPort
	RdbUser = maserUser
	RdbPassword = masterPass
	return SetMasterRdbServerConfigFromEnv()
}

func SetMasterRdbServerConfigFromEnv() *RdbServerMaster {
	rdbConfigOnce.Do(func() {

		RdbConfigMaster = &RdbServerMaster{
			RdbType:       rdb_helper.GetEnv("RDB_TYPE", RdbDatabasesType),
			User:          rdb_helper.GetEnv("RDB_USER", *RdbUser),
			Password:      rdb_helper.GetEnv("RDB_PW", *RdbPassword),
			Database:      rdb_helper.GetEnv("RDB_DBNAME", RdbMasterDatabaseName),
			DbType:        CodeRdbConfigDatabaseTypeMaster,
			ConnectionStr: GetRdbClusterConnectionStr(),
			Address: AddressInfo{
				Host: rdb_helper.GetEnv("RDB_HOST", *RdbHost),
				Port: rdb_helper.GetEnv("RDB_PORT", *RdbPort),
			},
			MaxSetting: RdbCommonConfig{
				Timeout:            rdb_helper.GetEnv("RDB_TIMEOUT", RdbTimeout),
				MaxIdleConnections: rdb_helper.GetEnv("RDB_MAX_IDLE_CNT", RdbMaxIdleCnt),
				MaxOpenConnections: rdb_helper.GetEnv("RDB_MAX_OPEN_CONN", RdbMaxOpenConn),
				Charset:            rdb_helper.GetEnv("RDB_CHARSET", RdbCharset),
				Timezone:           rdb_helper.GetEnv("RDB_TIMEZONE", RdbTimezone),
			},
			DebugLevel: rdb_helper.GetEnv("DEBUG_LEVEL", rdbDebugLevel),
		}
	})
	return RdbConfigMaster
}
