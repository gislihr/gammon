package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gislihr/gammon/graph"
	"github.com/gislihr/gammon/graph/generated"
	"github.com/gislihr/gammon/pkg/gammon/db"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Port       string `envconfig:"PORT" default:"8080"`
	DBHost     string `envconfig:"DB_HOST" required:"true"`
	DBPort     string `envconfig:"DB_PORT" required:"true"`
	DBUser     string `envconfig:"DB_USER" required:"true"`
	DBName     string `envconfig:"DB_NAME" required:"true"`
	DBPassword string `envconfig:"DB_PASS"`
	DBSSL      bool   `envconfig:"DB_USE_SSL"`
}

func main() {
	config := config{}
	envconfig.MustProcess("", &config)

	DB := db.MustConnect(db.Options{
		Host:     config.DBHost,
		Port:     config.DBPort,
		User:     config.DBUser,
		Password: config.DBPassword,
		Database: config.DBName,
		UseSSL:   config.DBSSL,
	})

	s := db.NewStore(DB)
	resolver := graph.NewResolver(s)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
