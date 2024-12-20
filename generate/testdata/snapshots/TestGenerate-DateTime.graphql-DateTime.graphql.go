// Code generated by github.com/valstro/genqlient, DO NOT EDIT.

package test

import (
	"time"

	"github.com/valstro/genqlient/graphql"
)

// __convertTimezoneInput is used internally by genqlient
type __convertTimezoneInput struct {
	Dt time.Time `js:"dt" json:"dt"`
	Tz string    `js:"tz" json:"tz"`
}

// GetDt returns __convertTimezoneInput.Dt, and is useful for accessing the field via an interface.
func (v *__convertTimezoneInput) GetDt() time.Time { return v.Dt }

// GetTz returns __convertTimezoneInput.Tz, and is useful for accessing the field via an interface.
func (v *__convertTimezoneInput) GetTz() string { return v.Tz }

// convertTimezoneResponse is returned by convertTimezone on success.
type convertTimezoneResponse struct {
	Convert time.Time `js:"convert" json:"convert"`
}

// GetConvert returns convertTimezoneResponse.Convert, and is useful for accessing the field via an interface.
func (v *convertTimezoneResponse) GetConvert() time.Time { return v.Convert }

// The query executed by convertTimezone.
const convertTimezone_Operation = `
query convertTimezone ($dt: DateTime!, $tz: String) {
	convert(dt: $dt, tz: $tz)
}
`

func convertTimezone(
	client_ graphql.Client,
	dt time.Time,
	tz string,
) (data_ *convertTimezoneResponse, err_ error) {
	req_ := &graphql.Request{
		OpName: "convertTimezone",
		Query:  convertTimezone_Operation,
		Variables: &__convertTimezoneInput{
			Dt: dt,
			Tz: tz,
		},
	}

	data_ = &convertTimezoneResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		nil,
		req_,
		resp_,
	)

	return data_, err_
}

