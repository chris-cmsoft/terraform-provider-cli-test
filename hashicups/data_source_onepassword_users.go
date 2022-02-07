package hashicups

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceOnePasswordUsers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOnePasswordUsersRead,
		Schema: map[string]*schema.Schema{
			"users": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"email": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceOnePasswordUsersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	// Call API or read file or whatever else

	users := []map[string]interface{}{
		{
			"id": "B85E8258-0A99-4ADF-87AD-09855EA4FDDC",
			"name": "John Doe",
			"email": "john.doe@example.com",
		},
		{
			"id": "B85E8258-0A99-4ADF-87AD-09855EA4FDDC",
			"name": "Jane Doe",
			"email": "jane.doe@example.com",
		},
	}

	if err := d.Set("users", users); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
