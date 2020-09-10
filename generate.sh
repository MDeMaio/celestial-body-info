#!/bin/bash

protoc planet/planetpb/planet.proto --go_out=plugins=grpc:.
protoc star/starpb/star.proto --go_out=plugins=grpc:.
protoc nasa/nasapb/nasa.proto --go_out=plugins=grpc:.