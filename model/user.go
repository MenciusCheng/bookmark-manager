package model

type User struct {
	Model
	Name string `json:"name" gorm:"column:name;type:varchar(255) not null default '';comment:名称"`
	Age  int    `json:"age" gorm:"column:age;type:int not null default 0;comment:年龄"`
	Sex  int    `json:"sex" gorm:"column:sex;type:tinyint not null default 0;comment:性别,0男,1女"`
}

func (u *User) TableName() string {
	return "user"
}
