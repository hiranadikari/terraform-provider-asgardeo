// Package clients constructs the Asgardeo API client from provider configuration.
package clients

import "github.com/asgardeo/terraform-provider-asgardeo/asgardeo"

// AsgardeoClient is an alias so that the rest of the provider imports one type.
type AsgardeoClient = asgardeo.Client

// Build creates a new AsgardeoClient from the three provider-level credentials.
func Build(orgName, clientID, clientSecret string) *AsgardeoClient {
	return asgardeo.NewClient(orgName, clientID, clientSecret)
}
