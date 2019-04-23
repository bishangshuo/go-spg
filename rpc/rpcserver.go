package rpc

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"go-spg/utils"
	"net/http"
)

type Args struct {
	A int
	B int
}

type RPCServer struct{}

func (s *RPCServer) Add(r *http.Request, args *Args, result *int) error {
	*result = args.A + args.B
	return nil
}

func (s *RPCServer) Subtract(r *http.Request, args *Args, result *int) error {
	*result = args.A - args.B
	return nil
}

func RunRPCServer() {
	server := rpc.NewServer()
	server.RegisterCodec(json.NewCodec(), "application/json")
	server.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	rpcs := &RPCServer{}
	server.RegisterService(rpcs, "")
	r := mux.NewRouter()
	r.Handle("/", server)

	rpcport := fmt.Sprintf(":%d", utils.GetRPCPort())

	http.ListenAndServe(rpcport, r)
}
