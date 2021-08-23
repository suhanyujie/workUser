# MySQL 表设计
* 数据库采用 `work_` 作为前缀，表示工作、协作应用的意思。
* 数据库的字符集 `utf8mb4`，对应的排序规则使用 `utf8mb4_bin`。

### 数据库
用户服务的数据库的库名是 `work_user`。

```sql
CREATE DATABASE `work_user` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_bin';
```

### 表结构
> 用户主表

* id
* user_name
* nick_name
* status
* pwd
* token
* ext
* avatar
* sex
* created_at
* updated_at
* deleted_at

```sql
CREATE TABLE `work_user`.`work_user`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_name` varchar(120) NOT NULL COMMENT '用户名，全局唯一',
  `nick_name` varchar(200) NOT NULL COMMENT '用户昵称',
  `status` smallint NOT NULL DEFAULT 1 COMMENT '用户状态',
  `pwd` varchar(128) NOT NULL DEFAULT '' COMMENT '加密后的密码',
  `token` varchar(50) NOT NULL DEFAULT '' COMMENT '加密的 token',
  `avatar` varchar(200) NOT NULL DEFAULT '' COMMENT '头像地址',
  `sex` varchar(10) NOT NULL DEFAULT '' COMMENT '性别',
  `created_at` datetime not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime default null COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `user_name_idx`(`user_name`) USING BTREE
);
```

## 使用 gorm
由于 go-zero 自带的 sqlx 不太好用，打算还是使用自己更熟悉的 [gorm](https://gorm.io/zh_CN/docs/)。

使用下面命令引入 gorm：

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```

声明表对应的 model：

```
type WorkUser struct {
    BaseModel
    
    Id         int64     `db:"id"`          // id
    UserName   string    `db:"userName"`    // 用户名
    Pwd        string    `db:"pwd"`         // 密码
    Token      string    `db:"token"`       // 密码加密的 token
    Nickname   string    `db:"nickName"`    // 昵称
    Sex        int64     `db:"sex"`         // 性别
    Avatar     string    `db:"avatar"`      // 头像
    Ext        string     `db:"ext"`         // 预留的扩展字段
}
```

其中 BaseModel 的结构是自己定义的类似于 `gorm.Model` 的结构体：

```
type BaseModel struct {
    ID         uint64 `gorm:"primarykey"`
    CreateTime time.Time
    UpdateTime time.Time
    DeleteTime gorm.DeletedAt `gorm:"index"`
}
```

为了让 WorkUser 的实例能获取到表名称，需要实现 gorm.Tabler 接口：

```
func (WorkUser) TableName() string {
	return "work_user"
}
```

随后，你只需按正常逻辑编写增、删、改、查的方法了。例如创建用户：

```
// Insert 创建
func (_this *WorkUser) Insert(user WorkUser) (*WorkUser, error) {
	err := DB.Create(&user).Error
	if err != nil {
		return nil, errors.Wrap(err, "select error. ")
	}
	return &user, nil
}
```

## 参考资料
* grom 文档 https://gorm.io/zh_CN/docs/index.html
* MySQL 編碼挑選與差異比較 https://khiav223577.github.io/blog/2019/06/30/MySQL-%E7%B7%A8%E7%A2%BC%E6%8C%91%E9%81%B8%E8%88%87%E5%B7%AE%E7%95%B0%E6%AF%94%E8%BC%83/
