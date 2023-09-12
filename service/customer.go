package service

import "github.com/shiva0612/banking/domain"

type CustomerService interface {
}

type DefaultCustomerService struct {
	repo domain.CustomerRepo
}

func NewDefaultCustomerService(repo domain.CustomerRepo) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
