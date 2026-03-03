# agent.md — AI Collaboration Guide

This file explains the project structure and conventions for an AI assistant
(Claude, Copilot, GPT-4, etc.) extending this provider with new resources.

---

## Quick Mental Model

```
Asgardeo REST API
      │
      ▼
asgardeo/ package          ← Pure Go, no Terraform dependency
  client.go (HTTP + OAuth2)
  models.go (API structs)
  applications.go (methods on *Client)
      │
      ▼
internal/clients/builder.go   ← Converts provider config → *asgardeo.Client
      │
      ▼
internal/provider/provider.go ← Registers resources + data sources
      │
      ├── internal/services/applications/
      │     registration.go
      │     application_resource.go        (CRUD)
      │     application_resource_schema.go (Schema)
      │     application_data_source.go
      │     application_data_source_schema.go
      └── internal/services/<new-service>/  ← Add new services here
```

---

## Adding a New Resource: Step-by-Step

### 1 — Add API structs to `asgardeo/models.go`

```go
type IdentityProviderCreateRequest struct {
    Name string `json:"name"`
    // ... fields from the Asgardeo API spec
}
```

### 2 — Add API methods to `asgardeo/<resource>.go`

Methods are defined directly on `*Client`:

```go
func (c *Client) CreateIdentityProvider(ctx context.Context, req IdentityProviderCreateRequest) (*IdentityProviderResponse, error) {
    // POST /identity-providers
}
```

Follow the pattern in `asgardeo/applications.go`.

### 3 — Create `internal/services/<name>/`

Required files:
- `registration.go` — exports `Resources()` and `DataSources()`
- `<name>_resource_schema.go` — schema only, no logic
- `<name>_resource.go` — implements `resource.Resource`; calls `asgardeo.Client` methods
- `<name>_data_source_schema.go` (if applicable)
- `<name>_data_source.go` (if applicable)

### 4 — Register in `internal/provider/provider.go`

```go
import "github.com/asgardeo/terraform-provider-asgardeo/internal/services/identityproviders"

func (p *AsgardeoProvider) Resources(_ context.Context) []func() resource.Resource {
    return append(applications.Resources(), identityproviders.Resources()...)
}
```

### 5 — Add examples and templates

```
examples/resources/asgardeo_identity_provider/resource.tf
templates/resources/identity_provider.md.tmpl
```

### 6 — Run `make docs` to regenerate docs.

---

## Key Conventions

| Convention | Rule |
|---|---|
| Client access | Inject via `Configure()` — never create a new client in resource methods |
| Nil = not found | `Get*` methods return `nil, nil` for 404 — callers must handle with `RemoveResource` |
| Schema separation | Schema files (`*_schema.go`) contain only `schema.*` definitions |
| Protocol separation | OIDC / SAML config is fetched separately after GET /applications/{id} |
| Computed secrets | `client_id` and `client_secret` use `UseStateForUnknown()` plan modifier |
| Error wrapping | Always use `fmt.Errorf("context: %w", err)` |
| No SDK in asgardeo/ | The `asgardeo/` package must not import any `terraform-plugin-*` packages |

---

## Asgardeo API Reference

| Resource | Endpoint |
|---|---|
| Applications | `POST/GET/PATCH/DELETE /api/server/v1/applications/{id}` |
| OIDC config | `GET/PUT /api/server/v1/applications/{id}/inbound-protocols/oidc` |
| SAML config | `GET/PUT /api/server/v1/applications/{id}/inbound-protocols/saml` |
| Identity Providers | `GET/POST /api/server/v1/identity-providers` |
| Groups | `GET/POST /api/server/v1/groups` |
| Users | `GET/POST /scim2/Users` |

Base URL: `https://api.asgardeo.io/t/{org_name}/api/server/v1`
Auth: OAuth2 `client_credentials` → `https://api.asgardeo.io/t/{org_name}/oauth2/token`

---

## Terraform Plugin Framework Cheat-Sheet

```go
// Read plan into model
var plan MyModel
resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

// Write state
resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)

// Resource not found in Read → remove from state
resp.State.RemoveResource(ctx)

// Computed + stable across updates
PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}

// List block with max 1 item (optional nested object pattern)
schema.ListNestedBlock{
    Validators: []validator.List{listvalidator.SizeAtMost(1)},
    ...
}
```

---

## Running Tests

```bash
make test                  # unit tests (no credentials needed)
TF_ACC=1 make testacc      # acceptance tests (requires ASGARDEO_* env vars)
```

Acceptance tests live alongside the resource in `*_resource_test.go` and must
use `resource.Test(t, resource.TestCase{...})` from `terraform-plugin-testing`.
