package app

import (
	"github.com/shiva0612/banking/service"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	service service.AccountService
}

func (ah AccountHandler) NewAccount(c *gin.Context) {

}

func (ah AccountHandler) MakeTransaction(c *gin.Context) {

}
