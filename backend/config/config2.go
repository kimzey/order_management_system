package config

//
//import (
//	"github.com/go-playground/validator/v10"
//	"strings"
//	"time"
//
//	"github.com/spf13/viper"
//)
//
//type (
//	Database struct {
//		Host     string `mapstructure:"host" validate:"required"`
//		Port     int    `mapstructure:"port" validate:"required"`
//		User     string `mapstructure:"user" validate:"required"`
//		Password string `mapstructure:"password" validate:"required"`
//		DBName   string `mapstructure:"dbname" validate:"required"`
//		SSLMode  string `mapstructure:"sslmode" validate:"required"`
//		Schema   string `mapstructure:"schema" validate:"required"`
//	}
//
//	Server struct {
//		Port         int           `mapstructure:"port" validate:"required"`
//		AllowOrigins []string      `mapstructure:"allowOrigins" validate:"required"`
//		Timeout      time.Duration `mapstructure:"timeout" validate:"required"`
//		BodyLimit    string        `mapstructure:"bodyLimit" validate:"required"`
//	}
//
//	Config struct {
//		Database *Database `mapstructure:"database" validate:"required"`
//		Server   *Server   `mapstructure:"server" validate:"required"`
//	}
//)
//
//func GettingConfig() *Config {
//	viper.SetConfigName("config")
//	viper.SetConfigType("yaml")
//	viper.AddConfigPath("./config")
//
//	//comment for migration database
//	//viper.AddConfigPath("../../config")
//
//	viper.AutomaticEnv()
//	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
//
//	configInstance := &Config{}
//
//	if err := viper.ReadInConfig(); err != nil {
//		panic(err)
//	}
//
//	if err := viper.Unmarshal(configInstance); err != nil {
//		panic(err)
//	}
//
//	validate := validator.New()
//
//	if err := validate.Struct(configInstance); err != nil {
//		panic(err)
//	}
//	return configInstance
//}
