package app

import (
	"github.com/shiva0612/banking/service"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch CustomerHandler) GetAllCustomers(c *gin.Context) {

}

func (ch CustomerHandler) GetCustomerById(c *gin.Context) {

}
