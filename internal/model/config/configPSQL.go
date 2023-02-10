package config

import "fmt"

// ConfigPSQL PSQL config info
type ConfigPSQL struct {
	Name         string `toml:"db_name"`
	Pass         string `toml:"db_password"`
	User         string `toml:"db_user"`
	Type         string `toml:"db_type"`
	Host         string `toml:"db_host"`
	Port         string `toml:"db_port"`
	MaxOpenConst int    `toml:"db_SetMaxOpenConst"`
	MaxIdleConst int    `toml:"db_SetMaxIdleConst"`
}

// GetPSQLUrl URL for PSQL connect
func (conf *ConfigPSQL) GetPSQLUrl() string {
	return fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", conf.Host, conf.User, conf.Name, conf.Pass)
}
