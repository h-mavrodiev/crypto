package configs

import (
	"github.com/spf13/viper"
)

type GateConfig struct {
	Host          string            `yaml:"host"`
	Prefix        string            `yaml:"prefix"`
	APIKey        string            `yaml:"apiKey"`
	APISecret     string            `yaml:"apiSecret"`
	Endpoints     GateEndpoints     `yaml:"endpoints"`
	CommonHeaders GateCommonHeaders `yaml:"commonHeaders"`
}

type GateEndpoints struct {
	Wallet string `yaml:"wallet"`
	Spot   string `yaml:"spot"`
	Margin string `yaml:"margin"`
}

type GateCommonHeaders struct {
	Accept      string `yaml:"accept"`
	ContentType string `yaml:"contentType"`
}

type StexConfig struct {
	Host          string            `yaml:"host"`
	APIKey        string            `yaml:"apiKey"`
	Endpoints     StexEndpoints     `yaml:"endpoints"`
	CommonHeaders StexCommonHeaders `yaml:"commonHeaders"`
}

type StexEndpoints struct {
	Public  string `yaml:"public"`
	Trading string `yaml:"trading"`
	Profile string `yaml:"profile"`
}

type StexCommonHeaders struct {
	Accept      string `yaml:"accept"`
	ContentType string `yaml:"contentType"`
}

// Needs updating when new platform is added
type Config struct {
	Gate GateConfig `yaml:"gate"`
	Stex StexConfig `yaml:"stex"`
}

var vp *viper.Viper

func LoadConfig(cn string, ct string, cp string) (Config, error) {
	vp = viper.New()
	var config Config

	vp.SetConfigName(cn)
	vp.SetConfigType(ct)
	vp.AddConfigPath(cp)
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
