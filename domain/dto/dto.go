package dto

type SuccessResponse[T any] struct {
	Success bool   `json:"success" binding:"default:true" example:"true"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

type ErrorResponse struct {
	Success bool        `json:"success" binding:"default:false" example:"false"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
}

type MessageDetailResponse struct {
	Header      string `json:"header"`
	Description string `json:"description"`
}

type ClientInfo struct {
	IPAddress string
	UserAgent string
	UserOS    string
	UserGeo   string
}
