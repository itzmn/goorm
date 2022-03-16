package dialect

import (
	"fmt"
	"reflect"
	"time"
)

// ORM框架对 sqlite3的支持

type sqlite3 struct{}

var _ Dialect = (*sqlite3)(nil)

func init() {
	RegisterDialect("sqlite3", &sqlite3{})
}

//=============针对不同点进行适配

func (s *sqlite3) DataTypeOf(value reflect.Value) string {

	switch value.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return "integer"
	case reflect.Int64, reflect.Uint64:
		return "bigint"
	case reflect.Float32, reflect.Float64:
		return "real"
	case reflect.String:
		return "text"
	case reflect.Array, reflect.Slice:
		return "blob"
	case reflect.Struct:
		if _, ok := value.Interface().(time.Time); ok {
			return "datetime"
		}
	}
	panic(fmt.Sprintf("invalid sql type %s (%s)", value.Type().Name(), value.Kind()))

}

func (s *sqlite3) TableExistSql(table string) (string, []interface{}) {
	args := []interface{}{table}
	return "SELECT name FROM sqlite_master WHERE type='table' and name = ?", args
}
