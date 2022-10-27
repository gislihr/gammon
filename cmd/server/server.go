package main

import (
	"encoding/json"
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
	DatabaseURL string `envconfig:"DATABASE_URL" required:"true"`
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST")

			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization")
			return
		} else {
			h.ServeHTTP(w, r)
		}
	})
}

func main() {
	config := config{}
	envconfig.MustProcess("", &config)

	database := db.MustConnect(config.DatabaseURL)

	s := db.NewStore(database)
	resolver := graph.NewResolver(s)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	srvWithDatalaoder := dataloader.Middleware(*s, srv)
	srvWIthCOrs := cors(srvWithDatalaoder)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srvWIthCOrs)
	http.HandleFunc("/healthy", func(w http.ResponseWriter, r *http.Request) {
		type response struct {
			Healthy bool `json:"healthy"`
		}

		res, err := s.Healthy()
		if err == nil && res == 1 {
			respond(response{Healthy: true}, w)
			return
		}

		respond(response{Healthy: false}, w)
		return
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}

//Respond marshals data as json and writes to response writer
func respond(data interface{}, w http.ResponseWriter) {
	bytes, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
