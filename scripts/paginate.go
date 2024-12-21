package scripts

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginate(Page string, Size string, ctx *gin.Context) func(db *gorm.DB) *gorm.DB {
	page, err := strconv.Atoi(Page)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return nil
	}

	size, err := strconv.Atoi(Size)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid size parameter"})
		return nil
	}

	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		switch {
		case size > 100:
			size = 100
		case size <= 0:
			size = 10
		}

		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}
