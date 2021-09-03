## 商城项目微服务

### 使用微服务框架
github.com/asim/go-micro/v3


### 快速开始

- [项目结构](#项目结构)
- [使用](#使用)

#### 项目结构

``` lua
newmicro-mall
├── cart -- 购物车服务
├── cart-api -- 前端调用购物车接口
├── category -- 商品分类模块服务
├── common -- 公共方法
├── order -- 订单模块服务
├── product -- 商品模块服务
└── user -- 用户模块服务
```


#### 使用
切换到各目录，根据 proto 自动生成
```
make proto
```

编译
```
make build
```

构建镜像
```
make docker
```

本地启动服务
```
make *-server
```

### 项目预览地址

[点击查看微服务商城](http://mall.hjlinfo.top)

### 关于本项目说明

本项目是根据微服务整体架构抽离的部分内容，供学习使用，不包含微服务商城的全部代码，如果想获取商城的全部代码，请给我发邮件:a516972602@gmail.com

