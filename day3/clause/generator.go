package clause

import (
	"fmt"
	"strings"
)

// 用于拼凑SQL语句的 模块

type generator func(values ...interface{}) (string, []interface{})

type Type int

const (
	INSERT Type = iota
	VALUES
	SELECT
	LIMIT
	WHERE
	ORDERBY
)

var generators map[Type]generator

// 根据参数个数 构建 ?, ?, ?
func genBindVars(num int) string {
	var vars []string
	for i := 0; i < num; i++ {
		vars = append(vars, "?")
	}
	return strings.Join(vars, ",")
}

func _insert(values ...interface{}) (string, []interface{}) {
	// insert into $table ($fields)
	tableName := values[0]
	sql := strings.Join(values[1].([]string), ",")
	insertSql := fmt.Sprintf("insert into %s (%v)", tableName, sql)
	return insertSql, []interface{}{}
}

func _values(values ...interface{}) (string, []interface{}) {
	// values ($1), ($2)
	var bindStr string
	var sql strings.Builder
	var vars []interface{}
	sql.WriteString("values ")
	for i, value := range values {
		v := value.([]interface{})
		if bindStr == "" {
			bindStr = genBindVars(len(v))
		}
		sql.WriteString(fmt.Sprintf("(%v)", bindStr))
		if i+1 != len(values) {
			sql.WriteString(", ")
		}
		// 保存 参数绑定的值
		vars = append(vars, v...)
	}
	return sql.String(), vars
}

func _select(values ...interface{}) (string, []interface{}) {
	// select $fields from $table
	tableName := values[0]
	fields := strings.Join(values[1].([]string), ",")
	selectSql := fmt.Sprintf("select %v from %s", fields, tableName)
	return selectSql, []interface{}{}
}

func _limit(values ...interface{}) (string, []interface{}) {
	// limit $n
	return "limit ?", values
}

func _where(values ...interface{}) (string, []interface{}) {
	// where $desc
	desc, args := values[0], values[1:]
	return fmt.Sprintf("where %s", desc), args
}

func _orderBy(values ...interface{}) (string, []interface{}) {
	// order by xxx
	return fmt.Sprintf("order by %s", values[0]), []interface{}{}
}

func init() {
	generators = make(map[Type]generator)
	generators[INSERT] = _insert
	generators[VALUES] = _values
	generators[SELECT] = _select
	generators[LIMIT] = _limit
	generators[WHERE] = _where
	generators[ORDERBY] = _orderBy
}
