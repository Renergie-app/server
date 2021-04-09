package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/handlers"
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
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
	)
	r.Use(cors)
	muxAdapter = gorillamux.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	rsp, err := muxAdapter.Proxy(req)
	if err != nil {
		log.Println(err)
	}
	//rsp.Headers["Access-Control-Allow-Origin"] = "*"
	//rsp.Headers["Access-Control-Allow-Methods"] = "GET, POST, OPTIONS,PUT"
	return rsp, err
}

func main() {
	lambda.Start(Handler)

}
