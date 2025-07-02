package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/IrusHunter/MicroserviceCalculator/types"
)

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

type JSONAPIServer struct {
	listenAddr string
	svc        ICalculator
}

func NewJSONAPIServer(listenAddr string, svc ICalculator) *JSONAPIServer {
	return &JSONAPIServer{
		listenAddr: listenAddr,
		svc:        svc,
	}
}

func (s *JSONAPIServer) Run() {
	http.HandleFunc("/", makeHTTPHandlerFunc(s.handleCalculate))

	http.ListenAndServe(s.listenAddr, nil)
}

func makeHTTPHandlerFunc(apiFunc APIFunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", rand.Intn(1_000_000))

	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFunc(ctx, w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}

func (s *JSONAPIServer) handleCalculate(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	a, err := strconv.ParseFloat(r.URL.Query().Get("a"), 32)
	if err != nil {
		return fmt.Errorf("invalid value for 'a': %v", err)
	}
	b, err := strconv.ParseFloat(r.URL.Query().Get("b"), 32)
	if err != nil {
		return fmt.Errorf("invalid value for 'b': %v", err)
	}
	operation := r.URL.Query().Get("operation")

	result, err := s.svc.Calculate(ctx, float32(a), float32(b), operation)
	if err != nil {
		return err
	}

	resultResponce := types.ResultResponce{
		A:         float32(a),
		B:         float32(b),
		Operation: operation,
		Result:    result,
	}

	return writeJSON(w, http.StatusOK, &resultResponce)
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
