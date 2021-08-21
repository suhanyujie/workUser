package model

import (
	"strings"
	"time"

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
	workUserRowsExpectAutoSet   = strings.Join(stringx.Remove(workUserFieldNames, "pwd", "create_time", "update_time"), ",")
	workUserRowsWithPlaceHolder = strings.Join(stringx.Remove(workUserFieldNames, "pwd", "create_time", "update_time"), "=?,") + "=?"

	cacheWorkUserInfoPrefix = "cache#AppUser#uid#"
)

type (
	WorkUserModel struct {
		sqlc.CachedConn
		table string
	}

	WorkUser struct {
		Uid        int64     `db:"uid"`         // 对应中台表中的id
		Avatar     string    `db:"avatar"`      // 头像
		CreateTime time.Time `db:"create_time"` // 创建时间
		UserName   string    `db:"nickname"`    // 用户名
		Pwd        string    `db:"pwd"`         // 密码
		Nickname   string    `db:"nickname"`    // 昵称
		Sex        int64     `db:"sex"`         // 性别
		UpdateTime time.Time `db:"update_time"`
	}
)

func NewWorkUserModel(conn sqlx.SqlConn, c cache.CacheConf) *WorkUserModel {
	return &WorkUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "work_user",
	}
}
