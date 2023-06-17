package config

import (
	"context"
	"strings"

	"github.com/kamva/mgm/v3"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnectionString() string {
	databaseUrl := strings.ReplaceAll(viper.GetString("DATABASEURL"), "mongodb://", "")
	return "mongodb://" + viper.GetString("DATABASEADMIN") + ":" + viper.GetString("DATABASEPASSWORD") + "@" + databaseUrl
}

func DatabaseName() string {
	return viper.GetString("DATABASE")
}
func Setup(envPath string) {
	viper.SetConfigFile(envPath)
	viper.AutomaticEnv()
	viper.ReadInConfig()

	cmdMonitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			// log.Print(evt.Command)
		},
	}

	db := DatabaseName()
	err := mgm.SetDefaultConfig(nil, db, options.Client().ApplyURI(MongoConnectionString()).SetMonitor(cmdMonitor))
	if err != nil {
		panic(err)
	}

}
