package main

import (
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"net/http"
)

func main() {
	server := rpc.NewServer()
	codec := json.NewCodec()
	server.RegisterCodec(codec, "application/json")
	server.RegisterCodec(codec, "application/json; charset=UTF-8") // For firefox 11 and other browsers which append the charset=UTF-8
	server.RegisterService(new(Service), "")
	http.Handle("/rpc", server)
	http.ListenAndServe("127.0.0.1:8080", nil)
}