package scripts

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SaveImage(ctx *gin.Context) string {
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get image"})
		return ""
	}

	uploadDir := "./public/images"
	err = os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create images directory"})
		return ""
	}

	ext := filepath.Ext(file.Filename)
	uniqueFilename := fmt.Sprintf("%d-%s%s", time.Now().UnixNano(), uuid.New().String(), ext)

	filePath := filepath.Join(uploadDir, uniqueFilename)
	err = ctx.SaveUploadedFile(file, filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return ""
	}

	return filePath
}
