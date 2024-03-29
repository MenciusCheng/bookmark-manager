package dao

import (
	"github.com/MenciusCheng/bookmark-manager/conf"
	"github.com/MenciusCheng/bookmark-manager/model"
)

var BookmarkDao bookmarkDao

type bookmarkDao struct {
}

func (b *bookmarkDao) Create(req *model.Bookmark) error {
	return conf.DB.Create(req).Error
}

func (b *bookmarkDao) Update(req *model.Bookmark) error {
	if _, err := b.Info(req.ID); err != nil {
		return err
	}

	return conf.DB.UpdateColumns(req).Error
}

func (b *bookmarkDao) Delete(id int64) error {
	if _, err := b.Info(id); err != nil {
		return err
	}

	return conf.DB.Delete(&model.Bookmark{}, id).Error
}

func (b *bookmarkDao) Info(id int64) (model.Bookmark, error) {
	res := model.Bookmark{}
	err := conf.DB.First(&res, id).Error
	return res, err
}

func (b *bookmarkDao) List(req model.BookmarkPageReq) (res []model.Bookmark, err error) {
	query := conf.DB.Table((&model.Bookmark{}).TableName())
	if req.ID > 0 {
		query = query.Where("id = ?", req.ID)
	}
	if req.FolderID > 0 {
		query = query.Where("folder_id = ?", req.FolderID)
	}

	if req.Offset > 0 || req.Limit > 0 {
		query = query.Offset(req.Offset).Limit(req.Limit)
	} else {
		query = query.Limit(model.DefaultLimit)
	}

	if err = query.Find(&res).Error; err != nil {
		return
	}
	return
}

func (b *bookmarkDao) Page(req model.BookmarkPageReq) (res []model.Bookmark, count int64, err error) {
	query := conf.DB.Table((&model.Bookmark{}).TableName())
	if req.ID > 0 {
		query = query.Where("id = ?", req.ID)
	}
	if req.FolderID > 0 {
		query = query.Where("folder_id = ?", req.FolderID)
	}

	if req.Offset > 0 || req.Limit > 0 {
		query = query.Offset(req.Offset).Limit(req.Limit)
	} else {
		query = query.Limit(model.DefaultLimit)
	}

	if err = query.Find(&res).Error; err != nil {
		return
	}

	if err = query.Offset(-1).Limit(-1).Count(&count).Error; err != nil {
		return
	}

	return
}
