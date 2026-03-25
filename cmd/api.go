package main

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type api struct {
	config config
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	//
}

// API body definitions
type CalcRequest struct {
	LiquidAmt float64 `json:"amt"`
}

type CaclResponse struct {
	LiquidAmt float64 `json:"LiquidAmt"`
	PowderAmt float64 `json:"PowderAmt"`
}

// API functions
func (app *api) mount() http.Handler {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.HandleFunc("/calc", PostCalc).Methods("POST", "OPTIONS")

	return r
}

func (app *api) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("Server has started at address %s", app.config.addr)
	return srv.ListenAndServe()
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("API called on", "URL", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func PostCalc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		return
	}

	var calcRequest CalcRequest
	err := json.NewDecoder(r.Body).Decode(&calcRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := CalcPowder(float64(calcRequest.LiquidAmt))
	json.NewEncoder(w).Encode(result)
}
