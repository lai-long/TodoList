# TodoList
*<h2>一.实现内容* 
<br><br>

1、使用token实现注册登录

2、增，删，改，查

* 增加 一条 新的待办事项
* 将 一条/所有 待办事项设为已完成
* 将 一条/所有 已完成事项设为待办
* 查看所有 已完成/未完成/所有 事项
* 输入关键词查询事项
* 删除 一条/所有已完成/所有待办/所有事项
* 使用apifox生成接口文档

*<h2>二.项目架构*

1、  技术栈
    
    语言：golang

    web框架：gin框架

    ORM框架： GORM

    数据库：mysql

    

2、目录
       
    
    Todolist/
    ├── config/          
    ├── doc/              # 接口文档
    ├── internal/
    │   ├── controller/   
    │   ├── middleware/   # 中间件
    │   ├── service/      # 业务层
    │   ├── dao/          # 数据库操作
    │   ├── model/
    │   │   ├── dto/      # 数据传输对象
    │   │   └── entity/   # 数据库实体
    │   
    ├── pkg/
    │   └── database/     # 数据库连接
    ├── router/           # 路由设置
    ├── main.go
    └── go.mod
