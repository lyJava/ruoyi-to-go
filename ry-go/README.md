# ry-go-echo

#### 介绍
{**框架相关**
1. echo：web端框架，生态中高效构建API的首选框架之一，适合追求性能与简洁性的开发需求。中间件支持前置与后置处理，提供了更精细的控制，原生集成了数据验证，原生支持websocket。
2. gorm：全功能 ORM（对象关系映射）框架，其核心优势在于简化数据库操作、提升开发效率，同时兼顾灵活性和扩展性。支持结构体变动自动同步到数据库。支持链式查询，关联关系自动化处理。默认事务及事务钩子机制。支持原生SQL查询，支持子查询及复用查询逻辑。支持批量与连接池。支持分页软删除。使用简单灵活，需要效率的首选。
3. postgresql：支持 CTE（公用表表达式）、窗口函数、多维数组、JSON 聚合，适合复杂分析。原生支持 JSONB（二进制 JSON），支持索引和高效查询。通过 PostGIS 扩展提供完整的 GIS 功能，支持空间索引(地图服务、物流路径优化、地理围栏。)MVCC（多版本并发控制）减少锁竞争，支持行级锁和乐观锁(电商秒杀、票务系统等高并发 OLTP)。支持全文搜索、词干提取、模糊搜索、权重控制。

#### 软件架构
软件架构说明
Gorm+Echo+Postgresql

#### 功能说明

1. 用户管理功能
2. 角色管理功能
3. 菜单管理功能
4. 部门管理功能
5. 岗位管理功能
6. 字典管理功能
7. 参数设置功能
8. 通知公告管理
9. 日志管理(登录日志与操作日志)

#### 使用说明

1. 部署好golang运行环境，具体操作请仔细参考网上答案
2. 首先初始化数据库，确保能正常连接
3. 安装好redis并配置好链接
4. 在配置文件中修改echo运行的端口号
5. 使用goland或者vscode中启动项目即可

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### 特技

1.  使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2.  Gitee 官方博客 [blog.gitee.com](https://blog.gitee.com)
3.  你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解 Gitee 上的优秀开源项目
4.  [GVP](https://gitee.com/gvp) 全称是 Gitee 最有价值开源项目，是综合评定出的优秀开源项目
5.  Gitee 官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6.  Gitee 封面人物是一档用来展示 Gitee 会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)
