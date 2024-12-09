package routes

import (
	"github.com/GiorgiTsukhishvili/BookShelf-Api/controllers"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/middlewares"
	"github.com/gin-gonic/gin"
)

func ApiRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		public := v1.Group("")
		{
			public.POST("/login", controllers.Login)
			public.POST("/register", controllers.Register)
			public.POST("/user-verify", controllers.VerifyUser)
			public.POST("/refresh-token", controllers.RefreshToken)
		}

		private := v1.Group("")
		private.Use(middlewares.AuthCheck)
		{
			private.POST("/logout", controllers.Logout)
			private.GET("/me", controllers.Me)

			user := private.Group("/user")
			{
				user.PUT("/:id", controllers.PutUser)
				user.DELETE("/:id", controllers.DeleteUser)
			}

		}
	}
}
