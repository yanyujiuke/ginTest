package models

type Focus struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	FocusType int    `json:"focus_type"` // 1pc 2app 3小程序
	FocusImg  string `json:"focus_img"`
	Link      string `json:"link"`
	Sort      int    `json:"sort"`
	Status    int    `json:"status"`
	AddTime   int    `json:"add_time"`
}

func (Focus) TableName() string {
	return "focus"
}
