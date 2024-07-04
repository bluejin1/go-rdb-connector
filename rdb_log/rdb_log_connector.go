package rdb_log

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"rdb/rdb_config"
	"strconv"
	"time"
)

var (
	rdbLogConfig      = &rdb_config.RdbServerLog{}
	IsRdbLogInit bool = false
)

var RdbConnLog *LogDB

type LogDB struct {
	DbConn       *gorm.DB
	IsConnect    bool
	DbType       *string
	DatabaseName string
}

func (r *LogDB) ConnectLog() (db *gorm.DB, err error) {
	r.DbType = &rdb_config.CodeRdbConfigDatabaseTypeLog
	r.DatabaseName = rdbLogConfig.Database
	connectionStr := rdb_config.GetRdbLogConnectionStr()

	// 만약 클러스터 url 이 있으면
	if rdbLogConfig.ConnectionStr != "" && len(rdbLogConfig.ConnectionStr) > 0 {
		connectionStr = rdbLogConfig.ConnectionStr
	}
	if connectionStr == "" {
		return nil, errors.New("connectionStr empty")
	}

	if rdbLogConfig.RdbType == rdb_config.RdbDatabasesTypeMariaDb {

	} else if rdbLogConfig.RdbType == rdb_config.RdbDatabasesTypeMysql {

	} else {
		return nil, errors.New("not support rdbMasterConfig.RdbType")
	}

	db, err = r.SetConnectGorm(connectionStr, &r.DatabaseName)
	if err != nil {
		fmt.Printf("SetConnectGorm err %v", err)
		return nil, errors.New("SetConnectGorm err")
	}
	return db, err
}

func (r *LogDB) SetConnectGorm(connection string, databaseName *string) (db *gorm.DB, err error) {
	fmt.Printf("DB connection databaseName: %s", *databaseName)

	loggerLevel := logger.Error
	if rdbLogConfig.DebugLevel == "TRACE" {
		loggerLevel = logger.Info
	}

	db, err = gorm.Open(mysql.New(mysql.Config{DSN: connection}), &gorm.Config{Logger: logger.Default.LogMode(loggerLevel)})
	if err != nil {
		//panic(err)
		fmt.Printf("SetConnectGorm gorm.Open err")
		return db, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		//panic(err)
		fmt.Printf("SetConnectGorm sqlDB err")
		return db, err
	}
	maxIdleCnt, _ := strconv.Atoi(rdbLogConfig.MaxSetting.MaxIdleConnections)
	maxOpenConn, _ := strconv.Atoi(rdbLogConfig.MaxSetting.MaxOpenConnections)
	timeout, _ := strconv.Atoi(rdbLogConfig.MaxSetting.Timeout)

	sqlDB.SetMaxIdleConns(maxIdleCnt)
	sqlDB.SetMaxOpenConns(maxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Duration(timeout))

	r.DbConn = db
	r.IsConnect = true
	RdbConnLog = r
	return db, nil
}

func (r *LogDB) Close() {
	if r.IsConnect == true {
		sqlDB, err := r.DbConn.DB()
		if err != nil {
			err = sqlDB.Close()
			if err != nil {
				fmt.Printf("sqlDB.Close err %v", err)
			}
		}
	}
}

func InitRdbLog(logHost *string, logPort *string, logUser *string, logPass *string) (*rdb_config.RdbServerLog, error) {
	if rdb_config.IsUseRdbLogDatabase() {
		rdbLogConfig = rdb_config.InitSetConfigLog(logHost, logPort, logUser, logPass)
		IsRdbLogInit = true
	} else {
		return rdbLogConfig, errors.New("log database not use")
	}
	return rdbLogConfig, nil
}
