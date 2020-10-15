package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/mdemaio/celestial-body-info/nasa/nasapb"
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

type apodItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	CopyRight   string             `json:"copyright" bson:"copy_right"`
	Date        string             `json:"date" bson:"date"`
	Explanation string             `json:"explanation" bson:"explanation"`
	HDUrl       string             `json:"hdurl" bson:"hd_url"`
	URL         string             `json:"url"`
	Title       string             `json:"title" bson:"title"`
	MediaType   string             `json:"media_type" bson:"media_type"`
}

func (*server) ListAPOD(ctx context.Context, req *nasapb.ListAPODRequest) (*nasapb.ListAPODResponse, error) {
	fmt.Println("List APOD Request")

	// create an empty struct
	filter := bson.M{}
	apods := []*nasapb.APOD{}

	cursor, err := collection.Find(context.Background(), filter, nil)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		data := &apodItem{}
		err := cursor.Decode(data)
		if err != nil {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding from mongodb: %v", err),
			)
		}

		apod, err := dataToAPODPb(data)
		if err != nil {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error occured while generating apod data: %v", err),
			)
		}

		apods = append(apods, apod)
	}

	return &nasapb.ListAPODResponse{
		Apod: apods,
	}, nil
}

func dataToAPODPb(data *apodItem) (*nasapb.APOD, error) {
	return &nasapb.APOD{
		ApodId:      data.ID.Hex(),
		CopyRight:   data.CopyRight,
		Date:        data.Date,
		Explanation: data.Explanation,
		HdUrl:       data.HDUrl,
		Title:       data.Title,
		MediaType:   data.MediaType,
	}, nil
}

func fetchAPOD() ([]*apodItem, error) {
	apiKey := os.Getenv("NASA_API_KEY") // Use ENV variable instead of hardcode.
	if apiKey == "" {
		apiKey = "DEMO_KEY"
	}
	startDate := time.Now()
	var apodList []*apodItem

	for i := 0; i < 7; i++ { // Get last 7 days worth of data.
		apod := &apodItem{}
		url := fmt.Sprintf("https://api.nasa.gov/planetary/apod?api_key=%s&date=%s", apiKey, startDate.Format("2006-01-02")) // Format date to what NASA API expects.
		res, err := http.Get(url)
		if err != nil {
			log.Fatalf("An error occured while fetching NASA APOD: %v", err)
			return nil, err
		}
		resByte, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatalf("An error occured while reading response body: %v", err)
			return nil, err
		}

		err = json.Unmarshal(resByte, apod)
		if err != nil {
			log.Fatalf("An error occured while unmarshalling into struct: %v", err)
			return nil, err
		}

		if apod.MediaType == "video" { // Videos dont appear to have an HDUrl, so instead we just update it to the basic URL.
			apod.HDUrl = apod.URL
		}

		apodList = append(apodList, apod)
		startDate = startDate.AddDate(0, 0, -1) // Shift back one day, we do this while looping 7 time, aka we get today and 6 previous days(1 week worth of data)
	}

	return apodList, nil
}

func insertAPOD(ctx context.Context, apodList []*apodItem) {
	filter := bson.M{}
	collection.DeleteMany(ctx, filter)

	var ts []interface{}            // ts = temp slice
	for _, apod := range apodList { // Need empty interface for mongodb insert.
		ts = append(ts, apod)
	}
	if _, err := collection.InsertMany(ctx, ts); err != nil {
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

	fmt.Println("NASA Service Started")
	collection = client.Database("celestial-body-info").Collection("apod")
	apod, err := fetchAPOD() //THIS needs to run in a cron job daily or scheduled task somehow.....
	if err != nil {
		log.Fatal(err)
	}
	insertAPOD(context.Background(), apod)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":50053"
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
	nasapb.RegisterNasaServiceServer(s, &server{})
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
