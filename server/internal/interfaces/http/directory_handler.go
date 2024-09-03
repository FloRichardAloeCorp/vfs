package http

import (
	"net/http"

	"github.com/FloRichardAloeCorp/vfs/server/internal/features/directory"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DirectoryHandler struct {
	directoryFeatures directory.DirectoryFeatures
}

func NewDirectoryHandler(directoryFeatures directory.DirectoryFeatures) *DirectoryHandler {
	return &DirectoryHandler{
		directoryFeatures: directoryFeatures,
	}
}

func (h *DirectoryHandler) RegisterRoutes(router *gin.Engine) {
	router.POST("/directory/*path", h.Post)
	router.GET("/directory/*path", h.Get)
	router.DELETE("/directory/*path", h.Delete)
}

func (h *DirectoryHandler) Post(c *gin.Context) {
	path := c.Param("path")
	if path == "" {
		log.Error("DirectoryHandler.Post fail", zap.String("error", "empty path"))
		c.JSON(http.StatusBadRequest, "Bad Request: empty path")
		return
	}

	err := h.directoryFeatures.Create(path)
	if err != nil {
		log.Error("DirectoryHandler.Post fail",
			zap.Error(err),
		)
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	c.JSON(http.StatusCreated, "Directory created")
}

func (h *DirectoryHandler) Get(c *gin.Context) {
	path := c.Param("path")
	if path == "" {
		log.Error("DirectoryHandler.Get fail", zap.String("error", "empty path"))
		c.JSON(http.StatusBadRequest, "Bad Request: empty path")
		return
	}

	files, err := h.directoryFeatures.ListFiles(path)
	if err != nil {
		log.Error("DirectoryHandler.Get fail list directory content:", zap.Error(err))
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	c.JSON(http.StatusOK, files)
}

func (h *DirectoryHandler) Delete(c *gin.Context) {
	path := c.Param("path")
	if path == "" {
		log.Error("DirectoryHandler.Delete fail", zap.String("error", "empty path"))
		c.JSON(http.StatusBadRequest, "Bad Request: empty path")
		return
	}

	err := h.directoryFeatures.Delete(path)
	if err != nil {
		log.Error("DirectoryHandler.Delete :", zap.Error(err))
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	c.JSON(http.StatusNoContent, "Directory deleted")
}
