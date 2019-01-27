//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Ewout Prangsma
//

// Package provider/v1 contains the API used to fetch provider & region information.
package v1

//go:generate protoc -I .:../../:../../vendor/:../../vendor/googleapis/:../../vendor/github.com/gogo/protobuf/protobuf/ --gofast_out=Mgithub.com/golang/protobuf/ptypes/timestamp/timestamp.proto=github.com/gogo/protobuf/types,plugins=grpc,paths=source_relative:. --grpc-gateway_out=paths=source_relative,logtostderr=true,allow_delete_body=true:. ./platform.proto
