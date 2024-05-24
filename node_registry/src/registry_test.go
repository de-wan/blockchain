package src

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/de-wan/blockchain/node_registry/db_sqlc"
	"github.com/stretchr/testify/assert"
)

// func TestMain(m *testing.M) {
// 	err := godotenv.Load(filepath.Join("..", ".env"))
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 		return
// 	}
// 	log.Println("Loaded .env file")
// }

func TestAddNode(t *testing.T) {
	db_sqlc.InitTestDB()

	addNodeData := AddNodeReq{
		Label:   "mynode",
		BaseUrl: "http://localhost:9001",
	}

	addNodeBytes, err := json.Marshal(addNodeData)
	assert.NoError(t, err, "Unexpected error returned")

	// Create a new request
	req, err := http.NewRequest("GET", "/test", strings.NewReader(string(addNodeBytes)))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleAddNode)
	handler.ServeHTTP(w, req)

	// Check the status code is what we expect
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect
	expected := `{"code":0,"msg":"success"}`
	received := strings.TrimSpace(w.Body.String())
	if received != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			received, expected)
	}
}

func TestListOnlineNodes(t *testing.T) {
	db_sqlc.InitTestDB()

	// Create a new request
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleListOnlineNodes)
	handler.ServeHTTP(w, req)

	// Check the status code is what we expect
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	received := strings.TrimSpace(w.Body.String())
	fmt.Println(received)
}

func TestNodeConnectionHeartbeat(t *testing.T) {

}

func TestGetSyncNode(t *testing.T) {

}
