package config

import (
	"fmt"
	"os"
	"time"

	"app.io/pkg/logHandler"
	"github.com/spf13/viper"
)

// func Init() error {
// 	fmt.Println("We can start config here")
// }

type Config struct {
	Database struct {
		Type   string `mapstructure:"type"`
		DbUri  string `mapstructure:"uri"`
		DbName string `mapstructure:"name"`
	} `mapstructure:"database"`

	Server struct {
		Env             string `mapstructure:"env"`
		ServerHost      string `mapstructure:"host"`
		ServerPort      string `mapstructure:"port"`
		Prefix          string `mapstructure:"prefix"`
		CronJobInterval int    `mapstructure:"cron_job_interval"`
		NetworkTitle    string `mapstructure:"network"`

		AuthenticationPlugin string `mapstructure:"authentication_plugin"`

		ListenLimit  int           `mapstructure:"listenLimit"`
		KeepAlive    time.Duration `mapstructure:"keepAlive"`
		ReadTimeout  time.Duration `mapstructure:"readTimeout"`
		WriteTimeout time.Duration `mapstructure:"writeTimeout"`

		TatumApiToken string `mapstructure:"tatum_api_token"`

		EnabledListeners []string      `mapstructure:"enabledListeners"`
		CleanupTimeout   time.Duration `mapstructure:"cleanupTimeout"`
		GracefulTimeout  time.Duration `mapstructure:"gracefulTimeout"`
		MaxHeaderSize    int           `mapstructure:"maxHeaderSize"`

		TlsHost           string        `mapstructure:"tlsHost"`
		TlsPort           int           `mapstructure:"tlsPort"`
		TlsListenLimit    int           `mapstructure:"tlsListenLimit"`
		TlsKeepAlive      time.Duration `mapstructure:"tlsKeepAlive"`
		TlsReadTimeout    time.Duration `mapstructure:"tlsReadTimeout"`
		TlsWriteTimeout   time.Duration `mapstructure:"tlsWriteTimeout"`
		TlsCertificate    string        `mapstructure:"tlsCertificateKey"`
		TlsCertificateKey string        `mapstructure:"tlsCertificateKey"`
		TlsCACertificate  string        `mapstructure:"tlsCACertificate"`
	} `mapstructure:"server"`
}

func NewConfig() *Config {
	cfg := &Config{}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		logHandler.Log(logHandler.ERROR, fmt.Sprintf("Fatal error config file: %v \n", err))
		os.Exit(1)
	}

	err = viper.Unmarshal(cfg)
	if err != nil {
		logHandler.Log(logHandler.ERROR, fmt.Sprintf("unable to decode into config struct, %v", err))
		os.Exit(1)
	}

	return cfg
}

func GetDbDSN(cfg Config) string {
	return fmt.Sprintf("%s/%s", cfg.Database.DbUri, cfg.Database.DbName)
}
