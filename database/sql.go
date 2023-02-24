package database

import (
	"database/sql"
	"fmt"
	"net/url"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/occmundial/consumer-abe-atreel-user-message/libs/logger"
	"github.com/occmundial/consumer-abe-atreel-user-message/models"
)

var (
	dataSourceName string
)

type SQLServerConnection struct {
	Connection *sql.DB
}

// NewSQLServer constructor used for the dependency injection container
func NewSQLServer(config *models.Configuration) *SQLServerConnection {
	return &SQLServerConnection{
		Connection: createConnection(config),
	}
}

func createConnection(config *models.Configuration) *sql.DB {
	dataSourceName = fmt.Sprintf(
		"sqlserver://%s:%s@%s?database=%s&connection+timeout=%d&dial+timeout=%d",
		config.DBUser,
		url.QueryEscape(config.DBPassword),
		config.DBHost,
		config.DBName,
		config.DBTimeout,
		config.DBTimeout,
	)

	connection, err := sql.Open("sqlserver", dataSourceName)
	if err != nil {
		logger.Fatal("repositories", "createConnection", err)
	}
	connection.SetMaxOpenConns(config.DBMaxOpenConns)
	connection.SetMaxIdleConns(config.DBMaxIdleConns)
	connection.SetConnMaxLifetime(time.Minute * time.Duration(config.DBConnMaxLifetime))
	return connection
}
