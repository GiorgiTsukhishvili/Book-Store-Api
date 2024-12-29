package requests

type NotificationGetRequest struct {
	Page string `form:"page" binding:"required"`
	Size string `form:"size" binding:"required"`
}

type NotificationPutRequest struct {
	IDs []string `json:"id" binding:"required"`
}
