package db

import (
	"database/sql"
	"errors"
)

// IBuilder 数据库迁移，暂不实现该里面的一些方法
type IBuilder interface {
	SetConn(connector IConnector) IBuilder
}

// ISchemaBuilder 数据库/数据表 构建器
type ISchemaBuilder interface {
	SetConn(conn IConnector) ISchemaBuilder
	GetConn() *sql.DB
	// CreateTable 创建数据表
	CreateTable(tableName string,call func(table IBlueprint))  error
	CreateTableIfNotExists(tableName string, call func(table IBlueprint)) error
	//更改表结构
	Table(tableName string, call func(table IBlueprint)) error
}

type SchemaBuilder struct {
	connector IConnector
}

func (builder *SchemaBuilder) SetConn(conn Connector) ISchemaBuilder {
	return nil
}
func (builder *SchemaBuilder) GetConn() *sql.DB {
	return nil
}
func (builder *SchemaBuilder) CreateTable(tableName string, call func(table IBlueprint)) (sql.Result, error) {
	return nil, errors.New("")
}

func (builder *SchemaBuilder) CreateTableIfNotExists(tableName string, call func(table IBlueprint)) error {
	return nil
}

type Builder struct {
}