//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// WorkspacesClient contains the methods for the Workspaces group.
// Don't use this type directly, use NewWorkspacesClient() instead.
type WorkspacesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkspacesClient creates a new instance of WorkspacesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWorkspacesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *WorkspacesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &WorkspacesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a workspace
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// workspaceInfo - Workspace create or update request properties
// options - WorkspacesClientBeginCreateOrUpdateOptions contains the optional parameters for the WorkspacesClient.BeginCreateOrUpdate
// method.
func (client *WorkspacesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, workspaceInfo Workspace, options *WorkspacesClientBeginCreateOrUpdateOptions) (WorkspacesClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, workspaceInfo, options)
	if err != nil {
		return WorkspacesClientCreateOrUpdatePollerResponse{}, err
	}
	result := WorkspacesClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("WorkspacesClient.CreateOrUpdate", "location", resp, client.pl)
	if err != nil {
		return WorkspacesClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &WorkspacesClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a workspace
// If the operation fails it returns an *azcore.ResponseError type.
func (client *WorkspacesClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, workspaceInfo Workspace, options *WorkspacesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, workspaceInfo, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkspacesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, workspaceInfo Workspace, options *WorkspacesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, workspaceInfo)
}

// BeginDelete - Deletes a workspace
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - WorkspacesClientBeginDeleteOptions contains the optional parameters for the WorkspacesClient.BeginDelete method.
func (client *WorkspacesClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspacesClientBeginDeleteOptions) (WorkspacesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return WorkspacesClientDeletePollerResponse{}, err
	}
	result := WorkspacesClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("WorkspacesClient.Delete", "location", resp, client.pl)
	if err != nil {
		return WorkspacesClientDeletePollerResponse{}, err
	}
	result.Poller = &WorkspacesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a workspace
// If the operation fails it returns an *azcore.ResponseError type.
func (client *WorkspacesClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspacesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WorkspacesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspacesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets a workspace
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - WorkspacesClientGetOptions contains the optional parameters for the WorkspacesClient.Get method.
func (client *WorkspacesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspacesClientGetOptions) (WorkspacesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return WorkspacesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspacesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspacesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkspacesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspacesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkspacesClient) getHandleResponse(resp *http.Response) (WorkspacesClientGetResponse, error) {
	result := WorkspacesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Workspace); err != nil {
		return WorkspacesClientGetResponse{}, err
	}
	return result, nil
}

// List - Returns a list of workspaces in a subscription
// If the operation fails it returns an *azcore.ResponseError type.
// options - WorkspacesClientListOptions contains the optional parameters for the WorkspacesClient.List method.
func (client *WorkspacesClient) List(options *WorkspacesClientListOptions) *WorkspacesClientListPager {
	return &WorkspacesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp WorkspacesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.WorkspaceInfoListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *WorkspacesClient) listCreateRequest(ctx context.Context, options *WorkspacesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Synapse/workspaces"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkspacesClient) listHandleResponse(resp *http.Response) (WorkspacesClientListResponse, error) {
	result := WorkspacesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceInfoListResult); err != nil {
		return WorkspacesClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Returns a list of workspaces in a resource group
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - WorkspacesClientListByResourceGroupOptions contains the optional parameters for the WorkspacesClient.ListByResourceGroup
// method.
func (client *WorkspacesClient) ListByResourceGroup(resourceGroupName string, options *WorkspacesClientListByResourceGroupOptions) *WorkspacesClientListByResourceGroupPager {
	return &WorkspacesClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp WorkspacesClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.WorkspaceInfoListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *WorkspacesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *WorkspacesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *WorkspacesClient) listByResourceGroupHandleResponse(resp *http.Response) (WorkspacesClientListByResourceGroupResponse, error) {
	result := WorkspacesClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceInfoListResult); err != nil {
		return WorkspacesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a workspace
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// workspacePatchInfo - Workspace patch request properties
// options - WorkspacesClientBeginUpdateOptions contains the optional parameters for the WorkspacesClient.BeginUpdate method.
func (client *WorkspacesClient) BeginUpdate(ctx context.Context, resourceGroupName string, workspaceName string, workspacePatchInfo WorkspacePatchInfo, options *WorkspacesClientBeginUpdateOptions) (WorkspacesClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, workspaceName, workspacePatchInfo, options)
	if err != nil {
		return WorkspacesClientUpdatePollerResponse{}, err
	}
	result := WorkspacesClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("WorkspacesClient.Update", "location", resp, client.pl)
	if err != nil {
		return WorkspacesClientUpdatePollerResponse{}, err
	}
	result.Poller = &WorkspacesClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates a workspace
// If the operation fails it returns an *azcore.ResponseError type.
func (client *WorkspacesClient) update(ctx context.Context, resourceGroupName string, workspaceName string, workspacePatchInfo WorkspacePatchInfo, options *WorkspacesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, workspaceName, workspacePatchInfo, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *WorkspacesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, workspacePatchInfo WorkspacePatchInfo, options *WorkspacesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, workspacePatchInfo)
}