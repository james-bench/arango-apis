//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

//go:generate protoc -I .:../../ --gofast_out=plugins=grpc,paths=source_relative:. ./iam.proto
