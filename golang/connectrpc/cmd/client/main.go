package main

import (
	"context"
	"io"
	greetv1 "mukezhz/connectrpc/gen/greet/v1"
	"mukezhz/connectrpc/gen/greet/v1/greetv1connect"
	"net/http"
	"strings"

	"connectrpc.com/connect"
)

func main() {
	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
		connect.WithGRPC(),
	)
	res, err := client.Greet(context.Background(), &connect.Request[greetv1.GreetRequest]{
		Msg: &greetv1.GreetRequest{
			Name: "Alice",
		},
	})
	if err != nil {
		panic(err)
	}

	greeting := res.Msg.Greeting
	println(greeting)

	// Send a raw HTTP request to the server.
	body := strings.NewReader(`{
		"name": "WOW"
	}`)
	request, _ := http.NewRequest("POST", "http://localhost:8080/greet.v1.GreetService/Greet", body)
	request.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	println(response.Status)
	defer response.Body.Close()

	data, _ := io.ReadAll(response.Body)
	println(string(data))

}
