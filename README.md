# workUser
任务管理-用户服务

主要提供用户的管理，包含用户的新增、信息修改、删除、三方用户接入等接口等

## 启动
```shell
go run service/workUser/cmd/api/workuser.go -f service/workUser/cmd/api/etc/workuser-api.yaml
```


## 功能 list
* [x] 用户新增
* [ ] 用户修改
* [ ] 用户删除
* [ ] 用户查询

## 设计


### 接口

```
service workUser-api {
	@handler CreateUser
	post /user/create(CreateUserReq) returns (CreateUserResp);
	@handler QueryUser
	post /user/list(Request) returns (Response);
	@handler FindUser
	get /user/find(Request) returns (Response);
	@handler UpdateUser
	post /user/update(Request) returns (Response);
	@handler DeleteUser
	delete /user/delete(Request) returns (Response);
}
```

#### 新增用户
#### 修改用户信息
#### 删除用户

## 遇到的问题
* 使用路由不确定如何获取参数，如 `/user/find/:userId`，如何获取 userId 呢？
* 没有 orm 的介绍，不知道如何 migration

## 参考资料
* go-zero 文档 https://legacy.go-zero.dev/cn/
* 文档 https://go-zero.dev/cn/
* https://github.com/tal-tech/zero-doc
* https://github.com/tal-tech/zero-doc/blob/main/doc/shorturl.md
