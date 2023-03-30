package config

type config struct {
	Debug         bool        `yaml:"Debug"`
	DB            DB          `yaml:"DB"`
	Valut         ValutConfig `yaml:"Valut"`
	ChainCfg      []ChainCfg  `yaml:"ChainCfg"`
	EtcdEndpoints []string    `yaml:"EtcdEndpoints"`
}

type DB struct {
	Driver string `yaml:"Driver"`
	Source string `yaml:"Source"`
	Debug  bool   `yaml:"Debug"`
}

type ChainCfg struct {
	NetworkId         int    `yaml:"NetworkId"`
	AdminAddress      string `yaml:"AdminAddress"`
	RpcUrl            string `yaml:"RpcUrl"`
	Graph             string `yaml:"Graph"`
	BlockStep         int    `yaml:"BlockStep"`
	BlockDelay        int    `yaml:"BlockDelay"`
	BridgeContract721 string `yaml:"BridgeContract721"`
	BridgeContract20  string `yaml:"BridgeContract20"`
}

type ValutConfig struct {
	Endpoint string `yaml:"Endpoint"`
	Role     string `yaml:"Role"`
	Secret   string `yaml:"Secret"`
}
