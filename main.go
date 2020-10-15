package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mdemaio/celestial-body-info/nasa/nasapb"
	"github.com/mdemaio/celestial-body-info/planet/planetpb"
	"github.com/mdemaio/celestial-body-info/star/starpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type planetListResponse struct {
	Planets           []*planetpb.Planet `json:"planets"`
	NumberOfDocuments int64              `json:"number_of_documents"`
}

type starListResponse struct {
	Stars             []*starpb.Star `json:"stars"`
	NumberOfDocuments int64          `json:"number_of_documents"`
}

type planetListTypeResponse struct {
	PlanetType []string `json:"planet_type"`
}

// Change logic for closing cc later, we only need it when a request is made to be open, otherwise closed.
func connectToGRPCPlanet() (*grpc.ClientConn, planetpb.PlanetServiceClient) {
	fmt.Println("Planet Client")

	opts := grpc.WithInsecure()

	ip := os.Getenv("PLANET_PORT")
	if ip == "" {
		ip = "0.0.0.0"
	}
	port := ip + ":50051"

	cc, err := grpc.Dial(port, opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	c := planetpb.NewPlanetServiceClient(cc)
	return cc, c
}

func connectToGRPCStar() (*grpc.ClientConn, starpb.StarServiceClient) {
	fmt.Println("Star Client")

	opts := grpc.WithInsecure()

	ip := os.Getenv("STAR_PORT")
	if ip == "" {
		ip = "0.0.0.0"
	}
	port := ip + ":50052"

	cc, err := grpc.Dial(port, opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	c := starpb.NewStarServiceClient(cc)
	return cc, c
}

func connectToGRPCNASA() (*grpc.ClientConn, nasapb.NasaServiceClient) {
	fmt.Println("NASA Client")

	opts := grpc.WithInsecure()

	ip := os.Getenv("NASA_PORT")
	if ip == "" {
		ip = "0.0.0.0"
	}
	port := ip + ":50053"

	cc, err := grpc.Dial(port, opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	c := nasapb.NewNasaServiceClient(cc)
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

	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["page"])
	if err != nil {
		fmt.Printf("Error happened while converting: %v \n", err)
	}

	filters := []*planetpb.ListPlanetRequestFilter{
		&planetpb.ListPlanetRequestFilter{
			Column: "basic_information.type",
			Value:  vars["type"],
		},
		&planetpb.ListPlanetRequestFilter{
			Column: "name",
			Value:  vars["name"],
		},
	}

	resList, err := c.ListPlanet(context.Background(), &planetpb.ListPlanetRequest{
		Skip:                    int64((page * 5) - 5),
		ListPlanetRequestFilter: filters,
	})
	if err != nil { // Handle our gRPC errors.
		fmt.Printf("Error happened while listing: %v \n", err)
		code, errJSON := handleGRPCErrors(err)
		w.WriteHeader(code)
		w.Write(errJSON)
		return
	}
	//fmt.Printf("Planets were listed: %v \n", resList)
	res := planetListResponse{
		Planets:           resList.GetPlanet(),
		NumberOfDocuments: resList.GetNumberOfDocuments(),
	}
	slcB, err := json.Marshal(res)
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
	resRead, err := c.ReadPlanet(context.Background(), &planetpb.ReadPlanetRequest{Name: vars["name"]})
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

func listPlanetTypeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Listing the planet types")
	cc, c := connectToGRPCPlanet()
	defer cc.Close()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	resList, err := c.ListPlanetType(context.Background(), &planetpb.ListPlanetTypeRequest{})
	if err != nil { // Handle our gRPC errors.
		fmt.Printf("Error happened while listing: %v \n", err)
		code, errJSON := handleGRPCErrors(err)
		w.WriteHeader(code)
		w.Write(errJSON)
		return
	}

	res := planetListTypeResponse{
		PlanetType: resList.GetPlanetType(),
	}
	slcB, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("Error happened while marshalling: %v \n", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(slcB)
}

func listStarHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Listing the stars")
	cc, c := connectToGRPCStar()
	defer cc.Close()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["page"])
	if err != nil {
		fmt.Printf("Error happened while converting: %v \n", err)
	}

	filters := []*starpb.ListStarRequestFilter{
		&starpb.ListStarRequestFilter{
			Column: "basic_information.classification",
			Value:  vars["classification"],
		},
		&starpb.ListStarRequestFilter{
			Column: "name",
			Value:  vars["name"],
		},
	}

	resList, err := c.ListStar(context.Background(), &starpb.ListStarRequest{
		Skip:                  int64((page * 5) - 5),
		ListStarRequestFilter: filters,
	})
	if err != nil { // Handle our gRPC errors.
		fmt.Printf("Error happened while listing: %v \n", err)
		code, errJSON := handleGRPCErrors(err)
		w.WriteHeader(code)
		w.Write(errJSON)
		return
	}
	//fmt.Printf("Planets were listed: %v \n", resList)
	res := starListResponse{
		Stars:             resList.GetStar(),
		NumberOfDocuments: resList.GetNumberOfDocuments(),
	}
	slcB, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("Error happened while marshalling: %v \n", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(slcB)
}

func readStarHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("reading a star")
	cc, c := connectToGRPCStar()
	defer cc.Close()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	resRead, err := c.ReadStar(context.Background(), &starpb.ReadStarRequest{Name: vars["name"]})
	if err != nil { // Handle our gRPC errors.
		fmt.Printf("Error happened while reading: %v \n", err)
		code, errJSON := handleGRPCErrors(err)
		w.WriteHeader(code)
		w.Write(errJSON)
		return
	}

	fmt.Printf("Star was read: %v \n", resRead)
	slcB, err := json.Marshal(resRead.GetStar())
	if err != nil {
		fmt.Printf("Error happened while marshalling: %v \n", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(slcB)
}

func listNASAAPODHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("listing all apods")
	cc, c := connectToGRPCNASA()
	defer cc.Close()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	resRead, err := c.ListAPOD(context.Background(), &nasapb.ListAPODRequest{})
	if err != nil { // Handle our gRPC errors.
		fmt.Printf("Error happened while reading: %v \n", err)
		code, errJSON := handleGRPCErrors(err)
		w.WriteHeader(code)
		w.Write(errJSON)
		return
	}

	fmt.Printf("apod was read: %v \n", resRead)
	slcB, err := json.Marshal(resRead.GetApod())
	if err != nil {
		fmt.Printf("Error happened while marshalling: %v \n", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(slcB)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/planet/{page}/{type}/{name}", listPlanetHandler).Methods(http.MethodGet)
	r.HandleFunc("/planet/{name}", readPlanetHandler).Methods(http.MethodGet)
	r.HandleFunc("/planettype", listPlanetTypeHandler).Methods(http.MethodGet)
	r.HandleFunc("/star/{page}/{classification}/{name}", listStarHandler).Methods(http.MethodGet)
	r.HandleFunc("/star/{name}", readStarHandler).Methods(http.MethodGet)
	r.HandleFunc("/apod", listNASAAPODHandler).Methods(http.MethodGet)
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
