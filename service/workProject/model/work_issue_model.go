package model

import userModel "github.com/suhanyujie/workUser/service/workUser/model"

type (
	WorkTask struct {
		userModel.BaseModel

		Title string `gorm:"comment:项目名称;size:100"`
		Desc  string `gorm:"comment:详情;size:1000"`
		Sort  int64  `gorm:"comment:排序值"`
		Ext   string `gorm:"comment:预留的扩展字段"`
	}
)
