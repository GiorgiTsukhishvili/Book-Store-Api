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
				author.GET("/:id", controllers.GetAuthor)
				author.GET("/", controllers.GetAuthors)

				admin := author.Group("")
				admin.Use(middlewares.AdminCheck)
				{
					admin.POST("/", controllers.PostAuthor)
					admin.PUT("/", controllers.PutAuthor)
					admin.DELETE("/:id", controllers.DeleteAuthor)
				}
			}

			book := private.Group("/book")
			{
				book.GET("/:id", controllers.GetBook)
				book.GET("/", controllers.GetBooks)

				business := book.Group("")
				business.Use(middlewares.BusinessCheck)
				{
					business.POST("/", controllers.PostBook)
					business.PUT("/", controllers.PutBook)
					business.DELETE("/:id", controllers.DeleteBook)
				}
			}

			review := private.Group("/review")
			{
				review.GET("/:id", controllers.GetReview)
				review.GET("/", controllers.GetReviews)
				review.POST("/", controllers.PostReview)
				review.PUT("/", controllers.PutReview)
			}

			favorite := private.Group("/favorite")
			{
				favorite.GET("/", controllers.GetUserFavorites)
				favorite.POST("/", controllers.PostFavorite)
				favorite.DELETE("/", controllers.DeleteFavorite)
			}

			notification := private.Group("/notification")
			{
				notification.GET("/", controllers.GetNotifications)
				notification.PUT("/", controllers.PutNotification)
			}
		}
	}
}
