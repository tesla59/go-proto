package main

import (
	"testing"

	"github.com/tesla59/go-proto/pb/V2"
	"github.com/tesla59/go-proto/pb/V2vtproto"
	"google.golang.org/protobuf/proto"
)

func BenchmarkSerialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(&V2.SimpleMessage{})
		if err != nil {
			b.Error("Error while marshalling", err)
		}
	}
}

func BenchmarkDeserialize(b *testing.B) {
	data, err := proto.Marshal(&V2.SimpleMessage{})
	if err != nil {
		b.Error("Error while marshalling", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sm V2.SimpleMessage
		err := proto.Unmarshal(data, &sm)
		if err != nil {
			b.Error("Error while unmarshalling", err)
		}
	}
}

func BenchmarkSerializeVT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := (&V2vtproto.SimpleMessage{}).MarshalVT()
		if err != nil {
			b.Error("Error while marshalling", err)
		}
	}
}

func BenchmarkDeserializeVT(b *testing.B) {
	data, err := (&V2vtproto.SimpleMessage{}).MarshalVT()
	if err != nil {
		b.Error("Error while marshalling", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sm V2vtproto.SimpleMessage
		err := sm.UnmarshalVT(data)
		if err != nil {
			b.Error("Error while unmarshalling", err)
		}
	}
}
