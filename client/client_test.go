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
	query := `
	query findTodos {
	  	todos {
	      text
	      done
	      user {
	        name
	      }
	    }
	}
	`

	response, err := makeGraphQlRequest(query)
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("Request %v", response)
}

func TestGraphQlMutation(t *testing.T) {
	query := `
mutation createTodo {
  createTodo(input:{text:"Newtodo", userId:"3"}) {
    user {
      id
    }
    done
    text
  }
}`

	response, err := makeGraphQlRequest(query)
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("Request %v", response)
}

func makeGraphQlRequest(query string) (*response, error) {
	client := graphql.NewClient("http://localhost:8080/query")
	request := graphql.NewRequest(query)
	ctx := context.Background()
	var response response
	if err := client.Run(ctx, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
