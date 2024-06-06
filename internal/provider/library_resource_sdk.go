// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	tfTypes "github.com/StyraInc/terraform-provider-styra/internal/provider/types"
	"github.com/StyraInc/terraform-provider-styra/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *LibraryResourceModel) ToSharedLibrariesV1CreateLibraryRequest() *shared.LibrariesV1CreateLibraryRequest {
	description := r.Description.ValueString()
	readOnly := r.ReadOnly.ValueBool()
	var sourceControl *shared.LibrariesV1SourceControlConfig
	if r.SourceControl != nil {
		var libraryOrigin *shared.GitV1GitRepoConfig
		if r.SourceControl.LibraryOrigin != nil {
			commit := r.SourceControl.LibraryOrigin.Commit.ValueString()
			credentials := r.SourceControl.LibraryOrigin.Credentials.ValueString()
			path := r.SourceControl.LibraryOrigin.Path.ValueString()
			reference := r.SourceControl.LibraryOrigin.Reference.ValueString()
			var sshCredentials *shared.GitV1SSHCredentials
			if r.SourceControl.LibraryOrigin.SSHCredentials != nil {
				passphrase := r.SourceControl.LibraryOrigin.SSHCredentials.Passphrase.ValueString()
				privateKey := r.SourceControl.LibraryOrigin.SSHCredentials.PrivateKey.ValueString()
				sshCredentials = &shared.GitV1SSHCredentials{
					Passphrase: passphrase,
					PrivateKey: privateKey,
				}
			}
			url := r.SourceControl.LibraryOrigin.URL.ValueString()
			libraryOrigin = &shared.GitV1GitRepoConfig{
				Commit:         commit,
				Credentials:    credentials,
				Path:           path,
				Reference:      reference,
				SSHCredentials: sshCredentials,
				URL:            url,
			}
		}
		useWorkspaceSettings := r.SourceControl.UseWorkspaceSettings.ValueBool()
		sourceControl = &shared.LibrariesV1SourceControlConfig{
			LibraryOrigin:        libraryOrigin,
			UseWorkspaceSettings: useWorkspaceSettings,
		}
	}
	out := shared.LibrariesV1CreateLibraryRequest{
		Description:   description,
		ReadOnly:      readOnly,
		SourceControl: sourceControl,
	}
	return &out
}

func (r *LibraryResourceModel) RefreshFromSharedLibrariesV1LibraryResponse(resp *shared.LibrariesV1LibraryResponse) {
	if resp != nil {
		r.Result.Datasources = []tfTypes.SystemsV1DatasourceConfig{}
		if len(r.Result.Datasources) > len(resp.Result.Datasources) {
			r.Result.Datasources = r.Result.Datasources[:len(resp.Result.Datasources)]
		}
		for datasourcesCount, datasourcesItem := range resp.Result.Datasources {
			var datasources1 tfTypes.SystemsV1DatasourceConfig
			datasources1.Category = types.StringValue(datasourcesItem.Category)
			datasources1.ID = types.StringValue(datasourcesItem.ID)
			datasources1.Optional = types.BoolPointerValue(datasourcesItem.Optional)
			if datasourcesItem.Status == nil {
				datasources1.Status = nil
			} else {
				datasources1.Status = &tfTypes.MetaV1Status{}
				datasources1.Status.Code = types.StringValue(datasourcesItem.Status.Code)
				datasources1.Status.Message = types.StringValue(datasourcesItem.Status.Message)
				datasources1.Status.Timestamp = types.StringValue(datasourcesItem.Status.Timestamp.Format(time.RFC3339Nano))
			}
			if datasourcesCount+1 > len(r.Result.Datasources) {
				r.Result.Datasources = append(r.Result.Datasources, datasources1)
			} else {
				r.Result.Datasources[datasourcesCount].Category = datasources1.Category
				r.Result.Datasources[datasourcesCount].ID = datasources1.ID
				r.Result.Datasources[datasourcesCount].Optional = datasources1.Optional
				r.Result.Datasources[datasourcesCount].Status = datasources1.Status
			}
		}
		r.Result.Description = types.StringValue(resp.Result.Description)
		r.Result.ID = types.StringValue(resp.Result.ID)
		if resp.Result.Metadata == nil {
			r.Result.Metadata = nil
		} else {
			r.Result.Metadata = &tfTypes.MetaV2ObjectMeta{}
			if resp.Result.Metadata.CreatedAt != nil {
				r.Result.Metadata.CreatedAt = types.StringValue(resp.Result.Metadata.CreatedAt.Format(time.RFC3339Nano))
			} else {
				r.Result.Metadata.CreatedAt = types.StringNull()
			}
			r.Result.Metadata.CreatedBy = types.StringPointerValue(resp.Result.Metadata.CreatedBy)
			r.Result.Metadata.CreatedThrough = types.StringPointerValue(resp.Result.Metadata.CreatedThrough)
			if resp.Result.Metadata.LastModifiedAt != nil {
				r.Result.Metadata.LastModifiedAt = types.StringValue(resp.Result.Metadata.LastModifiedAt.Format(time.RFC3339Nano))
			} else {
				r.Result.Metadata.LastModifiedAt = types.StringNull()
			}
			r.Result.Metadata.LastModifiedBy = types.StringPointerValue(resp.Result.Metadata.LastModifiedBy)
			r.Result.Metadata.LastModifiedThrough = types.StringPointerValue(resp.Result.Metadata.LastModifiedThrough)
		}
		r.Result.Policies = []tfTypes.SystemsV1PolicyConfig{}
		if len(r.Result.Policies) > len(resp.Result.Policies) {
			r.Result.Policies = r.Result.Policies[:len(resp.Result.Policies)]
		}
		for policiesCount, policiesItem := range resp.Result.Policies {
			var policies1 tfTypes.SystemsV1PolicyConfig
			policies1.Created = types.StringPointerValue(policiesItem.Created)
			policies1.Enforcement.Enforced = types.BoolValue(policiesItem.Enforcement.Enforced)
			policies1.Enforcement.Type = types.StringValue(policiesItem.Enforcement.Type)
			policies1.ID = types.StringValue(policiesItem.ID)
			policies1.Modules = []tfTypes.SystemsV1Module{}
			for modulesCount, modulesItem := range policiesItem.Modules {
				var modules1 tfTypes.SystemsV1Module
				modules1.Name = types.StringValue(modulesItem.Name)
				modules1.Placeholder = types.BoolPointerValue(modulesItem.Placeholder)
				modules1.ReadOnly = types.BoolValue(modulesItem.ReadOnly)
				if modulesItem.Rules == nil {
					modules1.Rules = nil
				} else {
					modules1.Rules = &tfTypes.PoliciesV1RuleCounts{}
					modules1.Rules.Allow = types.Int64Value(int64(modulesItem.Rules.Allow))
					modules1.Rules.Deny = types.Int64Value(int64(modulesItem.Rules.Deny))
					modules1.Rules.Enforce = types.Int64Value(int64(modulesItem.Rules.Enforce))
					modules1.Rules.Ignore = types.Int64Value(int64(modulesItem.Rules.Ignore))
					modules1.Rules.Monitor = types.Int64Value(int64(modulesItem.Rules.Monitor))
					modules1.Rules.Notify = types.Int64Value(int64(modulesItem.Rules.Notify))
					modules1.Rules.Other = types.Int64Value(int64(modulesItem.Rules.Other))
					modules1.Rules.Test = types.Int64Value(int64(modulesItem.Rules.Test))
					modules1.Rules.Total = types.Int64Value(int64(modulesItem.Rules.Total))
				}
				if modulesCount+1 > len(policies1.Modules) {
					policies1.Modules = append(policies1.Modules, modules1)
				} else {
					policies1.Modules[modulesCount].Name = modules1.Name
					policies1.Modules[modulesCount].Placeholder = modules1.Placeholder
					policies1.Modules[modulesCount].ReadOnly = modules1.ReadOnly
					policies1.Modules[modulesCount].Rules = modules1.Rules
				}
			}
			if policiesItem.Rules == nil {
				policies1.Rules = nil
			} else {
				policies1.Rules = &tfTypes.PoliciesV1RuleCounts{}
				policies1.Rules.Allow = types.Int64Value(int64(policiesItem.Rules.Allow))
				policies1.Rules.Deny = types.Int64Value(int64(policiesItem.Rules.Deny))
				policies1.Rules.Enforce = types.Int64Value(int64(policiesItem.Rules.Enforce))
				policies1.Rules.Ignore = types.Int64Value(int64(policiesItem.Rules.Ignore))
				policies1.Rules.Monitor = types.Int64Value(int64(policiesItem.Rules.Monitor))
				policies1.Rules.Notify = types.Int64Value(int64(policiesItem.Rules.Notify))
				policies1.Rules.Other = types.Int64Value(int64(policiesItem.Rules.Other))
				policies1.Rules.Test = types.Int64Value(int64(policiesItem.Rules.Test))
				policies1.Rules.Total = types.Int64Value(int64(policiesItem.Rules.Total))
			}
			policies1.Type = types.StringValue(policiesItem.Type)
			if policiesCount+1 > len(r.Result.Policies) {
				r.Result.Policies = append(r.Result.Policies, policies1)
			} else {
				r.Result.Policies[policiesCount].Created = policies1.Created
				r.Result.Policies[policiesCount].Enforcement = policies1.Enforcement
				r.Result.Policies[policiesCount].ID = policies1.ID
				r.Result.Policies[policiesCount].Modules = policies1.Modules
				r.Result.Policies[policiesCount].Rules = policies1.Rules
				r.Result.Policies[policiesCount].Type = policies1.Type
			}
		}
		r.Result.ReadOnly = types.BoolValue(resp.Result.ReadOnly)
		if resp.Result.SourceControl == nil {
			r.Result.SourceControl = nil
		} else {
			r.Result.SourceControl = &tfTypes.LibrariesV1SourceControlConfig{}
			if resp.Result.SourceControl.LibraryOrigin == nil {
				r.Result.SourceControl.LibraryOrigin = nil
			} else {
				r.Result.SourceControl.LibraryOrigin = &tfTypes.GitV1GitRepoConfig{}
				r.Result.SourceControl.LibraryOrigin.Commit = types.StringValue(resp.Result.SourceControl.LibraryOrigin.Commit)
				r.Result.SourceControl.LibraryOrigin.Credentials = types.StringValue(resp.Result.SourceControl.LibraryOrigin.Credentials)
				r.Result.SourceControl.LibraryOrigin.Path = types.StringValue(resp.Result.SourceControl.LibraryOrigin.Path)
				r.Result.SourceControl.LibraryOrigin.Reference = types.StringValue(resp.Result.SourceControl.LibraryOrigin.Reference)
				if resp.Result.SourceControl.LibraryOrigin.SSHCredentials == nil {
					r.Result.SourceControl.LibraryOrigin.SSHCredentials = nil
				} else {
					r.Result.SourceControl.LibraryOrigin.SSHCredentials = &tfTypes.GitV1SSHCredentials{}
					r.Result.SourceControl.LibraryOrigin.SSHCredentials.Passphrase = types.StringValue(resp.Result.SourceControl.LibraryOrigin.SSHCredentials.Passphrase)
					r.Result.SourceControl.LibraryOrigin.SSHCredentials.PrivateKey = types.StringValue(resp.Result.SourceControl.LibraryOrigin.SSHCredentials.PrivateKey)
				}
				r.Result.SourceControl.LibraryOrigin.URL = types.StringValue(resp.Result.SourceControl.LibraryOrigin.URL)
			}
			r.Result.SourceControl.UseWorkspaceSettings = types.BoolValue(resp.Result.SourceControl.UseWorkspaceSettings)
		}
		r.Result.UsedBy = []tfTypes.LibrariesV1SystemUsingLibrary{}
		if len(r.Result.UsedBy) > len(resp.Result.UsedBy) {
			r.Result.UsedBy = r.Result.UsedBy[:len(resp.Result.UsedBy)]
		}
		for usedByCount, usedByItem := range resp.Result.UsedBy {
			var usedBy1 tfTypes.LibrariesV1SystemUsingLibrary
			usedBy1.Bundles = []tfTypes.LibrariesV1SystemBundle{}
			for bundlesCount, bundlesItem := range usedByItem.Bundles {
				var bundles1 tfTypes.LibrariesV1SystemBundle
				bundles1.BundleID = types.StringValue(bundlesItem.BundleID)
				bundles1.Version = types.Int64Value(bundlesItem.Version)
				if bundlesCount+1 > len(usedBy1.Bundles) {
					usedBy1.Bundles = append(usedBy1.Bundles, bundles1)
				} else {
					usedBy1.Bundles[bundlesCount].BundleID = bundles1.BundleID
					usedBy1.Bundles[bundlesCount].Version = bundles1.Version
				}
			}
			usedBy1.SystemID = types.StringValue(usedByItem.SystemID)
			if usedByCount+1 > len(r.Result.UsedBy) {
				r.Result.UsedBy = append(r.Result.UsedBy, usedBy1)
			} else {
				r.Result.UsedBy[usedByCount].Bundles = usedBy1.Bundles
				r.Result.UsedBy[usedByCount].SystemID = usedBy1.SystemID
			}
		}
	}
}
