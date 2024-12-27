package scripts

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ConvertStringToInt(value string, ctx *gin.Context) int {
	number, err := strconv.Atoi(value)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return 0
	}

	return number
}
