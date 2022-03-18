# GOORM 项目

## day4
目标
1. ORM框架新增删除和更新函数
2. 新增链式操作 

数据介绍
- clause
 
   构建SQL语句的基础工具类，构建sql的各个模块，新增delete update orderby count等
- generator
  
   构建整条SQL， 内部通过Map保存 sql各个部位， 使用Build方法 将所有sql 拼接完成,新增delete update orderby count等
- record

  针对记录的增删改查类， 新增Delete update count函数，且新增链式调用函数




使用案例
```go
参考 test_sql_op.go
type User struct{
	Name string
	Age int
}
var u User{}
session.Where("Age=", 12).Find(&u)
测试案例见test_sql_op.go test_update_delete_and_chain()
```