package configs

type GateConfig struct {
	Host          string            `yaml:"host"`
	Prefix        string            `yaml:"prefix"`
	APIKey        string            `yaml:"apiKey"`
	APISecret     string            `yaml:"apiSecret"`
	Endpoints     GateEndpoints     `yaml:"endpoints"`
	CommonHeaders GateCommonHeaders `yaml:"commonHeaders"`
	Pair          string            `yaml:"pair"`
	MinTrade      float64           `yaml:"minTrade"`
	WS            WS                `yaml:"ws"`
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
	Pair          int               `yaml:"pair"`
	MinTrade      float64           `yaml:"minTrade"`
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

type WS struct {
	Schema string `yaml:"schema"`
	WSHost string `yaml:"wshost"`
	Path   string `yaml:"path"`
}
