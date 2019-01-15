//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Ewout Prangsma
//

// Package resourcemanager/v1 contains the API of ResourceManager services.
package v1

//go:generate protoc -I .:../../:../../vendor/ --gofast_out=Mgithub.com/golang/protobuf/ptypes/timestamp/timestamp.proto=github.com/gogo/protobuf/types,plugins=grpc,paths=source_relative:. ./resourcemanager.proto
