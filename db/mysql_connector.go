package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// MysqlConnector mysql连接器
type MysqlConnector struct {
	Connector
	connection   *sql.DB
	username     string
	password     string
	host         string
	port         string
	databaseName string
	charset      string
}

func (m *MysqlConnector) SetConn(config map[string]string) {
	if n, ok := config["db_username"]; ok {
		m.username = n
	} else {
		panic("the mysql connector need username")
	}
	if n, ok := config["db_password"]; ok {
		m.password = n
	} else {
		panic("the mysql connector need password")
	}
	if n, ok := config["db_host"]; ok {
		m.host = n
	} else {
		panic("the mysql connector need host")
	}
	if n, ok := config["db_port"]; ok {
		m.port = n
	} else {
		panic("the mysql connector need port")
	}
	if n, ok := config["db_name"]; ok {
		m.databaseName = n
	} else {
		panic("the mysql connector need db_name")
	}
	if n, ok := config["db_charset"]; ok {
		m.charset = n
	} else {
		panic("the mysql connector need charset")
	}

	db, err := sql.Open("mysql", m.username+":"+m.password+"@tcp("+m.host+":"+m.port+")/"+m.databaseName+"?charset="+m.charset)
	if err != nil {
		panic(err)
	}
	m.connection = db
}

func (m *MysqlConnector) GetConn() *sql.DB {
	return m.connection
}

func (m *MysqlConnector) Table(tableName string) IBuilder {
	return nil
}

func (m *MysqlConnector) Schema() ISchemaBuilder {
	return NewMysqlSchemaBuilder().SetConn(m)
}

func NewConnector(config map[string]string) IConnector {
	c := &MysqlConnector{}
	c.SetConn(config)
	return c
}
