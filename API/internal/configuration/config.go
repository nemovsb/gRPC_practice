package configuration

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type AppConfig struct {
	HTTP    `mapstructure:"http"`
	GRPC    `mapstructure:"grpc"`
	ZapMode string `mapstructure:"zap_mode"`
}

type HTTP struct {
	Port string `mapstructure:"port"`
}

type GRPC struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Configurator interface {
	NewAppConfig() *AppConfig
}

const EnvProduction = "Production"

var ErrUnmarshalConfig = errors.New("viper failed to unmarshal app config")

func ViperConfigurationProvider(env string, writeConfig bool) (cfg *AppConfig, err error) {
	var filename string

	switch env {
	case "Production":
		filename = "config"
	default:
		filename = "config"
	}

	v := NewViper(filename)

	cfg, err = NewConfiguration(v)
	if err != nil {
		return
	}

	if writeConfig {
		if err = v.WriteConfig(); err != nil {
			log.Println("viper failed to write app config file:", err)
		}
	}

	return cfg, nil
}

func NewViper(filename string) *viper.Viper {
	v := viper.New()

	if filename != "" {
		v.SetConfigName(filename)
		v.AddConfigPath(".")
		v.AddConfigPath(filepath.FromSlash("./build/cfg/"))
	}

	v.SetDefault("HTTP.port", "8081")
	v.SetDefault("GRPC.port", "50051")
	v.SetDefault("GRPC.host", "0.0.0.0")

	v.SetDefault("zapMode", "production")

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetEnvPrefix("MYAPP")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		log.Println("viper failed to read app config file:", err)
	}

	return v
}

func NewConfiguration(v *viper.Viper) (*AppConfig, error) {
	var c AppConfig
	if err := v.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrUnmarshalConfig, err)
	}

	fmt.Printf("My config: %+v\n", c)

	return &c, nil
}
