package repository

import "base-gin/storage"

var (
	accountRepo *AccountRepository
	personRepo  *PersonRepository
	authorRepo  *AuthorRepository
)

func SetupRepositories() {
	db := storage.GetDB()
	accountRepo = NewAccountRepository(db)
	personRepo = NewPersonRepository(db)
	authorRepo = NewAuthorRepository(db)
}

func GetAccountRepo() *AccountRepository {
	return accountRepo
}

func GetPersonRepo() *PersonRepository {
	return personRepo
}

func GetAuthorRepo() *AuthorRepository {
	return authorRepo
}
