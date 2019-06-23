package main

import (
	"context"
	"fmt"
	"hamburguer/handlers"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	adapter "github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
)

var gorillalambda *adapter.GorillaMuxAdapter

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print(w, `{"teste" : "ok"}`)
	})
	r.HandleFunc("/company", handlers.CreateCompanyHandler).Methods("POST")

	gorillalambda = adapter.New(r)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return gorillalambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(handler)
}
