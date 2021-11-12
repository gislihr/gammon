package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gislihr/gammon/graph"
	"github.com/gislihr/gammon/graph/generated"
	"github.com/gislihr/gammon/pkg/gammon/dataloader"
	"github.com/gislihr/gammon/pkg/gammon/db"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Port        string `envconfig:"PORT" default:"8080"`
	DatabaseURL string `envconfig:"DATABASE_URL" required:"true'`
}

func main() {
	config := config{}
	envconfig.MustProcess("", &config)

	DB := db.MustConnect(config.DatabaseURL)

	s := db.NewStore(DB)
	resolver := graph.NewResolver(s)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	srvWithDatalaoder := dataloader.Middleware(*s, srv)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srvWithDatalaoder)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
