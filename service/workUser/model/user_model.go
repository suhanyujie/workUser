package model

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

// user mode 的实现可以参考 https://github.com/jackluo2012/datacenter/blob/main/user/model/appusermodel.go

var (
	workUserFieldNames          = builderx.FieldNames(&WorkUser{})
	workUserRows                = strings.Join(workUserFieldNames, ",")
	workUserRowsExpectAutoSet   = strings.Join(stringx.Remove(workUserFieldNames, "pwd", "created_at", "updated_at", "deleted_at"), ",")
	workUserRowsWithPlaceHolder = strings.Join(stringx.Remove(workUserFieldNames, "pwd", "created_at", "updated_at", "deleted_at"), "=?,") + "=?"

	cacheWorkUserInfoPrefix = "cache:WorkUser:id:"
)

type (
	WorkUserModelIf interface {
		Insert(data WorkUser) (*WorkUser, error)
		FindOne(id int64) (*WorkUser, error)
		FindOneByUserName(user string) (*WorkUser, error)
		Update(data WorkUser) error
		Delete(id int64) error
	}

	WorkUser struct {
		BaseModel

		UserName string `gorm:"comment:用户名;size:50"`         // 用户名
		Pwd      string `gorm:"comment:密码;size:128"`         // 密码
		Token    string `gorm:"comment:密码加密的 token;size:50"` // 密码加密的 token
		Nickname string `gorm:"comment:昵称;size:200"`         // 昵称
		Sex      int    `gorm:"comment:性别"`                  // 性别
		Avatar   string `gorm:"comment:头像;size:200"`         // 头像
		Ext      string `gorm:"comment:预留的扩展字段"`             // 预留的扩展字段
	}
)

func (WorkUser) TableName() string {
	return "work_user"
}

func NewWorkUserModel() *WorkUser {
	return &WorkUser{}
}

// Insert 创建
func (_this *WorkUser) Insert(user WorkUser) (*WorkUser, error) {
	err := DB.Create(&user).Error
	if err != nil {
		return nil, errors.Wrap(err, "select error. ")
	}
	return &user, nil
}

// gorm 实现 WorkUser 的修改 todo

// FindOne 通过 id 查询一条数据
func (_this *WorkUser) FindOne(id int64) (*WorkUser, error) {
	var user *WorkUser
	if err := DB.Where("id=?", id).Find(&user).Error; err != nil {
		return nil, errors.Wrap(err, "findOne error. ")
	}
	return user, nil
}

// FindOneByUserName 通过用户名模糊匹配用户
func (_this *WorkUser) FindOneByUserName(userName string) ([]WorkUser, error) {
	users := make([]WorkUser, 0)
	likeFmt := fmt.Sprintf("%%%s%%", userName)
	if err := DB.Table(_this.TableName()).Where("user_name like ?", likeFmt).Find(&users).Error; err != nil {
		return nil, errors.Wrap(err, "findOne error. ")
	}
	return users, nil
}

// todo
func (_this *WorkUser) Update(data WorkUser) error {
	return nil
}

// Delete 通过主键删除
func (_this *WorkUser) Delete(id int64) error {
	if err := DB.Table(_this.TableName()).Where("id=?", id).Delete(NewWorkUserModel()).Error; err != nil {
		return errors.Wrap(err, "delete error. ")
	}
	return nil
}
