package model

// 书签
type Bookmark struct {
	Model
	Name     string `json:"name" gorm:"column:name;type:varchar(255) not null default '';comment:名称"`
	Url      string `json:"url" gorm:"column:url;type:varchar(5000) not null default '';comment:网址"`
	FolderID int64  `json:"folderID" gorm:"column:folder_id;type:bigint not null default '0';comment:所属文件夹ID"`
}

func (b *Bookmark) TableName() string {
	return "bookmark"
}

type BookmarkReq struct {
	ID int64 `json:"id"` // 主键
}

type BookmarkPageReq struct {
	PageReq
	ID       int64 `json:"id"`       // 主键
	FolderID int64 `json:"folderID"` // 所属文件夹ID
}

type BookmarkPageRes struct {
	List  []Bookmark `json:"list"`
	Count int64      `json:"count"`
}
