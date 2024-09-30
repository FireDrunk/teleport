// Code generated by _gen/main.go DO NOT EDIT
/*
Copyright 2015-2024 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import (
	"context"
	"fmt"
	apitypes "github.com/gravitational/teleport/api/types"

	userprovisioningv2 "github.com/gravitational/teleport/api/gen/proto/go/teleport/userprovisioning/v2"
	
	"github.com/gravitational/teleport/integrations/lib/backoff"
	"github.com/gravitational/trace"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/jonboulle/clockwork"

	schemav1 "github.com/gravitational/teleport/integrations/terraform/tfschema/userprovisioning/v2"
)

// resourceTeleportStaticHostUserType is the resource metadata type
type resourceTeleportStaticHostUserType struct{}

// resourceTeleportStaticHostUser is the resource
type resourceTeleportStaticHostUser struct {
	p Provider
}

// GetSchema returns the resource schema
func (r resourceTeleportStaticHostUserType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return schemav1.GenSchemaStaticHostUser(ctx)
}

// NewResource creates the empty resource
func (r resourceTeleportStaticHostUserType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceTeleportStaticHostUser{
		p: *(p.(*Provider)),
	}, nil
}

// Create creates the StaticHostUser
func (r resourceTeleportStaticHostUser) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var err error
	if !r.p.IsConfigured(resp.Diagnostics) {
		return
	}

	var plan types.Object
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	staticHostUser := &userprovisioningv2.StaticHostUser{}
	diags = schemav1.CopyStaticHostUserFromTerraform(ctx, plan, staticHostUser)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	
	staticHostUserResource := staticHostUser

	staticHostUserResource.Kind = apitypes.KindStaticHostUser

	id := staticHostUserResource.Metadata.Name

	_, err = r.p.Client.StaticHostUserClient().GetStaticHostUser(ctx, id)
	if !trace.IsNotFound(err) {
		if err == nil {
			existErr := fmt.Sprintf("StaticHostUser exists in Teleport. Either remove it (tctl rm static_host_user/%v)"+
				" or import it to the existing state (terraform import teleport_static_host_user.%v %v)", id, id, id)

			resp.Diagnostics.Append(diagFromErr("StaticHostUser exists in Teleport", trace.Errorf(existErr)))
			return
		}

		resp.Diagnostics.Append(diagFromWrappedErr("Error reading StaticHostUser", trace.Wrap(err), "static_host_user"))
		return
	}

	_, err = r.p.Client.StaticHostUserClient().CreateStaticHostUser(ctx, staticHostUserResource)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error creating StaticHostUser", trace.Wrap(err), "static_host_user"))
		return
	}
		var staticHostUserI *userprovisioningv2.StaticHostUser
	tries := 0
	backoff := backoff.NewDecorr(r.p.RetryConfig.Base, r.p.RetryConfig.Cap, clockwork.NewRealClock())
	for {
		tries = tries + 1
		staticHostUserI, err = r.p.Client.StaticHostUserClient().GetStaticHostUser(ctx, id)
		if trace.IsNotFound(err) {
			if bErr := backoff.Do(ctx); bErr != nil {
				resp.Diagnostics.Append(diagFromWrappedErr("Error reading StaticHostUser", trace.Wrap(bErr), "static_host_user"))
				return
			}
			if tries >= r.p.RetryConfig.MaxTries {
				diagMessage := fmt.Sprintf("Error reading StaticHostUser (tried %d times) - state outdated, please import resource", tries)
				resp.Diagnostics.AddError(diagMessage, "static_host_user")
			}
			continue
		}
		break
	}

	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading StaticHostUser", trace.Wrap(err), "static_host_user"))
		return
	}

	staticHostUserResource = staticHostUserI
	
	staticHostUser = staticHostUserResource

	diags = schemav1.CopyStaticHostUserToTerraform(ctx, staticHostUser, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Attrs["id"] = types.String{Value: staticHostUser.Metadata.Name}

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read reads teleport StaticHostUser
func (r resourceTeleportStaticHostUser) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state types.Object
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var id types.String
	diags = req.State.GetAttribute(ctx, path.Root("metadata").AtName("name"), &id)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	staticHostUserI, err := r.p.Client.StaticHostUserClient().GetStaticHostUser(ctx, id.Value)
	if trace.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}

	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading StaticHostUser", trace.Wrap(err), "static_host_user"))
		return
	}
	staticHostUser := staticHostUserI
	diags = schemav1.CopyStaticHostUserToTerraform(ctx, staticHostUser, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates teleport StaticHostUser
func (r resourceTeleportStaticHostUser) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	if !r.p.IsConfigured(resp.Diagnostics) {
		return
	}

	var plan types.Object
	diags := req.Plan.Get(ctx, &plan)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	staticHostUser := &userprovisioningv2.StaticHostUser{}
	diags = schemav1.CopyStaticHostUserFromTerraform(ctx, plan, staticHostUser)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	staticHostUserResource := staticHostUser


	
	name := staticHostUserResource.Metadata.Name

	staticHostUserBefore, err := r.p.Client.StaticHostUserClient().GetStaticHostUser(ctx, name)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading StaticHostUser", err, "static_host_user"))
		return
	}

	_, err = r.p.Client.StaticHostUserClient().UpsertStaticHostUser(ctx, staticHostUserResource)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error updating StaticHostUser", err, "static_host_user"))
		return
	}
		var staticHostUserI *userprovisioningv2.StaticHostUser

	tries := 0
	backoff := backoff.NewDecorr(r.p.RetryConfig.Base, r.p.RetryConfig.Cap, clockwork.NewRealClock())
	for {
		tries = tries + 1
		staticHostUserI, err = r.p.Client.StaticHostUserClient().GetStaticHostUser(ctx, name)
		if err != nil {
			resp.Diagnostics.Append(diagFromWrappedErr("Error reading StaticHostUser", err, "static_host_user"))
			return
		}
		if staticHostUserBefore.GetMetadata().Revision != staticHostUserI.GetMetadata().Revision || false {
			break
		}

		if err := backoff.Do(ctx); err != nil {
			resp.Diagnostics.Append(diagFromWrappedErr("Error reading StaticHostUser", trace.Wrap(err), "static_host_user"))
			return
		}
		if tries >= r.p.RetryConfig.MaxTries {
			diagMessage := fmt.Sprintf("Error reading StaticHostUser (tried %d times) - state outdated, please import resource", tries)
			resp.Diagnostics.AddError(diagMessage, "static_host_user")
			return
		}
	}

	staticHostUserResource = staticHostUserI
	
	diags = schemav1.CopyStaticHostUserToTerraform(ctx, staticHostUser, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes Teleport StaticHostUser
func (r resourceTeleportStaticHostUser) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var id types.String
	diags := req.State.GetAttribute(ctx, path.Root("metadata").AtName("name"), &id)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.p.Client.StaticHostUserClient().DeleteStaticHostUser(ctx, id.Value)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error deleting StaticHostUser", trace.Wrap(err), "static_host_user"))
		return
	}

	resp.State.RemoveResource(ctx)
}

// ImportState imports StaticHostUser state
func (r resourceTeleportStaticHostUser) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	staticHostUser, err := r.p.Client.StaticHostUserClient().GetStaticHostUser(ctx, req.ID)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading StaticHostUser", trace.Wrap(err), "static_host_user"))
		return
	}

	staticHostUserResource := staticHostUser
	

	var state types.Object

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = schemav1.CopyStaticHostUserToTerraform(ctx, staticHostUserResource, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	id := staticHostUser.Metadata.Name

	state.Attrs["id"] = types.String{Value: id}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
