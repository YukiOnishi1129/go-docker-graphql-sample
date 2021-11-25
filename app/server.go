package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/YukiOnishi1129/go-docker-graphql-sample/app/db"
	"github.com/YukiOnishi1129/go-docker-graphql-sample/app/graphql/generated"
	"github.com/YukiOnishi1129/go-docker-graphql-sample/app/graphql/resolver"
)

const defaultPort = "4000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbCon := db.Init()
	defer db.CloseDB(dbCon)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{DB: dbCon}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
