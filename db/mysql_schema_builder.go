package db

import (
	"database/sql"
)

type MysqlSchemaBuilder struct {
	SchemaBuilder
	connector IConnector
}

const (
	SchemaDefaultEngine = "innodb"
)

func NewMysqlSchemaBuilder() *MysqlSchemaBuilder {
	return &MysqlSchemaBuilder{}
}

func (msb *MysqlSchemaBuilder) SetConn(conn IConnector) ISchemaBuilder {
	msb.connector = conn
	return msb
}
func (msb *MysqlSchemaBuilder) GetConn() *sql.DB {
	return msb.connector.GetConn()
}
func (msb *MysqlSchemaBuilder) CreateTable(tableName string, call func(table IBlueprint)) error {
	schema := new(MysqlBlueprint)
	schema.engine = SchemaDefaultEngine
	schema.operator = AlterTable
	schema.name = tableName
	call(schema)

	sql := Assembly(CreateDefaultType, schema)
	stmt, err := msb.GetConn().Prepare(sql)
	if err != nil {
		return err
	}
	stmt.Exec()
	return nil
}

func (msb *MysqlSchemaBuilder) CreateTableIfNotExists(tableName string, call func(table IBlueprint)) error {
	schema := new(MysqlBlueprint)
	schema.engine = SchemaDefaultEngine
	schema.name = tableName
	schema.operator = CreateTable
	call(schema)
	sql := Assembly(CreateIfNotExists, schema)
	stmt, err := msb.GetConn().Prepare(sql)
	if err != nil {
		return err
	}
	stmt.Exec()
	return nil
}
func (msb *MysqlSchemaBuilder) Table(tableName string, call func(table IBlueprint)) error {
	schema := new(MysqlBlueprint)
	schema.name = tableName
	schema.operator = AlterTable
	call(schema)

	//拼接更改表结构的sql
	s := alterAssembly(schema)

	for k := range s {
		stmt, err := msb.GetConn().Prepare(s[k])
		if err != nil {
			return err
		}
		_, err = stmt.Exec()
		if err != nil {
			return err
		}
	}
	return nil
}
