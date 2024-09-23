// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	tfTypes "github.com/StyraInc/terraform-provider-styra/internal/provider/types"
	"github.com/StyraInc/terraform-provider-styra/internal/sdk"
	"github.com/StyraInc/terraform-provider-styra/internal/sdk/models/operations"
	"github.com/StyraInc/terraform-provider-styra/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"regexp"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &LibraryResource{}
var _ resource.ResourceWithImportState = &LibraryResource{}

func NewLibraryResource() resource.Resource {
	return &LibraryResource{}
}

// LibraryResource defines the resource implementation.
type LibraryResource struct {
	client *sdk.StyraDas
}

// LibraryResourceModel describes the resource data model.
type LibraryResourceModel struct {
	Description   types.String                             `tfsdk:"description"`
	ID            types.String                             `tfsdk:"id"`
	ReadOnly      types.Bool                               `tfsdk:"read_only"`
	Result        tfTypes.LibrariesV1LibraryEntityExpanded `tfsdk:"result"`
	SourceControl *tfTypes.LibrariesV1SourceControlConfig  `tfsdk:"source_control"`
}

func (r *LibraryResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_library"
}

func (r *LibraryResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Library Resource",
		Attributes: map[string]schema.Attribute{
			"description": schema.StringAttribute{
				Required: true,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `id`,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`.*`), "must match pattern "+regexp.MustCompile(`.*`).String()),
				},
			},
			"read_only": schema.BoolAttribute{
				Required: true,
			},
			"result": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"datasources": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"category": schema.StringAttribute{
									Computed:    true,
									Description: `datasource category`,
								},
								"id": schema.StringAttribute{
									Computed:    true,
									Description: `datasource ID`,
								},
								"optional": schema.BoolAttribute{
									Computed:    true,
									Description: `optional datasources can be deleted without being recreated automatically`,
								},
								"status": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"code": schema.StringAttribute{
											Computed: true,
										},
										"message": schema.StringAttribute{
											Computed: true,
										},
										"timestamp": schema.StringAttribute{
											Computed: true,
											Validators: []validator.String{
												validators.IsRFC3339(),
											},
										},
									},
								},
							},
						},
					},
					"description": schema.StringAttribute{
						Computed: true,
					},
					"id": schema.StringAttribute{
						Computed: true,
					},
					"metadata": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"created_at": schema.StringAttribute{
								Computed: true,
								Validators: []validator.String{
									validators.IsRFC3339(),
								},
							},
							"created_by": schema.StringAttribute{
								Computed: true,
							},
							"created_through": schema.StringAttribute{
								Computed: true,
							},
							"last_modified_at": schema.StringAttribute{
								Computed: true,
								Validators: []validator.String{
									validators.IsRFC3339(),
								},
							},
							"last_modified_by": schema.StringAttribute{
								Computed: true,
							},
							"last_modified_through": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"policies": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"created": schema.StringAttribute{
									Computed:    true,
									Description: `policy on when to (re)generate the policy`,
								},
								"enforcement": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"enforced": schema.BoolAttribute{
											Computed:    true,
											Description: `true if the policy is enforced`,
										},
										"type": schema.StringAttribute{
											Computed:    true,
											Description: `enforcement type e.g. opa, test, mask`,
										},
									},
								},
								"id": schema.StringAttribute{
									Computed:    true,
									Description: `policy ID (path)`,
								},
								"modules": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Computed:    true,
												Description: `module name`,
											},
											"placeholder": schema.BoolAttribute{
												Computed:    true,
												Default:     booldefault.StaticBool(false),
												Description: `module is a placeholder. Default: false`,
											},
											"read_only": schema.BoolAttribute{
												Computed:    true,
												Description: `true if module is read-only`,
											},
											"rules": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"allow": schema.Int64Attribute{
														Computed:    true,
														Description: `number of allow rules`,
													},
													"deny": schema.Int64Attribute{
														Computed:    true,
														Description: `number of deny rules`,
													},
													"enforce": schema.Int64Attribute{
														Computed:    true,
														Description: `number of enforce rules`,
													},
													"ignore": schema.Int64Attribute{
														Computed:    true,
														Description: `number of ignore rules`,
													},
													"monitor": schema.Int64Attribute{
														Computed:    true,
														Description: `number of monitor rules`,
													},
													"notify": schema.Int64Attribute{
														Computed:    true,
														Description: `number of notify rules`,
													},
													"other": schema.Int64Attribute{
														Computed:    true,
														Description: `number of unclassified rules`,
													},
													"test": schema.Int64Attribute{
														Computed:    true,
														Description: `number of test rules`,
													},
													"total": schema.Int64Attribute{
														Computed:    true,
														Description: `total number of rules`,
													},
												},
											},
										},
									},
									Description: `rego modules policy consists of`,
								},
								"rules": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"allow": schema.Int64Attribute{
											Computed:    true,
											Description: `number of allow rules`,
										},
										"deny": schema.Int64Attribute{
											Computed:    true,
											Description: `number of deny rules`,
										},
										"enforce": schema.Int64Attribute{
											Computed:    true,
											Description: `number of enforce rules`,
										},
										"ignore": schema.Int64Attribute{
											Computed:    true,
											Description: `number of ignore rules`,
										},
										"monitor": schema.Int64Attribute{
											Computed:    true,
											Description: `number of monitor rules`,
										},
										"notify": schema.Int64Attribute{
											Computed:    true,
											Description: `number of notify rules`,
										},
										"other": schema.Int64Attribute{
											Computed:    true,
											Description: `number of unclassified rules`,
										},
										"test": schema.Int64Attribute{
											Computed:    true,
											Description: `number of test rules`,
										},
										"total": schema.Int64Attribute{
											Computed:    true,
											Description: `total number of rules`,
										},
									},
								},
								"type": schema.StringAttribute{
									Computed:    true,
									Description: `policy type e.g. validating/rules`,
								},
							},
						},
					},
					"read_only": schema.BoolAttribute{
						Computed: true,
					},
					"source_control": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"library_origin": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"commit": schema.StringAttribute{
										Computed:    true,
										Description: `Commit SHA. Only one of reference or commit can be set at any time`,
									},
									"credentials": schema.StringAttribute{
										Computed:    true,
										Description: `Credentials are looked under the key <name>/<creds>`,
									},
									"path": schema.StringAttribute{
										Computed:    true,
										Description: `Path to limit the import to`,
									},
									"reference": schema.StringAttribute{
										Computed:    true,
										Description: `Remote reference. Only one of reference or commit can be set at any time`,
									},
									"ssh_credentials": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"passphrase": schema.StringAttribute{
												Computed:    true,
												Description: `Passphrase is looked under the key passphrase/<pass>`,
											},
											"private_key": schema.StringAttribute{
												Computed:    true,
												Description: `PrivateKey is looked under the key private-key/<key>`,
											},
										},
									},
									"url": schema.StringAttribute{
										Computed:    true,
										Description: `Repository URL`,
									},
								},
							},
							"use_workspace_settings": schema.BoolAttribute{
								Computed: true,
							},
						},
					},
					"used_by": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"bundles": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"bundle_id": schema.StringAttribute{
												Computed: true,
											},
											"version": schema.Int64Attribute{
												Computed: true,
											},
										},
									},
								},
								"system_id": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
			"source_control": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"library_origin": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"commit": schema.StringAttribute{
								Required:    true,
								Description: `Commit SHA. Only one of reference or commit can be set at any time`,
							},
							"credentials": schema.StringAttribute{
								Required:    true,
								Description: `Credentials are looked under the key <name>/<creds>`,
							},
							"path": schema.StringAttribute{
								Required:    true,
								Description: `Path to limit the import to`,
							},
							"reference": schema.StringAttribute{
								Required:    true,
								Description: `Remote reference. Only one of reference or commit can be set at any time`,
							},
							"ssh_credentials": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"passphrase": schema.StringAttribute{
										Required:    true,
										Description: `Passphrase is looked under the key passphrase/<pass>`,
									},
									"private_key": schema.StringAttribute{
										Required:    true,
										Description: `PrivateKey is looked under the key private-key/<key>`,
									},
								},
							},
							"url": schema.StringAttribute{
								Required:    true,
								Description: `Repository URL`,
							},
						},
					},
					"use_workspace_settings": schema.BoolAttribute{
						Required: true,
					},
				},
			},
		},
	}
}

func (r *LibraryResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.StyraDas)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.StyraDas, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *LibraryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *LibraryResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	var id string
	id = data.ID.ValueString()

	librariesV1CreateLibraryRequest := *data.ToSharedLibrariesV1CreateLibraryRequest()
	request := operations.LibrariesUpdateRequest{
		ID:                              id,
		LibrariesV1CreateLibraryRequest: librariesV1CreateLibraryRequest,
	}
	res, err := r.client.Libraries.LibrariesUpdate(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.LibrariesV1LibraryResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedLibrariesV1LibraryResponse(res.LibrariesV1LibraryResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var id1 string
	id1 = data.ID.ValueString()

	// create.library.dependant_bundlescreate.library.dependant_bundles impedance mismatch: string != classtrace=["Library#create,update","Library#create,update.req"]
	var dependantBundles *string
	request1 := operations.LibrariesGetRequest{
		ID:               id1,
		DependantBundles: dependantBundles,
	}
	res1, err := r.client.Libraries.LibrariesGet(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if !(res1.LibrariesV1LibraryResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedLibrariesV1LibraryResponse(res1.LibrariesV1LibraryResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LibraryResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *LibraryResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	var id string
	id = data.ID.ValueString()

	// read.library.dependant_bundlesread.library.dependant_bundles impedance mismatch: string != classtrace=["Library#create,update.req","Library#create,update"]
	var dependantBundles *string
	request := operations.LibrariesGetRequest{
		ID:               id,
		DependantBundles: dependantBundles,
	}
	res, err := r.client.Libraries.LibrariesGet(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.LibrariesV1LibraryResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedLibrariesV1LibraryResponse(res.LibrariesV1LibraryResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LibraryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *LibraryResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	var id string
	id = data.ID.ValueString()

	librariesV1CreateLibraryRequest := *data.ToSharedLibrariesV1CreateLibraryRequest()
	request := operations.LibrariesUpdateRequest{
		ID:                              id,
		LibrariesV1CreateLibraryRequest: librariesV1CreateLibraryRequest,
	}
	res, err := r.client.Libraries.LibrariesUpdate(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.LibrariesV1LibraryResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedLibrariesV1LibraryResponse(res.LibrariesV1LibraryResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var id1 string
	id1 = data.ID.ValueString()

	// update.library.dependant_bundlesupdate.library.dependant_bundles impedance mismatch: string != classtrace=["Library#create,update.req","Library#create,update"]
	var dependantBundles *string
	request1 := operations.LibrariesGetRequest{
		ID:               id1,
		DependantBundles: dependantBundles,
	}
	res1, err := r.client.Libraries.LibrariesGet(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if !(res1.LibrariesV1LibraryResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedLibrariesV1LibraryResponse(res1.LibrariesV1LibraryResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LibraryResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *LibraryResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	var id string
	id = data.ID.ValueString()

	request := operations.LibrariesDeleteRequest{
		ID: id,
	}
	res, err := r.client.Libraries.LibrariesDelete(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *LibraryResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
