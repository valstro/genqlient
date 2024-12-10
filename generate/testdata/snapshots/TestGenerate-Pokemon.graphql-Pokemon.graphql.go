// Code generated by github.com/valstro/genqlient, DO NOT EDIT.

package test

import (
	"github.com/valstro/genqlient/graphql"
	"github.com/valstro/genqlient/internal/testutil"
)

// GetPokemonSiblingsResponse is returned by GetPokemonSiblings on success.
type GetPokemonSiblingsResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User GetPokemonSiblingsUser `js:"user" json:"user"`
}

// GetUser returns GetPokemonSiblingsResponse.User, and is useful for accessing the field via an interface.
func (v *GetPokemonSiblingsResponse) GetUser() GetPokemonSiblingsUser { return v.User }

// GetPokemonSiblingsUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type GetPokemonSiblingsUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id               string                                   `js:"id" json:"id"`
	Roles            []string                                 `js:"roles" json:"roles"`
	Name             string                                   `js:"name" json:"name"`
	Pokemon          []testutil.Pokemon                       `js:"pokemon" json:"pokemon"`
	GenqlientPokemon []GetPokemonSiblingsUserGenqlientPokemon `js:"genqlientPokemon" json:"genqlientPokemon"`
}

// GetId returns GetPokemonSiblingsUser.Id, and is useful for accessing the field via an interface.
func (v *GetPokemonSiblingsUser) GetId() string { return v.Id }

// GetRoles returns GetPokemonSiblingsUser.Roles, and is useful for accessing the field via an interface.
func (v *GetPokemonSiblingsUser) GetRoles() []string { return v.Roles }

// GetName returns GetPokemonSiblingsUser.Name, and is useful for accessing the field via an interface.
func (v *GetPokemonSiblingsUser) GetName() string { return v.Name }

// GetPokemon returns GetPokemonSiblingsUser.Pokemon, and is useful for accessing the field via an interface.
func (v *GetPokemonSiblingsUser) GetPokemon() []testutil.Pokemon { return v.Pokemon }

// GetGenqlientPokemon returns GetPokemonSiblingsUser.GenqlientPokemon, and is useful for accessing the field via an interface.
func (v *GetPokemonSiblingsUser) GetGenqlientPokemon() []GetPokemonSiblingsUserGenqlientPokemon {
	return v.GenqlientPokemon
}

// GetPokemonSiblingsUserGenqlientPokemon includes the requested fields of the GraphQL type Pokemon.
type GetPokemonSiblingsUserGenqlientPokemon struct {
	Species string `js:"species" json:"species"`
	Level   int    `js:"level" json:"level"`
}

// GetSpecies returns GetPokemonSiblingsUserGenqlientPokemon.Species, and is useful for accessing the field via an interface.
func (v *GetPokemonSiblingsUserGenqlientPokemon) GetSpecies() string { return v.Species }

// GetLevel returns GetPokemonSiblingsUserGenqlientPokemon.Level, and is useful for accessing the field via an interface.
func (v *GetPokemonSiblingsUserGenqlientPokemon) GetLevel() int { return v.Level }

// __GetPokemonSiblingsInput is used internally by genqlient
type __GetPokemonSiblingsInput struct {
	Input testutil.Pokemon `js:"input" json:"input"`
}

// GetInput returns __GetPokemonSiblingsInput.Input, and is useful for accessing the field via an interface.
func (v *__GetPokemonSiblingsInput) GetInput() testutil.Pokemon { return v.Input }

// The query executed by GetPokemonSiblings.
const GetPokemonSiblings_Operation = `
query GetPokemonSiblings ($input: PokemonInput!) {
	user(query: {hasPokemon:$input}) {
		id
		roles
		name
		pokemon {
			species
			level
		}
		genqlientPokemon: pokemon {
			species
			level
		}
	}
}
`

func GetPokemonSiblings(
	client_ graphql.Client,
	input testutil.Pokemon,
) (data_ *GetPokemonSiblingsResponse, err_ error) {
	req_ := &graphql.Request{
		OpName: "GetPokemonSiblings",
		Query:  GetPokemonSiblings_Operation,
		Variables: &__GetPokemonSiblingsInput{
			Input: input,
		},
	}

	data_ = &GetPokemonSiblingsResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		nil,
		req_,
		resp_,
	)

	return data_, err_
}

