// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SystemsV1VerificationConfig struct {
	ExcludeFiles []types.String                `tfsdk:"exclude_files"`
	Keyid        types.String                  `tfsdk:"keyid"`
	PublicKeys   map[string]SystemsV1KeyConfig `tfsdk:"public_keys"`
	Scope        types.String                  `tfsdk:"scope"`
}
