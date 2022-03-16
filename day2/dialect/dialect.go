package dialect

import "reflect"

// 封装该ORM框架对不同 数据库的抽象

type Dialect interface {
	DataTypeOf (value reflect.Value) string // 获取数据类型在数据库对应的数据类型
	TableExistSql(table string) (string, []interface{}) // 获取表是否存在的sql
}
// 存储多个数据库实例的 差异部分
var dialectsMap = map[string]Dialect{}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}

func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}




