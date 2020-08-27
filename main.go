package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mdemaiocwg/celestial-body-info/planet/planetpb"
	"google.golang.org/grpc"
)

// Change logic for closing cc later, we only need it when a request is made to be open, otherwise closed.
func connectToGRPCPlanet() (*grpc.ClientConn, planetpb.PlanetServiceClient) {
	fmt.Println("Planet Client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	c := planetpb.NewPlanetServiceClient(cc)
	return cc, c
}

// todo: Define an http server to handle requests from the client.
// This http server should fetch the data first from our microservices using gRPC.
// Afterwards we return the data back to the frontend user and display it formatted.
func main() {

	cc, c := connectToGRPCPlanet()
	defer cc.Close()
	// read Planet
	// fmt.Println("Reading the planet")

	// resRead, err := c.ReadPlanet(context.Background(), &planetpb.ReadPlanetRequest{PlanetId: "5f47c1ad093cc8a6577e0918"})
	// if err != nil {
	// 	fmt.Printf("Error happened while reading: %v \n", err)
	// }

	// fmt.Printf("Planet was read: %v \n", resRead)

	//list Planet
	// fmt.Println("Listing the planets")

	// resList, err := c.ListPlanet(context.Background(), &planetpb.ListPlanetRequest{})
	// if err != nil {
	// 	fmt.Printf("Error happened while listing: %v \n", err)
	// }

	// fmt.Printf("Planets were listed: %v \n", resList.GetPlanet())

	http.HandleFunc("/planet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Listing the planets")

		resList, err := c.ListPlanet(context.Background(), &planetpb.ListPlanetRequest{})
		if err != nil {
			fmt.Printf("Error happened while listing: %v \n", err)
		}
		fmt.Printf("Planets were listed: %v \n", resList)
		slcB, err := json.Marshal(resList.GetPlanet())
		if err != nil {
			fmt.Printf("Error happened while marshalling: %v \n", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(slcB)

	})
	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
