package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

// user mode 的实现可以参考 https://github.com/jackluo2012/datacenter/blob/main/user/model/appusermodel.go

var (
	workUserFieldNames          = builderx.FieldNames(&WorkUser{})
	workUserRows                = strings.Join(workUserFieldNames, ",")
	workUserRowsExpectAutoSet   = strings.Join(stringx.Remove(workUserFieldNames, "pwd", "create_time", "update_time", "delete_time"), ",")
	workUserRowsWithPlaceHolder = strings.Join(stringx.Remove(workUserFieldNames, "pwd", "create_time", "update_time", "delete_time"), "=?,") + "=?"

	cacheWorkUserInfoPrefix = "cache#AppUser#uid#"
)

type (
	WorkUserModel interface {
		Insert(data WorkUser) (sql.Result, error)
		FindOne(id int64) (*WorkUser, error)
		FindOneByUserName(user string) (*WorkUser, error)
		Update(data WorkUser) error
		Delete(id int64) error
	}

	defaultWorkUserModel struct {
		conn  sqlx.SqlConn
		table string
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

func NewWorkUserModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultWorkUserModel {
	return &defaultWorkUserModel{
		conn:  conn,
		table: "work_user",
	}
}

func (rc *defaultWorkUserModel) Insert(data WorkUser) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", rc.table, workUserRows)
	ret, err := rc.conn.Exec(query, data.UserName, data.Pwd)

	return ret, errors.Wrap(err, "insert error")
}

func (rc *defaultWorkUserModel) FindOne(id int64) (*WorkUser, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", workUserRows, rc.table)
	var resp WorkUser
	err := rc.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (rc *defaultWorkUserModel) FindOneByUserName(userName string) (*WorkUser, error) {
	query := fmt.Sprintf("select %s from %s where `user_name` LIKE ? limit 1", workUserRows, rc.table)
	var resp WorkUser
	userNameLikeStr := fmt.Sprintf("%%%s%%", userName)
	err := rc.conn.QueryRow(&resp, query, userNameLikeStr)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (rc *defaultWorkUserModel) Update(data WorkUser) error {
	query := fmt.Sprintf("update %s SET %s where id=?", rc.table, workUserRowsWithPlaceHolder)
	_, err := rc.conn.Exec(query, data.Nickname, data.ID)
	return err
}

func (rc *defaultWorkUserModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id=?", rc.table)
	_, err := rc.conn.Exec(query, id)
	return err
}
