package db

import "database/sql"

// IConnector 连接器接口
type IConnector interface {
	SetConn(config map[string]string) 	//设置配置信息
	GetConn() *sql.DB                   //获取连接器
	//Table(tableName string) IBuilder    //表句柄，用来增删改查
	Schema() ISchemaBuilder             //数据库句柄，用来操作数据表，删表，增表，该表名
}

// Connector 基础连接器
type Connector struct {
	db *sql.DB
}

func (c *Connector) SetConn(config map[string]string) {

}

func (c *Connector) GetConn() *sql.DB {
	return nil
}

func (c *Connector) Table(tableName string) IBuilder {
	return nil
}

func (c *Connector) Schema() IBuilder {
	return nil
}
