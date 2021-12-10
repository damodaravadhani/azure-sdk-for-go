//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataprotection

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ExportJobsClient contains the methods for the ExportJobs group.
// Don't use this type directly, use NewExportJobsClient() instead.
type ExportJobsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewExportJobsClient creates a new instance of ExportJobsClient with the specified values.
func NewExportJobsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ExportJobsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ExportJobsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginTrigger - Triggers export of jobs and returns an OperationID to track.
// If the operation fails it returns the *CloudError error type.
func (client *ExportJobsClient) BeginTrigger(ctx context.Context, resourceGroupName string, vaultName string, options *ExportJobsBeginTriggerOptions) (ExportJobsTriggerPollerResponse, error) {
	resp, err := client.trigger(ctx, resourceGroupName, vaultName, options)
	if err != nil {
		return ExportJobsTriggerPollerResponse{}, err
	}
	result := ExportJobsTriggerPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExportJobsClient.Trigger", "location", resp, client.pl, client.triggerHandleError)
	if err != nil {
		return ExportJobsTriggerPollerResponse{}, err
	}
	result.Poller = &ExportJobsTriggerPoller{
		pt: pt,
	}
	return result, nil
}

// Trigger - Triggers export of jobs and returns an OperationID to track.
// If the operation fails it returns the *CloudError error type.
func (client *ExportJobsClient) trigger(ctx context.Context, resourceGroupName string, vaultName string, options *ExportJobsBeginTriggerOptions) (*http.Response, error) {
	req, err := client.triggerCreateRequest(ctx, resourceGroupName, vaultName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.triggerHandleError(resp)
	}
	return resp, nil
}

// triggerCreateRequest creates the Trigger request.
func (client *ExportJobsClient) triggerCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, options *ExportJobsBeginTriggerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/exportBackupJobs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// triggerHandleError handles the Trigger error response.
func (client *ExportJobsClient) triggerHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}