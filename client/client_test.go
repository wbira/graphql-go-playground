package client

import (
	"context"
	"fmt"
	"testing"
)
import "github.com/machinebox/graphql"

type todo struct {
	Text string
	Done bool
	User struct {
		Name string
	}
}

type response struct {
	Todos []todo
}

func TestGraphQlQuery(t *testing.T) {
	client := graphql.NewClient("http://localhost:8080/query")

	request := graphql.NewRequest(`
query findTodos {
  	todos {
      text
      done
      user {
        name
      }
    }
}
	`)

	ctx := context.Background()
	var response response
	if err := client.Run(ctx, request, &response); err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("Request %v", response)
}
