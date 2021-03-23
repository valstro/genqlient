package test

// Code generated by github.com/Khan/genql, DO NOT EDIT.

import (
	"context"

	"github.com/Khan/genql/graphql"
)

type QueryWithSlicesResponse struct {
	User QueryWithSlicesUser `json:"user"`
}

type QueryWithSlicesUser struct {
	Emails                []string `json:"emails"`
	EmailsOrNull          []string `json:"emailsOrNull"`
	EmailsWithNulls       []string `json:"emailsWithNulls"`
	EmailsWithNullsOrNull []string `json:"emailsWithNullsOrNull"`
}

func QueryWithSlices(client *graphql.Client) (*QueryWithSlicesResponse, error) {
	var retval QueryWithSlicesResponse
	err := client.MakeRequest(context.Background(), `
query QueryWithSlices {
	user {
		emails
		emailsOrNull
		emailsWithNulls
		emailsWithNullsOrNull
	}
}
`, &retval, nil)
	return &retval, err
}
