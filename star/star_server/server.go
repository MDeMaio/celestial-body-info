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

	"github.com/mdemaio/celestial-body-info/logger"
	"github.com/mdemaio/celestial-body-info/star/starpb"
	"github.com/mdemaio/celestial-body-info/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
)

var collection *mongo.Collection

type server struct {
}

type facts struct {
	Title string `bson:"title"`
	Fact  string `bson:"fact"`
}

type basicInformation struct {
	Mass           string `bson:"mass"`
	Classification string `bson:"classification"`
	Radius         string `bson:"radius"`
	System         string `bson:"system"`
	Temperature    string `bson:"temperature"`
	Age            string `bson:"age"`
}
type starItem struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	Name             string             `bson:"name"`
	Facts            []facts            `bson:"facts"`
	Image            string             `bson:"image"`
	BasicInformation basicInformation   `bson:"basic_information"`
}

func (*server) ReadStar(ctx context.Context, req *starpb.ReadStarRequest) (*starpb.ReadStarResponse, error) {

	fmt.Println("Read star request")

	starName := req.GetName()

	// create an empty struct
	data := &starItem{}
	filter := bson.M{"name": starName}

	res := collection.FindOne(context.Background(), filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find star with specified Name: %v", err),
		)
	}

	star, err := dataToStarPb(data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error occured while generating star data: %v", err),
		)
	}

	return &starpb.ReadStarResponse{
		Star: star,
	}, nil
}

func (*server) ListStar(ctx context.Context, req *starpb.ListStarRequest) (*starpb.ListStarResponse, error) {
	var l = &logger.Logger{
		Msg:     "List-Star request",
		Funtion: "ListStar",
		Service: "Star",
		LogType: "INFORMATION",
	}

	l.LogToFile("/logs/logs.txt")
	fmt.Println("List star request")

	filter := bson.M{} // Nested filter.
	for _, v := range req.GetListStarRequestFilter() {
		if v.GetValue() != "All" {
			if v.GetColumn() == "name" { // Need a better way to handle regex.
				filter["name"] = primitive.Regex{Pattern: fmt.Sprintf("^.*%s.*", v.GetValue()), Options: "i"}
			} else {
				filter[v.GetColumn()] = v.GetValue()
			}
		}

	}

	options := options.Find()
	stars := []*starpb.Star{}

	// Fetch total documents for pagination.
	itemCount, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}

	options.SetLimit(5)
	options.SetSort(bson.M{"name": 1})
	options.SetSkip(req.GetSkip())
	cursor, err := collection.Find(context.Background(), filter, options)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		data := &starItem{}
		err := cursor.Decode(data)
		if err != nil {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding from mongodb: %v", err),
			)
		}

		star, err := dataToStarPb(data)
		if err != nil {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error occured while generating star data: %v", err),
			)
		}

		stars = append(stars, star)
	}

	return &starpb.ListStarResponse{
		Star:              stars,
		NumberOfDocuments: itemCount,
	}, nil
}

func dataToStarPb(data *starItem) (*starpb.Star, error) {
	facts := []*starpb.Facts{} // Not sure if this is correct.
	for _, v := range data.Facts {
		fact := &starpb.Facts{
			Title: v.Title,
			Fact:  v.Fact,
		}

		facts = append(facts, fact)
	}

	basicInformation := &starpb.BasicInformation{
		Mass:           data.BasicInformation.Mass,
		Classification: data.BasicInformation.Classification,
		Radius:         data.BasicInformation.Radius,
		System:         data.BasicInformation.System,
		Temperature:    data.BasicInformation.Temperature,
		Age:            data.BasicInformation.Age,
	}

	img, err := util.EncodeImgToBase64(data.Image)
	if err != nil {
		return nil, err
	}

	return &starpb.Star{
		StarId:           data.ID.Hex(),
		Name:             data.Name,
		Facts:            facts,
		Image:            img,
		BasicInformation: basicInformation,
	}, nil
}

func insertTestData(ctx context.Context) {
	filter := bson.M{}
	collection.DeleteMany(ctx, filter) // Empty collection.

	jsonFile, err := os.Open("star.json")
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

	fmt.Println("Star Service Started")
	collection = client.Database("celestial-body-info").Collection("star")
	insertTestData(context.Background())

	port := os.Getenv("PORT")
	if port == "" {
		port = ":50052"
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
	starpb.RegisterStarServiceServer(s, &server{})
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
