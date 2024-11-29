// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package queries

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// ListInputQueryResponse is returned by ListInputQuery on success.
type ListInputQueryResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User ListInputQueryUser `js:"user" json:"user"`
}

// GetUser returns ListInputQueryResponse.User, and is useful for accessing the field via an interface.
func (v *ListInputQueryResponse) GetUser() ListInputQueryUser { return v.User }

// ListInputQueryUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type ListInputQueryUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id string `js:"id" json:"id"`
}

// GetId returns ListInputQueryUser.Id, and is useful for accessing the field via an interface.
func (v *ListInputQueryUser) GetId() string { return v.Id }

// QueryWithSlicesResponse is returned by QueryWithSlices on success.
type QueryWithSlicesResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User QueryWithSlicesUser `js:"user" json:"user"`
}

// GetUser returns QueryWithSlicesResponse.User, and is useful for accessing the field via an interface.
func (v *QueryWithSlicesResponse) GetUser() QueryWithSlicesUser { return v.User }

// QueryWithSlicesUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type QueryWithSlicesUser struct {
	Emails                []string `js:"emails" json:"emails"`
	EmailsOrNull          []string `js:"emailsOrNull" json:"emailsOrNull"`
	EmailsWithNulls       []string `js:"emailsWithNulls" json:"emailsWithNulls"`
	EmailsWithNullsOrNull []string `js:"emailsWithNullsOrNull" json:"emailsWithNullsOrNull"`
}

// GetEmails returns QueryWithSlicesUser.Emails, and is useful for accessing the field via an interface.
func (v *QueryWithSlicesUser) GetEmails() []string { return v.Emails }

// GetEmailsOrNull returns QueryWithSlicesUser.EmailsOrNull, and is useful for accessing the field via an interface.
func (v *QueryWithSlicesUser) GetEmailsOrNull() []string { return v.EmailsOrNull }

// GetEmailsWithNulls returns QueryWithSlicesUser.EmailsWithNulls, and is useful for accessing the field via an interface.
func (v *QueryWithSlicesUser) GetEmailsWithNulls() []string { return v.EmailsWithNulls }

// GetEmailsWithNullsOrNull returns QueryWithSlicesUser.EmailsWithNullsOrNull, and is useful for accessing the field via an interface.
func (v *QueryWithSlicesUser) GetEmailsWithNullsOrNull() []string { return v.EmailsWithNullsOrNull }

// __ListInputQueryInput is used internally by genqlient
type __ListInputQueryInput struct {
	Names []string `js:"names" json:"names"`
}

// GetNames returns __ListInputQueryInput.Names, and is useful for accessing the field via an interface.
func (v *__ListInputQueryInput) GetNames() []string { return v.Names }

// The query executed by ListInputQuery.
const ListInputQuery_Operation = `
query ListInputQuery ($names: [String]) {
	user(query: {names:$names}) {
		id
	}
}
`

func ListInputQuery(
	ctx_ context.Context,
	client_ graphql.Client,
	names []string,
) (data_ *ListInputQueryResponse, err_ error) {
	req_ := &graphql.Request{
		OpName: "ListInputQuery",
		Query:  ListInputQuery_Operation,
		Variables: &__ListInputQueryInput{
			Names: names,
		},
	}

	data_ = &ListInputQueryResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return data_, err_
}

// The query executed by QueryWithSlices.
const QueryWithSlices_Operation = `
query QueryWithSlices {
	user {
		emails
		emailsOrNull
		emailsWithNulls
		emailsWithNullsOrNull
	}
}
`

func QueryWithSlices(
	ctx_ context.Context,
	client_ graphql.Client,
) (data_ *QueryWithSlicesResponse, err_ error) {
	req_ := &graphql.Request{
		OpName: "QueryWithSlices",
		Query:  QueryWithSlices_Operation,
	}

	data_ = &QueryWithSlicesResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return data_, err_
}

