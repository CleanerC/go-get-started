package main

import (
	"fmt"
	v1_endpoint "BasicUsedHTTP/pkg/v1/endpoint"
	v1_service "BasicUsedHTTP/pkg/v1/service"
	v1_transport "BasicUsedHTTP/pkg/v1/transport"
	"net/http"
)

func main() {
	server := v1_service.NewService()
	endpoint := v1_endpoint.NewEndpointServer(server)
	httpHandler := v1_transport.NewHttpHandle(endpoint)
	fmt.Println("Server started at http://localhost:8888")
	_ = http.ListenAndServe("localhost:8888", httpHandler)
}


