package service

import "github.com/shiva0612/banking/domain"

type AccountService interface {
}

type DefaultAccountService struct {
	repo domain.AccountRepo
}

func NewDefaultAccountService(repo domain.AccountRepo) DefaultAccountService {
	return DefaultAccountService{repo}
}
