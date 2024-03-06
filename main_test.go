package main

import (
	"testing"

	"github.com/tesla59/go-proto/pb/simple"
	"google.golang.org/protobuf/proto"
)

var SimpleMessage = simple.SimpleMessage{
	Id:      1234,
	Message: "This is a simple message",
}

func BenchmarkSerialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(&SimpleMessage)
		if err != nil {
			b.Error("Error while marshalling", err)
		}
	}
}
