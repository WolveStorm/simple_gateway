# simple_gateway

#### 介绍
该项目有三个目录
- gateway_server
- manager_desktop
- real_server

gateway_server：网关代理服务核心服务器

manager_desktop：网关管理后台接口

real_server：服务测试样例

#### 软件架构

## 后台管理前端界面

采用github上现成的vue-element-admin。

```python
https://github.com/PanJiaChen/vue-element-admin
```

将路由和api请求路径参数改一下，修改成网关后台的管理页面。

## 后台管理后端接口

后端主要技术：

- Golang
- Consul
- Redis
- Mysql
- Gin

### 为什么选择golang？

- Golang的语法简单明了，易于学习和使用。
- Golang具有很好的性能，可以通过静态编译和内存分配优化来提升效率。
- Golang提供了简单易用的并发模型，可以帮助开发人员更方便地编写高性能的并发程序。
- Golang有庞大的社区支持，可以为开发人员提供帮助和支持。
- Golang拥有比较完善的库，适合开发中间件。

#### 使用了golang的哪些开源库？

[jwt-go](https://github.com/dgrijalva/jwt-go): 用于网关鉴权层的token生成与校验。

[go-redis](https://github.com/go-redis/redis)：一个golang集成的redis客户端用来与redis的交互。

[gin](https://github.com/gin-gonic/gin)：一个封装了golang网络库的操作的web框架。

[sessions](https://github.com/gorilla/sessions)：http是无状态的，用于存储管理员的信息到session。

[zap](https://github.com/uber-go/zap)：效率很高的日志库，用于生成网关的运行日志。

[lumberjack](https://github.com/natefinch/lumberjack)：用于做日志的切分，过期清除功能。

[consul](https://github.com/hashicorp/consul)：一个golang集成的zookeeper客户端用来与zookeeper的交互。

[viper](https://github.com/spf13/viper)：一个简约的配置文件读取库。(后期可以考虑采用nacos来统一配置)。

[gorm](https://github.com/go-gorm/gorm)：拥有不错的文档的ORM库，学习简单，用于与mysql交互。

[validator](https://github.com/go-playground/validator)：用于做参数校验。

### 为什么选择Consul

我们主要**用到其配置存储和分布式同步的功能实现动态获取下游ip列表**。前者可以理解成具有一致性的KV存储，后者提供了Consul特有的watcher注册于异步通知机制，Consul能将节点的状态实时异步通知给ZooKeeper客户端。

*Consul*是一个分布式的，开放源码的分布式应用程序协调服务

他有以下功能：

（1）作为配置信息的存储的中心服务器
（2）命名服务
（3）分布式同步
（4）分组服务

Consul不仅实现了对cusumer和provider的灵活管理，平滑过渡功能，而且还内置了负载均衡、主动通知等功能，使我们能够几乎实时的感应到服务的状态。

选择Consul主要是因为他的实时性。

### 为什么选择Redis

项目中采用了Redis的原子自增IncryBy命令，主要使用了其中的kv数据结构，可以通过key唯一标识某一服务在某一特定时段的信息，通过val记录流量统计的数值。

由于Redis的性能极高，在写入的时候不会占用过多的时间，并且后台从Redis拉取数据的效率也极高。

Remote DIctionary Server(Redis) 是一个由 Salvatore Sanfilippo 写的 key-value 存储系统，是跨平台的非关系型数据库。

Redis的特点：

- **内存数据库，速度快，也支持数据的持久化**，可以将内存中的数据保存在磁盘中，重启的时候可以再次加载进行使用。
- Redis不仅仅支持简单的key-value类型的数据，同时还提供list，set，zset，hash等数据结构的存储。
- Reids**性能极高** – Redis能读的速度是110000次/s,写的速度是81000次/s 。

选择Redis主要是由于他的性能，以及强大的功能。

### 为什么选择Mysql

项目中主要通过mysql来存储管理员的一些信息，和网关http/tcp/grpc服务的具体配置持久化到硬盘。

Mysql的优势：

- 运行速度快，MySQL体积小，命令执行的速度快。
- 使用容易。与其他大型数据库的设置和管理相比，其复杂程度较低，易于使用。
- 可移植性强。MySQL**能够运行与多种系统平台上**，如windouws，Linux，Unix等。
- 支持事务，索引，锁等高级特性。

选择mysql主要是因为mysql是老牌的数据库，经过时间的检验，[Mysql](https://www.lsjlt.com/tag/Mysql/)性能卓越，服务稳定，很少出现异常宕机，从学习成本来看，比较容易学习sql语句。

### 为什么选择Gin

项目主要采用gin做业务层和前端的Http通信。

Gin是一个golang的微框架，封装比较优雅，API友好, 源代码比较明确。 具有快速灵活，容错方便等特点。其实对于golang而言，web框架的依赖远比Python，Java之类的要小。自身的net/http足够简单，性能也非常不错。框架更像是一个常用函数或者工具的集合。借助框架开发，不仅可以省去很多常用的封装带来的时间，也有助于团队的编码风格和形成规范。

Gin的优势

- 速度： Gin之所以被很多企业和团队使用，第一个原因是因为速度快，性能表现出众
- 中间件: 和iris类似, gin在处理请求时,支持中间件操作, 方便编码处理
- 路由: Gin中可以非常简单的实现路由解析功能，并包含路由组解析功能
- 内置渲染: Gin支持JSON、XML和HTML等多种数据格式的渲染， 并提供了方便的操作API

选择Gin，因为Gin是一个轻量级的 Web框架，可以做到开箱即用，并且性能非常优秀。可以很方便的开发。


#### 安装教程
运行后端代码
首先git clone 本项目

git clone https://gitee.com/sekiro-phm/simple_gateway.git

确保本地环境安装了Go 1.12+版本
go version
go version go1.12.15 darwin/amd64
下载类库依赖
export GO111MODULE=on && export GOPROXY=https://goproxy.cn 
cd gateway_demo
go mod tidy
在相应功能文件夹下，执行 go run main.go 即可。


#### 使用说明

1.  在后台管理界面注册代理服务，目前已上线： http://43.143.169.111:9527/disk#/login?redirect=%2Fdashboard
2.  根据提示配置相关配置。
3.  访问页面提示的路径。

