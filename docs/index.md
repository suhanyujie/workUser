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

