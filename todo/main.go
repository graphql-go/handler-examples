package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql/examples/todo/schema"
	"github.com/graphql-go/handler"
)

func main() {
	h := handler.New(&handler.Config{
		Schema:     &schema.TodoSchema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	http.Handle("/graphql", h)
	log.Println("GraphQL Playground running on localhost:8080/graphql")
	http.ListenAndServe(":8080", nil)
}
