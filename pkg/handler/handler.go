package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	var router = gin.New()
	h.initIndex(router)
	list := router.Group("/todo")
	{
		list.GET("/", h.getAllItem)
		list.POST("/", h.createItem)
		list.PUT("/:id", h.putItem)
		list.DELETE("/:id", h.deleteItem)
	}
	return router
}

func (h *Handler) initIndex(router *gin.Engine) {
	index := router.Group("/")
	{
		index.GET("/", h.index)
	}
}
