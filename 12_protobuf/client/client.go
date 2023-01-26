package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"protobuf/proto/echo"

	"github.com/golang/protobuf/proto"
)

const (
	endpoint string = "http://127.0.0.1:8080/v1/echo"
)

func requestProtobuf(request *echo.EchoRequest) *echo.EchoResponse {

	req, err := proto.Marshal(request)
	if err != nil {
		log.Fatalf("Unable to marshal request : %v", err)
	}

	resp, err := http.Post(endpoint, "application/x-binary", bytes.NewReader(req))
	if err != nil {
		log.Fatalf("Unable to read from the server : %v", err)
	}
	respBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Unable to read bytes from request : %v", err)
	}

	respObj := &echo.EchoResponse{}
	proto.Unmarshal(respBytes, respObj)
	return respObj

}

func requestJson(request *echo.EchoRequest) *echo.EchoResponse {
	data, _ := request.MarshalJSON()
	resp, _ := http.Post(endpoint, "application/json", bytes.NewReader(data))

	bytes, _ := io.ReadAll(resp.Body)
	result := echo.EchoResponse{}
	result.UnmarshalJSON(bytes)

	return &result
}

func main() {

	protoRequest := &echo.EchoRequest{Name: "Proto"}
	protoResp := requestProtobuf(protoRequest)
	fmt.Printf("Response from Proto API is : %v\n", protoResp.GetMessage())

	jsonRequest := &echo.EchoRequest{Name: "Json"}
	jsonResp := requestJson(jsonRequest)
	fmt.Printf("Response from Json API is : %v\n", jsonResp.GetMessage())
}
