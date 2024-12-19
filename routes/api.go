package routes

import (
	"github.com/GiorgiTsukhishvili/BookShelf-Api/controllers"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/middlewares"
	"github.com/gin-gonic/gin"
)

func ApiRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	v1.Use(middlewares.LangCheck)
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
				user.POST("/update-email", controllers.PutUserPassword)
				user.PUT("/", controllers.PutUser)
				user.PUT("/password-update", controllers.UserEmailUpdate)
				user.PUT("/update-email-verify", controllers.UserEmailUpdateVerify)
				user.DELETE("/:id", controllers.DeleteUser)
			}

			author := private.Group("/author")
			{
				author.GET("/:id")
				author.GET("/")

				admin := author.Group("")

				{
					admin.POST("/")
					admin.PUT("/")
					admin.DELETE("/:id")
				}
			}
		}
	}
}
