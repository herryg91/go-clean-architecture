package mysql

import (
	"fmt"
	"time"

	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type mysql struct {
	DBHost                        string
	DBPort                        string
	DBUserName                    string
	DBPass                        string
	DBDatabaseName                string
	DBLogMode                     logger.LogLevel
	maxIdleConnection             int
	maxOpenConnection             int
	connectionMaxLifetimeInSecond int
}
type mysqlOption func(*mysql)

func Connect(DBHost string, DBPort string, DBUserName string, DBPass string, DBDatabaseName string, DBLogMode logger.LogLevel, options ...mysqlOption) (*gorm.DB, error) {
	db := &mysql{
		DBHost:                        DBHost,
		DBPort:                        DBPort,
		DBUserName:                    DBUserName,
		DBPass:                        DBPass,
		DBDatabaseName:                DBDatabaseName,
		DBLogMode:                     DBLogMode,
		maxIdleConnection:             5,
		maxOpenConnection:             10,
		connectionMaxLifetimeInSecond: 60,
	}

	for _, o := range options {
		o(db)
	}

	return connect(db)
}

func connect(param *mysql) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?loc=Local&parseTime=true",
		param.DBUserName, param.DBPass, param.DBHost, param.DBPort, param.DBDatabaseName)

	db, err := gorm.Open(gormMysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(param.DBLogMode)),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	// set configuration pooling connection
	mysqlDb, _ := db.DB()
	mysqlDb.SetMaxOpenConns(param.maxOpenConnection)
	mysqlDb.SetConnMaxLifetime(time.Duration(param.connectionMaxLifetimeInSecond) * time.Minute)
	mysqlDb.SetMaxIdleConns(param.maxIdleConnection)

	return db, nil
}

func SetMaxIdleConns(conns int) mysqlOption {
	return func(c *mysql) {
		if conns > 0 {
			c.maxIdleConnection = conns
		}
	}
}

func SetMaxOpenConns(conns int) mysqlOption {
	return func(c *mysql) {
		if conns > 0 {
			c.maxOpenConnection = conns
		}
	}
}

func SetConnMaxLifetime(conns int) mysqlOption {
	return func(c *mysql) {
		if conns > 0 {
			c.connectionMaxLifetimeInSecond = conns
		}
	}
}
