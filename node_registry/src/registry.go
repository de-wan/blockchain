package src

import "net/http"

type AddNodeReq struct {
	Label   string `json:"label"`
	BaseUrl string `json:"base_url"`
}

func HandleAddNode(w http.ResponseWriter, r *http.Request) {

}
