package mysql

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type mysql struct {
	DBHost                        string
	DBPort                        string
	DBUserName                    string
	DBPass                        string
	DBDatabaseName                string
	DBLogMode                     bool
	maxIdleConnection             int
	maxOpenConnection             int
	connectionMaxLifetimeInSecond int
}
type mysqlOption func(*mysql)

func Connect(DBHost string, DBPort string, DBUserName string, DBPass string, DBDatabaseName string, DBLogMode bool, options ...mysqlOption) (*gorm.DB, error) {
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
	dbDialect := "mysql"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?loc=Local&parseTime=true",
		param.DBUserName, param.DBPass, param.DBHost, param.DBPort, param.DBDatabaseName)

	db, err := gorm.Open(dbDialect, dsn)
	if err != nil {
		return nil, err
	}

	db.LogMode(param.DBLogMode)
	// setup db pool connections
	db.DB().SetMaxIdleConns(param.maxIdleConnection)
	db.DB().SetMaxOpenConns(param.maxOpenConnection)
	db.DB().SetConnMaxLifetime(time.Second * time.Duration(param.connectionMaxLifetimeInSecond))
	return db, err
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
