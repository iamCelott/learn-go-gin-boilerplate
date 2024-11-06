package service

import (
	"base-gin/config"
	"base-gin/repository"
)

var (
	accountService *AccountService
	personService  *PersonService
	authorService  *AuthorService
)

func SetupServices(cfg *config.Config) {
	accountService = NewAccountService(cfg, repository.GetAccountRepo())
	personService = NewPersonService(repository.GetPersonRepo())
	authorService = NewAuthorService(repository.GetAuthorRepo())
}

func GetAccountService() *AccountService {
	if accountService == nil {
		panic("account service is not initialised")
	}

	return accountService
}

func GetPersonService() *PersonService {
	return personService
}

func GetAuthorService() *AuthorService {
	return authorService
}
