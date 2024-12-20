// Code generated by github.com/valstro/genqlient, DO NOT EDIT.

package test

import (
	"github.com/valstro/genqlient/graphql"
)

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
	client_ graphql.Client,
) (data_ *QueryWithSlicesResponse, err_ error) {
	req_ := &graphql.Request{
		OpName: "QueryWithSlices",
		Query:  QueryWithSlices_Operation,
	}

	data_ = &QueryWithSlicesResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		nil,
		req_,
		resp_,
	)

	return data_, err_
}

