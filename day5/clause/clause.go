package clause

import (
	"strings"
)

// 拼凑SQL的模块

type Clause struct {
	sql     map[Type]string
	sqlVars map[Type][]interface{}
}

// Set 根据传入的每一个 sql的阶段 编写出 具体的sql 语句
func (c *Clause) Set(name Type, values ...interface{}) {
	if c.sql == nil {
		c.sql = make(map[Type]string)
		c.sqlVars = make(map[Type][]interface{})
	}

	// 根据 每个函数 调用函数方法 得到最终sql 和参数
	sql, vars := generators[name](values...)
	c.sql[name] = sql
	c.sqlVars[name] = vars
}

// Build 根据传入的sql阶段顺序，组装最终的sql
func (c *Clause) Build(orders ...Type) (string, []interface{}) {

	var sqls []string
	var vars []interface{}
	for _, order := range orders {
		if _, ok := c.sql[order]; ok {
			sqls = append(sqls, c.sql[order])
			vars = append(vars, c.sqlVars[order]...)
		}
	}
	return strings.Join(sqls, " "), vars
}
