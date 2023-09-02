package model

import (
	"github.com/MenciusCheng/bookmark-manager/conf"
	"testing"
)

func TestAutoMigrate(t *testing.T) {
	conf.Init("")

	if err := conf.DB.Set("gorm:table_options", "COMMENT='用户'").AutoMigrate(&User{}); err != nil {
		t.Error(err)
		return
	}
	if err := conf.DB.Set("gorm:table_options", "COMMENT='书签'").AutoMigrate(&Bookmark{}); err != nil {
		t.Error(err)
		return
	}
	if err := conf.DB.Set("gorm:table_options", "COMMENT='书签文件夹'").AutoMigrate(&BookmarkFolder{}); err != nil {
		t.Error(err)
		return
	}
}
