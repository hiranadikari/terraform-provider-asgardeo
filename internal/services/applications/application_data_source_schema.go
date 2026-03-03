package applications

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// applicationDataSourceSchema returns the schema for the asgardeo_application data source.
func applicationDataSourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Retrieves information about an existing Asgardeo application by name or ID.\n\n" +
			"Use this data source to look up applications created outside of Terraform " +
			"(e.g. the Asgardeo built-in console applications) and reference their `client_id`.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Asgardeo application ID. Used as the lookup key when provided.",
				Optional:            true,
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Application display name. Used as the lookup key when `id` is not provided.",
				Optional:            true,
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the application.",
				Computed:            true,
			},
			"access_url": schema.StringAttribute{
				MarkdownDescription: "URL users are directed to when launching the application.",
				Computed:            true,
			},
			"logout_return_url": schema.StringAttribute{
				MarkdownDescription: "Post-global-logout redirect URL.",
				Computed:            true,
			},
			"application_enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether the application is currently enabled.",
				Computed:            true,
			},
			"client_id": schema.StringAttribute{
				MarkdownDescription: "OAuth2 / OIDC client ID (if an OIDC protocol is configured).",
				Computed:            true,
			},
			// Inbound protocol references returned by the list API.
			"inbound_protocol_types": schema.ListAttribute{
				MarkdownDescription: "List of inbound protocol types configured for the application (e.g. `oidc`, `saml`).",
				Computed:            true,
				ElementType:         types.StringType,
			},
		},
	}
}
