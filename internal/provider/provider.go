// Package provider implements the Asgardeo Terraform provider using
// terraform-plugin-framework.
package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/asgardeo/terraform-provider-asgardeo/internal/clients"
	"github.com/asgardeo/terraform-provider-asgardeo/internal/services/applications"
)

// Ensure AsgardeoProvider satisfies the provider.Provider interface.
var _ provider.Provider = &AsgardeoProvider{}

// AsgardeoProvider is the root provider implementation.
type AsgardeoProvider struct {
	// version is set by the release build process (via -ldflags).
	version string
}

// New returns a factory function that creates a new provider instance.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &AsgardeoProvider{version: version}
	}
}

// ─── provider.Provider interface ─────────────────────────────────────────────

func (p *AsgardeoProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "asgardeo"
	resp.Version = p.version
}

// Schema defines the provider-level configuration attributes.
func (p *AsgardeoProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The **Asgardeo** provider manages applications and identity resources in " +
			"[Asgardeo](https://wso2.com/asgardeo/), WSO2's cloud-native identity platform.\n\n" +
			"Authentication uses OAuth 2.0 client credentials. Create an M2M application in the " +
			"Asgardeo console and authorize it with the required management API scopes.",

		Attributes: map[string]schema.Attribute{
			"org_name": schema.StringAttribute{
				MarkdownDescription: "Asgardeo organisation name (the subdomain part of your Asgardeo URL). " +
					"Can also be set via the `ASGARDEO_ORG_NAME` environment variable.",
				Optional: true,
			},
			"client_id": schema.StringAttribute{
				MarkdownDescription: "Client ID of the M2M application used to authenticate against the " +
					"Asgardeo Management API. Can also be set via `ASGARDEO_CLIENT_ID`.",
				Optional:  true,
				Sensitive: false,
			},
			"client_secret": schema.StringAttribute{
				MarkdownDescription: "Client secret of the M2M application. Can also be set via `ASGARDEO_CLIENT_SECRET`.",
				Optional:            true,
				Sensitive:           true,
			},
		},
	}
}

// providerModel maps the HCL provider block to Go types.
type providerModel struct {
	OrgName      types.String `tfsdk:"org_name"`
	ClientID     types.String `tfsdk:"client_id"`
	ClientSecret types.String `tfsdk:"client_secret"`
}

// Configure builds the API client and stores it in resp.ResourceData /
// resp.DataSourceData so resources and data sources can retrieve it.
func (p *AsgardeoProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var cfg providerModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &cfg)...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgName := coalesce(cfg.OrgName.ValueString(), os.Getenv("ASGARDEO_ORG_NAME"))
	clientID := coalesce(cfg.ClientID.ValueString(), os.Getenv("ASGARDEO_CLIENT_ID"))
	clientSecret := coalesce(cfg.ClientSecret.ValueString(), os.Getenv("ASGARDEO_CLIENT_SECRET"))

	if orgName == "" {
		resp.Diagnostics.AddError(
			"Missing org_name",
			"Set org_name in the provider block or the ASGARDEO_ORG_NAME environment variable.",
		)
	}
	if clientID == "" {
		resp.Diagnostics.AddError(
			"Missing client_id",
			"Set client_id in the provider block or the ASGARDEO_CLIENT_ID environment variable.",
		)
	}
	if clientSecret == "" {
		resp.Diagnostics.AddError(
			"Missing client_secret",
			"Set client_secret in the provider block or the ASGARDEO_CLIENT_SECRET environment variable.",
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	client := clients.Build(orgName, clientID, clientSecret)
	resp.ResourceData = client
	resp.DataSourceData = client
}

// Resources returns the list of managed resources provided by this provider.
func (p *AsgardeoProvider) Resources(_ context.Context) []func() resource.Resource {
	return applications.Resources()
}

// DataSources returns the list of data sources provided by this provider.
func (p *AsgardeoProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return applications.DataSources()
}

// ─── helpers ─────────────────────────────────────────────────────────────────

// coalesce returns the first non-empty string.
func coalesce(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}
