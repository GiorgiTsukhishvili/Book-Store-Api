package scripts

import "github.com/gin-gonic/gin"

func InvalidateJwtCookies(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "", true, true)
	ctx.SetCookie("refreshToken", "", -1, "/", "", true, true)
}
