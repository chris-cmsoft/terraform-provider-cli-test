package hashicups

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"cli_onepassword_version": dataSourceOnePasswordVersion(),
			"cli_onepassword_users": dataSourceOnePasswordUsers(),
		},
	}
}
