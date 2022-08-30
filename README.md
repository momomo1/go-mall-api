# go-mall

## 项目介绍
项目将采用当前流行的技术来实现一套完整的规范

## 编写规范
+ 使用 **RESTFul** 的方式定义API

## 目录组织
```

```

## 热加载
```shell
安装 air
使用以下命令来安装 air ：
$ GO111MODULE=on  go install github.com/cosmtrek/air@lates
Windows 下也可以手动安装，进入 github.com/cosmtrek/air/releases 下载后放入 Go 安装目录下的 bin 目录，重命名为 air.exe。
```

## 代码提交规范
```
<type>(<scope>): <subject>
// 注意冒号 : 后有空格
// 如 feat(商品): 新增商品添加功能
```

**scope** 选填表示commit的作用范围，比如数据层、控制层、视图层等等，视项目不同而不同

**subject** 必填用于对commit进行简短的描述

**type** 必填表示提交类型，值有以下几种：
+ feat  - 新功能 feature
+ fix   - 修复 bug
+ test  - 增加测试
+ docs  - 文档注释
+ style - 代码格式(不影响代码运行的变动)
+ refactor - 重构、优化(既不增加新功能，也不是修复bug)

## 第三方依赖
使用到的开源库：
- [gin](https://github.com/gin-gonic/gin) —— 路由、路由组、中间件
- [zap](https://github.com/gin-contrib/zap) —— 高性能日志方案
- [gorm](https://github.com/go-gorm/gorm) —— ORM 数据操作
- [cobra](https://github.com/spf13/cobra) —— 命令行结构
- [viper](https://github.com/spf13/viper) —— 配置信息
- [cast](https://github.com/spf13/cast) —— 类型转换
- [redis](https://github.com/go-redis/redis/v8) —— Redis 操作
- [jwt](https://github.com/golang-jwt/jwt) —— JWT 操作
- [base64Captcha](https://github.com/mojocn/base64Captcha) —— 图片验证码
- [govalidator](https://github.com/thedevsaddam/govalidator) —— 请求验证器
- [limiter](https://github.com/ulule/limiter/v3) —— 限流器
- [email](https://github.com/jordan-wright/email) —— SMTP 邮件发送
- [aliyun-communicate](https://github.com/KenmyZhang/aliyun-communicate) —— 发送阿里云短信
- [ansi](https://github.com/mgutz/ansi) —— 终端高亮输出
- [strcase](https://github.com/iancoleman/strcase) —— 字符串大小写操作
- [pluralize](https://github.com/gertd/go-pluralize) —— 英文字符单数复数处理


## 自定义的包

自建的库：

- app —— 应用对象
- auth —— 用户授权
- cache —— 缓存
- captcha —— 图片验证码
- config —— 配置信息
- console —— 终端
- database —— 数据库操作
- file —— 文件处理
- hash —— 哈希
- helpers —— 辅助方法
- jwt —— JWT 认证
- limiter —— API 限流
- logger —— 日志记录
- mail —— 邮件发送
- mq —— rabbitmq
- migrate —— 数据库迁移
- paginator —— 分页器
- redis —— Redis 数据库操作
- response —— 响应处理
- seed —— 数据填充
- sms —— 发送短信
- str —— 字符串处理
- verifycode —— 数字验证码

## 所有命令

```
$ go run main.go -h
Default will run "serve" command, you can use "-h" flag to see all subcommands
Usage:
   [command]
Available Commands:
  cache       Cache management
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  key         Generate App Key, will print the generated Key
  make        Generate file and code
  play        Likes the Go Playground, but running at our application context
  seed        Insert fake data to the database
  serve       Start web server

Flags:
  -e, --env string   load .env file, example: --env=testing will use .env.testing file
  -h, --help         help for this command
Use " [command] --help" for more information about a command.
```

make 命令：

```
$ go run main.go make -h
Generate file and code
Usage:
   make [command]
Available Commands:
  apicontroller Create api controller，exmaple: make apicontroller v1/user
  cmd           Create a command, should be snake_case, exmaple: make cmd buckup_database
  factory       Create model's factory file, example: make factory user
  model         Create model file, example: make model user
  request       Create request file, example make request user
  seeder        Create seeder file, example:  make seeder user

Flags:
  -h, --help   help for make
Global Flags:
  -e, --env string   load .env file, example: --env=testing will use .env.testing file
Use " make [command] --help" for more information about a command.
```
