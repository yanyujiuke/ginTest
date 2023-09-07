package models

type Access struct {
	Id         int      `json:"id"`
	ModuleName string   `json:"module_name"`
	Type       int      `json:"type"`        // 节点类型 1模块 2菜单 3操作
	ActionName string   `json:"action_name"` // 操作名称
	Url        string   `json:"url"`
	ModuleId   int      `json:"module_id"` // 和当前模型的id关联
	Sort       int      `json:"sort"`
	Desc       string   `json:"desc"`
	AddTime    int      `json:"add_time"`
	Status     int      `json:"status"`
	AccessItem []Access `gorm:"foreignKey:ModuleId"`
	Checked    bool     `json:"checked" gorm:"-"` // 忽略本字段
}

func (Access) TableName() string {
	return "access"
}
