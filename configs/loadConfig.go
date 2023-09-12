package configs

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/shiva0612/banking-lib/logger"

	"github.com/spf13/viper"
)

var (
	Cfg *Config
)

type Config struct {
	Server string      `json:"server"`
	Port   string      `json:"port"`
	DbCfg  DatabaseCfg `json:"db_cfg"`
}

type DatabaseCfg struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func LoadConfigJson(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic("error while opening the config file: " + err.Error())
	}

	Cfg = new(Config)
	err = json.NewDecoder(file).Decode(Cfg)
	if err != nil {
		panic("error un-marshalling the config: " + err.Error())
	}
}
func LoadConfig(path string) {
	viper.AddConfigPath(filepath.Dir(path))     //dir path
	viper.SetConfigName(filepath.Base(path))    //filename
	viper.SetConfigType(filepath.Ext(path)[1:]) //.json

	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal("error reading the config file at path = %s, err = %s", path, err.Error())
	}

	Cfg = new(Config)
	err = viper.Unmarshal(Cfg)
	if err != nil {
		logger.Fatal("error unmarshalling config, err = ", err.Error())
	}
}
