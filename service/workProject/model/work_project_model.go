package model

import userModel "github.com/suhanyujie/workUser/service/workUser/model"

type (
	WorkProject struct {
		userModel.BaseModel

		Title    string `gorm:"comment:项目名称;size:100"`
		Desc     string `gorm:"comment:简介;size:100"`
		Settings string `gorm:"comment:一些设置数据;size:1000"`
		Sort     int64  `gorm:"comment:排序值"`
		Ext      string `gorm:"comment:预留的扩展字段"`
	}
)
