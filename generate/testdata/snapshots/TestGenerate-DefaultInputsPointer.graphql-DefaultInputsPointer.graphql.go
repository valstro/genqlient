// Code generated by github.com/valstro/genqlient, DO NOT EDIT.

package test

import (
	"github.com/valstro/genqlient/graphql"
)

// DefaultInputsResponse is returned by DefaultInputs on success.
type DefaultInputsResponse struct {
	Default bool `js:"default" json:"default"`
}

// GetDefault returns DefaultInputsResponse.Default, and is useful for accessing the field via an interface.
func (v *DefaultInputsResponse) GetDefault() bool { return v.Default }

type InputWithDefaults struct {
	Field         string  `js:"field" json:"field"`
	NullableField *string `js:"nullableField" json:"nullableField"`
}

// GetField returns InputWithDefaults.Field, and is useful for accessing the field via an interface.
func (v *InputWithDefaults) GetField() string { return v.Field }

// GetNullableField returns InputWithDefaults.NullableField, and is useful for accessing the field via an interface.
func (v *InputWithDefaults) GetNullableField() *string { return v.NullableField }

// __DefaultInputsInput is used internally by genqlient
type __DefaultInputsInput struct {
	Input InputWithDefaults `js:"input" json:"input"`
}

// GetInput returns __DefaultInputsInput.Input, and is useful for accessing the field via an interface.
func (v *__DefaultInputsInput) GetInput() InputWithDefaults { return v.Input }

// The query executed by DefaultInputs.
const DefaultInputs_Operation = `
query DefaultInputs ($input: InputWithDefaults!) {
	default(input: $input)
}
`

// The `InputWithDefaults.field` cannot be `pointer: true`, together with implicit `omitempty: false`, as `null` is
// not a valid value there. However, nullableField should still be ok
// (this will send null, overwriting the server's default)
func DefaultInputs(
	client_ graphql.Client,
	input InputWithDefaults,
) (data_ *DefaultInputsResponse, err_ error) {
	req_ := &graphql.Request{
		OpName: "DefaultInputs",
		Query:  DefaultInputs_Operation,
		Variables: &__DefaultInputsInput{
			Input: input,
		},
	}

	data_ = &DefaultInputsResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		nil,
		req_,
		resp_,
	)

	return data_, err_
}

