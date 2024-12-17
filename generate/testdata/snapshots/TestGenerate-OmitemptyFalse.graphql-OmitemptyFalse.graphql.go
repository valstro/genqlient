// Code generated by github.com/valstro/genqlient, DO NOT EDIT.

package test

import (
	"github.com/valstro/genqlient/graphql"
)

// OmitemptyFalseResponse is returned by OmitemptyFalse on success.
type OmitemptyFalseResponse struct {
	Omitempty bool `js:"omitempty" json:"omitempty"`
}

// GetOmitempty returns OmitemptyFalseResponse.Omitempty, and is useful for accessing the field via an interface.
func (v *OmitemptyFalseResponse) GetOmitempty() bool { return v.Omitempty }

type OmitemptyInput struct {
	Field         string `js:"field" json:"field"`
	NullableField string `js:"nullableField" json:"nullableField,omitempty"`
}

// GetField returns OmitemptyInput.Field, and is useful for accessing the field via an interface.
func (v *OmitemptyInput) GetField() string { return v.Field }

// GetNullableField returns OmitemptyInput.NullableField, and is useful for accessing the field via an interface.
func (v *OmitemptyInput) GetNullableField() string { return v.NullableField }

// __OmitemptyFalseInput is used internally by genqlient
type __OmitemptyFalseInput struct {
	Input OmitemptyInput `js:"input" json:"input,omitempty"`
}

// GetInput returns __OmitemptyFalseInput.Input, and is useful for accessing the field via an interface.
func (v *__OmitemptyFalseInput) GetInput() OmitemptyInput { return v.Input }

// The query executed by OmitemptyFalse.
const OmitemptyFalse_Operation = `
query OmitemptyFalse ($input: OmitemptyInput) {
	omitempty(input: $input)
}
`

func OmitemptyFalse(
	client_ graphql.Client,
	input OmitemptyInput,
) (data_ *OmitemptyFalseResponse, err_ error) {
	req_ := &graphql.Request{
		OpName: "OmitemptyFalse",
		Query:  OmitemptyFalse_Operation,
		Variables: &__OmitemptyFalseInput{
			Input: input,
		},
	}

	data_ = &OmitemptyFalseResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		nil,
		req_,
		resp_,
	)

	return data_, err_
}
