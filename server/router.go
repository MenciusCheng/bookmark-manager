package server

import (
	"github.com/MenciusCheng/bookmark-manager/server/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// 路由
	v1 := r.Group("/api/v1")
	v1.Any("/ping", handler.Ping)

	v1.POST("bookmark/create", handler.BookmarkCreate)
	v1.POST("bookmark/update", handler.BookmarkUpdate)
	v1.POST("bookmark/delete", handler.BookmarkDelete)
	v1.POST("bookmark/info", handler.BookmarkInfo)
	v1.POST("bookmark/page", handler.BookmarkPage)

	v1.POST("bookmarkFolder/create", handler.BookmarkFolderCreate)
	v1.POST("bookmarkFolder/update", handler.BookmarkFolderUpdate)
	v1.POST("bookmarkFolder/delete", handler.BookmarkFolderDelete)
	v1.POST("bookmarkFolder/info", handler.BookmarkFolderInfo)
	v1.POST("bookmarkFolder/page", handler.BookmarkFolderPage)

	return r
}
