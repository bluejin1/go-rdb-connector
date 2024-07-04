package rdb_master

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
	rdbMasterConfig      = &rdb_config.RdbServerMaster{}
	IsRdbMasterInit bool = false
)

var RdbConnMaster *MasterDB

type MasterDB struct {
	DbConn       *gorm.DB
	IsConnect    bool
	DbType       *string
	DatabaseName string
}

func (r *MasterDB) ConnectMaster() (db *gorm.DB, err error) {
	fmt.Printf(">>> ConnectMaster <<<")
	r.DbType = &rdb_config.CodeRdbConfigDatabaseTypeMaster
	r.DatabaseName = rdbMasterConfig.Database
	connectionStr := rdb_config.GetRdbConnectionStr()

	// 만약 클러스터 url 이 있으면
	if rdbMasterConfig.ConnectionStr != "" && len(rdbMasterConfig.ConnectionStr) > 0 {
		connectionStr = rdbMasterConfig.ConnectionStr
	}
	if connectionStr == "" {
		return nil, errors.New("connectionStr empty")
	}

	if rdbMasterConfig.RdbType == rdb_config.RdbDatabasesTypeMariaDb {

	} else if rdbMasterConfig.RdbType == rdb_config.RdbDatabasesTypeMysql {

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

func (r *MasterDB) SetConnectGorm(connection string, databaseName *string) (db *gorm.DB, err error) {
	fmt.Printf("DB connection databaseName: %s", *databaseName)
	//fmt.Printf("connection %v", connection)

	loggerLevel := logger.Error
	if rdbMasterConfig.DebugLevel == "TRACE" {
		loggerLevel = logger.Info
	}
	db, err = gorm.Open(mysql.New(mysql.Config{DSN: connection}), &gorm.Config{Logger: logger.Default.LogMode(loggerLevel)})
	if err != nil {
		//panic(err)
		fmt.Printf("SetConnectGorm err")
		return db, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		//panic(err)
		fmt.Printf("SetConnectGorm err")
		return db, err
	}
	maxIdleCnt, _ := strconv.Atoi(rdbMasterConfig.MaxSetting.MaxIdleConnections)
	maxOpenConn, _ := strconv.Atoi(rdbMasterConfig.MaxSetting.MaxOpenConnections)
	timeout, _ := strconv.Atoi(rdbMasterConfig.MaxSetting.Timeout)

	sqlDB.SetMaxIdleConns(maxIdleCnt)
	sqlDB.SetMaxOpenConns(maxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Duration(timeout))

	r.DbConn = db
	r.IsConnect = true
	RdbConnMaster = r
	return db, nil
}

func (r *MasterDB) Close() {
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

func InitRdbMaster(masterHost *string, masterPort *string, maserUser *string, masterPass *string) (*rdb_config.RdbServerMaster, error) {
	if rdb_config.IsUseRdbMasterDatabase() {
		rdbMasterConfig = rdb_config.InitSetConfigMaster(masterHost, masterPort, maserUser, masterPass)
		IsRdbMasterInit = true
		//fmt.Printf("init store_rdb master env : %v", rdbMasterConfig)
	} else {
		return rdbMasterConfig, errors.New("master database not use")
	}
	return rdbMasterConfig, nil
}
