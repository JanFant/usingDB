package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"time"
	"usingDB/internal/model/config"
	"usingDB/internal/model/dbases/mongodb"
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

	/*
		//PSQL
			fmt.Println("PSQL")
			connPSQL, err := postgreSQL.ConnectPSQLD()
			if err != nil {
				fmt.Println("PSQL db - err: ", err.Error())
				os.Exit(1)
			}
			defer func() {
				if err := connPSQL.Close(); err != nil {
					fmt.Println("close db err", err.Error())
					panic(err)
				}
			}()
			postgreSQL.PSQLExamples(connPSQL)
	*/

	fmt.Println("mongo")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongodb.ConnectMongo(ctx)
	if err != nil {
		fmt.Println("Mongo db -err", err.Error())
		os.Exit(1)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			fmt.Println("close db err", err.Error())
			panic(err)
		}
	}()
	mongodb.MongoExamples(client)

	fmt.Println("done")
}
