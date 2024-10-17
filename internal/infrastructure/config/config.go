package config

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var Config appConfig

type appConfig struct {
	Env            string        `mapstructure:"env" validate:"required"`
	LogLevel       string        `mapstructure:"log_level"`
	ServerAddress  string        `mapstructure:"server_address" validate:"required"`
	AllowedOrigins []string      `mapstructure:"allowed_origins" validate:"required"`
	JWTSecretKey   string        `mapstructure:"jwt_secret_key" validate:"required"`
	RequestTimeout time.Duration `mapstructure:"request_timeout" validate:"required"`
}

func init() {
	v := viper.New()
	v.SetDefault("env", "dev")
	v.SetDefault("log_level", "INFO")
	v.SetDefault("server_address", ":8081")
	v.SetDefault("allowed_origins", []string{"*"})
	v.SetDefault("jwt_secret_key", "jwt-secret")
	v.SetDefault("request_timeout", 180*time.Second)

	viper.AutomaticEnv()
	if err := v.Unmarshal(&Config); err != nil {
		panic(err)
	}
}

// Validate validates the config values.
func (c *appConfig) Validate() error {
	return validator.New().Struct(c)
}
