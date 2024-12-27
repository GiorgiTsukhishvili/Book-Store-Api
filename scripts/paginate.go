package scripts

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginate(Page string, Size string, ctx *gin.Context) func(db *gorm.DB) *gorm.DB {
	page := ConvertStringToInt(Page, ctx)

	size := ConvertStringToInt(Size, ctx)

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
