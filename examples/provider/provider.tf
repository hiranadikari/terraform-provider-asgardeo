terraform {
  required_providers {
    asgardeo = {
      source  = "asgardeo/asgardeo"
      version = "~> 0.1"
    }
  }
}

# Credentials can be passed as arguments or via environment variables:
#   ASGARDEO_ORG_NAME, ASGARDEO_CLIENT_ID, ASGARDEO_CLIENT_SECRET
provider "asgardeo" {
  org_name      = var.asgardeo_org_name
  client_id     = var.asgardeo_client_id
  client_secret = var.asgardeo_client_secret
}
