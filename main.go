package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/mdemaio/celestial-body-info/planet/planetpb"
	"google.golang.org/grpc"
)

// Change logic for closing cc later, we only need it when a request is made to be open, otherwise closed.
func connectToGRPCPlanet() (*grpc.ClientConn, planetpb.PlanetServiceClient) {
	fmt.Println("Planet Client")

	opts := grpc.WithInsecure()

	port := os.Getenv("PORT")
	if port == "" {
		port = "0.0.0.0:50051"
	} else {
		port = "0.0.0.0:80"
	}

	cc, err := grpc.Dial(port, opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	c := planetpb.NewPlanetServiceClient(cc)
	return cc, c
}

func listPlanetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Listing the planets")
	cc, c := connectToGRPCPlanet()
	defer cc.Close()

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
}

func readPlanetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("reading a planet")
	cc, c := connectToGRPCPlanet()
	defer cc.Close()

	vars := mux.Vars(r)
	resRead, err := c.ReadPlanet(context.Background(), &planetpb.ReadPlanetRequest{PlanetId: vars["id"]})
	if err != nil {
		fmt.Printf("Error happened while reading: %v \n", err)
	}
	fmt.Printf("Planet was read: %v \n", resRead)
	slcB, err := json.Marshal(resRead.GetPlanet())
	if err != nil {
		fmt.Printf("Error happened while marshalling: %v \n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(slcB)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/planet", listPlanetHandler)
	r.HandleFunc("/planet/{id}", readPlanetHandler)

	log.Println("Listening for http requests.")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
