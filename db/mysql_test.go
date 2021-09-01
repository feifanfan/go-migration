package db

import (
	"testing"
)

func TestSchemaBuilder_GetConn(t *testing.T) {
	// 输入数据
	config := make(map[string]string)
	config["db_username"] = "root"
	config["db_password"] = "rootroot"
	config["db_host"] = "127.0.0.1"
	config["db_port"] = "3306"
	config["db_name"] = "migrationtest"
	config["db_charset"] = "utf8"

	connector := NewConnector(config)
	////创建表
	//connector.Schema().CreateTable("test", func(table IBlueprint) {
	//	table.String("username",255)
	//})
	//创建
	connector.Schema().CreateTable("test_num", func(table IBlueprint) {
		table.String("username",255).Comment("用户名")
		table.Increments("id",11)
		table.Date("created_at").Default("2020-01-01")
		table.Float("money",10,2)
		table.Boolean("isBig")
	})
	//修改表字段
	connector.Schema().Table("test_num", func(table IBlueprint) {
		//table.String("username").Comment("测试备注").Change()
		table.String("created_time").RenameColumn("created_at")
	})

}
