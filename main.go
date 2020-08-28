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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func handleGRPCErrors(err error) (int, []byte) {
	if e, ok := status.FromError(err); ok {
		switch e.Code() {
		case codes.Internal:
			return http.StatusInternalServerError, []byte(fmt.Sprintf(`{"error": "%v"}`, e.Message()))
		case codes.InvalidArgument:
			return http.StatusBadRequest, []byte(fmt.Sprintf(`{"error": "%v"}`, e.Message()))
		case codes.NotFound:
			return http.StatusNotFound, []byte(fmt.Sprintf(`{"error": "%v"}`, e.Message()))
		default:
			return http.StatusInternalServerError, []byte(fmt.Sprintf(`{"error": "%v"}`, e.Message()))

		}
	} else {
		return http.StatusInternalServerError, []byte(fmt.Sprintf(`{"error": "%v"}`, "Unknown internal error occured."))
	}
}

func listPlanetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Listing the planets")
	cc, c := connectToGRPCPlanet()
	defer cc.Close()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	resList, err := c.ListPlanet(context.Background(), &planetpb.ListPlanetRequest{})
	if err != nil { // Handle our gRPC errors.
		fmt.Printf("Error happened while listing: %v \n", err)
		code, errJSON := handleGRPCErrors(err)
		w.WriteHeader(code)
		w.Write(errJSON)
		return
	}
	fmt.Printf("Planets were listed: %v \n", resList)
	slcB, err := json.Marshal(resList.GetPlanet())
	if err != nil {
		fmt.Printf("Error happened while marshalling: %v \n", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(slcB)
}

func readPlanetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("reading a planet")
	cc, c := connectToGRPCPlanet()
	defer cc.Close()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	resRead, err := c.ReadPlanet(context.Background(), &planetpb.ReadPlanetRequest{PlanetId: vars["id"]})
	if err != nil { // Handle our gRPC errors.
		fmt.Printf("Error happened while reading: %v \n", err)
		code, errJSON := handleGRPCErrors(err)
		w.WriteHeader(code)
		w.Write(errJSON)
		return
	}

	fmt.Printf("Planet was read: %v \n", resRead)
	slcB, err := json.Marshal(resRead.GetPlanet())
	if err != nil {
		fmt.Printf("Error happened while marshalling: %v \n", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(slcB)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/planet", listPlanetHandler).Methods(http.MethodGet)
	r.HandleFunc("/planet/{id}", readPlanetHandler).Methods(http.MethodGet)
	r.Use(mux.CORSMethodMiddleware(r))

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
