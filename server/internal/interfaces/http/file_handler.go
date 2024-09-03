package http

import (
	"io"
	"net/http"

	"github.com/FloRichardAloeCorp/vfs/server/internal/features/file"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FileHandler struct {
	fileFeatures file.FileFeatures
}

func NewFileHandler(fileFeatures file.FileFeatures) *FileHandler {
	return &FileHandler{
		fileFeatures: fileFeatures,
	}
}

func (h *FileHandler) RegisterRoutes(router *gin.Engine) {
	router.POST("/file/*path", h.Post)
	router.DELETE("/file/*path", h.Delete)
	router.GET("/file/info/*path", h.GetInfo)
	router.GET("/file/content/*path", h.Get)
}

func (h *FileHandler) Post(c *gin.Context) {
	path := c.Param("path")
	if path == "" {
		log.Error("FileHandler.Post fail", zap.String("error", "empty path"))
		c.JSON(http.StatusBadRequest, "Bad Request: empty path")
		return
	}

	content, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("FileHandler.Post fail", zap.Error(err))
		c.JSON(http.StatusBadRequest, "Bad Request: can't read request body")
		return
	}

	defer c.Request.Body.Close()

	err = h.fileFeatures.Create(path, content)
	if err != nil {
		log.Error("FileHandler.Post fail",
			zap.Error(err),
		)
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	c.JSON(http.StatusCreated, "File created")
}

func (h *FileHandler) Get(c *gin.Context) {
	path := c.Param("path")
	if path == "" {
		log.Error("FileHandler.Get fail", zap.String("error", "empty path"))
		c.JSON(http.StatusBadRequest, "Bad Request: empty path")
		return
	}

	file, err := h.fileFeatures.Read(path)
	if err != nil {
		log.Error("FileHandler.Get fail to get file:", zap.Error(err))
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	c.Data(http.StatusOK, "application/octet-stream", file)
}

func (h *FileHandler) GetInfo(c *gin.Context) {
	path := c.Param("path")
	if path == "" {
		log.Error("FileHandler.GetInfo fail", zap.String("error", "empty path"))
		c.JSON(http.StatusBadRequest, "Bad Request: empty path")
		return
	}

	info, err := h.fileFeatures.ReadInfo(path)
	if err != nil {
		log.Error("FileHandler.GetInfo fail to get file info:", zap.Error(err))
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	c.JSON(http.StatusOK, info)
}

func (h *FileHandler) Delete(c *gin.Context) {
	path := c.Param("path")
	if path == "" {
		log.Error("FileHandler.Delete fail", zap.String("error", "empty path"))
		c.JSON(http.StatusBadRequest, "Bad Request: empty path")
		return
	}

	err := h.fileFeatures.Delete(path)
	if err != nil {
		log.Error("FileHandler.Delete :", zap.Error(err))
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	c.JSON(http.StatusNoContent, "File deleted")
}
