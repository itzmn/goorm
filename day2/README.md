# GOORM 项目

## day1
目标
1. ORM框架 对于不同数据库支持不同的操作，需要将每个数据库的不同 抽象出来 封装成为dialect
2. 实现Sqlite3的 方言
3. 定义schema 用于处理类和表的关系
4. 将表相关的操作封装起来

数据介绍
- dialect
  
  对于不同数据库的 抽象，存储对于不同数据的数据类型转换的方法， 以及整个系统所有的方言
- sqlite3

  sqlite3的方言具体实现，包含方言注册，数据类型GO语言的转换
- schema

  GO类和表的映射关系，包含字段列表，表名，通过反射的方式 通过传入的GO对象解析成schema对象
- table

  通过schema中保存的数据，用于对表的增删改API

调用案例

```go
session := engine.NewSession()
session = session.Model(struct)
session.CreateTable
sesssion.Raw(sql).Exec
```