package account

import store "github.com/dpurbosakti/go-native/internal/repositories"

type AccountService struct {
	Store store.Store
}

func NewAccountService(s store.Store) *AccountService {
	return &AccountService{
		Store: s,
	}
}
