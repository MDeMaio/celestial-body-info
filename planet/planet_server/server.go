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

	"github.com/mdemaiocwg/celestial-body-info/planet/planetpb"
	"google.golang.org/grpc"
)

var collection *mongo.Collection

type server struct {
}

type inclination struct {
	Ecliptic        float64 `bson:"ecliptic"`
	SunsEquator     float64 `bson:"suns_equator"`
	InvariablePlane float64 `bson:"invariable_plane"`
}

type orbitalInfo struct {
	Aphelion                 float64     `bson:"aphelion"`
	Perihelion               float64     `bson:"perihelion"`
	SemiMajorAxis            float64     `bson:"semi_major_axis"`
	Eccentricity             float64     `bson:"eccentricity"`
	OrbitalPeriod            float64     `bson:"orbital_period"`
	SynodicPeriod            float64     `bson:"synodic_period"`
	AvgOrbitalSpeed          float64     `bson:"avg_orbital_speed"`
	MeanAnomaly              float64     `bson:"mean_anomaly"`
	Inclination              inclination `bson:"inclination"`
	LongitudeOfAscendingNode float64     `bson:"longitude_of_ascending_node"`
	Satelites                uint32      `bson:"satelites"`
}

type albedo struct {
	Geometric float64 `bson:"geometric"`
	Bond      float64 `bson:"bond"`
}

type surfaceTemp struct {
	Min  float64 `bson:"min"`
	Max  float64 `bson:"max"`
	Mean float64 `bson:"mean"`
}

type apparentMagnitude struct {
	Min float64 `bson:"min"`
	Max float64 `bson:"max"`
}

type physicalInfo struct {
	MeanRadius                 string            `bson:"mean_radius"`
	EquatorialRadius           string            `bson:"equatorial_radius"`
	PolarRadius                string            `bson:"polar_radius"`
	Flattening                 string            `bson:"flattening"`
	SurfaceArea                float64           `bson:"surface_area"`
	Volume                     float64           `bson:"volume"`
	Mass                       float64           `bson:"mass"`
	MeanDensity                float64           `bson:"mean_density"`
	SurfaceGravity             float64           `bson:"surface_gravity"`
	MomentOfInertiaFactor      string            `bson:"moment_of_inertia_factor"`
	EscapeVelocity             float64           `bson:"escape_velocity"`
	SiderealRotationPeriod     float64           `bson:"sidereal_rotation_period"`
	EquatorialRotationVelocity float64           `bson:"equatorial_rotation_velocity"`
	AxialTilt                  float64           `bson:"axial_tilt"`
	NorthpoleRightAscension    float64           `bson:"northpole_right_ascension"`
	NorthpoleDeclination       float64           `bson:"northpole_declination"`
	Albedo                     albedo            `bson:"albedo"`
	SurfaceTemp                surfaceTemp       `bson:"surface_temp"`
	ApparentMagnitude          apparentMagnitude `bson:"apparent_magnitude"`
}

type element struct {
	Name             string  `bson:"name"`
	PercentAsDecimal float64 `bson:"percent_as_decimal"`
}

type atmosphereInfo struct {
	SurfacePressure float64   `bson:"surface_pressure"`
	Elements        []element `bson:"element"` // Many elements make up a planets atmospheric composition.
}

type planetItem struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Name           string             `bson:"name"`
	OrbitalInfo    orbitalInfo        `bson:"orbital_info"`
	PhysicalInfo   physicalInfo       `bson:"physical_info"`
	AtmosphereInfo atmosphereInfo     `bson:"atmosphere_info"`
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
	inclination := &planetpb.Inclination{
		Ecliptic:        data.OrbitalInfo.Inclination.Ecliptic,
		Sunsequator:     data.OrbitalInfo.Inclination.SunsEquator,
		Invariableplane: data.OrbitalInfo.Inclination.InvariablePlane,
	}
	orbitalInfo := &planetpb.OrbitalInfo{
		Aphelion:                 data.OrbitalInfo.Aphelion,
		Perihelion:               data.OrbitalInfo.Perihelion,
		Semimajoraxis:            data.OrbitalInfo.SemiMajorAxis,
		Eccentricity:             data.OrbitalInfo.Eccentricity,
		Orbitalperiod:            data.OrbitalInfo.OrbitalPeriod,
		Synodicperiod:            data.OrbitalInfo.SynodicPeriod,
		Avgorbitalspeed:          data.OrbitalInfo.AvgOrbitalSpeed,
		Meananomaly:              data.OrbitalInfo.MeanAnomaly,
		Inclination:              inclination,
		Longitudeofascendingnode: data.OrbitalInfo.LongitudeOfAscendingNode,
		Satelites:                data.OrbitalInfo.Satelites,
	}

	albedo := &planetpb.Albedo{
		Geometric: data.PhysicalInfo.Albedo.Geometric,
		Bond:      data.PhysicalInfo.Albedo.Bond,
	}
	surfaceTemp := &planetpb.SurfaceTemp{
		Min:  data.PhysicalInfo.SurfaceTemp.Min,
		Max:  data.PhysicalInfo.SurfaceTemp.Max,
		Mean: data.PhysicalInfo.SurfaceTemp.Mean,
	}
	apparentMagnitude := &planetpb.ApparentMagnitude{
		Min: data.PhysicalInfo.ApparentMagnitude.Min,
		Max: data.PhysicalInfo.ApparentMagnitude.Max,
	}
	physicalInfo := &planetpb.PhysicalInfo{
		Meanradius:                 data.PhysicalInfo.MeanRadius,
		Equatorialradius:           data.PhysicalInfo.EquatorialRadius,
		Polarradius:                data.PhysicalInfo.PolarRadius,
		Flattening:                 data.PhysicalInfo.Flattening,
		Surfacearea:                data.PhysicalInfo.SurfaceArea,
		Volume:                     data.PhysicalInfo.Volume,
		Mass:                       data.PhysicalInfo.Mass,
		Meandensity:                data.PhysicalInfo.MeanDensity,
		Surfacegravity:             data.PhysicalInfo.SurfaceGravity,
		Momentofinertiafactor:      data.PhysicalInfo.MomentOfInertiaFactor,
		Escapevelocity:             data.PhysicalInfo.EscapeVelocity,
		Siderealrotationperiod:     data.PhysicalInfo.SiderealRotationPeriod,
		Equatorialrotationvelocity: data.PhysicalInfo.EquatorialRotationVelocity,
		Axialtilt:                  data.PhysicalInfo.AxialTilt,
		Northpolerightascension:    data.PhysicalInfo.NorthpoleRightAscension,
		Northpoledeclination:       data.PhysicalInfo.NorthpoleDeclination,
		Albedo:                     albedo,
		Surfacetemp:                surfaceTemp,
		Apparentmagnitude:          apparentMagnitude,
	}

	elements := []*planetpb.Element{} // Not sure if this is correct.
	for _, v := range data.AtmosphereInfo.Elements {
		element := &planetpb.Element{
			Name:             v.Name,
			Percentasdecimal: v.PercentAsDecimal,
		}

		elements = append(elements, element)
	}

	atmosphereInfo := &planetpb.AtmosphereInfo{
		Surfacepressure: data.AtmosphereInfo.SurfacePressure,
		Element:         elements,
	}

	return &planetpb.Planet{
		PlanetId:       data.ID.Hex(),
		Name:           data.Name,
		Orbitalinfo:    orbitalInfo,
		Physicalinfo:   physicalInfo,
		Atmosphereinfo: atmosphereInfo,
	}
}

func main() {
	// if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Connecting to MongoDB")
	// connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Planet Service Started")
	collection = client.Database("celestial-body-info").Collection("planet")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

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
