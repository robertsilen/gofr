package gofr

import (
	"encoding/json"
	"fmt"
	"github.com/graph-gophers/graphql-go"
	"log"
	"net/http"
	"strings"
)

type GraphQLHandler struct {
	Schema *graphql.Schema
}

func (s *httpServer) RegisterGraphQLSchema(schemaString string, resolver interface{}) error {
	if s.graphQLHandler != nil && s.graphQLHandler.Schema != nil {
		return fmt.Errorf("GraphQL schema is already registered")
	}

	schema, err := graphql.ParseSchema(schemaString, resolver)
	if err != nil {
		return fmt.Errorf("failed to parse schema: %w", err)
	}

	s.graphQLHandler = &GraphQLHandler{
		Schema: schema,
	}

	log.Printf("GraphQL schema registered: %s", schemaString)

	// Add the GraphQL handler to the router
	s.router.Handle("/graphql", s.graphQLHandler)

	return nil
}

func (h *GraphQLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Check for GET request
	if r.Method == http.MethodGet {
		query := r.URL.Query().Get("query")
		variables := r.URL.Query().Get("variables")
		operationName := r.URL.Query().Get("operationName")

		operationName = strings.TrimSpace(operationName)

		// Decode variables if present
		var vars map[string]interface{}
		if variables != "" {
			if err := json.Unmarshal([]byte(variables), &vars); err != nil {
				http.Error(w, `{"error": "Invalid variables format"}`, http.StatusBadRequest)
				return
			}
		}

		req := struct {
			Query         string                 `json:"query"`
			Variables     map[string]interface{} `json:"variables"`
			OperationName string                 `json:"operationName"`
		}{
			Query:         query,
			Variables:     vars,
			OperationName: operationName,
		}

		// Execute the query
		response := h.Schema.Exec(r.Context(), req.Query, req.OperationName, req.Variables)
		if len(response.Errors) > 0 {
			log.Printf("GraphQL execution errors: %v", response.Errors)
		}

		// Marshal the response into JSON
		data, err := json.Marshal(response)
		if err != nil {
			http.Error(w, `{"error": "Failed to marshal response"}`, http.StatusInternalServerError)
			return
		}

		// Write the response
		w.Write(data)
		return
	}

	// Handle POST requests (existing logic)
	// Read the request body
	var req struct {
		Query         string                 `json:"query"`
		Variables     map[string]interface{} `json:"variables"`
		OperationName string                 `json:"operationName"`
	}

	// Check if the request body is empty
	if r.Body == nil {
		http.Error(w, `{"error": "Request body is missing"}`, http.StatusBadRequest)
		return
	}

	// Decode the JSON request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	log.Printf("Received GraphQL request: %+v", req)

	// Check for empty query
	if req.Query == "" {
		http.Error(w, `{"error": "GraphQL query is missing"}`, http.StatusBadRequest)
		return
	}

	// Check if schema is registered
	if h.Schema == nil {
		http.Error(w, `{"error": "Schema not registered"}`, http.StatusInternalServerError)
		return
	}

	// Execute the query
	response := h.Schema.Exec(r.Context(), req.Query, req.OperationName, req.Variables)
	if len(response.Errors) > 0 {
		log.Printf("GraphQL execution errors: %v", response.Errors)
	}

	// Marshal the response into JSON
	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, `{"error": "Failed to marshal response"}`, http.StatusInternalServerError)
		return
	}

	// Write the response
	w.Write(data)
}
