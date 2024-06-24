package server

import (
	"encoding/json"
	"net/http"

	_ "net/http/pprof"

	"github.com/rahul1804/kvstore/internal/store"
	"github.com/rahul1804/kvstore/pkg/logger"
)

var kvStore = store.NewStore()

// Start initializes and starts the HTTP server.
func Start() {
	logger.Logger.Info("Starting server on :8080")
	go func() {
		logger.Logger.Fatal(http.ListenAndServe(":8081", nil)) // pprof server
	}()
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/set", setHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/show", showHandler)
	logger.Logger.Fatal(http.ListenAndServe(":8080", nil))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value, ok := kvStore.Get(key)
	if !ok {
		logger.Logger.Warnf("Key not found: %s", key)
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}
	logger.Logger.Infof("Retrieved key: %s, value: %s", key, value)
	json.NewEncoder(w).Encode(map[string]string{"value": value})
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	var req map[string]string
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Logger.Errorf("Invalid request: %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	key, okKey := req["key"]
	value, okValue := req["value"]
	if !okKey || !okValue {
		logger.Logger.Warn("Missing key or value in request")
		http.Error(w, "Missing key or value", http.StatusBadRequest)
		return
	}
	kvStore.Set(key, value)
	logger.Logger.Infof("Set key: %s, value: %s", key, value)
	w.WriteHeader(http.StatusOK)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	var req map[string]string
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Logger.Errorf("Invalid request: %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	key, ok := req["key"]
	if !ok {
		logger.Logger.Warn("Missing key in request")
		http.Error(w, "Missing key", http.StatusBadRequest)
		return
	}
	kvStore.Delete(key)
	logger.Logger.Infof("Deleted key: %s", key)
	w.WriteHeader(http.StatusOK)
}

func showHandler(w http.ResponseWriter, r *http.Request) {
	data := kvStore.GetAll()
	json.NewEncoder(w).Encode(data)
	logger.Logger.Info("Displayed all key-value pairs")
}
