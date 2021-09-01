package db

type IBlueprint interface {
	//数据库列的几种类型
	Boolean(columnName string) IBlueprint
	Char(columnName string,length int) IBlueprint
	Date(columnName string) IBlueprint
	DateTime(columnName string) IBlueprint
	Decimal(columnName string) IBlueprint
	Float(columnName string,length ...int) IBlueprint
	Increments(columnName string,length int) IBlueprint //自增主键
	Integer(columnName string,length int) IBlueprint
	LongText(columnName string,length int) IBlueprint
	Text(columnName string,length int) IBlueprint
	String(columnName string, length ...int) IBlueprint
	TinyInt(columnName string,length int) IBlueprint

	Default(def interface{}) IBlueprint
	Comment(com string) IBlueprint
	Nullable() IBlueprint

	Drop() IBlueprint                            // 删除列
	Change() IBlueprint                          // 修改列
	RenameColumn(originColumn string) IBlueprint // 重命名列
}

type Blueprint struct {

}
