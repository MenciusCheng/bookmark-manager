package handler

import (
	"github.com/MenciusCheng/bookmark-manager/dao"
	"github.com/MenciusCheng/bookmark-manager/model"
	"github.com/MenciusCheng/bookmark-manager/service"
	"github.com/gin-gonic/gin"
)

func BookmarkCreate(c *gin.Context) {
	bookmark := model.Bookmark{}
	if err := c.BindJSON(&bookmark); err != nil {
		JSONErr(c, ErrCodeParam, err)
		return
	}

	err := dao.CreateBookmark(bookmark)
	if err != nil {
		JSONErr(c, ErrCodeInternal, err)
		return
	}

	JSON(c, nil)
}

func BookmarkList(c *gin.Context) {
	res, err := service.ListBookmark()
	if err != nil {
		JSONErr(c, ErrCodeParam, err)
		return
	}
	JSON(c, res)
}
