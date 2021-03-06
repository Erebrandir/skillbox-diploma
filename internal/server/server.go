package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"skillbox-diploma/internal/config"
	"skillbox-diploma/internal/result"
)

func listenAndServeHTTP() {
	router := mux.NewRouter()
	router.HandleFunc("/api", handleAPI).Methods("GET")

	fileServer := http.FileServer(http.Dir(config.GlobalConfig.WebDir))
	router.PathPrefix("/").Handler(http.StripPrefix("/", fileServer))
	log.Fatal(http.ListenAndServe(config.GlobalConfig.Addr, router))
}

func handleAPI(w http.ResponseWriter, r *http.Request) {
	resultT := result.ResultT{Status: false, Error: "Error on collect data"}

	resultSetT := result.GetResultData()
	if result.CheckResult(resultSetT) {
		resultT.Status = true
		resultT.Data = resultSetT
		resultT.Error = ""
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	response, err := json.Marshal(resultT)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte{})
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}

func StartServer() {
	listenAndServeHTTP()
}
