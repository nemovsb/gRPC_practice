package configuration

type AppConfig struct {
	HTTP
	GRPC
	zapMode string
}

type HTTP struct {
	port string `mapstructure:"port"`
}

type GRPC struct {
	port string `mapstructure:"port"`
}

type Configurator interface {
	NewAppConfig() *AppConfig
}
