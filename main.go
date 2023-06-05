package main

import (
	"github.com/koksmat-com/koksmat/cmd"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.ReadInConfig()
	cmd.Execute()
}
