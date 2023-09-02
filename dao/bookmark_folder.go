package dao

import (
	"github.com/MenciusCheng/bookmark-manager/conf"
	"github.com/MenciusCheng/bookmark-manager/model"
)

var BookmarkFolderDao bookmarkFolderDao

type bookmarkFolderDao struct {
}

func (b *bookmarkFolderDao) Create(req *model.BookmarkFolder) error {
	return conf.DB.Create(req).Error
}

func (b *bookmarkFolderDao) Update(req *model.BookmarkFolder) error {
	if _, err := b.Info(req.ID); err != nil {
		return err
	}

	return conf.DB.UpdateColumns(req).Error
}

func (b *bookmarkFolderDao) Delete(id int64) error {
	if _, err := b.Info(id); err != nil {
		return err
	}

	return conf.DB.Delete(&model.BookmarkFolder{}, id).Error
}

func (b *bookmarkFolderDao) Info(id int64) (model.BookmarkFolder, error) {
	res := model.BookmarkFolder{}
	err := conf.DB.First(&res, id).Error
	return res, err
}

func (b *bookmarkFolderDao) List(req model.BookmarkFolderPageReq) (res []model.BookmarkFolder, err error) {
	query := conf.DB.Table((&model.BookmarkFolder{}).TableName())
	if req.ID > 0 {
		query = query.Where("id = ?", req.ID)
	}
	if req.ParentID > 0 {
		query = query.Where("parent_id = ?", req.ParentID)
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

func (b *bookmarkFolderDao) Page(req model.BookmarkFolderPageReq) (res []model.BookmarkFolder, count int64, err error) {
	query := conf.DB.Table((&model.BookmarkFolder{}).TableName())
	if req.ID > 0 {
		query = query.Where("id = ?", req.ID)
	}
	if req.ParentID > 0 {
		query = query.Where("parent_id = ?", req.ParentID)
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
