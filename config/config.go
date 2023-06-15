package config

import (
	"context"
	"log"
	"strings"

	"github.com/kamva/mgm/v3"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Setup(envPath string) {
	viper.SetConfigFile(envPath)
	viper.AutomaticEnv()
	viper.ReadInConfig()

	cmdMonitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			//log.Print(evt.Command)
		},
	}

	databaseUrl := strings.ReplaceAll(viper.GetString("DATABASEURL"), "mongodb://", "")

	connectionString := "mongodb://" + viper.GetString("DATABASEADMIN") + ":" + viper.GetString("DATABASEPASSWORD") + "@" + databaseUrl
	db := viper.GetString("DATABASE")
	err := mgm.SetDefaultConfig(nil, db, options.Client().ApplyURI(connectionString).SetMonitor(cmdMonitor))
	if err != nil {
		panic(err)
	}
	log.Println("Database", db)

}
