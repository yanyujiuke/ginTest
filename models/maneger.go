package models

type Manager struct {
	Id       int    `json:"id" gorm:"id"`
	Username string `json:"username" gorm:"username"`
	Password string `json:"password" gorm:"password"`
	Mobile   string `json:"mobile" gorm:"mobile"`
	Email    string `json:"email" gorm:"email"`
	Status   int    `json:"status" gorm:"status"`
	RoleId   int    `json:"roleId" gorm:"roleId"`
	AddTime  int    `json:"addTime" gorm:"addTime"`
	IsSuper  int    `json:"isSuper" gorm:"isSuper"`
	Role     Role   `gorm:"foreignKey:RoleId"`
}

func (Manager) TableName() string {
	return "manager"
}
