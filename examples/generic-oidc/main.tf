terraform {
  required_providers {
    asgardeo = {
      source  = "asgardeo/asgardeo"
      version = "~> 0.1"
    }
  }
}

provider "asgardeo" {
  org_name      = var.asgardeo_org_name
  client_id     = var.asgardeo_client_id
  client_secret = var.asgardeo_client_secret
}

# ──────────────────────────────────────────────────────────────────────────────
# Generic OIDC application with Authorization Code + PKCE flow.
# Suitable for any web application that supports standard OIDC discovery.
# ──────────────────────────────────────────────────────────────────────────────
resource "asgardeo_application" "generic_oidc" {
  name        = var.app_name
  description = "Generic OIDC application managed by Terraform."
  access_url  = var.app_callback_url

  oidc {
    grant_types          = ["authorization_code", "refresh_token"]
    callback_urls        = [var.app_callback_url]
    allowed_origins      = [replace(var.app_callback_url, "/callback", "")]
    logout_redirect_urls = var.app_logout_url != "" ? [var.app_logout_url] : []

    pkce {
      mandatory                         = true
      support_plain_transform_algorithm = false
    }

    access_token {
      type                             = "JWT"
      user_access_token_expiry_seconds = 3600
    }

    refresh_token {
      expiry_seconds      = 86400
      renew_refresh_token = true
    }
  }

  advanced {
    skip_login_consent  = false
    skip_logout_consent = false
  }
}
