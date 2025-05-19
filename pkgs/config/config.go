package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var config *viper.Viper

func init() {
	println("-------config init---------------------")
	Load("local")
	println("-----config init---------------------")
}
func Load(env string) {
	config = viper.New()
	config.SetConfigFile("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../../")
	config.AddConfigPath("app")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file ", err)
	}
	if env == "server" {
		log.Println("server running in prod getting values from env")
		for _, configKey := range config.AllKeys() {
			if envVal, ok := os.LookupEnv(config.GetString(configKey)); ok {
				log.Println("updating config value with env value for key=", configKey, " value=", envVal)
				config.Set(configKey, envVal)
			} else {
				log.Println("config value not found in env. key= ", configKey)
				os.Exit(1)
			}
		}
	}
	if err == nil {
		log.Println("successfully read server config. values are :", config.AllSettings())
	}

}

func GetConfig() *viper.Viper {
	return config
}
