package configs

import (
	"github.com/spf13/viper"
)

// ApplicationConfig holds the application configuration
type ApplicationConfig struct {
	Port int
}

// appCfg is the default application configuration
var appCfg ApplicationConfig

// App returns the default application configuration
func App() ApplicationConfig {
	return appCfg
}

// LoadApp loads application configuration
func LoadApp() {
	appCfg = ApplicationConfig{
		Port: viper.GetInt("app.port"),
	}
}
