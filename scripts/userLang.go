package scripts

import (
	"github.com/gin-gonic/gin"
)

func GetUserLang(ctx *gin.Context) string {
	claims, _ := ctx.Get("lang")

	lang, _ := claims.(string)

	return lang
}
