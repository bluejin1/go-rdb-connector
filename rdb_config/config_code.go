package rdb_config

import "sync"

const (
	RdbTimeout              string = "600"
	RdbMaxIdleCnt           string = "1000"
	RdbMaxOpenConn          string = "1000"
	RdbCharset              string = "utf8"
	RdbTimezone             string = "UTC"
	RdbDatabasesType        string = "mariadb" // mariadb or mysql
	RdbDatabasesTypeMariaDb string = "mariadb"
	RdbDatabasesTypeMysql   string = "mysql"
	rdbDebugLevel           string = "TRACE"
)

var (
	EnvServerType *string = nil
)

var (
	RdbConfigMaster                 *RdbServerMaster
	rdbConfigOnce                   sync.Once
	RdbHost                         *string
	RdbPort                         *string
	RdbUser                         *string
	RdbPassword                     *string
	CodeRdbConfigDatabaseTypeMaster string = "master"
)

var (
	RdbConfigLog                 *RdbServerLog
	rdbConfigLogOnce             sync.Once
	RdbLogHost                   *string
	RdbLogUser                   *string
	RdbLogPort                   *string
	RdbLogPassword               *string
	CodeRdbConfigDatabaseTypeLog string = "log"
)

var (
	RdbConfigStatistics                 *RdbServerStatistics
	rdbConfigStatisticsOnce             sync.Once
	RdbStatisticsHost                   *string
	RdbStatisticsUser                   *string
	RdbStatisticsPort                   *string
	RdbStatisticsPassword               *string
	CodeRdbConfigDatabaseTypeStatistics string = "log"
)
