package rest

import (
	"base-gin/server"
	"base-gin/service"

	"github.com/gin-gonic/gin"
)

type AuthorHandler struct {
	hr            *server.Handler
	service       *service.AuthorService
	authorService *service.AuthorService
}

func NewAuthorHandler(
	hr *server.Handler,
	authorService *service.AuthorService,
) *AuthorHandler {
	return &AuthorHandler{hr: hr, service: authorService}
}

func (h *AuthorHandler) Route(app *gin.Engine) {
	// grp := app.Group(server.RootAuthor)
}