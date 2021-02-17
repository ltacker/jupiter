package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	// this line is used by starport scaffolding # 1
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers laugh-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("/laugh/hohos/{id}", getHohoHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/laugh/hohos", listHohoHandler(clientCtx)).Methods("GET")

	r.HandleFunc("/laugh/hihis/{id}", getHihiHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/laugh/hihis", listHihiHandler(clientCtx)).Methods("GET")

	r.HandleFunc("/laugh/hahas/{id}", getHahaHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/laugh/hahas", listHahaHandler(clientCtx)).Methods("GET")

	r.HandleFunc("/laugh/hohosents/{id}", getHohosentHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/laugh/hohosents", listHohosentHandler(clientCtx)).Methods("GET")

	r.HandleFunc("/laugh/hihisents/{id}", getHihisentHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/laugh/hihisents", listHihisentHandler(clientCtx)).Methods("GET")

	r.HandleFunc("/laugh/hahasents/{id}", getHahasentHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/laugh/hahasents", listHahasentHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	r.HandleFunc("/laugh/hohos", createHohoHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/laugh/hohos/{id}", updateHohoHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/laugh/hohos/{id}", deleteHohoHandler(clientCtx)).Methods("POST")

	r.HandleFunc("/laugh/hihis", createHihiHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/laugh/hihis/{id}", updateHihiHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/laugh/hihis/{id}", deleteHihiHandler(clientCtx)).Methods("POST")

	r.HandleFunc("/laugh/hahas", createHahaHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/laugh/hahas/{id}", updateHahaHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/laugh/hahas/{id}", deleteHahaHandler(clientCtx)).Methods("POST")

	r.HandleFunc("/laugh/hohosents", createHohosentHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/laugh/hohosents/{id}", updateHohosentHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/laugh/hohosents/{id}", deleteHohosentHandler(clientCtx)).Methods("POST")

	r.HandleFunc("/laugh/hihisents", createHihisentHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/laugh/hihisents/{id}", updateHihisentHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/laugh/hihisents/{id}", deleteHihisentHandler(clientCtx)).Methods("POST")

	r.HandleFunc("/laugh/hahasents", createHahasentHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/laugh/hahasents/{id}", updateHahasentHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/laugh/hahasents/{id}", deleteHahasentHandler(clientCtx)).Methods("POST")

}
