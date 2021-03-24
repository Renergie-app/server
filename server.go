package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"log"
	"renergie-server/graph"
	"renergie-server/graph/generated"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gorilla/mux"
)

var muxAdapter *gorillamux.GorillaMuxAdapter

func init() {
	r := mux.NewRouter()

	// From server.go
	schema := generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})
	server := handler.NewDefaultServer(schema)
	r.Handle("/query", server)
	r.Handle("/", playground.Handler("GraphQL playground", "/query"))

	muxAdapter = gorillamux.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	rsp, err := muxAdapter.Proxy(req)
	if err != nil {
		log.Println(err)
	}
	return rsp, err
}

func main() {
	lambda.Start(Handler)

}
