package model

import (
	"time"
)

const (
	DefaultLimit = 100000
)

type Model struct {
	ID        int64     `json:"id" gorm:"primarykey;column:id;type:bigint not null auto_increment;comment:主键"`
	CreatedAt time.Time `json:"createdAt" gorm:"<-:false;column:created_at;type:datetime not null default current_timestamp;comment:创建时间"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"<-:false;column:updated_at;type:datetime not null default current_timestamp on update current_timestamp;comment:修改时间"`
}

type PageReq struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
