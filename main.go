package main

import (
	"simple-golang-ewallet/api"
	"simple-golang-ewallet/config"
	"sync"
)

func main() {
	//Init Config
	config.SetConfig()

	var wg sync.WaitGroup
	wg.Add(1)

	//Start HTTP / REST Server
	go api.StartHttpServer()

	wg.Wait()
}
