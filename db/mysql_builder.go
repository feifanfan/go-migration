package db

const (
	InnerJoinType  = "inner join"
	LeftJoinType   = "left join"
	RightJoinType  = "right join"
	FullJoinType   = "full join"
	ModelTag       = "json"
	SelectOperator = iota
	UpdateOperator
	DeleteOperator
	InsertOperator
)

type MysqlBuilder struct {
	connector       IConnector
	Builder                             // 塞入一个构建器
	bindings        map[string][]string // 绑定的操作符与列名之间的映射
	columns         []string            // 列名()
	distinct        bool                // 是否用到了去重查询
	distinctColumns []string            // 唯一的列
	from            string              // 表名
	joinType        string              // 连接类型
	joinTable       string              // 连接的表名
	//joinOn          []*joinOnCondition
	//wheres          []*whereCondition // where的数组
	groups          []string   // 组
	havings         []string   // group by 之后的操作
	orders          []string   // 排序
	limit           int        // 限制
	offset          int        // 偏移
	unions          []IBuilder // 联合
	unionLimit      string
	unionOffset     string
	unionOrders     string
	lock            bool
	operator        int // 操作符
}

func (m *MysqlBuilder) SetConn(connector IConnector) IBuilder  {
	m.connector = connector
	return m
}

func NewMysqlBuilder() IBuilder {
	m := &MysqlBuilder{}
	m.operator = SelectOperator
	return m
}


