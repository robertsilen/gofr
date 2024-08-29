package main

import (
	"log"

	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/container"
)

type User struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserResolver struct {
	user *User
}

func (r *UserResolver) ID() int32 {
	return r.user.ID
}

func (r *UserResolver) Name() string {
	return r.user.Name
}

func (r *UserResolver) Email() string {
	return r.user.Email
}

type MyResolver struct {
	Container *container.Container
}

// Resolver for the user query
func (r *MyResolver) User(args struct{ ID int32 }) (*UserResolver, error) {
	row := r.Container.SQL.QueryRow("SELECT name, email FROM users WHERE id = ?", args.ID)
	if row == nil {
		return nil, nil
	}

	var user User

	err := row.Scan(&user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	user.ID = int32(args.ID)
	return &UserResolver{user: &user}, nil
}

// Resolver for the addUser mutation
func (r *MyResolver) AddUser(args struct{ Name, Email string }) (*UserResolver, error) {
	result, err := r.Container.SQL.Exec("INSERT INTO users (name, email) VALUES (?, ?)", args.Name, args.Email)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	user := &User{ID: int32(int(id)), Name: args.Name, Email: args.Email}
	return &UserResolver{user: user}, nil
}

const schemaString = `
	type User {
		id: Int!
		name: String!
		email: String!
	}

	type Query {
		user(id: Int!): User!
	}

	type Mutation {
		addUser(name: String!, email: String!): User!
	}
`

func main() {
	// Create a new application
	a := gofr.New()

	// Register the updated GraphQL schema
	if err := a.RegisterGraphQLSchema(schemaString, func(container *container.Container) interface{} {
		return &MyResolver{Container: container}
	}); err != nil {
		log.Fatalf("Failed to register GraphQL schema: %v", err)
	}

	// Add all the routes
	a.GET("/hello", HelloHandler)

	// Run the application
	a.Run()
}

func HelloHandler(c *gofr.Context) (interface{}, error) {
	return "Hello, World!", nil
}
