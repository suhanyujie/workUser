# workUser
任务管理-用户服务

主要提供用户的管理，包含用户的新增、信息修改、删除、三方用户接入等接口等

## 功能 list
* [ ] 用户新增
* [ ] 用户修改
* [ ] 用户删除
* [ ] 三方用户新增/同步
* [ ] 三方用户修改
* [ ] 三方用户删除

## 设计
### 表结构
> 用户主表

* id
* username
* nick_name
* status
* pwd
* token
* ext
* third_flag
* create_time
* update_time
* is_delete

>用户信息表

>部门表
>用户部门表
>组织表
* id, org_name

>用户组织关联表
