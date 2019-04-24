package rpc

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"go-spg/utils"
	"net/http"
)

type WalletArgs struct {
	Account string
}

type Wallet struct{}

func (w *Wallet) GetNewAddress(r *http.Request, args *WalletArgs, address *string) error {
	*address = "0x29138902384239090432"
	return nil
}

func RunRPCServer() {
	server := rpc.NewServer()
	server.RegisterCodec(json.NewCodec(), "application/json")
	server.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	rpcs := &Wallet{}
	server.RegisterService(rpcs, "")

	r := mux.NewRouter()
	r.Handle("/", server)

	rpcport := fmt.Sprintf(":%d", utils.GetRPCPort())

	http.ListenAndServe(rpcport, r)
}
