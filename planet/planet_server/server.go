package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	"github.com/mdemaio/celestial-body-info/util"
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
	AlternateName        string  `bson:"alternate_name"`
	Type                 string  `bson:"type"`
	NumberOfSatelites    int32   `bson:"number_of_satelites"`
	StarSystem           string  `bson:"star_system"`
	MostAbundantResource string  `bson:"most_abundant_resource"`
	SurfaceGravity       float64 `bson:"surface_gravity"`
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

	planetName := req.GetName()

	// create an empty struct
	data := &planetItem{}
	filter := bson.M{"name": primitive.Regex{Pattern: fmt.Sprintf("^.*%s.*", planetName), Options: "i"}} // i for case insensitive.

	res := collection.FindOne(context.Background(), filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find planet with specified Name: %v", err),
		)
	}

	planet, err := dataToPlanetPb(data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error occured while generating planet data: %v", err),
		)
	}

	return &planetpb.ReadPlanetResponse{
		Planet: planet,
	}, nil
}

func (*server) ListPlanet(ctx context.Context, req *planetpb.ListPlanetRequest) (*planetpb.ListPlanetResponse, error) {
	fmt.Println("List planet request")

	filter := bson.M{} // Nested filter.
	for _, v := range req.GetListPlanetRequestFilter() {
		if v.GetValue() != "All" {
			if v.GetColumn() == "name" { // Need a better way to handle regex.
				filter["name"] = primitive.Regex{Pattern: fmt.Sprintf("^.*%s.*", v.GetValue()), Options: "i"}
			} else {
				filter[v.GetColumn()] = v.GetValue()
			}
		}

	}

	options := options.Find()
	planets := []*planetpb.Planet{}

	// Fetch total documents for pagination.
	itemCount, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}

	options.SetLimit(5)                // Return 5 documents when called.
	options.SetSort(bson.M{"name": 1}) // Sort by name.
	options.SetSkip(req.GetSkip())     // Set skip allows us to use pagination, return 5 documents, but the 5 are different depending on the skip.
	cursor, err := collection.Find(context.Background(), filter, options)
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

		planet, err := dataToPlanetPb(data)
		if err != nil {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error occured while generating planet data: %v", err),
			)
		}

		planets = append(planets, planet)
	}

	return &planetpb.ListPlanetResponse{
		Planet:            planets,
		NumberOfDocuments: itemCount,
	}, nil
}

func (*server) ListPlanetType(ctx context.Context, req *planetpb.ListPlanetTypeRequest) (*planetpb.ListPlanetTypeResponse, error) {
	fmt.Println("List planet type request")

	filter := bson.M{} // Nested filter.

	data, err := collection.Distinct(context.Background(), "basic_information.type", filter) // Grab distinct planet types to populate frontend dropdown.
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error fetching planet type data: %v", err),
		)
	}

	types := []string{}
	for _, value := range data {
		types = append(types, value.(string))
	}

	return &planetpb.ListPlanetTypeResponse{
		PlanetType: types,
	}, nil
}

func dataToPlanetPb(data *planetItem) (*planetpb.Planet, error) { // Helper function to format golang struct data into gRPC data.
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
		Type:                 data.BasicInformation.Type,
		NumberOfSatelites:    data.BasicInformation.NumberOfSatelites,
		StarSystem:           data.BasicInformation.StarSystem,
		MostAbundantResource: data.BasicInformation.MostAbundantResource,
		SurfaceGravity:       data.BasicInformation.SurfaceGravity,
	}

	img, err := util.EncodeImgToBase64(data.Image)
	if err != nil {
		return nil, err
	}

	return &planetpb.Planet{
		PlanetId:         data.ID.Hex(),
		Name:             data.Name,
		Facts:            facts,
		Image:            img,
		BasicInformation: basicInformation,
	}, nil
}

func insertTestData(ctx context.Context) {
	filter := bson.M{}
	collection.DeleteMany(ctx, filter) // Empty collection.

	jsonFile, err := os.Open("planet.json")
	if err != nil {
		log.Fatal(err)
	}

	byteValues, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		// Print any IO errors with the .json file
		log.Fatal("ioutil.ReadFile ERROR:", err)
	}

	var v []interface{}
	if err := json.Unmarshal(byteValues, &v); err != nil {
		log.Fatal(err)
	}
	if _, err := collection.InsertMany(ctx, v); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Connecting to MongoDB")
	// connect to MongoDB
	ip := os.Getenv("MONGODB_URI")
	if ip == "" {
		ip = "localhost"
	}
	uri := "mongodb://" + ip + ":27017"
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
	insertTestData(context.Background())

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
