//
// DISCLAIMER
//
// Copyright 2019 ArangoDB GmbH, Cologne, Germany
//
// Author Robert Stam
//

// Package backup/v1 contains the API of Backup services.
package v1

//go:generate protoc -I .:../../:../../vendor/:../../vendor/googleapis/:../../vendor/github.com/gogo/protobuf/protobuf/ --gofast_out=Mgithub.com/golang/protobuf/ptypes/duration/duration.proto=github.com/gogo/protobuf/types,Mgithub.com/golang/protobuf/ptypes/timestamp/timestamp.proto=github.com/gogo/protobuf/types,plugins=grpc,paths=source_relative:. --grpc-gateway_out=paths=source_relative,logtostderr=true,allow_delete_body=true:. ./backup.proto
