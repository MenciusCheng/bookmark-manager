package handler

import (
	"github.com/MenciusCheng/bookmark-manager/dao"
	"github.com/MenciusCheng/bookmark-manager/model"
	"github.com/gin-gonic/gin"
)

func BookmarkFolderCreate(c *gin.Context) {
	req := model.BookmarkFolder{}
	if err := c.BindJSON(&req); err != nil {
		JSONErr(c, ErrCodeParam, err)
		return
	}

	if req.ParentID > 0 {
		if _, err := dao.BookmarkFolderDao.Info(req.ParentID); err != nil {
			JSONErr(c, ErrCodeParam, err)
			return
		}
	}

	err := dao.BookmarkFolderDao.Create(&req)
	if err != nil {
		JSONErr(c, ErrCodeInternal, err)
		return
	}

	JSON(c, nil)
}

func BookmarkFolderUpdate(c *gin.Context) {
	req := model.BookmarkFolder{}
	if err := c.BindJSON(&req); err != nil {
		JSONErr(c, ErrCodeParam, err)
		return
	}

	if req.ParentID > 0 {
		if _, err := dao.BookmarkFolderDao.Info(req.ParentID); err != nil {
			JSONErr(c, ErrCodeParam, err)
			return
		}
	}

	err := dao.BookmarkFolderDao.Update(&req)
	if err != nil {
		JSONErr(c, ErrCodeInternal, err)
		return
	}

	JSON(c, nil)
}

func BookmarkFolderDelete(c *gin.Context) {
	req := model.BookmarkFolderReq{}
	if err := c.BindJSON(&req); err != nil {
		JSONErr(c, ErrCodeParam, err)
		return
	}

	err := dao.BookmarkFolderDao.Delete(req.ID)
	if err != nil {
		JSONErr(c, ErrCodeInternal, err)
		return
	}

	JSON(c, nil)
}

func BookmarkFolderInfo(c *gin.Context) {
	req := model.BookmarkFolderReq{}
	if err := c.BindJSON(&req); err != nil {
		JSONErr(c, ErrCodeParam, err)
		return
	}

	res, err := dao.BookmarkFolderDao.Info(req.ID)
	if err != nil {
		JSONErr(c, ErrCodeInternal, err)
		return
	}

	JSON(c, res)
}

func BookmarkFolderPage(c *gin.Context) {
	req := model.BookmarkFolderPageReq{}
	if err := c.BindJSON(&req); err != nil {
		JSONErr(c, ErrCodeParam, err)
		return
	}

	list, count, err := dao.BookmarkFolderDao.Page(req)
	if err != nil {
		JSONErr(c, ErrCodeInternal, err)
		return
	}
	res := model.BookmarkFolderPageRes{
		List:  list,
		Count: count,
	}

	JSON(c, res)
}
