package models

type GoodsCate struct {
	Id             int         `json:"id"`
	Title          string      `json:"title"`
	CateImg        string      `json:"cate_img"`
	Link           string      `json:"link"`
	Template       string      `json:"template"`
	Pid            int         `json:"pid"`
	SubTitle       string      `json:"sub_title"`
	Keywords       string      `json:"keywords"`
	Desc           string      `json:"desc"`
	Status         int         `json:"status"`
	Sort           int         `json:"sort"`
	AddTime        int         `json:"add_time"`
	GoodsCateItems []GoodsCate `gorm:"foreignKey:pid;references:id"`
}

func (GoodsCate) TableName() string {
	return "goods_cate"
}
