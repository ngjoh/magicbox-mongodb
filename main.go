package main

import (
	"github.com/koksmat-com/koksmat/cmd"
	"github.com/koksmat-com/koksmat/config"
)

func main() {

	config.Setup(".env")
	cmd.Execute()
}
