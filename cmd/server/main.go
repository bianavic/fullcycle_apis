package main

import "github.com/bianavic/fullcycle_apis/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
