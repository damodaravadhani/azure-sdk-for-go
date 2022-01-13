//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragecache

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// AscOperationsClientGetResponse contains the response from method AscOperationsClient.Get.
type AscOperationsClientGetResponse struct {
	AscOperationsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AscOperationsClientGetResult contains the result from method AscOperationsClient.Get.
type AscOperationsClientGetResult struct {
	AscOperation
}

// CachesClientCreateOrUpdatePollerResponse contains the response from method CachesClient.CreateOrUpdate.
type CachesClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *CachesClientCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l CachesClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (CachesClientCreateOrUpdateResponse, error) {
	respType := CachesClientCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Cache)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a CachesClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *CachesClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *CachesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("CachesClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &CachesClientCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// CachesClientCreateOrUpdateResponse contains the response from method CachesClient.CreateOrUpdate.
type CachesClientCreateOrUpdateResponse struct {
	CachesClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CachesClientCreateOrUpdateResult contains the result from method CachesClient.CreateOrUpdate.
type CachesClientCreateOrUpdateResult struct {
	Cache
}

// CachesClientDebugInfoPollerResponse contains the response from method CachesClient.DebugInfo.
type CachesClientDebugInfoPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *CachesClientDebugInfoPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l CachesClientDebugInfoPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (CachesClientDebugInfoResponse, error) {
	respType := CachesClientDebugInfoResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a CachesClientDebugInfoPollerResponse from the provided client and resume token.
func (l *CachesClientDebugInfoPollerResponse) Resume(ctx context.Context, client *CachesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("CachesClient.DebugInfo", token, client.pl)
	if err != nil {
		return err
	}
	poller := &CachesClientDebugInfoPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// CachesClientDebugInfoResponse contains the response from method CachesClient.DebugInfo.
type CachesClientDebugInfoResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CachesClientDeletePollerResponse contains the response from method CachesClient.Delete.
type CachesClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *CachesClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l CachesClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (CachesClientDeleteResponse, error) {
	respType := CachesClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a CachesClientDeletePollerResponse from the provided client and resume token.
func (l *CachesClientDeletePollerResponse) Resume(ctx context.Context, client *CachesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("CachesClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &CachesClientDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// CachesClientDeleteResponse contains the response from method CachesClient.Delete.
type CachesClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CachesClientFlushPollerResponse contains the response from method CachesClient.Flush.
type CachesClientFlushPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *CachesClientFlushPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l CachesClientFlushPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (CachesClientFlushResponse, error) {
	respType := CachesClientFlushResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a CachesClientFlushPollerResponse from the provided client and resume token.
func (l *CachesClientFlushPollerResponse) Resume(ctx context.Context, client *CachesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("CachesClient.Flush", token, client.pl)
	if err != nil {
		return err
	}
	poller := &CachesClientFlushPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// CachesClientFlushResponse contains the response from method CachesClient.Flush.
type CachesClientFlushResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CachesClientGetResponse contains the response from method CachesClient.Get.
type CachesClientGetResponse struct {
	CachesClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CachesClientGetResult contains the result from method CachesClient.Get.
type CachesClientGetResult struct {
	Cache
}

// CachesClientListByResourceGroupResponse contains the response from method CachesClient.ListByResourceGroup.
type CachesClientListByResourceGroupResponse struct {
	CachesClientListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CachesClientListByResourceGroupResult contains the result from method CachesClient.ListByResourceGroup.
type CachesClientListByResourceGroupResult struct {
	CachesListResult
}

// CachesClientListResponse contains the response from method CachesClient.List.
type CachesClientListResponse struct {
	CachesClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CachesClientListResult contains the result from method CachesClient.List.
type CachesClientListResult struct {
	CachesListResult
}

// CachesClientStartPollerResponse contains the response from method CachesClient.Start.
type CachesClientStartPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *CachesClientStartPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l CachesClientStartPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (CachesClientStartResponse, error) {
	respType := CachesClientStartResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a CachesClientStartPollerResponse from the provided client and resume token.
func (l *CachesClientStartPollerResponse) Resume(ctx context.Context, client *CachesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("CachesClient.Start", token, client.pl)
	if err != nil {
		return err
	}
	poller := &CachesClientStartPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// CachesClientStartResponse contains the response from method CachesClient.Start.
type CachesClientStartResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CachesClientStopPollerResponse contains the response from method CachesClient.Stop.
type CachesClientStopPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *CachesClientStopPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l CachesClientStopPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (CachesClientStopResponse, error) {
	respType := CachesClientStopResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a CachesClientStopPollerResponse from the provided client and resume token.
func (l *CachesClientStopPollerResponse) Resume(ctx context.Context, client *CachesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("CachesClient.Stop", token, client.pl)
	if err != nil {
		return err
	}
	poller := &CachesClientStopPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// CachesClientStopResponse contains the response from method CachesClient.Stop.
type CachesClientStopResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CachesClientUpdateResponse contains the response from method CachesClient.Update.
type CachesClientUpdateResponse struct {
	CachesClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CachesClientUpdateResult contains the result from method CachesClient.Update.
type CachesClientUpdateResult struct {
	Cache
}

// CachesClientUpgradeFirmwarePollerResponse contains the response from method CachesClient.UpgradeFirmware.
type CachesClientUpgradeFirmwarePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *CachesClientUpgradeFirmwarePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l CachesClientUpgradeFirmwarePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (CachesClientUpgradeFirmwareResponse, error) {
	respType := CachesClientUpgradeFirmwareResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a CachesClientUpgradeFirmwarePollerResponse from the provided client and resume token.
func (l *CachesClientUpgradeFirmwarePollerResponse) Resume(ctx context.Context, client *CachesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("CachesClient.UpgradeFirmware", token, client.pl)
	if err != nil {
		return err
	}
	poller := &CachesClientUpgradeFirmwarePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// CachesClientUpgradeFirmwareResponse contains the response from method CachesClient.UpgradeFirmware.
type CachesClientUpgradeFirmwareResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsClientListResult contains the result from method OperationsClient.List.
type OperationsClientListResult struct {
	APIOperationListResult
}

// SKUsClientListResponse contains the response from method SKUsClient.List.
type SKUsClientListResponse struct {
	SKUsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SKUsClientListResult contains the result from method SKUsClient.List.
type SKUsClientListResult struct {
	ResourceSKUsResult
}

// StorageTargetClientFlushPollerResponse contains the response from method StorageTargetClient.Flush.
type StorageTargetClientFlushPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *StorageTargetClientFlushPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l StorageTargetClientFlushPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (StorageTargetClientFlushResponse, error) {
	respType := StorageTargetClientFlushResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a StorageTargetClientFlushPollerResponse from the provided client and resume token.
func (l *StorageTargetClientFlushPollerResponse) Resume(ctx context.Context, client *StorageTargetClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("StorageTargetClient.Flush", token, client.pl)
	if err != nil {
		return err
	}
	poller := &StorageTargetClientFlushPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// StorageTargetClientFlushResponse contains the response from method StorageTargetClient.Flush.
type StorageTargetClientFlushResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StorageTargetClientResumePollerResponse contains the response from method StorageTargetClient.Resume.
type StorageTargetClientResumePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *StorageTargetClientResumePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l StorageTargetClientResumePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (StorageTargetClientResumeResponse, error) {
	respType := StorageTargetClientResumeResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a StorageTargetClientResumePollerResponse from the provided client and resume token.
func (l *StorageTargetClientResumePollerResponse) Resume(ctx context.Context, client *StorageTargetClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("StorageTargetClient.Resume", token, client.pl)
	if err != nil {
		return err
	}
	poller := &StorageTargetClientResumePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// StorageTargetClientResumeResponse contains the response from method StorageTargetClient.Resume.
type StorageTargetClientResumeResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StorageTargetClientSuspendPollerResponse contains the response from method StorageTargetClient.Suspend.
type StorageTargetClientSuspendPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *StorageTargetClientSuspendPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l StorageTargetClientSuspendPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (StorageTargetClientSuspendResponse, error) {
	respType := StorageTargetClientSuspendResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a StorageTargetClientSuspendPollerResponse from the provided client and resume token.
func (l *StorageTargetClientSuspendPollerResponse) Resume(ctx context.Context, client *StorageTargetClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("StorageTargetClient.Suspend", token, client.pl)
	if err != nil {
		return err
	}
	poller := &StorageTargetClientSuspendPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// StorageTargetClientSuspendResponse contains the response from method StorageTargetClient.Suspend.
type StorageTargetClientSuspendResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StorageTargetsClientCreateOrUpdatePollerResponse contains the response from method StorageTargetsClient.CreateOrUpdate.
type StorageTargetsClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *StorageTargetsClientCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l StorageTargetsClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (StorageTargetsClientCreateOrUpdateResponse, error) {
	respType := StorageTargetsClientCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.StorageTarget)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a StorageTargetsClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *StorageTargetsClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *StorageTargetsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("StorageTargetsClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &StorageTargetsClientCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// StorageTargetsClientCreateOrUpdateResponse contains the response from method StorageTargetsClient.CreateOrUpdate.
type StorageTargetsClientCreateOrUpdateResponse struct {
	StorageTargetsClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StorageTargetsClientCreateOrUpdateResult contains the result from method StorageTargetsClient.CreateOrUpdate.
type StorageTargetsClientCreateOrUpdateResult struct {
	StorageTarget
}

// StorageTargetsClientDNSRefreshPollerResponse contains the response from method StorageTargetsClient.DNSRefresh.
type StorageTargetsClientDNSRefreshPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *StorageTargetsClientDNSRefreshPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l StorageTargetsClientDNSRefreshPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (StorageTargetsClientDNSRefreshResponse, error) {
	respType := StorageTargetsClientDNSRefreshResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a StorageTargetsClientDNSRefreshPollerResponse from the provided client and resume token.
func (l *StorageTargetsClientDNSRefreshPollerResponse) Resume(ctx context.Context, client *StorageTargetsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("StorageTargetsClient.DNSRefresh", token, client.pl)
	if err != nil {
		return err
	}
	poller := &StorageTargetsClientDNSRefreshPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// StorageTargetsClientDNSRefreshResponse contains the response from method StorageTargetsClient.DNSRefresh.
type StorageTargetsClientDNSRefreshResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StorageTargetsClientDeletePollerResponse contains the response from method StorageTargetsClient.Delete.
type StorageTargetsClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *StorageTargetsClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l StorageTargetsClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (StorageTargetsClientDeleteResponse, error) {
	respType := StorageTargetsClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a StorageTargetsClientDeletePollerResponse from the provided client and resume token.
func (l *StorageTargetsClientDeletePollerResponse) Resume(ctx context.Context, client *StorageTargetsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("StorageTargetsClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &StorageTargetsClientDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// StorageTargetsClientDeleteResponse contains the response from method StorageTargetsClient.Delete.
type StorageTargetsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StorageTargetsClientGetResponse contains the response from method StorageTargetsClient.Get.
type StorageTargetsClientGetResponse struct {
	StorageTargetsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StorageTargetsClientGetResult contains the result from method StorageTargetsClient.Get.
type StorageTargetsClientGetResult struct {
	StorageTarget
}

// StorageTargetsClientListByCacheResponse contains the response from method StorageTargetsClient.ListByCache.
type StorageTargetsClientListByCacheResponse struct {
	StorageTargetsClientListByCacheResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StorageTargetsClientListByCacheResult contains the result from method StorageTargetsClient.ListByCache.
type StorageTargetsClientListByCacheResult struct {
	StorageTargetsResult
}

// UsageModelsClientListResponse contains the response from method UsageModelsClient.List.
type UsageModelsClientListResponse struct {
	UsageModelsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// UsageModelsClientListResult contains the result from method UsageModelsClient.List.
type UsageModelsClientListResult struct {
	UsageModelsResult
}