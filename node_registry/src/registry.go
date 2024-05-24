package src

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/de-wan/blockchain/node_registry/db_sqlc"
)

type AddNodeReq struct {
	Label   string `json:"label"`
	BaseUrl string `json:"base_url"`
}

func HandleAddNode(w http.ResponseWriter, r *http.Request) {
	var req AddNodeReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c := context.Background()
	queries := db_sqlc.New(db_sqlc.DB)

	err = queries.AddNode(c, db_sqlc.AddNodeParams{
		Label:    req.Label,
		BaseUrl:  req.BaseUrl,
		IsOnline: true,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"code": 0,
		"msg":  "success",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func HandleListOnlineNodes(w http.ResponseWriter, r *http.Request) {
	c := context.Background()
	queries := db_sqlc.New(db_sqlc.DB)

	onlineNodes, err := queries.ListOnlineNodes(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"code": 0,
		"msg":  "success",
		"data": onlineNodes,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
