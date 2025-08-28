package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SaiTejaBandamidi/go-graphql-sqlc-api/config"
	"github.com/SaiTejaBandamidi/go-graphql-sqlc-api/db"
	"github.com/SaiTejaBandamidi/go-graphql-sqlc-api/graph"
	"github.com/SaiTejaBandamidi/go-graphql-sqlc-api/internal"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	// Load environment variables from .env
	config.LoadEnv()

	// Build the DB connection string using individual .env params
	dbSource := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Connect to the database
	dbConn := internal.ConnectDB(dbSource)

	// Initialize queries with the database connection
	queries := db.New(dbConn)
	if queries == nil {
		log.Fatal("Failed to create database queries")
	}

	// Create a new GraphQL resolver with DB connection and queries
	resolver := &graph.Resolver{
		DB:      dbConn,
		Queries: queries,
	}

	// Create the GraphQL server
	server := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: resolver,
	}))

	// Setup GraphQL handlers
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", server)

	// Start the server
	port := os.Getenv("PORT")
	log.Printf("🚀 GraphQL Playground ready at http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
