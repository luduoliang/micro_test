package main

import (
	"context"
	"fmt"

	proto "micro_test/proto"
	micro_client "github.com/micro/go-micro/client"
)


func main() {
	// Create a new service. Optionally include some options here.
	//service := micro.NewService(micro.Name("server.client"))
	/*serviceName := "server.client"
	service := micro.NewService(
		micro.Name(serviceName),
		micro.Server(
			server.NewServer(
				server.Name(serviceName),
				server.Address(":7777"),
			),
		),
	)

	service.Init()*/
	var cli  micro_client.Client
	// Create new greeter client
	server := proto.NewWaiterService("server", cli)

	// Call the greeter
	//rsp, err := server.GetPddSessionsInfo(context.TODO(), &proto.RequestGetPddSessionsInfo{Id: 20})
	rsp, err := server.GetPddSessionsList(context.TODO(), &proto.RequestGetPddSessionsList{Page: 1, PerPage:2})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(fmt.Sprintf("%+v", rsp))

}