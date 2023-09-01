package service

import (
	"github.com/MenciusCheng/bookmark-manager/dao"
	"github.com/MenciusCheng/bookmark-manager/model"
)

type ListBookmarkRes struct {
	List  []model.Bookmark `json:"list"`
	Count int64            `json:"count"`
}

func ListBookmark() (ListBookmarkRes, error) {
	bookmarks, count, err := dao.ListBookmark()
	if err != nil {
		return ListBookmarkRes{}, err
	}
	return ListBookmarkRes{
		List:  bookmarks,
		Count: count,
	}, nil
}
