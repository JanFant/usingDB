package config

// GlobalConfig decode config toml
var GlobalConfig *Config

// Config struct all config toml element
type Config struct {
	PSQLConfig  ConfigPSQL  `toml:"PSQL"`
	MongoConfig ConfigMongo `toml:"Mongo"`
}

// NewConfig create GlobalConfig
func NewConfig() *Config {
	return &Config{}
}
