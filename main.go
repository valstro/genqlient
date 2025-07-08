// genqlient is a GraphQL client generator for Go.
//
// To run genqlient:
//
//	go run github.com/valstro/genqlient
//
// For programmatic access, see the "generate" package, below.  For
// user documentation, see the project [GitHub].
//
// [GitHub]: https://github.com/valstro/genqlient
package main

import (
	"fmt"

	"github.com/valstro/genqlient/generate"
)

func main() {
	fmt.Println("Starting genqlient code generation...")
	generate.Main()
	fmt.Println("Code generation completed successfully")
}
