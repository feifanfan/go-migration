## go-migration

---
一款用于golang项目的数据库迁移工具，仿照laravel迁移工具，使用简单

###用法
* 1.配置数据库连接信息
```
config := make(map[string]string)
config["db_username"] = "root"
config["db_password"] = ""
config["db_host"] = "127.0.0.1"
config["db_port"] = "3306"
config["db_name"] = "migration"
config["db_charset"] = "utf8"
```
* 2 获取连接数据库connector
```
    connector := NewConnector(config)
```
* 3 获取操作数据库（数据表）schema
```
    schema := connector.Schema()
```
* 4 写数据库迁移函数
```
func CreateXxxTable(schema ISchemaBuilder)  {
	schema.CreateTable("Xxx", func(table IBlueprint) {
		table.Increments("id",11)
		table.String("username",255).Comment("用户名")
		table.Date("created_at").Default("2020-01-01")
		table.Float("money",10,2)
		table.Boolean("isBig")
	})
}
```

### 其它用法
* 修改表

```
connector.Schema().Table("test_num", func(table IBlueprint) {
        table.String("username").Comment("测试备注").Change()
	table.String("created_time").RenameColumn("created_at")
	table.String("username").Drop()
})
```

### 后续功能
* 删除数据表等表级操作
* 数据填充
