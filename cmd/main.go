package main

import (
	"lessons/casbin/api"

	"github.com/casbin/casbin/v2"
)

func main() {

	// casbin with CSV -------------------------------------------
	casbinEnforcer, err := casbin.NewEnforcer("./config/auth.conf", "./config/auth.csv")
	if err != nil {
		println("casbin enforcer error")
		return
	}

	router := api.New(api.Option{
		Enforcer: casbinEnforcer,
	})

	router.Run(":8000")

}
