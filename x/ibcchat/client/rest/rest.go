package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	// this line is used by starport scaffolding # 1
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers ibcchat-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("/ibcchat/messages/{id}", getMessageHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/ibcchat/messages", listMessageHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	r.HandleFunc("/ibcchat/messages", createMessageHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/ibcchat/messages/{id}", updateMessageHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/ibcchat/messages/{id}", deleteMessageHandler(clientCtx)).Methods("POST")

}
