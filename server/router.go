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
	v1.POST("bookmark/update", handler.BookmarkCreate)
	v1.POST("bookmark/delete", handler.BookmarkCreate)
	v1.GET("bookmark/list", handler.BookmarkList)
	v1.GET("bookmark/info", handler.BookmarkList)

	return r
}
