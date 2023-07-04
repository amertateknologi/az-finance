package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBSource1         string `mapstructure:"DB_SOURCE_1"`
	DBSource2         string `mapstructure:"DB_SOURCE_2"`
	DolphinAddress    string `mapstructure:"62DOLPHIN_ADDRESS"`
	WhaleAddress      string `mapstructure:"62WHALE_ADDRESS"`
}

var Data Config

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string, data *Config) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.SetDefault("HTTP_SERVER_ADDRESS", "0.0.0.0:8080")
	viper.SetDefault("DB_DRIVER", "mysql")
	viper.SetDefault("DB_SOURCE_1", "root@tcp(127.0.0.1:3306)/az_finance_local?charset=utf8mb4&parseTime=True&loc=Local")
	viper.SetDefault("DB_SOURCE_2", "")

	viper.SetDefault("62DOLPHIN_ADDRESS", "http://0.0.0.0:8090")
	viper.SetDefault("62WHALE_ADDRESS", "http://0.0.0.0:10081")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	err = viper.Unmarshal(data)
	if err != nil {
		fmt.Println(err)
	}

	return *data, err
}
