package authorization

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// RoleManagementPoliciesClient is the client for the RoleManagementPolicies methods of the Authorization service.
type RoleManagementPoliciesClient struct {
	BaseClient
}

// NewRoleManagementPoliciesClient creates an instance of the RoleManagementPoliciesClient client.
func NewRoleManagementPoliciesClient(subscriptionID string) RoleManagementPoliciesClient {
	return NewRoleManagementPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRoleManagementPoliciesClientWithBaseURI creates an instance of the RoleManagementPoliciesClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewRoleManagementPoliciesClientWithBaseURI(baseURI string, subscriptionID string) RoleManagementPoliciesClient {
	return RoleManagementPoliciesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Delete delete a role management policy
// Parameters:
// scope - the scope of the role management policy to upsert.
// roleManagementPolicyName - the name (guid) of the role management policy to upsert.
func (client RoleManagementPoliciesClient) Delete(ctx context.Context, scope string, roleManagementPolicyName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleManagementPoliciesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, scope, roleManagementPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleManagementPoliciesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "authorization.RoleManagementPoliciesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleManagementPoliciesClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RoleManagementPoliciesClient) DeletePreparer(ctx context.Context, scope string, roleManagementPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"roleManagementPolicyName": autorest.Encode("path", roleManagementPolicyName),
		"scope":                    scope,
	}

	const APIVersion = "2020-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/roleManagementPolicies/{roleManagementPolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RoleManagementPoliciesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RoleManagementPoliciesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get the specified role management policy for a resource scope
// Parameters:
// scope - the scope of the role management policy.
// roleManagementPolicyName - the name (guid) of the role management policy to get.
func (client RoleManagementPoliciesClient) Get(ctx context.Context, scope string, roleManagementPolicyName string) (result RoleManagementPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleManagementPoliciesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, scope, roleManagementPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleManagementPoliciesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.RoleManagementPoliciesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleManagementPoliciesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client RoleManagementPoliciesClient) GetPreparer(ctx context.Context, scope string, roleManagementPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"roleManagementPolicyName": autorest.Encode("path", roleManagementPolicyName),
		"scope":                    scope,
	}

	const APIVersion = "2020-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/roleManagementPolicies/{roleManagementPolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RoleManagementPoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RoleManagementPoliciesClient) GetResponder(resp *http.Response) (result RoleManagementPolicy, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListForScope gets role management policies for a resource scope.
// Parameters:
// scope - the scope of the role management policy.
func (client RoleManagementPoliciesClient) ListForScope(ctx context.Context, scope string) (result RoleManagementPolicyListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleManagementPoliciesClient.ListForScope")
		defer func() {
			sc := -1
			if result.rmplr.Response.Response != nil {
				sc = result.rmplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listForScopeNextResults
	req, err := client.ListForScopePreparer(ctx, scope)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleManagementPoliciesClient", "ListForScope", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListForScopeSender(req)
	if err != nil {
		result.rmplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.RoleManagementPoliciesClient", "ListForScope", resp, "Failure sending request")
		return
	}

	result.rmplr, err = client.ListForScopeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleManagementPoliciesClient", "ListForScope", resp, "Failure responding to request")
		return
	}
	if result.rmplr.hasNextLink() && result.rmplr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListForScopePreparer prepares the ListForScope request.
func (client RoleManagementPoliciesClient) ListForScopePreparer(ctx context.Context, scope string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope": scope,
	}

	const APIVersion = "2020-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/roleManagementPolicies", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListForScopeSender sends the ListForScope request. The method will close the
// http.Response Body if it receives an error.
func (client RoleManagementPoliciesClient) ListForScopeSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListForScopeResponder handles the response to the ListForScope request. The method always
// closes the http.Response Body.
func (client RoleManagementPoliciesClient) ListForScopeResponder(resp *http.Response) (result RoleManagementPolicyListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listForScopeNextResults retrieves the next set of results, if any.
func (client RoleManagementPoliciesClient) listForScopeNextResults(ctx context.Context, lastResults RoleManagementPolicyListResult) (result RoleManagementPolicyListResult, err error) {
	req, err := lastResults.roleManagementPolicyListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.RoleManagementPoliciesClient", "listForScopeNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListForScopeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization.RoleManagementPoliciesClient", "listForScopeNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListForScopeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleManagementPoliciesClient", "listForScopeNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListForScopeComplete enumerates all values, automatically crossing page boundaries as required.
func (client RoleManagementPoliciesClient) ListForScopeComplete(ctx context.Context, scope string) (result RoleManagementPolicyListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleManagementPoliciesClient.ListForScope")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListForScope(ctx, scope)
	return
}

// Update update a role management policy
// Parameters:
// scope - the scope of the role management policy to upsert.
// roleManagementPolicyName - the name (guid) of the role management policy to upsert.
// parameters - parameters for the role management policy.
func (client RoleManagementPoliciesClient) Update(ctx context.Context, scope string, roleManagementPolicyName string, parameters RoleManagementPolicy) (result RoleManagementPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleManagementPoliciesClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, scope, roleManagementPolicyName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleManagementPoliciesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.RoleManagementPoliciesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleManagementPoliciesClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client RoleManagementPoliciesClient) UpdatePreparer(ctx context.Context, scope string, roleManagementPolicyName string, parameters RoleManagementPolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"roleManagementPolicyName": autorest.Encode("path", roleManagementPolicyName),
		"scope":                    scope,
	}

	const APIVersion = "2020-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.ID = nil
	parameters.Name = nil
	parameters.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/roleManagementPolicies/{roleManagementPolicyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client RoleManagementPoliciesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client RoleManagementPoliciesClient) UpdateResponder(resp *http.Response) (result RoleManagementPolicy, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}