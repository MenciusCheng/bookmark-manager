package model

// 文件夹
type Folder struct {
	Model
	Name     string `json:"name" gorm:"column:name;type:varchar(255) not null default '';comment:名称"`
	ParentID int64  `json:"parentID" gorm:"column:parent_id;type:bigint not null default '0';comment:父文件夹ID"`
}

func (f *Folder) TableName() string {
	return "folder"
}
