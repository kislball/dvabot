package gateway

import (
	"fmt"
	"github.com/kislball/dva/pkg/config"
)

func Start() {
	cfg := config.Config{}
	err := cfg.Load("./configs/config.json")
	if err != nil {
		panic(err)
	}
	fmt.Println("hello, world!", cfg.Database.Driver, "lol")
}
