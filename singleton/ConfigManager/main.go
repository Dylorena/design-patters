package main

import (
	"config-manager/infra"
	"fmt"
)

func main() {
	config1 := infra.GetConfigManager()
	fmt.Println(config1.Get("DB_HOST"))

	config2 := infra.GetConfigManager()
	fmt.Println(config2.Get("APP_ENV"))

	fmt.Println("Same instance?", config1 == config2) // Need to be true

}
