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

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
	})
}

func main() {
	config := config{}
	envconfig.MustProcess("", &config)

	DB := db.MustConnect(config.DatabaseURL)

	s := db.NewStore(DB)
	resolver := graph.NewResolver(s)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	srvWithDatalaoder := dataloader.Middleware(*s, srv)
	srvWIthCOrs := cors(srvWithDatalaoder)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srvWIthCOrs)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
