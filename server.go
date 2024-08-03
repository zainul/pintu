package main

import (
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/zainul/pintu/graph"
)

const defaultPort = "8089"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//plugin.Generate("company")

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	srv.Use(extension.FixedComplexityLimit(100)) // This line is key
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
