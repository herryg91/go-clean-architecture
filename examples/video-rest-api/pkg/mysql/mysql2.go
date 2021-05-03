package mysql

// // We're need this because we're using docker, in actual situation we're gonna use shared-lib

// import (
// 	"fmt"
// 	"time"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// 	"gorm.io/gorm/schema"
// )

// type mysqlConfig struct {
// 	// SSLMode to enable/disable SSL connection
// 	SSLMode bool `envconfig:"MYSQL_SSL_MODE" default:"true"`
// 	// MaxIdleConnection to set max idle connection pooling
// 	MaxIdleConnection int `envconfig:"MYSQL_MAX_IDLE_CONNECTION" default:"5"`
// 	// MaxOpenConnection to set max open connection pooling
// 	MaxOpenConnection int `envconfig:"MYSQL_MAX_OPEN_CONNECTION" default:"10"`
// 	// MaxLifetimeConnectionn to set max lifetime of pooling | minutes unit
// 	MaxLifetimeConnection int `envconfig:"MYSQL_MAX_LIFETIME_CONNECTION" default:"10"`
// 	// Host is host of mysql service
// 	Host string `envconfig:"MYSQL_HOST" default:""`
// 	// Port is port of mysql service
// 	Port string `envconfig:"MYSQL_PORT" default:""`
// 	// Username is name of registered user in mysql service
// 	Username string `envconfig:"MYSQL_USERNAME" default:""`
// 	// DBName is name of registered database in mysql service
// 	DBName string `envconfig:"MYSQL_DB_NAME" default:""`
// 	// Password is password of used Username in mysql service
// 	Password string `envconfig:"MYSQL_PASSWORD" default:""`
// 	// LogMode is toggle to enable/disable log query in your service by default false
// 	LogMode logger.LogLevel `envconfig:"MYSQL_LOG_MODE" default:"0"`
// 	// SingularTable to activate singular table if you are using eloquent query
// 	SingularTable bool `envconfig:"MYSQL_SINGULAR_TABLE" default:"true"`
// 	// ParseTime to parse to local time
// 	ParseTime bool `envconfig:"MYSQL_PARSE_TIME" default:"true"`
// 	// Charset to define charset of database
// 	Charset string `envconfig:"MYSQL_CHARSET" default:"utf8mb4"`
// 	// Charset to define charset of database
// 	Loc string `envconfig:"MYSQL_LOC" default:"Local"`
// }

// func NewMysqlCli(host, port, username, password, dbname string, logmode logger.LogLevel) mysqlConfig {
// 	return mysqlConfig{
// 		SSLMode:               false,
// 		MaxIdleConnection:     5,
// 		MaxOpenConnection:     10,
// 		MaxLifetimeConnection: 10,
// 		Host:                  host,
// 		Port:                  port,
// 		Username:              username,
// 		DBName:                dbname,
// 		Password:              password,
// 		LogMode:               logmode,
// 		SingularTable:         true,
// 		ParseTime:             true,
// 		Charset:               "utf8mb4",
// 		Loc:                   "Local",
// 	}
// }

// func (cfg mysqlConfig) Connect() *gorm.DB {
// 	// construct connection string
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%+v&loc=%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.Charset, cfg.ParseTime, cfg.Loc)
// 	// open mysql connecion
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.LogLevel(cfg.LogMode)),
// 		NamingStrategy: schema.NamingStrategy{
// 			SingularTable: true,
// 		},
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	// set configuration pooling connection
// 	mysqlDb, _ := db.DB()
// 	mysqlDb.SetMaxOpenConns(cfg.MaxOpenConnection)
// 	mysqlDb.SetConnMaxLifetime(time.Duration(cfg.MaxLifetimeConnection) * time.Minute)
// 	mysqlDb.SetMaxIdleConns(cfg.MaxIdleConnection)

// 	return db

// }
