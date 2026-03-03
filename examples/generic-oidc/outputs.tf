output "application_id" {
  description = "Asgardeo application ID."
  value       = asgardeo_application.generic_oidc.id
}

output "client_id" {
  description = "OAuth2 client ID — use this in your application's OIDC config."
  value       = asgardeo_application.generic_oidc.client_id
}

output "client_secret" {
  description = "OAuth2 client secret — treat as a secret."
  value       = asgardeo_application.generic_oidc.client_secret
  sensitive   = true
}

output "oidc_discovery_url" {
  description = "OIDC discovery endpoint for this organisation."
  value       = "https://api.asgardeo.io/t/${var.asgardeo_org_name}/oauth2/token/.well-known/openid-configuration"
}
