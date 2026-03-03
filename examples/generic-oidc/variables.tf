variable "asgardeo_org_name" {
  description = "Asgardeo organisation name (e.g. mycompany)."
  type        = string
}

variable "asgardeo_client_id" {
  description = "Client ID of the M2M application used to authenticate."
  type        = string
  sensitive   = true
}

variable "asgardeo_client_secret" {
  description = "Client secret of the M2M application."
  type        = string
  sensitive   = true
}

variable "app_name" {
  description = "Display name for the new application."
  type        = string
  default     = "generic-oidc-app"
}

variable "app_callback_url" {
  description = "OAuth2 redirect URI for the application."
  type        = string
}

variable "app_logout_url" {
  description = "Post-logout redirect URI."
  type        = string
  default     = ""
}
