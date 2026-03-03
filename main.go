package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/asgardeo/terraform-provider-asgardeo/internal/provider"
)

// version is set at release time by goreleaser via -ldflags.
var version = "dev"

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "Enable provider debug mode (for use with Terraform debugger).")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/asgardeo/asgardeo",
		Debug:   debug,
	}

	if err := providerserver.Serve(context.Background(), provider.New(version), opts); err != nil {
		log.Fatal(err)
	}
}
