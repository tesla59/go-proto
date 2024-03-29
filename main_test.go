package main

import (
	"testing"

	"github.com/tesla59/go-proto/pb/V2"
	"github.com/tesla59/go-proto/pb/V2gogo"
	"github.com/tesla59/go-proto/pb/V2vtproto"
	"google.golang.org/protobuf/proto"
)

func BenchmarkSerialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(&V2.WriteRequest{})
		if err != nil {
			b.Error("Error while marshalling", err)
		}
	}
}

func BenchmarkSerializeVT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := (&V2vtproto.WriteRequest{}).MarshalVT()
		if err != nil {
			b.Error("Error while marshalling", err)
		}
	}
}

func BenchmarkSerializeGogo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := (&V2gogo.WriteRequest{}).Marshal()
		if err != nil {
			b.Error("Error while marshalling", err)
		}
	}

}

func BenchmarkDeserialize(b *testing.B) {
	data, err := proto.Marshal(&V2.WriteRequest{})
	if err != nil {
		b.Error("Error while marshalling", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sm V2.WriteRequest
		err := proto.Unmarshal(data, &sm)
		if err != nil {
			b.Error("Error while unmarshalling", err)
		}
	}
}

func BenchmarkDeserializeVT(b *testing.B) {
	data, err := (&V2vtproto.WriteRequest{}).MarshalVT()
	if err != nil {
		b.Error("Error while marshalling", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sm V2vtproto.WriteRequest
		err := sm.UnmarshalVT(data)
		if err != nil {
			b.Error("Error while unmarshalling", err)
		}
	}
}

func BenchmarkDeserializeGogo(b *testing.B) {
	data, err := (&V2gogo.WriteRequest{}).Marshal()
	if err != nil {
		b.Error("Error while marshalling", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sm V2gogo.WriteRequest
		err := sm.Unmarshal(data)
		if err != nil {
			b.Error("Error while unmarshalling", err)
		}
	}
}
