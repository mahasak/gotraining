// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./exercise1

// Sample program that creates a pachyderm data repository.
package main

import (
	"log"

	"github.com/pachyderm/pachyderm/src/client"
)

func main() {

	// Connect to Pachyderm on our localhost.  By default
	// Pachyderm will be exposed on port 30650.
	c, err := client.NewFromAddress("0.0.0.0:30650")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Create a data repository called "diabetes."

	// Now, we will list all the current data repositories
	// again on the Pachyderm cluster.  We should now have
	// a single data repository called iris.

	// Check that the number of repos is what we expect.

	// Check that the name of the repo is what we expect.
}
