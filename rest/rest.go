package rest

import (
	"base-gin/server"
	"base-gin/service"

	"github.com/gin-gonic/gin"
)

var (
	accountHandler *AccountHandler
	personHandler  *PersonHandler
	authorHandler  *AuthorHandler
)

func SetupRestHandlers(app *gin.Engine) {
	handler := server.GetHandler()

	accountHandler = NewAccountHandler(
		handler, service.GetAccountService(), service.GetPersonService())
	personHandler = NewPersonHandler(handler, service.GetPersonService())
	authorHandler = NewAuthorHandler(handler, service.GetAuthorService())

	setupRoutes(app)
}

func setupRoutes(app *gin.Engine) {
	accountHandler.Route(app)
	personHandler.Route(app)
	authorHandler.Route(app)
}
