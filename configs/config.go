package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
	Database struct {
		Driver   string `mapstructure:"driver"`
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
		Charset  string `mapstructure:"charset"`
		Location string `mapstructure:"location"`
	} `mapstructure:"database"`
	Secret struct {
		TokenKey string `mapstructure:"tokenKey"`
	} `mapstructure:"secret"`
	Log struct {
		Level string `mapstructure:"level"`
		File  string `mapstructure:"file"`
	} `mapstructure:"log"`
}

// 加载配置文件
func LoadConfig(env string) (*Config, error) {
	// 根据环境变量或命令行参数确定要加载的配置文件
	configFileName := fmt.Sprintf("config.%s.yaml", env)

	viper.SetConfigName(configFileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %w", err)
	}
	return &config, nil
}
