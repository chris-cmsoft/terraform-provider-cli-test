package hashicups

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceOnePasswordVersion() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOnePasswordVersionRead,
		Schema: map[string]*schema.Schema{
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				//Elem: &schema.Resource{
				//	Schema: map[string]*schema.Schema{
				//		"name": &schema.Schema{
				//			Type:     schema.TypeString,
				//			Computed: true,
				//		},
				//	},
				//},
			},
		},
	}
}

func dataSourceOnePasswordVersionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	output, err := RunCommandWithEnvReturnOutput(
		"op",
		map[string]string{},
		"--version")

	if err != nil {
		return diag.FromErr(err)
	}
	if len(output) == 0 {
		return diag.FromErr(errors.New("Packer did not output anything"))
	}

	version := strings.TrimSpace(string(output))

	if err := d.Set("version", version); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

//
//func (r dataSourceVersion) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
//	resourceState := dataSourceVersionType{}
//	exe, _ := os.Executable()
//	output, err := RunCommandWithEnvReturnOutput(
//		exe,
//		map[string]string{packer_interop.TPPRunPacker: "true"},
//		"version")
//	if err != nil {
//		resp.Diagnostics.AddError("Failed to run packer", err.Error())
//		return
//	}
//
//	if len(output) == 0 {
//		resp.Diagnostics.AddError("Unexpected output", "Packer did not output anything")
//		return
//	}
//
//	resourceState.Version = strings.TrimPrefix(
//		strings.TrimSpace(strings.TrimPrefix(string(output), "Packer")), "v")
//
//	diags := resp.State.Set(ctx, &resourceState)
//	resp.Diagnostics.Append(diags...)
//	if resp.Diagnostics.HasError() {
//		return
//	}
//}
