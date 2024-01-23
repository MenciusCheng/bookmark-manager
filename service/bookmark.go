package service

import (
	"github.com/MenciusCheng/bookmark-manager/dao"
	"github.com/MenciusCheng/bookmark-manager/model"
)

type ListBookmarkRes struct {
	List  []model.Bookmark `json:"list"`
	Count int64            `json:"count"`
}

func ListBookmark(req model.BookmarkPageReq) (ListBookmarkRes, error) {
	bookmarks, count, err := dao.BookmarkDao.Page(req)
	if err != nil {
		return ListBookmarkRes{}, err
	}
	return ListBookmarkRes{
		List:  bookmarks,
		Count: count,
	}, nil
}
