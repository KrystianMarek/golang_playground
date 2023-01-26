package main

import (
	"github.com/gorilla/mux"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"net/http"
	"protobuf/proto/echo"
	"time"
)

func Echo(resp http.ResponseWriter, req *http.Request) {
	contentLength := req.ContentLength
	log.Printf("Content Length Received : %v\n", contentLength)
	log.Printf("Headers: %v", req.Header)
	request := &echo.EchoRequest{}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Unable to read message from request : %v", err)
	}
	proto.Unmarshal(data, request)
	name := request.GetName()
	result := &echo.EchoResponse{Message: "Hello " + name}
	response, err := proto.Marshal(result)
	if err != nil {
		log.Fatalf("Unable to marshal response : %v", err)
	}

	resp.Header().Add("Access-Control-Allow-Origin", "*")
	resp.Header().Add("Access-Control-Allow-Method", "*")
	resp.Header().Add("Access-Control-Allow-Headers", "*")
	resp.Write(response)
}

func EchoJson(resp http.ResponseWriter, req *http.Request) {
	contentLength := req.ContentLength
	log.Printf("Content Length Received : %v\n", contentLength)
	log.Printf("Headers: %v", req.Header)
	request := &echo.EchoRequest{}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Unable to read message from request : %v", err)
	}
	request.UnmarshalJSON(data)
	name := request.Name
	result := &echo.EchoResponse{Message: "Hello " + name}
	response, err := result.MarshalJSON()
	if err != nil {
		log.Fatalf("Unable to marshal response : %v", err)
	}

	resp.Header().Add("Access-Control-Allow-Origin", "*")
	resp.Header().Add("Access-Control-Allow-Method", "*")
	resp.Header().Add("Access-Control-Allow-Headers", "*")
	resp.Write(response)
}

func EchoJsonGeneric(resp http.ResponseWriter, req *http.Request) {
	result := &echo.EchoResponse{Message: "Hello Stranger"}
	response, err := result.MarshalJSON()
	if err != nil {
		log.Fatalf("Unable to marshal response : %v", err)
	}

	resp.Header().Add("Access-Control-Allow-Origin", "*")
	resp.Header().Add("Access-Control-Allow-Method", "*")
	resp.Header().Add("Access-Control-Allow-Headers", "*")
	resp.Write(response)
}

func EchoGeneric(resp http.ResponseWriter, req *http.Request) {
	result := &echo.EchoResponse{Message: "Hello Stranger"}
	response, err := proto.Marshal(result)
	if err != nil {
		log.Fatalf("Unable to marshal response : %v", err)
	}

	resp.Header().Add("Access-Control-Allow-Origin", "*")
	resp.Header().Add("Access-Control-Allow-Method", "*")
	resp.Header().Add("Access-Control-Allow-Headers", "*")
	resp.Write(response)
}

func main() {
	log.Println("Starting the API server...")
	root := mux.NewRouter()

	root.Path("/v1/echo").Methods("POST").Headers("Content-Type", "application/x-binary").HandlerFunc(Echo)
	root.Path("/v1/echo").Methods("POST").HeadersRegexp("Content-Type", "application/(text|json)").HandlerFunc(EchoJson)
	root.Path("/v1/echo").Methods("GET").Headers("Content-Type", "application/x-binary").HandlerFunc(EchoGeneric)
	root.Path("/v1/echo").Methods("GET").HeadersRegexp("Content-Type", "application/(text|json)").HandlerFunc(EchoJsonGeneric)

	//https://stackoverflow.com/questions/43871637/no-access-control-allow-origin-header-is-present-on-the-requested-resource-whe
	root.Path("/v1/echo").Methods("OPTIONS").HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Access-Control-Allow-Origin", "*")
		resp.Header().Add("Access-Control-Allow-Method", "*")
		resp.Header().Add("Access-Control-Allow-Headers", "*")
		resp.Write(make([]byte, 0))
	})

	server := &http.Server{
		Handler:      root,
		Addr:         ":8080",
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
