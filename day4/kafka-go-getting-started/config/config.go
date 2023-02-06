package config

import "github.com/spf13/viper"

var AppConfig Config

type Config struct {
	KafkaURL      string `mapstructure:"KAFKA_URL"`
	KafkaUser     string `mapstructure:"KAFKA_USER"`
	KafkaPassword string `mapstructure:"KAFKA_PASSWORD"`
}

func LoadConfig(path string) (config Config) {
	viper.SetDefault("KAFKA_URL", "localhost:9092")
	viper.SetDefault("KAFKA_USER", "user")
	viper.SetDefault("KAFKA_PASSWORD", "password")

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.ReadInConfig()
	viper.Unmarshal(&config)
	AppConfig = config
	return AppConfig
}
