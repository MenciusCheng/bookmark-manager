package dao

import (
	"github.com/MenciusCheng/bookmark-manager/conf"
	"github.com/MenciusCheng/bookmark-manager/model"
)

func CreateBookmark(bookmark model.Bookmark) error {
	err := conf.DB.Table("bookmark").Create(&bookmark).Error
	return err
}

func GetAllBookmark() ([]model.Bookmark, error) {
	res := make([]model.Bookmark, 0)
	err := conf.DB.Table("bookmark").Find(&res).Error
	return res, err
}

func ListBookmark() ([]model.Bookmark, int64, error) {
	res := make([]model.Bookmark, 0)
	err := conf.DB.Table("bookmark").Find(&res).Error

	var count int64
	conf.DB.Table("bookmark").Count(&count)
	return res, count, err
}
