package config

import (
	"time"
)

type System struct {
	Name string `env:"SYSTEM_NAME" env-default:"Unchained" yaml:"name"`
	Log  string `env:"SYSTEM_LOG"  env-default:"info"      yaml:"log"`
}

type RPC struct {
	Name  string   `yaml:"name"`
	Nodes []string `yaml:"nodes"`
}

type Uniswap struct {
	Schedule    map[string]time.Duration `yaml:"schedule"`
	Tokens      []Token                  `yaml:"tokens"`
	Correctness []string                 `yaml:"correctness"`
}

type EthLog struct {
	Schedule    map[string]time.Duration `yaml:"schedule"`
	Events      []Event                  `yaml:"events"`
	Correctness []string                 `yaml:"correctness"`
}

type Plugins struct {
	EthLog  *EthLog  `yaml:"logs"`
	Uniswap *Uniswap `yaml:"uniswap"`
}

type Event struct {
	Name          string  `yaml:"name"`
	Chain         string  `yaml:"chain"`
	Abi           string  `yaml:"abi"`
	Event         string  `yaml:"event"`
	Address       string  `yaml:"address"`
	From          *uint64 `yaml:"from"`
	Step          uint64  `yaml:"step"`
	Confirmations uint64  `yaml:"confirmations"`
	Store         bool    `yaml:"store"`
	Send          bool    `yaml:"send"`
}

type Token struct {
	Name   string `yaml:"name"`
	Pair   string `yaml:"pair"`
	Chain  string `yaml:"chain"`
	Delta  int64  `yaml:"delta"`
	Invert bool   `yaml:"invert"`
	Unit   string `yaml:"unit"`
	Send   bool   `yaml:"send"`
	Store  bool   `yaml:"store"`
}

type ProofOfStack struct {
	Chain   string `env:"POS_CHAIN"   env-default:"arbitrum_sepolia"                           yaml:"chain"`
	Address string `env:"POS_ADDRESS" env-default:"0x965e364987356785b7E89e2Fe7B70f5E5107332d" yaml:"address"`
	Base    int64  `env:"POS_BASE"    env-default:"1"                                          yaml:"base"`
}

type Broker struct {
	Bind string `env:"BROKER_BIND" env-default:"0.0.0.0:9123"                    yaml:"bind"`
	URI  string `env:"BROKER_URI"  env-default:"wss://shinobi.brokers.kenshi.io" yaml:"uri"`
}

type Postgres struct {
	URL string `env:"DATABASE_URL" yaml:"url"`
}

type Secret struct {
	Address   string `env:"ADDRESS"    yaml:"address"`
	EvmWallet string `env:"ENM_WALLET" yaml:"evm-wallet"`
	SecretKey string `env:"SECRET_KEY" yaml:"secret-key"`
	PublicKey string `env:"PUBLIC_KEY" yaml:"public-key"`
}

type Config struct {
	System       System       `yaml:"system"`
	Broker       Broker       `yaml:"broker"`
	RPC          []RPC        `yaml:"rpc"`
	Postgres     Postgres     `yaml:"postgres"`
	ProofOfStack ProofOfStack `yaml:"pos"`
	Plugins      Plugins      `yaml:"plugins"`
	Secret       Secret       `yaml:"secret"`
}
