package domain

import (
	"github.com/shiva0612/banking/configs"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var (
	db *sqlx.DB
)

func ConnectToDb() {
	var err error
	dbUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configs.Cfg.DbCfg.Host,
		configs.Cfg.DbCfg.Port,
		configs.Cfg.DbCfg.User,
		configs.Cfg.DbCfg.Password,
		configs.Cfg.DbCfg.Database)
	db, err = sqlx.Open("postgres", dbUrl)
	if err != nil {
		panic("db connection error: " + err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic("error in ping: " + err.Error())
	}
}
