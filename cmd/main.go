package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/jmoiron/sqlx"
	"os"
	"usingDB/internal/model/config"
	"usingDB/internal/model/dbases/postgreSQL"
)

var err error

func init() {
	var configPath string
	flag.StringVar(&configPath, "config-path", "configs/config.toml", "path to config file")

	config.GlobalConfig = config.NewConfig()

	if _, err = toml.DecodeFile(configPath, &config.GlobalConfig); err != nil {
		fmt.Println("Can't load config file : ", err.Error())
		os.Exit(1)
	}

}

func main() {

	connPSQL, err := postgreSQL.ConnectPSQLD()
	if err != nil {
		fmt.Println("PSQL sb - err: ", err.Error())
		os.Exit(1)
	}
	defer func(connPSQL *sqlx.DB) {
		if err := connPSQL.Close(); err != nil {
			fmt.Println("close db err", err.Error())
		}
	}(connPSQL)

	postgreSQL.PSQLExamples(connPSQL)

	fmt.Println("done")
}
