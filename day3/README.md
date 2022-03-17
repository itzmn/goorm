# GOORM 项目

## day3
目标
1. 实现ORM框架最根本的功能，对象和表的转换
2. 实现构建SQL语句的基础工具类，构建sql的各个部位
3. 实现构建SQL语句的上层工具，反馈完整的SQL
4. 实现INSERT功能，可以通过INSERT(struct)方式插入数据库
5. 实现Find功能，可以通过Find([]struct) 方式读取到所有的数据

数据介绍
- clause
 
   构建SQL语句的基础工具类，构建sql的各个模块，select where value insert等
- generator
  
   构建整条SQL， 内部通过Map保存 sql各个部位， 使用Build方法 将所有sql 拼接完成
- schema

  schema中新增函数，将传入的struct实体，获取到实体的字段的值的列表， 用于Insert value中
- record

  针对记录的增删改查类， 新增Find和 Insert函数




使用案例
```go
参考 test_sql_op.go
type User struct{
	Name string
	Age int
}
var us []User
session.Find(&us)
session.Insert(&User{Name: "test", "Age": 12})
```