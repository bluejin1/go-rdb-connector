package rdb_statistics

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"rdb/rdb_config"
	"strconv"
	"time"
)

var (
	RdbStatisticsConfig      = &rdb_config.RdbServerStatistics{}
	IsRdbStatisticsInit bool = false
)

var RdbConnCollection *CollectionDB

type CollectionDB struct {
	DbConn       *gorm.DB
	IsConnect    bool
	DbType       *string
	DatabaseName string
}

func (r *CollectionDB) ConnectCollection() (db *gorm.DB, err error) {
	r.DbType = &rdb_config.CodeRdbConfigDatabaseTypeStatistics
	r.DatabaseName = RdbStatisticsConfig.Database
	connectionStr := rdb_config.GetRdbStatisticsConnectionStr()

	// 만약 클러스터 url 이 있으면
	if RdbStatisticsConfig.ConnectionStr != "" && len(RdbStatisticsConfig.ConnectionStr) > 0 {
		connectionStr = RdbStatisticsConfig.ConnectionStr
	}
	if connectionStr == "" {
		return nil, errors.New("connectionStr empty")
	}

	if RdbStatisticsConfig.RdbType == rdb_config.RdbDatabasesTypeMariaDb {

	} else if RdbStatisticsConfig.RdbType == rdb_config.RdbDatabasesTypeMysql {

	} else {
		return nil, errors.New("not support rdbMasterConfig.RdbType")
	}
	fmt.Printf("ConnectCollection connectionStr %v", connectionStr)
	fmt.Printf("rrrrrr %v", r)

	db, err = r.SetConnectGorm(connectionStr, &r.DatabaseName)
	if err != nil {
		return nil, errors.New("DB create perror")
	}
	return db, err
}

func (r *CollectionDB) SetConnectGorm(connection string, databaseName *string) (db *gorm.DB, err error) {
	log.Printf("DB connection databaseName: %s", databaseName)

	loggerLevel := logger.Error
	if RdbStatisticsConfig.DebugLevel == "TRACE" {
		loggerLevel = logger.Info
	}

	db, err = gorm.Open(mysql.New(mysql.Config{DSN: connection}), &gorm.Config{Logger: logger.Default.LogMode(loggerLevel)})
	if err != nil {
		//panic(err)
		log.Printf("SetConnectGorm  gorm.Open err %v", err)
		return db, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		//panic(err)
		log.Printf("SetConnectGorm err %v", err)
		return db, err
	}

	maxIdleCnt, _ := strconv.Atoi(RdbStatisticsConfig.MaxSetting.MaxIdleConnections)
	maxOpenConn, _ := strconv.Atoi(RdbStatisticsConfig.MaxSetting.MaxOpenConnections)
	timeout, _ := strconv.Atoi(RdbStatisticsConfig.MaxSetting.Timeout)
	sqlDB.SetMaxIdleConns(maxIdleCnt)
	sqlDB.SetMaxOpenConns(maxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Duration(timeout))

	r.DbConn = db
	r.IsConnect = true
	RdbConnCollection = r
	return db, nil
}

func (r *CollectionDB) Close() {
	if r.IsConnect == true {
		sqlDB, err := r.DbConn.DB()
		if err != nil {
			err = sqlDB.Close()
			if err != nil {
				log.Printf("sqlDB.Close err %v", err)
			}
		}
	}
}

func InitRdbCollection(logHost *string, logPort *string, logUser *string, logPass *string) (*rdb_config.RdbServerStatistics, error) {
	if rdb_config.IsUseRdbLogDatabase() {
		RdbStatisticsConfig = rdb_config.InitSetConfigStatistics(logHost, logPort, logUser, logPass)
		IsRdbStatisticsInit = true
		log.Printf("init store_rdb Statistics env : %v", RdbStatisticsConfig)
	} else {
		return RdbStatisticsConfig, errors.New("collection database not use")
	}
	return RdbStatisticsConfig, nil
}
