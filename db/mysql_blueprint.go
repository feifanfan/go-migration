package db

import (
	"strconv"
	"strings"
)

const (
	TypeInteger              = "int"
	TypeBigInteger           = "bigint"
	TypeString               = "varchar"
	TypeChar                 = "char"
	TypeBlob                 = "blob"
	TypeBool                 = "boolean"
	TypeDate                 = "date"
	TypeDateTime             = "datetime"
	TypeTimestamp            = "timestamp"
	TypeTime                 = "time"
	TypeDecimal              = "decimal"
	TypeDouble               = "double"
	TypeFloat                = "float"
	TypeGeometry             = "geometry"
	TypeGeometryCollection   = "geometrycollection"
	TypeJson                 = "json"
	TypeLineString           = "linestring"
	TypeLongText             = "longtext"
	TypeMediumInteger        = "mediumint"
	TypeTinyInteger          = "tinyint"
	TypeSmallInteger         = "smallint"
	TypeEnum                 = "enum"
	TypeMediumText           = "mediumtext"
	TypeText                 = "text"
	TypeMultiLineString      = "multilinestring"
	TypeMultiPoint           = "multipoint"
	TypeMultiPolygon         = "multipolygon"
	TypePoint                = "point"
	TypePolygon              = "polygon"
	TypeYear                 = "year"
	TypeStringDefaultLength  = 255
	TypeIntegerDefaultLength = 11
	TypeBooleanDefaultValue  = false

	CreateDefaultType = iota
	CreateIfNotExists
	AlterTable
	CreateTable
	DropTable
	ChangeColumn
	DropColumn
	AddColumn
	RenameColumn
)

// mysql实现 blueprint
type MysqlBlueprint struct {
	Blueprint
	name              string // 表名
	engine            string
	columns           []*MysqlColumn
	currentCol        *MysqlColumn
	charset           string
	indexList         []string   // 普通索引
	combineIndexList  [][]string // 组合索引数组
	uniqueIndexList   []string   // 唯一索引
	fulltextIndexList []string   // 全文索引
	primaryKey        string     // 主键索引
	operator          int        // 操作类型，是创建表还是更改表结构
}

type MysqlColumn struct {
	name             string // 列名
	columnType       string // 列类型
	defaultValue     string // 默认值
	defaultType      bool   // 默认值的类型 , true 是int，false是字符串
	comment          string // 备注
	nullable         bool   // 是否允许为null
	length           int    // 列的长度
	unsigned         bool   // 是否有符号
	autoIncrement    bool   // 是否自增
	originColumnName string // 原列名（只在重命名列时用到）
	operator         int    // 列操作符
}

func NewColumn(mb *MysqlBlueprint, name string) *MysqlColumn {
	c := new(MysqlColumn)
	c.nullable = false
	c.name = name
	c.operator = AddColumn
	c.unsigned = false // 默认都是有符号的
	c.autoIncrement = false
	c.defaultType = false // 默认值是 字符串
	mb.columns = append(mb.columns, c)
	mb.currentCol = c
	return c
}

func (mb *MysqlBlueprint) String(columnName string, length ...int) IBlueprint {
	column := NewColumn(mb, columnName)
	column.length = TypeStringDefaultLength
	if len(length) == 1 {
		column.length = length[0]
	}
	column.columnType = TypeString

	return mb
}

func (mb *MysqlBlueprint) Boolean(columnName string) IBlueprint {
	column := NewColumn(mb, columnName)
	column.columnType = TypeBool
	return mb
}

func (mb *MysqlBlueprint) Char(columnName string, length int) IBlueprint {
	column := NewColumn(mb, columnName)
	column.length = length
	column.columnType = TypeChar
	return mb
}

func (mb *MysqlBlueprint) Date(columnName string) IBlueprint {
	column := NewColumn(mb, columnName)
	column.columnType = TypeDate
	return mb
}

func (mb *MysqlBlueprint) DateTime(columnName string) IBlueprint {
	column := NewColumn(mb, columnName)
	column.columnType = TypeDateTime
	return mb
}

func (mb *MysqlBlueprint) Decimal(columnName string) IBlueprint {
	column := NewColumn(mb, columnName)
	column.columnType = TypeDecimal
	return mb
}

func (mb *MysqlBlueprint) Float(columnName string, length ...int) IBlueprint {
	column := NewColumn(mb, columnName)
	column.columnType = TypeFloat
	switch len(length) {
	case 0:
	case 2:
		column.columnType += `(` + strconv.Itoa(length[0]) + `,` + strconv.Itoa(length[1]) + `)`
	default:
		panic("float need 2 length args")
	}
	return mb
}

func (mb *MysqlBlueprint) Increments(columnName string, length int) IBlueprint {
	c := NewColumn(mb, columnName)
	c.columnType = TypeInteger
	c.length = length
	c.unsigned = true
	c.autoIncrement = true
	mb.PrimaryKey(columnName)
	return mb
}

func (mb *MysqlBlueprint) Integer(columnName string, length int) IBlueprint {
	c := NewColumn(mb, columnName)
	c.length = length
	c.columnType = TypeInteger
	return mb
}

func (mb *MysqlBlueprint) LongText(columnName string, length int) IBlueprint {
	c := NewColumn(mb, columnName)
	c.columnType = TypeLongText
	return mb
}

func (mb *MysqlBlueprint) Text(columnName string, length int) IBlueprint {
	c := NewColumn(mb, columnName)
	c.columnType = TypeText
	return mb
}

func (mb *MysqlBlueprint) TinyInt(columnName string, length int) IBlueprint {
	c := NewColumn(mb, columnName)
	c.columnType = TypeTinyInteger
	c.length = length
	return mb
}

func (mb *MysqlBlueprint) Default(def interface{}) IBlueprint {
	if mb.currentCol.defaultType {
		if v, ok := def.(int); ok {
			mb.currentCol.defaultValue = strconv.Itoa(v)
		} else {
			panic(mb.currentCol.name + " wrong default type(need int)")
		}
	} else {
		if v, ok := def.(string); ok {
			mb.currentCol.defaultValue = v
		} else {
			panic(mb.currentCol.name + " wrong defalut type (need string)")
		}
	}
	return mb
}

func (mb *MysqlBlueprint) Comment(common string) IBlueprint {
	mb.currentCol.comment = common
	return mb
}

func (mb *MysqlBlueprint) Nullable() IBlueprint {
	mb.currentCol.nullable = true
	return mb
}

func (mb *MysqlBlueprint) Drop() IBlueprint {
	mb.currentCol.operator = DropColumn
	return mb
}

func (mb *MysqlBlueprint) Change() IBlueprint {
	mb.currentCol.operator = ChangeColumn
	return mb
}

func (mb *MysqlBlueprint) RenameColumn(originColumn string) IBlueprint {
	mb.currentCol.operator = RenameColumn
	mb.currentCol.originColumnName = originColumn // 原列名
	return mb
}

func (mb *MysqlBlueprint) PrimaryKey(column string) {
	mb.primaryKey = column
}


// Assembly 组装创建表的sql
func Assembly(createType int, schema *MysqlBlueprint) string {
	var sql string
	if createType == CreateDefaultType {
		sql = "CREATE TABLE " + schema.name + "("
	} else if createType == CreateIfNotExists {
		sql = "CREATE TABLE IF NOT EXISTS " + schema.name + "("
	}
	//列sql
	var columnSql []string
	for k := range schema.columns {
		columnSql = append(columnSql, columnAssembly(schema.columns[k]))
	}
	//普通索引
	for k := range schema.indexList {
		columnSql = append(columnSql, indexAssembly(schema.indexList[k]))
	}
	//组合索引
	for k := range schema.combineIndexList {
		columnSql = append(columnSql, combineIndexAssembly(schema.combineIndexList[k]))
	}
	//主键
	if schema.primaryKey != "" {
		columnSql = append(columnSql, primaryKeyAssembly(schema.primaryKey))
	}
	//唯一索引
	for k := range schema.uniqueIndexList {
		columnSql = append(columnSql, uniqueIndexAssembly(schema.uniqueIndexList[k]))
	}
	//全文索引
	for k := range schema.fulltextIndexList {
		columnSql = append(columnSql, fulltextIndexAssembly(schema.fulltextIndexList[k]))
	}
	//拼接sql
	sql += strings.Join(columnSql, ",") + ") ENGINE=" + schema.engine

	//字符
	if schema.charset != "" {
		sql += " DEFAULT CHARSET = " + schema.charset
	}
	return sql
}
func columnAssembly(column *MysqlColumn) string {
	var sql string
	sql = column.name
	sql += ` ` + column.columnType
	//长度
	if column.length != 0 {
		sql += `(` + strconv.Itoa(column.length) + `)`
	}
	//符号
	if column.unsigned {
		sql += ` unsigned `
	}
	//是否允许为空
	if !column.nullable {
		sql += ` not null `
	}
	//是否自动递增
	if column.autoIncrement {
		sql += ` auto_increment `
	}
	//默认值
	if column.defaultValue != "" {
		sql += ` default `
		if column.defaultType {
			sql += column.defaultValue
		} else {
			sql += `"` + column.defaultValue + `"`
		}
	}
	if column.comment != "" {
		sql += ` comment "` + column.comment + `"`
	}
	return sql
}

func primaryKeyAssembly(name string) string {
	return `primary key (` + name + `)`
}
func uniqueIndexAssembly(index string) string {
	return `unique (` + index + `)`
}
func indexAssembly(index string) string {
	return `index (` + index + `)`
}
func combineIndexAssembly(columns []string) string {
	return `index (` + strings.Join(columns, ",") + ")"
}
func fulltextIndexAssembly(index string) string {
	return `fulltext (` + index + ` )`
}

// 修改表结构的方法
func alterAssembly(schema *MysqlBlueprint) []string {
	var sql []string
	for k := range schema.columns {
		switch schema.columns[k].operator {
		case AddColumn:
			sql = append(sql, `alter table `+schema.name+` add `+columnAssembly(schema.columns[k]))
		case ChangeColumn:
			sql = append(sql, `alter table `+schema.name+` modify `+columnAssembly(schema.columns[k]))
		case DropColumn:
			sql = append(sql, `alter table `+schema.name+` drop `+schema.columns[k].name)
		case RenameColumn:
			sql = append(sql, `alter table `+schema.name+` change `+schema.columns[k].originColumnName+` `+columnAssembly(schema.columns[k]))
		}
	}

	return sql
}