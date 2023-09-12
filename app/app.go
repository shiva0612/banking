package app

import (
	"github.com/shiva0612/banking/configs"
	"github.com/shiva0612/banking/domain"
	"github.com/shiva0612/banking/service"

	"github.com/gin-gonic/gin"
)

func Start() {
	configs.LoadConfigJson("/Users/kamarapushivachandra/Desktop/banking/banking/app-config.json")
	domain.ConnectToDb()

	ah := AccountHandler{service.NewDefaultAccountService(domain.NewAccountRepoDb())}
	ch := CustomerHandler{service.NewDefaultCustomerService(domain.NewCustomerRepoDb())}

	r := gin.Default()
	r.GET("/customers", ch.GetAllCustomers)
	r.GET("/customer/:customer_id", ch.GetCustomerById)
	r.POST("/customer/:customer_id/accounts", ah.NewAccount)
	r.POST("/customer/:customer_id/accounts/:account_id", ah.MakeTransaction)

	host := configs.Cfg.Server
	port := configs.Cfg.Port
	r.Run(host + ":" + port)

}
