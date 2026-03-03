# Minimal OIDC application — used by terraform-plugin-docs to generate docs.
resource "asgardeo_application" "example" {
  name        = "my-web-app"
  description = "Example OIDC web application"
  access_url  = "https://app.example.com/login"

  oidc {
    grant_types    = ["authorization_code", "refresh_token"]
    callback_urls  = ["https://app.example.com/callback"]
    allowed_origins = ["https://app.example.com"]

    pkce {
      mandatory = true
    }
  }

  advanced {
    skip_login_consent = true
  }
}

output "client_id" {
  value = asgardeo_application.example.client_id
}

output "client_secret" {
  value     = asgardeo_application.example.client_secret
  sensitive = true
}
