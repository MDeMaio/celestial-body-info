package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"gopkg.in/mgo.v2/bson"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/mdemaio/celestial-body-info/planet/planetpb"
	"google.golang.org/grpc"
)

var collection *mongo.Collection

type server struct {
}

type facts struct {
	Title string `bson:"title"`
	Fact  string `bson:"fact"`
}

type basicInformation struct {
	AlternateName        string `bson:"alternate_name"`
	NumberOfSatelites    int32  `bson:"number_of_satelites"`
	StarSystem           string `bson:"star_system"`
	MostAbundantResource string `bson:"most_abundant_resource"`
}

type planetItem struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	Name             string             `bson:"name"`
	Facts            []facts            `bson:"facts"`
	Image            string             `bson:"image"`
	BasicInformation basicInformation   `bson:"basic_information"`
}

func (*server) ReadPlanet(ctx context.Context, req *planetpb.ReadPlanetRequest) (*planetpb.ReadPlanetResponse, error) {
	fmt.Println("Read planet request")

	planetID := req.GetPlanetId()
	oid, err := primitive.ObjectIDFromHex(planetID)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	// create an empty struct
	data := &planetItem{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(context.Background(), filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find planet with specified ID: %v", err),
		)
	}

	return &planetpb.ReadPlanetResponse{
		Planet: dataToPlanetPb(data),
	}, nil
}

func (*server) ListPlanet(ctx context.Context, req *planetpb.ListPlanetRequest) (*planetpb.ListPlanetResponse, error) {
	fmt.Println("List planet request")

	filter := bson.M{}
	planets := []*planetpb.Planet{}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		data := &planetItem{}
		err := cursor.Decode(data)
		if err != nil {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding from mongodb: %v", err),
			)
		}
		planets = append(planets, dataToPlanetPb(data))
	}

	return &planetpb.ListPlanetResponse{
		Planet: planets,
	}, nil
}

func dataToPlanetPb(data *planetItem) *planetpb.Planet {
	facts := []*planetpb.Facts{} // Not sure if this is correct.
	for _, v := range data.Facts {
		fact := &planetpb.Facts{
			Title: v.Title,
			Fact:  v.Fact,
		}

		facts = append(facts, fact)
	}

	basicInformation := &planetpb.BasicInformation{
		AlternateName:        data.BasicInformation.AlternateName,
		NumberOfSatelites:    data.BasicInformation.NumberOfSatelites,
		StarSystem:           data.BasicInformation.StarSystem,
		MostAbundantResource: data.BasicInformation.MostAbundantResource,
	}

	return &planetpb.Planet{
		PlanetId:         data.ID.Hex(),
		Name:             data.Name,
		Facts:            facts,
		Image:            data.Image,
		BasicInformation: basicInformation,
	}
}

func main() {
	// if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Connecting to MongoDB")
	// connect to MongoDB
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}
	fmt.Println(uri)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Planet Service Started")
	collection = client.Database("celestial-body-info").Collection("planet")

	port := os.Getenv("PORT")
	if port == "" {
		port = ":50051"
	} else {
		port = ":" + port
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Println("Up and running on: " + port)

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	planetpb.RegisterPlanetServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch
	// First we close the connection with MongoDB:
	fmt.Println("Closing MongoDB Connection")
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatalf("Error on disconnection with MongoDB : %v", err)
	}
	// Second step : closing the listener
	fmt.Println("Closing the listener")
	if err := lis.Close(); err != nil {
		log.Fatalf("Error on closing the listener : %v", err)
	}
	// Finally, we stop the server
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("End of Program")
}
