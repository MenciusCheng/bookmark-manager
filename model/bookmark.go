package model

type Bookmark struct {
	Model
	Name string `json:"name" gorm:"column:name;type:varchar(255) not null default '';comment:名称"`
	Url  string `json:"url" gorm:"column:url;type:varchar(5000) not null default '';comment:网址"`
}

func (b *Bookmark) TableName() string {
	return "bookmark"
}
