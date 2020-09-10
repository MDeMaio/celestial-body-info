package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gopkg.in/mgo.v2/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

type server struct {
}

type apodItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	CopyRight   string             `json:"copyright" bson:"copy_right"`
	Date        int64              `json:"date" bson:"date"`
	Explanation string             `json:"explanation" bson:"explanation"`
	HDUrl       string             `json:"hdurl" bson:"hd_url"`
	Title       string             `json:"title" bson:"title"`
}

// func (*server) ReadAPOD(ctx context.Context, req *nasapb.ReadAPODRequest) (*nasapb.ReadAPODResponse, error) {
// 	fmt.Println("Read APOD Request")

// 	// create an empty struct
// 	data := &apodItem{}
// 	filter := bson.M{}

// 	res := collection.FindOne(context.Background(), filter)
// 	if err := res.Decode(data); err != nil {
// 		return nil, status.Errorf(
// 			codes.NotFound,
// 			fmt.Sprintf("Cannot find planet with specified Name: %v", err),
// 		)
// 	}

// 	planet, err := dataToPlanetPb(data)
// 	if err != nil {
// 		return nil, status.Errorf(
// 			codes.Internal,
// 			fmt.Sprintf("Error occured while generating planet data: %v", err),
// 		)
// 	}

// 	return &nasapb.ReadAPODResponse{
// 		Apod: planet,
// 	}, nil
// }

// func dataToPlanetPb(data *apodItem) (*nasapb.APOD, error) {
// 	facts := []*planetpb.Facts{} // Not sure if this is correct.
// 	for _, v := range data.Facts {
// 		fact := &planetpb.Facts{
// 			Title: v.Title,
// 			Fact:  v.Fact,
// 		}

// 		facts = append(facts, fact)
// 	}

// 	basicInformation := &planetpb.BasicInformation{
// 		AlternateName:        data.BasicInformation.AlternateName,
// 		Type:                 data.BasicInformation.Type,
// 		NumberOfSatelites:    data.BasicInformation.NumberOfSatelites,
// 		StarSystem:           data.BasicInformation.StarSystem,
// 		MostAbundantResource: data.BasicInformation.MostAbundantResource,
// 		SurfaceGravity:       data.BasicInformation.SurfaceGravity,
// 	}

// 	img, err := util.EncodeImgToBase64(data.Image)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &planetpb.Planet{
// 		PlanetId:         data.ID.Hex(),
// 		Name:             data.Name,
// 		Facts:            facts,
// 		Image:            img,
// 		BasicInformation: basicInformation,
// 	}, nil
// }

func fetchAPOD() (*apodItem, error) { // This function is incomplete, need to parse the date somehow into an int64 before unmarshalling.
	apod := &apodItem{}

	// SWITCH TO ENV VARIABLE BEFORE PUSHING THIS!!!!!!!
	// Default date is today, no need to pass it in.
	url := "https://api.nasa.gov/planetary/apod?api_key=eghT6nR439KFj7tnC8ShZZ6tR77GirCYV19P9DhM"
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
	fmt.Printf("%s", resByte)

	err = json.Unmarshal(resByte, apod)
	if err != nil {
		log.Fatalf("An error occured while unmarshalling into struct: %v", err)
		return nil, err
	}

	fmt.Print(apod)

	return apod, nil
}

func insertAPOD(ctx context.Context) {
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
	fetchAPOD()

	// // if we crash the go code, we get the file name and line number
	// log.SetFlags(log.LstdFlags | log.Lshortfile)

	// fmt.Println("Connecting to MongoDB")
	// // connect to MongoDB
	// uri := os.Getenv("MONGODB_URI")
	// if uri == "" {
	// 	uri = "mongodb://localhost:27017"
	// }
	// fmt.Println(uri)

	// client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = client.Connect(context.TODO())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("NASA Service Started")
	// collection = client.Database("celestial-body-info").Collection("apod")
	// insertTestData(context.Background())

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = ":50051"
	// } else {
	// 	port = ":" + port
	// }
	// lis, err := net.Listen("tcp", port)
	// if err != nil {
	// 	log.Fatalf("Failed to listen: %v", err)
	// }
	// fmt.Println("Up and running on: " + port)

	// opts := []grpc.ServerOption{}
	// s := grpc.NewServer(opts...)
	// nasapb.RegisterNasaServiceServer(s, &server{})
	// // Register reflection service on gRPC server.
	// reflection.Register(s)

	// go func() {
	// 	fmt.Println("Starting Server...")
	// 	if err := s.Serve(lis); err != nil {
	// 		log.Fatalf("failed to serve: %v", err)
	// 	}
	// }()

	// // Wait for Control C to exit
	// ch := make(chan os.Signal, 1)
	// signal.Notify(ch, os.Interrupt)

	// // Block until a signal is received
	// <-ch
	// // First we close the connection with MongoDB:
	// fmt.Println("Closing MongoDB Connection")
	// if err := client.Disconnect(context.TODO()); err != nil {
	// 	log.Fatalf("Error on disconnection with MongoDB : %v", err)
	// }
	// // Second step : closing the listener
	// fmt.Println("Closing the listener")
	// if err := lis.Close(); err != nil {
	// 	log.Fatalf("Error on closing the listener : %v", err)
	// }
	// // Finally, we stop the server
	// fmt.Println("Stopping the server")
	// s.Stop()
	// fmt.Println("End of Program")
}
