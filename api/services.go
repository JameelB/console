package api

import (
	"net/http"

	"github.com/coreos-inc/bridge/Godeps/_workspace/src/github.com/gorilla/mux"
)

func registerServices(router *mux.Router) {
	router.HandleFunc("/services", serviceList).Methods("GET")
	router.HandleFunc("/services/{id}", serviceGet).Methods("GET")
}

// Get Service api endpoint.
func serviceGet(w http.ResponseWriter, r *http.Request) {
	err := k8proxy.DoAndRespond(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
	}
}

// List Services api endpoint.
func serviceList(w http.ResponseWriter, r *http.Request) {
	err := k8proxy.DoAndRespond(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
	}
}