package main

import (
	"fmt"

	"github.com/bvaledev/go-expert-commerce-api/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	fmt.Println(config.DBDriver)
}
