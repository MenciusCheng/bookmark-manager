package handler

import (
	"github.com/MenciusCheng/bookmark-manager/dao"
	"github.com/MenciusCheng/bookmark-manager/model"
	"github.com/gin-gonic/gin"
)

func BookmarkCreate(c *gin.Context) {
	req := model.Bookmark{}
	if err := c.BindJSON(&req); err != nil {
		JSONErr(c, ErrCodeParam, err)
		return
	}

	err := dao.BookmarkDao.Create(&req)
	if err != nil {
		JSONErr(c, ErrCodeInternal, err)
		return
	}

	JSON(c, nil)
}

func BookmarkUpdate(c *gin.Context) {
	req := model.Bookmark{}
	if err := c.BindJSON(&req); err != nil {
		JSONErr(c, ErrCodeParam, err)
		return
	}

	err := dao.BookmarkDao.Update(&req)
	if err != nil {
		JSONErr(c, ErrCodeInternal, err)
		return
	}

	JSON(c, nil)
}

func BookmarkDelete(c *gin.Context) {
	req := model.BookmarkReq{}
	if err := c.BindJSON(&req); err != nil {
		JSONErr(c, ErrCodeParam, err)
		return
	}

	err := dao.BookmarkDao.Delete(req.ID)
	if err != nil {
		JSONErr(c, ErrCodeInternal, err)
		return
	}

	JSON(c, nil)
}

func BookmarkInfo(c *gin.Context) {
	req := model.BookmarkReq{}
	if err := c.BindJSON(&req); err != nil {
		JSONErr(c, ErrCodeParam, err)
		return
	}

	res, err := dao.BookmarkDao.Info(req.ID)
	if err != nil {
		JSONErr(c, ErrCodeInternal, err)
		return
	}

	JSON(c, res)
}

func BookmarkPage(c *gin.Context) {
	req := model.BookmarkPageReq{}
	if err := c.BindJSON(&req); err != nil {
		JSONErr(c, ErrCodeParam, err)
		return
	}

	list, count, err := dao.BookmarkDao.Page(req)
	if err != nil {
		JSONErr(c, ErrCodeInternal, err)
		return
	}
	res := model.BookmarkPageRes{
		List:  list,
		Count: count,
	}

	JSON(c, res)
}
