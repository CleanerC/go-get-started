package v1_transport

import (
	v1_endpoint "BasicUsedHTTP/pkg/v1/endpoint"
	v1_service "BasicUsedHTTP/pkg/v1/service"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type errorWrapper struct {
	Error string `json:"error"`
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

func decodeHTTPADDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var(
		in v1_service.Add
		err error
	)
	in.A, _ = strconv.Atoi(r.FormValue("a"))
	in.B, err = strconv.Atoi(r.FormValue("b"))

	if err != nil {
		return in, err
	}

	return in, nil
}

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func NewHttpHandle(endpoint v1_endpoint.EndPointServer) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
	}

	m := http.NewServeMux()

	m.Handle("/add", httptransport.NewServer(
		endpoint.AddEndPoint,
		decodeHTTPADDRequest,
		encodeHTTPGenericResponse,
		options...,
	))

	return m
}