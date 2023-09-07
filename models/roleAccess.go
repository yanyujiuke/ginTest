package models

type RoleAccess struct {
	// Id       int `json:"id"`
	RoleId   int `json:"role_id"`
	AccessId int `json:"access_id"`
}

func (RoleAccess) TableName() string {
	return "role_access"
}
