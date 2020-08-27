#!/bin/bash

protoc planet/planetpb/planet.proto --go_out=plugins=grpc:.
