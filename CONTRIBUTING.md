# Contributing to terraform-provider-asgardeo

Thank you for your interest in contributing! This guide explains how to set up
your environment, add new resources, and get your PR merged.

## Table of Contents

- [Development Setup](#development-setup)
- [Project Structure](#project-structure)
- [Adding a New Resource](#adding-a-new-resource)
- [Testing Requirements](#testing-requirements)
- [PR Process](#pr-process)
- [Code Style](#code-style)

---

## Development Setup

```bash
# Prerequisites: Go ≥ 1.21, Terraform ≥ 1.5
git clone https://github.com/asgardeo/terraform-provider-asgardeo
cd terraform-provider-asgardeo

make tools   # install golangci-lint, tfplugindocs, misspell
make build   # compile provider binary
make test    # run unit tests
```

For acceptance tests you need real Asgardeo credentials:

```bash
export ASGARDEO_ORG_NAME=myorg
export ASGARDEO_CLIENT_ID=<m2m-client-id>
export ASGARDEO_CLIENT_SECRET=<m2m-client-secret>
make testacc
```

---

## Project Structure

```
asgardeo/           Pure-Go API client (no Terraform dependency)
  client.go         HTTP client, OAuth2 token management
  models.go         API request/response structs
  applications.go   Application CRUD + protocol methods

internal/
  clients/          Wires asgardeo.Client into the provider
    builder.go
  provider/
    provider.go     Plugin-framework provider definition
  services/
    applications/   asgardeo_application resource + data source
      registration.go            Register Resources() / DataSources()
      application_resource.go    CRUD logic
      application_resource_schema.go  Schema definition
      application_data_source.go
      application_data_source_schema.go
```

---

## Adding a New Resource

Follow this checklist when adding, for example, an `asgardeo_identity_provider` resource:

1. **API layer** — Add Go structs to `asgardeo/models.go` and methods to a new
   `asgardeo/identity_providers.go` file.

2. **New service package** — Create `internal/services/identity_providers/`:
   ```
   registration.go
   identity_provider_resource.go
   identity_provider_resource_schema.go
   identity_provider_data_source.go          (if applicable)
   identity_provider_data_source_schema.go   (if applicable)
   ```

3. **Register** — Add the new service's `Resources()` and `DataSources()` to
   `internal/provider/provider.go`.

4. **Examples** — Add `examples/resources/asgardeo_identity_provider/resource.tf`
   and `examples/data-sources/asgardeo_identity_provider/data-source.tf`.

5. **Templates** — Add `templates/resources/identity_provider.md.tmpl`.

6. **Tests** — Write at minimum one acceptance test in
   `internal/services/identity_providers/identity_provider_resource_test.go`.

7. **Docs** — Run `make docs` and commit the generated `docs/` files.

---

## Testing Requirements

| Test type | Requirement |
|---|---|
| Unit tests | Must pass (`make test`) for all PRs |
| Acceptance tests | Must pass for any resource adding/changing CRUD logic |
| Docs lint | Must pass (`make docs-lint`) |

Acceptance tests must:
- Be guarded by `if os.Getenv("TF_ACC") == ""` or use `resource.Test(t, ...)`.
- Clean up all created resources (use `CheckDestroy`).
- Cover Create + Read + Update + Delete + Import.

---

## PR Process

1. Fork the repository and create a feature branch from `main`.
2. Make your changes, following the checklist above for new resources.
3. Run `make fmt lint vet test` — all must pass.
4. Push and open a PR against `main`.
5. Fill in the PR template: describe the change, link the relevant Asgardeo API
   docs, and list the acceptance test results.
6. A maintainer will review within 5 business days.

### Commit and PR title format

We follow the **WSO2 imperative sentence** convention — a plain present-tense
verb phrase with no type prefix. The first word is capitalised; no trailing period.

```
Add asgardeo_identity_provider resource
Fix nil pointer dereference in OIDC config flattener
Update refresh token expiry default to 86400 seconds
Remove deprecated accessTokenAttributes field
Improve error messages for 401 responses
```

For changes scoped to a subsystem, prefix with the area in square brackets:

```
[OIDC] Add support for hybrid flow response types
[SAML] Handle metadata URL configuration
[Docs] Update Rancher OIDC example with correct callback URL
```

**Do not use** `feat:`, `fix:`, `chore:` or other conventional-commit prefixes.

Commit messages on the branch should follow the same convention. Squash or
rebase to a clean history before requesting review — avoid commits like
`wip`, `fix review comments`, or `update`.

---

## Code Style

- Run `make fmt` before committing.
- No global variables; pass `*asgardeo.Client` via dependency injection.
- Schema files (`*_schema.go`) must not contain CRUD logic.
- Errors must use `%w` for wrapping; no `log.Fatal` outside `main.go`.
- Keep the `asgardeo/` package free of Terraform SDK imports.
