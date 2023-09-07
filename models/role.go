package models

type Role struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Status  int    `json:"status"`
	AddTime int    `json:"add_time"`
}

func (Role) TableName() string {
	return "role"
}
