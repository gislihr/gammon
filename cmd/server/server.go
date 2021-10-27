package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gislihr/gammon/graph"
	"github.com/gislihr/gammon/graph/generated"
	"github.com/gislihr/gammon/store"
	"github.com/jmoiron/sqlx"
)

const defaultPort = "8080"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := sqlx.Connect("postgres", "host=localhost port=5432 sslmode=disable user=user password=pass database=gammon")
	if err != nil {
		panic(err)
	}

	s, err := store.NewStore(db)
	if err != nil {
		panic(err)
	}

	resolver := graph.NewResolver(s)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
