package main

import (
	"fmt"

	"github.com/JREAMLU/core/consul"
	consulApi "github.com/hashicorp/consul/api"
)

func main() {
	address := consul.SetAddress("10.211.55.5:8500")
	register := consul.SetRegister(&consulApi.AgentServiceRegistration{
		Name: "jream",
		Tags: []string{"luj"},
		Port: 8080,
		Check: &consulApi.AgentServiceCheck{
			TTL: "15s",
		},
	})
	c, err := consul.NewClient(address, register)
	if err != nil {
		panic(err)
	}
	err = c.Put("mysql", "select")
	// res, err := c.Get("mulu1/key")
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println("res", res)
}
