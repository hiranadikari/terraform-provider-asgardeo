variable "asgardeo_org_name" {
  description = "Asgardeo organisation name."
  type        = string
}

variable "asgardeo_client_id" {
  description = "Client ID of the Asgardeo M2M app used by this provider."
  type        = string
  sensitive   = true
}

variable "asgardeo_client_secret" {
  description = "Client secret of the Asgardeo M2M app."
  type        = string
  sensitive   = true
}

variable "rancher_url" {
  description = "Base URL of the Rancher Manager instance (e.g. https://rancher.example.com)."
  type        = string
}

variable "app_name" {
  description = "Display name for the Rancher OIDC application in Asgardeo."
  type        = string
  default     = "rancher-oidc"
}

variable "skip_consent" {
  description = "Skip the Asgardeo login/logout consent screens (recommended for machine accounts)."
  type        = bool
  default     = true
}
