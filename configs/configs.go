package configs

import (
	"github.com/spf13/viper"
)

type GateConfig struct {
	Host          string        `yaml:"host"`
	Prefix        string        `yaml:"prefix"`
	APIKey        string        `yaml:"apiKey"`
	GateEndpoints GateEndpoints `yaml:"gate-endpoints"`
}

type GateEndpoints struct {
	Wallet string `yaml:"wallet"`
	Spot   string `yaml:"spot"`
	Margin string `yaml:"margin"`
}

type StexConfig struct {
	Host          string        `yaml:"host"`
	Prefix        string        `yaml:"prefix"`
	APIKey        string        `yaml:"apiKey"`
	StexEndpoints StexEndpoints `yaml:"stex-endpoints"`
}

type StexEndpoints struct {
	Public  string `yaml:"public"`
	Trading string `yaml:"trading"`
}

// Needs updating when new platform is added
type Config struct {
	Gate GateConfig `yaml:"gate"`
	Stex StexConfig `yaml:"stex"`
}

var vp *viper.Viper

func LoadConfig() (Config, error) {
	vp = viper.New()
	var config Config

	vp.SetConfigName("config")
	vp.SetConfigType("yml")
	vp.AddConfigPath("./configs")
	vp.AddConfigPath(".")
	err := vp.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	err = vp.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
