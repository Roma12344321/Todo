package handler

import (
	"HttpServer"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var items = []HttpServer.Todo{
	{1, "First", "First"},
	{2, "second", "second"},
	{3, "third", "third"},
}

func (h *Handler) index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": "success"})
}

func (h *Handler) createItem(c *gin.Context) {
	var newItem HttpServer.Todo
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	items = append(items, newItem)
	c.JSON(http.StatusOK, newItem)
}

func (h *Handler) getAllItem(c *gin.Context) {
	c.JSON(http.StatusOK, items)
}

func (h *Handler) putItem(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	var newItem HttpServer.Todo
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	found := false
	for i, item := range items {
		if item.Id == id {
			items[i].Id = newItem.Id
			items[i].Title = newItem.Title
			items[i].Description = newItem.Description
			found = true
			break
		}
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Item putted successfully"})
}

func (h *Handler) deleteItem(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	found := false
	for i, item := range items {
		if item.Id == id {
			items = append(items[:i], items[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}
