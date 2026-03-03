# output "application_id" {
#   description = "Asgardeo application ID for the Rancher OIDC app."
#   value       = asgardeo_application.rancher.id
# }

# output "client_id" {
#   description = "OAuth2 client ID — paste into Rancher's OIDC provider settings."
#   value       = asgardeo_application.rancher.client_id
# }

# output "client_secret" {
#   description = "OAuth2 client secret — paste into Rancher's OIDC provider settings."
#   value       = asgardeo_application.rancher.client_secret
#   sensitive   = true
# }

output "issuer_url" {
  description = "OIDC issuer URL to configure in Rancher."
  value       = "https://api.asgardeo.io/t/${var.asgardeo_org_name}/oauth2/token"
}

output "discovery_url" {
  description = "OIDC discovery endpoint — use to auto-populate Rancher OIDC settings."
  value       = "https://api.asgardeo.io/t/${var.asgardeo_org_name}/oauth2/token/.well-known/openid-configuration"
}
