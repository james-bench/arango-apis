//
// DISCLAIMER
//
// Copyright 2019 ArangoDB GmbH, Cologne, Germany
//
// Author Ewout Prangsma
//

// Package common/v1 contains the common types used in API of ArangoDB Managed Service
package v1

//go:generate protoc -I .:../../:../../vendor/ --gofast_out=Mgithub.com/golang/protobuf/ptypes/timestamp/timestamp.proto=github.com/gogo/protobuf/types,plugins=grpc,paths=source_relative:. ./common.proto
