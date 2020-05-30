package main

import (
	"fmt"
	"github.com/micro/go-micro/server"
	micro "github.com/micro/go-micro"
	"micro_test/config"
	"micro_test/controllers"
	proto "micro_test/proto"
)

func main() {
	// Create a new service. Optionally include some options here.
	/*service := micro.NewService(
		micro.Name("greeter"),
	)*/

	serviceName := "server"
	service := micro.NewService(
		micro.Name(serviceName),
		micro.Server(
			server.NewServer(
				server.Name(serviceName),
				server.Address(":" +config.Default("APP_PORT", "7777")),
			),
		),
	)


	// Init will parse the command line flags.
	service.Init()

	// Register handler
	proto.RegisterWaiterHandler(service.Server(), new(controllers.Server))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
