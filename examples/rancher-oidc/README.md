# Rancher OIDC Federation via Asgardeo

This example registers a Rancher Manager instance as an OIDC relying party in
Asgardeo and outputs the `client_id` / `client_secret` needed to complete the
Rancher **Generic OIDC** authentication setup.

## Prerequisites

| What | Where |
|---|---|
| Asgardeo M2M application | Asgardeo Console → Applications → New Application → M2M |
| Scopes authorised | `internal_application_mgt_create`, `…_view`, `…_update`, `…_delete` |
| Rancher Manager running | `https://rancher.example.com` (any reachable URL) |

## Usage

```bash
cd examples/rancher-oidc

terraform init

terraform apply \
  -var="asgardeo_org_name=myorg" \
  -var="asgardeo_client_id=<m2m-client-id>" \
  -var="asgardeo_client_secret=<m2m-client-secret>" \
  -var="rancher_url=https://rancher.example.com"
```

After apply, note the outputs:

```
client_id     = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
client_secret = <sensitive>
issuer_url    = "https://api.asgardeo.io/t/myorg/oauth2/token"
```

## Rancher Configuration

In **Rancher → ☰ → Users & Authentication → Auth Provider → Generic OIDC**:

| Field | Value |
|---|---|
| Issuer | `https://api.asgardeo.io/t/<org_name>/oauth2/token` |
| Client ID | from `terraform output client_id` |
| Client Secret | from `terraform output -raw client_secret` |
| Scopes | `openid profile email groups` |
| Username Attribute | `email` |
| Display Name Attribute | `name` |

Rancher will auto-discover the authorization, token, and JWKS endpoints via
the well-known discovery URL.

## Integration with harvester-dc-terraform

If you are using the `harvester-dc-terraform` repo, place this call in a new
phase (e.g. `04-asgardeo-auth`) and read the Rancher URL from the existing
remote state output of `02-management`:

```hcl
data "terraform_remote_state" "management" {
  backend = "local"
  config  = { path = "../02-management/terraform.tfstate" }
}

module "asgardeo_rancher_oidc" {
  source                 = "../../modules/asgardeo-rancher-oidc"
  asgardeo_org_name      = var.asgardeo_org_name
  asgardeo_client_id     = var.asgardeo_client_id
  asgardeo_client_secret = var.asgardeo_client_secret
  rancher_url            = data.terraform_remote_state.management.outputs.rancher_url
}
```
