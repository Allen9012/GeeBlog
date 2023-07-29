# GeeBlog
GeeBlog是由Gin框架和Vue前端写的一个博客项目

后期可能会考虑更新Gee框架和结合替换Gin框架
GeeBlog使用的技术栈是：Gin+Gorm+ES+Websocket
GeeBlog包含下面的功能
1. 在线群聊
2. 用户一对一单独聊天
3. 用户管理后台系统
4. 游客账号
5. 日志系统
6. 配置方式更新
7. 网站数据的统计
8. 基本博客功能的实现

使用go-api-practice和结合gee-init的脚手架模式进行一些脚手架的改造和格式的修改

## 配置方式
配置默认使用环境变量的方式和配置文件的方式来实现读取配置

提供loadEnv和loadConfig