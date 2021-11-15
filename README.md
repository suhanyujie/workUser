# workUser
任务管理-用户服务

主要提供用户的管理，包含用户的新增、信息修改、删除、三方用户接入等接口等


## 开发三部曲
### 创建网关
创建 xxx.api 文件，并运行 `goctl api go -api xxx.api -dir .`。本项目写的是用户服务，对应的命令就是 `goctl api go -api workUser.api -dir .`

### 创建服务的 rpc
### 编写服务逻辑
### 启动服务并测试

## 启动服务
* `go run workuser.go -f etc/workuser-api.yaml`

## 功能 list
* [*] 用户新增
* [*] 用户修改
* [*] 用户删除
* [ ] 三方用户新增/同步
* [ ] 三方用户修改
* [ ] 三方用户删除

### 接口

```
service workUser-api {
	@handler CreateUser
	post /user/create(CreateUserReq) returns(CreateUserResp)
	@handler FindUser
	get /user/find/:userId () returns(FindUserResp)
	@handler UserList
	post /user/list(UserListReq) returns(UserListRespVo)
	@handler UserDelete
	delete /user/delete/:userId () returns(ResponseVo)
}
```

#### 新增用户

#### 修改用户信息
#### 删除用户

## 遇到的问题
* 使用路由不确定如何获取参数，如 `/user/find/:userId`，如何获取 userId 呢？
* 没有 orm 的介绍，不知道如何 migration

## 参考资料
* https://github.com/tal-tech/zero-doc
* https://github.com/tal-tech/zero-doc/blob/main/doc/shorturl.md
* 文档 https://go-zero.dev/cn/quick-start.html 
* go-zero 实现的中台 https://github.com/jackluo2012/datacenter
* go-zero 实践（多图详解万星 Restful 框架原理与实现） https://mp.weixin.qq.com/s/0cJj_H5kUJjYdz2btBbGWA
