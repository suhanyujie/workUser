# 工作管理-用户服务
## 目录结构
* 根据 [go-zero](https://go-zero.dev/cn/) 的文档中描述的最佳实践，服务的目录结构如下所示：

```
mall // 工程名称
├── common // 通用库
│   ├── randx
│   └── stringx
├── go.mod
├── go.sum
└── service // 服务存放目录
    ├── afterSale
    │   ├── cmd
    │   │   ├── api
    │   │   └── rpc
    │   └── model
    ├── cart
    │   ├── cmd
    │   │   ├── api
    │   │   └── rpc
    │   └── model
    ├── order
    │   ├── cmd
    │   │   ├── api
    │   │   └── rpc
    │   └── model
```

model 目录位于 cmd 的**同级**目录。

## 开发三部曲
### 创建网关
创建 xxx.api 文件，并运行 `goctl api go -api xxx.api -dir .`。

需要注意的是，如果文件已经存在，则不会进行覆盖操作，而是忽略。
所以一旦你的 api 文件有更新，则需要删除原先已经生成的接口 handler 和 logic 文件。


### 创建服务的 rpc
### 编写服务逻辑
### 启动服务并测试