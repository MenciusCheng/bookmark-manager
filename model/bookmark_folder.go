package model

// 书签文件夹
type BookmarkFolder struct {
	Model
	Name     string `json:"name" gorm:"column:name;type:varchar(255) not null default '';comment:名称"`
	ParentID int64  `json:"parentID" gorm:"column:parent_id;type:bigint not null default '0';comment:父文件夹ID"`
}

func (f *BookmarkFolder) TableName() string {
	return "bookmark_folder"
}

type BookmarkFolderReq struct {
	ID int64 `json:"id"` // 主键
}

type BookmarkFolderPageReq struct {
	PageReq
	ID       int64 `json:"id"`       // 主键
	ParentID int64 `json:"parentID"` // 所属文件夹ID
}

type BookmarkFolderPageRes struct {
	List  []BookmarkFolder `json:"list"`
	Count int64            `json:"count"`
}
