package configs

import "github.com/spf13/viper"

type Config struct {
	ServerConfig struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	DatabaseConfig struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		User string `yaml:"user"`
	} `yaml:"database"`
	JWTConfig struct {
		SecretKey  string `yaml:"secret_key"`
		Expiration int64  `yaml:"expiration"`
	} `yaml:"jwt"`
}

func NewConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
