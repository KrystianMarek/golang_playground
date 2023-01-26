package main

import (
	"google.golang.org/protobuf/proto"
	"protobuf/proto/echo"
	"testing"
)

func TestProtoMarshaling(t *testing.T) {
	expectedValue := "Protobuf"

	req := &echo.EchoRequest{Name: expectedValue}
	data, err := proto.Marshal(req)
	if err != nil {
		t.Errorf("Error while marshalling the object : %v", err)
	}

	res := &echo.EchoRequest{}
	err = proto.Unmarshal(data, res)
	if err != nil {
		t.Errorf("Error while un-marshalling the object : %v", err)
	}

	if expectedValue != res.GetName() {
		t.Errorf("Got: %q, expected: %q", res.GetName(), expectedValue)
	}

}
