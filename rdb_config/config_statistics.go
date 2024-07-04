package rdb_config

import (
	"rdb/rdb_helper"
)

func IsUseRdbStatisticsDatabase() bool {
	useStatisticsDatabase := rdb_helper.GetEnvAsInt("RDB_USE_STATISTICS_DB", 1)
	if useStatisticsDatabase == 1 {
		return true
	}
	return false
}

func GetRdbStatisticsClusterConnectionStr() string {
	clusterUrl := rdb_helper.GetEnv("RDB_COLLECTION_CLUSTER_URL", "")
	return clusterUrl
}

func GetRdbStatisticsConnectionStr() string {
	var dbConnection string = ""
	if RdbConfigStatistics != nil {
		if RdbConfigStatistics.User == "" || RdbConfigStatistics.Password == "" || RdbConfigStatistics.Address.Host == "" || RdbConfigStatistics.Database == "" {
			return dbConnection
		}
		dbConnection = "" + RdbConfigStatistics.User + ":" + RdbConfigStatistics.Password + "@tcp(" + RdbConfigStatistics.Address.Host + ":" + RdbConfigStatistics.Address.Port + ")/" + RdbConfigStatistics.Database + "?charset=" + RdbConfigStatistics.MaxSetting.Charset + "&parseTime=True&loc=" + RdbConfigStatistics.MaxSetting.Timezone
	}

	return dbConnection
}

func InitSetConfigStatistics(statisticsHost *string, statisticsPort *string, statisticsUser *string, statisticsPass *string) *RdbServerStatistics {
	RdbStatisticsHost = statisticsHost
	RdbStatisticsPort = statisticsPort
	RdbStatisticsUser = statisticsUser
	RdbStatisticsPassword = statisticsPass
	return SetStatisticsRdbServerConfigFromEnv()
}

func SetStatisticsRdbServerConfigFromEnv() *RdbServerStatistics {
	rdbConfigStatisticsOnce.Do(func() {
		RdbConfigStatistics = &RdbServerStatistics{
			RdbType:       rdb_helper.GetEnv("RDB_COLLECTION_TYPE", RdbDatabasesType),
			User:          rdb_helper.GetEnv("RDB_COLLECTION_USER", *RdbStatisticsUser),
			Password:      rdb_helper.GetEnv("RDB_COLLECTION_PW", *RdbStatisticsPassword),
			Database:      rdb_helper.GetEnv("RDB_COLLECTION_DBNAME", RdbStatisticsDatabaseName),
			DbType:        CodeRdbConfigDatabaseTypeLog,
			ConnectionStr: GetRdbStatisticsClusterConnectionStr(),
			Address: AddressInfo{
				Host: rdb_helper.GetEnv("RDB_COLLECTION_HOST", *RdbStatisticsHost),
				Port: rdb_helper.GetEnv("RDB_COLLECTION_PORT", *RdbStatisticsPort),
			},
			MaxSetting: RdbCommonConfig{
				Timeout:            rdb_helper.GetEnv("RDB_COLLECTION_TIMEOUT", RdbTimeout),
				MaxIdleConnections: rdb_helper.GetEnv("RDB_COLLECTION_MAX_IDLE_CNT", RdbMaxIdleCnt),
				MaxOpenConnections: rdb_helper.GetEnv("RDB_COLLECTION_MAX_OPEN_CONN", RdbMaxOpenConn),
				Charset:            rdb_helper.GetEnv("RDB_COLLECTION_CHARSET", RdbCharset),
				Timezone:           rdb_helper.GetEnv("RDB_COLLECTION_TIMEZONE", RdbTimezone),
			},
			DebugLevel: rdb_helper.GetEnv("DEBUG_LEVEL", rdbDebugLevel),
		}
	})
	return RdbConfigStatistics
}
