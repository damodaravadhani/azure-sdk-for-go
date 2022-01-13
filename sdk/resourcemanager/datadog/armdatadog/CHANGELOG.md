# Release History

## 0.2.0 (2022-01-13)
### Breaking Changes

- Function `*MonitorsClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, *MonitorsBeginCreateOptions)` to `(context.Context, string, string, *MonitorsClientBeginCreateOptions)`
- Function `*MonitorsClient.BeginCreate` return value(s) have been changed from `(MonitorsCreatePollerResponse, error)` to `(MonitorsClientCreatePollerResponse, error)`
- Function `*TagRulesClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, *TagRulesCreateOrUpdateOptions)` to `(context.Context, string, string, string, *TagRulesClientCreateOrUpdateOptions)`
- Function `*TagRulesClient.CreateOrUpdate` return value(s) have been changed from `(TagRulesCreateOrUpdateResponse, error)` to `(TagRulesClientCreateOrUpdateResponse, error)`
- Function `*MonitorsClient.ListAPIKeys` parameter(s) have been changed from `(string, string, *MonitorsListAPIKeysOptions)` to `(string, string, *MonitorsClientListAPIKeysOptions)`
- Function `*MonitorsClient.ListAPIKeys` return value(s) have been changed from `(*MonitorsListAPIKeysPager)` to `(*MonitorsClientListAPIKeysPager)`
- Function `*MonitorsClient.ListHosts` parameter(s) have been changed from `(string, string, *MonitorsListHostsOptions)` to `(string, string, *MonitorsClientListHostsOptions)`
- Function `*MonitorsClient.ListHosts` return value(s) have been changed from `(*MonitorsListHostsPager)` to `(*MonitorsClientListHostsPager)`
- Function `*SingleSignOnConfigurationsClient.BeginCreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, *SingleSignOnConfigurationsBeginCreateOrUpdateOptions)` to `(context.Context, string, string, string, *SingleSignOnConfigurationsClientBeginCreateOrUpdateOptions)`
- Function `*SingleSignOnConfigurationsClient.BeginCreateOrUpdate` return value(s) have been changed from `(SingleSignOnConfigurationsCreateOrUpdatePollerResponse, error)` to `(SingleSignOnConfigurationsClientCreateOrUpdatePollerResponse, error)`
- Function `*MonitorsClient.SetDefaultKey` parameter(s) have been changed from `(context.Context, string, string, *MonitorsSetDefaultKeyOptions)` to `(context.Context, string, string, *MonitorsClientSetDefaultKeyOptions)`
- Function `*MonitorsClient.SetDefaultKey` return value(s) have been changed from `(MonitorsSetDefaultKeyResponse, error)` to `(MonitorsClientSetDefaultKeyResponse, error)`
- Function `*MonitorsClient.GetDefaultKey` parameter(s) have been changed from `(context.Context, string, string, *MonitorsGetDefaultKeyOptions)` to `(context.Context, string, string, *MonitorsClientGetDefaultKeyOptions)`
- Function `*MonitorsClient.GetDefaultKey` return value(s) have been changed from `(MonitorsGetDefaultKeyResponse, error)` to `(MonitorsClientGetDefaultKeyResponse, error)`
- Function `*MonitorsClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, *MonitorsBeginUpdateOptions)` to `(context.Context, string, string, *MonitorsClientBeginUpdateOptions)`
- Function `*MonitorsClient.BeginUpdate` return value(s) have been changed from `(MonitorsUpdatePollerResponse, error)` to `(MonitorsClientUpdatePollerResponse, error)`
- Function `*MarketplaceAgreementsClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, *MarketplaceAgreementsCreateOrUpdateOptions)` to `(context.Context, *MarketplaceAgreementsClientCreateOrUpdateOptions)`
- Function `*MarketplaceAgreementsClient.CreateOrUpdate` return value(s) have been changed from `(MarketplaceAgreementsCreateOrUpdateResponse, error)` to `(MarketplaceAgreementsClientCreateOrUpdateResponse, error)`
- Function `*MonitorsClient.Get` parameter(s) have been changed from `(context.Context, string, string, *MonitorsGetOptions)` to `(context.Context, string, string, *MonitorsClientGetOptions)`
- Function `*MonitorsClient.Get` return value(s) have been changed from `(MonitorsGetResponse, error)` to `(MonitorsClientGetResponse, error)`
- Function `*SingleSignOnConfigurationsClient.List` parameter(s) have been changed from `(string, string, *SingleSignOnConfigurationsListOptions)` to `(string, string, *SingleSignOnConfigurationsClientListOptions)`
- Function `*SingleSignOnConfigurationsClient.List` return value(s) have been changed from `(*SingleSignOnConfigurationsListPager)` to `(*SingleSignOnConfigurationsClientListPager)`
- Function `*TagRulesClient.List` parameter(s) have been changed from `(string, string, *TagRulesListOptions)` to `(string, string, *TagRulesClientListOptions)`
- Function `*TagRulesClient.List` return value(s) have been changed from `(*TagRulesListPager)` to `(*TagRulesClientListPager)`
- Function `*MonitorsClient.ListLinkedResources` parameter(s) have been changed from `(string, string, *MonitorsListLinkedResourcesOptions)` to `(string, string, *MonitorsClientListLinkedResourcesOptions)`
- Function `*MonitorsClient.ListLinkedResources` return value(s) have been changed from `(*MonitorsListLinkedResourcesPager)` to `(*MonitorsClientListLinkedResourcesPager)`
- Function `*MarketplaceAgreementsClient.List` parameter(s) have been changed from `(*MarketplaceAgreementsListOptions)` to `(*MarketplaceAgreementsClientListOptions)`
- Function `*MarketplaceAgreementsClient.List` return value(s) have been changed from `(*MarketplaceAgreementsListPager)` to `(*MarketplaceAgreementsClientListPager)`
- Function `*TagRulesClient.Get` parameter(s) have been changed from `(context.Context, string, string, string, *TagRulesGetOptions)` to `(context.Context, string, string, string, *TagRulesClientGetOptions)`
- Function `*TagRulesClient.Get` return value(s) have been changed from `(TagRulesGetResponse, error)` to `(TagRulesClientGetResponse, error)`
- Function `*MonitorsClient.RefreshSetPasswordLink` parameter(s) have been changed from `(context.Context, string, string, *MonitorsRefreshSetPasswordLinkOptions)` to `(context.Context, string, string, *MonitorsClientRefreshSetPasswordLinkOptions)`
- Function `*MonitorsClient.RefreshSetPasswordLink` return value(s) have been changed from `(MonitorsRefreshSetPasswordLinkResponse, error)` to `(MonitorsClientRefreshSetPasswordLinkResponse, error)`
- Function `*MonitorsClient.ListMonitoredResources` parameter(s) have been changed from `(string, string, *MonitorsListMonitoredResourcesOptions)` to `(string, string, *MonitorsClientListMonitoredResourcesOptions)`
- Function `*MonitorsClient.ListMonitoredResources` return value(s) have been changed from `(*MonitorsListMonitoredResourcesPager)` to `(*MonitorsClientListMonitoredResourcesPager)`
- Function `*OperationsClient.List` parameter(s) have been changed from `(*OperationsListOptions)` to `(*OperationsClientListOptions)`
- Function `*OperationsClient.List` return value(s) have been changed from `(*OperationsListPager)` to `(*OperationsClientListPager)`
- Function `*MonitorsClient.List` parameter(s) have been changed from `(*MonitorsListOptions)` to `(*MonitorsClientListOptions)`
- Function `*MonitorsClient.List` return value(s) have been changed from `(*MonitorsListPager)` to `(*MonitorsClientListPager)`
- Function `*MonitorsClient.BeginDelete` parameter(s) have been changed from `(context.Context, string, string, *MonitorsBeginDeleteOptions)` to `(context.Context, string, string, *MonitorsClientBeginDeleteOptions)`
- Function `*MonitorsClient.BeginDelete` return value(s) have been changed from `(MonitorsDeletePollerResponse, error)` to `(MonitorsClientDeletePollerResponse, error)`
- Function `*SingleSignOnConfigurationsClient.Get` parameter(s) have been changed from `(context.Context, string, string, string, *SingleSignOnConfigurationsGetOptions)` to `(context.Context, string, string, string, *SingleSignOnConfigurationsClientGetOptions)`
- Function `*SingleSignOnConfigurationsClient.Get` return value(s) have been changed from `(SingleSignOnConfigurationsGetResponse, error)` to `(SingleSignOnConfigurationsClientGetResponse, error)`
- Function `*MonitorsClient.ListByResourceGroup` parameter(s) have been changed from `(string, *MonitorsListByResourceGroupOptions)` to `(string, *MonitorsClientListByResourceGroupOptions)`
- Function `*MonitorsClient.ListByResourceGroup` return value(s) have been changed from `(*MonitorsListByResourceGroupPager)` to `(*MonitorsClientListByResourceGroupPager)`
- Type of `MonitorProperties.DatadogOrganizationProperties` has been changed from `*DatadogOrganizationProperties` to `*OrganizationProperties`
- Function `*MonitorsListAPIKeysPager.NextPage` has been removed
- Function `SingleSignOnConfigurationsCreateOrUpdatePollerResponse.PollUntilDone` has been removed
- Function `DatadogHostListResponse.MarshalJSON` has been removed
- Function `*MonitorsCreatePoller.ResumeToken` has been removed
- Function `*MonitorsDeletePoller.ResumeToken` has been removed
- Function `*MonitorsListLinkedResourcesPager.Err` has been removed
- Function `*MarketplaceAgreementsListPager.Err` has been removed
- Function `*DatadogAgreementProperties.UnmarshalJSON` has been removed
- Function `DatadogSingleSignOnResourceListResponse.MarshalJSON` has been removed
- Function `*MonitorsListByResourceGroupPager.NextPage` has been removed
- Function `*MonitorsUpdatePoller.ResumeToken` has been removed
- Function `*MonitorsListByResourceGroupPager.PageResponse` has been removed
- Function `*OperationsListPager.PageResponse` has been removed
- Function `*TagRulesListPager.NextPage` has been removed
- Function `*SingleSignOnConfigurationsCreateOrUpdatePollerResponse.Resume` has been removed
- Function `*MonitorsDeletePoller.FinalResponse` has been removed
- Function `DatadogAgreementResourceListResponse.MarshalJSON` has been removed
- Function `MonitorsUpdatePollerResponse.PollUntilDone` has been removed
- Function `*MonitorsCreatePoller.Poll` has been removed
- Function `DatadogHost.MarshalJSON` has been removed
- Function `*MonitorsListAPIKeysPager.Err` has been removed
- Function `DatadogMonitorResourceUpdateParameters.MarshalJSON` has been removed
- Function `MonitorsCreatePollerResponse.PollUntilDone` has been removed
- Function `*MarketplaceAgreementsListPager.PageResponse` has been removed
- Function `*TagRulesListPager.Err` has been removed
- Function `*MonitorsListPager.PageResponse` has been removed
- Function `*MonitorsListPager.NextPage` has been removed
- Function `*MonitorsDeletePollerResponse.Resume` has been removed
- Function `DatadogAPIKeyListResponse.MarshalJSON` has been removed
- Function `*MonitorsListMonitoredResourcesPager.PageResponse` has been removed
- Function `*SingleSignOnConfigurationsListPager.PageResponse` has been removed
- Function `*SingleSignOnConfigurationsCreateOrUpdatePoller.ResumeToken` has been removed
- Function `*MonitorsListHostsPager.NextPage` has been removed
- Function `*MonitorsListMonitoredResourcesPager.NextPage` has been removed
- Function `*MonitorsUpdatePoller.Poll` has been removed
- Function `MonitorsDeletePollerResponse.PollUntilDone` has been removed
- Function `*SingleSignOnConfigurationsListPager.Err` has been removed
- Function `*MonitorsListMonitoredResourcesPager.Err` has been removed
- Function `*TagRulesListPager.PageResponse` has been removed
- Function `*MonitorsUpdatePoller.Done` has been removed
- Function `*MonitorsListLinkedResourcesPager.PageResponse` has been removed
- Function `DatadogAgreementProperties.MarshalJSON` has been removed
- Function `*MonitorsListByResourceGroupPager.Err` has been removed
- Function `*MarketplaceAgreementsListPager.NextPage` has been removed
- Function `ErrorResponse.Error` has been removed
- Function `*SingleSignOnConfigurationsListPager.NextPage` has been removed
- Function `*MonitorsCreatePoller.FinalResponse` has been removed
- Function `*MonitorsDeletePoller.Done` has been removed
- Function `*SingleSignOnConfigurationsCreateOrUpdatePoller.Poll` has been removed
- Function `*MonitorsListHostsPager.PageResponse` has been removed
- Function `*MonitorsCreatePoller.Done` has been removed
- Function `*MonitorsListAPIKeysPager.PageResponse` has been removed
- Function `*SingleSignOnConfigurationsCreateOrUpdatePoller.FinalResponse` has been removed
- Function `*OperationsListPager.Err` has been removed
- Function `*MonitorsListLinkedResourcesPager.NextPage` has been removed
- Function `*MonitorsDeletePoller.Poll` has been removed
- Function `*SingleSignOnConfigurationsCreateOrUpdatePoller.Done` has been removed
- Function `*MonitorsCreatePollerResponse.Resume` has been removed
- Function `*OperationsListPager.NextPage` has been removed
- Function `DatadogMonitorResourceListResponse.MarshalJSON` has been removed
- Function `DatadogMonitorResource.MarshalJSON` has been removed
- Function `*MonitorsUpdatePollerResponse.Resume` has been removed
- Function `*MonitorsListHostsPager.Err` has been removed
- Function `*MonitorsUpdatePoller.FinalResponse` has been removed
- Function `*MonitorsListPager.Err` has been removed
- Struct `DatadogAPIKey` has been removed
- Struct `DatadogAPIKeyListResponse` has been removed
- Struct `DatadogAgreementProperties` has been removed
- Struct `DatadogAgreementResource` has been removed
- Struct `DatadogAgreementResourceListResponse` has been removed
- Struct `DatadogHost` has been removed
- Struct `DatadogHostListResponse` has been removed
- Struct `DatadogHostMetadata` has been removed
- Struct `DatadogInstallMethod` has been removed
- Struct `DatadogLogsAgent` has been removed
- Struct `DatadogMonitorResource` has been removed
- Struct `DatadogMonitorResourceListResponse` has been removed
- Struct `DatadogMonitorResourceUpdateParameters` has been removed
- Struct `DatadogOrganizationProperties` has been removed
- Struct `DatadogSetPasswordLink` has been removed
- Struct `DatadogSingleSignOnProperties` has been removed
- Struct `DatadogSingleSignOnResource` has been removed
- Struct `DatadogSingleSignOnResourceListResponse` has been removed
- Struct `MarketplaceAgreementsCreateOrUpdateOptions` has been removed
- Struct `MarketplaceAgreementsCreateOrUpdateResponse` has been removed
- Struct `MarketplaceAgreementsCreateOrUpdateResult` has been removed
- Struct `MarketplaceAgreementsListOptions` has been removed
- Struct `MarketplaceAgreementsListPager` has been removed
- Struct `MarketplaceAgreementsListResponse` has been removed
- Struct `MarketplaceAgreementsListResult` has been removed
- Struct `MonitorsBeginCreateOptions` has been removed
- Struct `MonitorsBeginDeleteOptions` has been removed
- Struct `MonitorsBeginUpdateOptions` has been removed
- Struct `MonitorsCreatePoller` has been removed
- Struct `MonitorsCreatePollerResponse` has been removed
- Struct `MonitorsCreateResponse` has been removed
- Struct `MonitorsCreateResult` has been removed
- Struct `MonitorsDeletePoller` has been removed
- Struct `MonitorsDeletePollerResponse` has been removed
- Struct `MonitorsDeleteResponse` has been removed
- Struct `MonitorsGetDefaultKeyOptions` has been removed
- Struct `MonitorsGetDefaultKeyResponse` has been removed
- Struct `MonitorsGetDefaultKeyResult` has been removed
- Struct `MonitorsGetOptions` has been removed
- Struct `MonitorsGetResponse` has been removed
- Struct `MonitorsGetResult` has been removed
- Struct `MonitorsListAPIKeysOptions` has been removed
- Struct `MonitorsListAPIKeysPager` has been removed
- Struct `MonitorsListAPIKeysResponse` has been removed
- Struct `MonitorsListAPIKeysResult` has been removed
- Struct `MonitorsListByResourceGroupOptions` has been removed
- Struct `MonitorsListByResourceGroupPager` has been removed
- Struct `MonitorsListByResourceGroupResponse` has been removed
- Struct `MonitorsListByResourceGroupResult` has been removed
- Struct `MonitorsListHostsOptions` has been removed
- Struct `MonitorsListHostsPager` has been removed
- Struct `MonitorsListHostsResponse` has been removed
- Struct `MonitorsListHostsResult` has been removed
- Struct `MonitorsListLinkedResourcesOptions` has been removed
- Struct `MonitorsListLinkedResourcesPager` has been removed
- Struct `MonitorsListLinkedResourcesResponse` has been removed
- Struct `MonitorsListLinkedResourcesResult` has been removed
- Struct `MonitorsListMonitoredResourcesOptions` has been removed
- Struct `MonitorsListMonitoredResourcesPager` has been removed
- Struct `MonitorsListMonitoredResourcesResponse` has been removed
- Struct `MonitorsListMonitoredResourcesResult` has been removed
- Struct `MonitorsListOptions` has been removed
- Struct `MonitorsListPager` has been removed
- Struct `MonitorsListResponse` has been removed
- Struct `MonitorsListResult` has been removed
- Struct `MonitorsRefreshSetPasswordLinkOptions` has been removed
- Struct `MonitorsRefreshSetPasswordLinkResponse` has been removed
- Struct `MonitorsRefreshSetPasswordLinkResult` has been removed
- Struct `MonitorsSetDefaultKeyOptions` has been removed
- Struct `MonitorsSetDefaultKeyResponse` has been removed
- Struct `MonitorsUpdatePoller` has been removed
- Struct `MonitorsUpdatePollerResponse` has been removed
- Struct `MonitorsUpdateResponse` has been removed
- Struct `MonitorsUpdateResult` has been removed
- Struct `OperationsListOptions` has been removed
- Struct `OperationsListPager` has been removed
- Struct `OperationsListResponse` has been removed
- Struct `OperationsListResult` has been removed
- Struct `SingleSignOnConfigurationsBeginCreateOrUpdateOptions` has been removed
- Struct `SingleSignOnConfigurationsCreateOrUpdatePoller` has been removed
- Struct `SingleSignOnConfigurationsCreateOrUpdatePollerResponse` has been removed
- Struct `SingleSignOnConfigurationsCreateOrUpdateResponse` has been removed
- Struct `SingleSignOnConfigurationsCreateOrUpdateResult` has been removed
- Struct `SingleSignOnConfigurationsGetOptions` has been removed
- Struct `SingleSignOnConfigurationsGetResponse` has been removed
- Struct `SingleSignOnConfigurationsGetResult` has been removed
- Struct `SingleSignOnConfigurationsListOptions` has been removed
- Struct `SingleSignOnConfigurationsListPager` has been removed
- Struct `SingleSignOnConfigurationsListResponse` has been removed
- Struct `SingleSignOnConfigurationsListResult` has been removed
- Struct `TagRulesCreateOrUpdateOptions` has been removed
- Struct `TagRulesCreateOrUpdateResponse` has been removed
- Struct `TagRulesCreateOrUpdateResult` has been removed
- Struct `TagRulesGetOptions` has been removed
- Struct `TagRulesGetResponse` has been removed
- Struct `TagRulesGetResult` has been removed
- Struct `TagRulesListOptions` has been removed
- Struct `TagRulesListPager` has been removed
- Struct `TagRulesListResponse` has been removed
- Struct `TagRulesListResult` has been removed
- Field `InnerError` of struct `ErrorResponse` has been removed

### Features Added

- New function `*MonitorsClientUpdatePoller.Done() bool`
- New function `*MonitorsClientListLinkedResourcesPager.NextPage(context.Context) bool`
- New function `*MonitorsClientListMonitoredResourcesPager.PageResponse() MonitorsClientListMonitoredResourcesResponse`
- New function `*MonitorsClientCreatePoller.ResumeToken() (string, error)`
- New function `*MonitorsClientUpdatePoller.Poll(context.Context) (*http.Response, error)`
- New function `*MarketplaceAgreementsClientListPager.PageResponse() MarketplaceAgreementsClientListResponse`
- New function `*MonitorsClientListLinkedResourcesPager.Err() error`
- New function `Host.MarshalJSON() ([]byte, error)`
- New function `MonitorsClientUpdatePollerResponse.PollUntilDone(context.Context, time.Duration) (MonitorsClientUpdateResponse, error)`
- New function `*MarketplaceAgreementsClientListPager.NextPage(context.Context) bool`
- New function `*TagRulesClientListPager.Err() error`
- New function `*MonitorsClientListPager.NextPage(context.Context) bool`
- New function `*SingleSignOnConfigurationsClientCreateOrUpdatePoller.FinalResponse(context.Context) (SingleSignOnConfigurationsClientCreateOrUpdateResponse, error)`
- New function `*MonitorsClientListLinkedResourcesPager.PageResponse() MonitorsClientListLinkedResourcesResponse`
- New function `MonitorResource.MarshalJSON() ([]byte, error)`
- New function `*MonitorsClientDeletePoller.FinalResponse(context.Context) (MonitorsClientDeleteResponse, error)`
- New function `*SingleSignOnConfigurationsClientListPager.PageResponse() SingleSignOnConfigurationsClientListResponse`
- New function `*MonitorsClientUpdatePoller.FinalResponse(context.Context) (MonitorsClientUpdateResponse, error)`
- New function `*MonitorsClientCreatePoller.Poll(context.Context) (*http.Response, error)`
- New function `*MonitorsClientCreatePoller.Done() bool`
- New function `AgreementResourceListResponse.MarshalJSON() ([]byte, error)`
- New function `*AgreementProperties.UnmarshalJSON([]byte) error`
- New function `*MonitorsClientListHostsPager.Err() error`
- New function `*MonitorsClientListHostsPager.PageResponse() MonitorsClientListHostsResponse`
- New function `*MonitorsClientCreatePoller.FinalResponse(context.Context) (MonitorsClientCreateResponse, error)`
- New function `*MonitorsClientUpdatePollerResponse.Resume(context.Context, *MonitorsClient, string) error`
- New function `MonitorsClientDeletePollerResponse.PollUntilDone(context.Context, time.Duration) (MonitorsClientDeleteResponse, error)`
- New function `*MonitorsClientListAPIKeysPager.PageResponse() MonitorsClientListAPIKeysResponse`
- New function `*SingleSignOnConfigurationsClientListPager.NextPage(context.Context) bool`
- New function `*MonitorsClientListAPIKeysPager.NextPage(context.Context) bool`
- New function `*MonitorsClientDeletePoller.Done() bool`
- New function `MonitorResourceUpdateParameters.MarshalJSON() ([]byte, error)`
- New function `*SingleSignOnConfigurationsClientCreateOrUpdatePollerResponse.Resume(context.Context, *SingleSignOnConfigurationsClient, string) error`
- New function `*MonitorsClientListAPIKeysPager.Err() error`
- New function `HostListResponse.MarshalJSON() ([]byte, error)`
- New function `*OperationsClientListPager.Err() error`
- New function `*MonitorsClientListByResourceGroupPager.Err() error`
- New function `*SingleSignOnConfigurationsClientListPager.Err() error`
- New function `*OperationsClientListPager.PageResponse() OperationsClientListResponse`
- New function `SingleSignOnConfigurationsClientCreateOrUpdatePollerResponse.PollUntilDone(context.Context, time.Duration) (SingleSignOnConfigurationsClientCreateOrUpdateResponse, error)`
- New function `*MarketplaceAgreementsClientListPager.Err() error`
- New function `SingleSignOnResourceListResponse.MarshalJSON() ([]byte, error)`
- New function `APIKeyListResponse.MarshalJSON() ([]byte, error)`
- New function `*MonitorsClientDeletePoller.ResumeToken() (string, error)`
- New function `*MonitorsClientCreatePollerResponse.Resume(context.Context, *MonitorsClient, string) error`
- New function `*MonitorsClientListMonitoredResourcesPager.Err() error`
- New function `MonitorsClientCreatePollerResponse.PollUntilDone(context.Context, time.Duration) (MonitorsClientCreateResponse, error)`
- New function `*SingleSignOnConfigurationsClientCreateOrUpdatePoller.ResumeToken() (string, error)`
- New function `*MonitorsClientListByResourceGroupPager.PageResponse() MonitorsClientListByResourceGroupResponse`
- New function `*OperationsClientListPager.NextPage(context.Context) bool`
- New function `*MonitorsClientListPager.Err() error`
- New function `*MonitorsClientListPager.PageResponse() MonitorsClientListResponse`
- New function `*SingleSignOnConfigurationsClientCreateOrUpdatePoller.Poll(context.Context) (*http.Response, error)`
- New function `*MonitorsClientUpdatePoller.ResumeToken() (string, error)`
- New function `*MonitorsClientDeletePollerResponse.Resume(context.Context, *MonitorsClient, string) error`
- New function `*SingleSignOnConfigurationsClientCreateOrUpdatePoller.Done() bool`
- New function `*MonitorsClientListHostsPager.NextPage(context.Context) bool`
- New function `*TagRulesClientListPager.PageResponse() TagRulesClientListResponse`
- New function `MonitorResourceListResponse.MarshalJSON() ([]byte, error)`
- New function `AgreementProperties.MarshalJSON() ([]byte, error)`
- New function `*MonitorsClientListByResourceGroupPager.NextPage(context.Context) bool`
- New function `*TagRulesClientListPager.NextPage(context.Context) bool`
- New function `*MonitorsClientDeletePoller.Poll(context.Context) (*http.Response, error)`
- New function `*MonitorsClientListMonitoredResourcesPager.NextPage(context.Context) bool`
- New struct `APIKey`
- New struct `APIKeyListResponse`
- New struct `AgreementProperties`
- New struct `AgreementResource`
- New struct `AgreementResourceListResponse`
- New struct `Host`
- New struct `HostListResponse`
- New struct `HostMetadata`
- New struct `InstallMethod`
- New struct `LogsAgent`
- New struct `MarketplaceAgreementsClientCreateOrUpdateOptions`
- New struct `MarketplaceAgreementsClientCreateOrUpdateResponse`
- New struct `MarketplaceAgreementsClientCreateOrUpdateResult`
- New struct `MarketplaceAgreementsClientListOptions`
- New struct `MarketplaceAgreementsClientListPager`
- New struct `MarketplaceAgreementsClientListResponse`
- New struct `MarketplaceAgreementsClientListResult`
- New struct `MonitorResource`
- New struct `MonitorResourceListResponse`
- New struct `MonitorResourceUpdateParameters`
- New struct `MonitorsClientBeginCreateOptions`
- New struct `MonitorsClientBeginDeleteOptions`
- New struct `MonitorsClientBeginUpdateOptions`
- New struct `MonitorsClientCreatePoller`
- New struct `MonitorsClientCreatePollerResponse`
- New struct `MonitorsClientCreateResponse`
- New struct `MonitorsClientCreateResult`
- New struct `MonitorsClientDeletePoller`
- New struct `MonitorsClientDeletePollerResponse`
- New struct `MonitorsClientDeleteResponse`
- New struct `MonitorsClientGetDefaultKeyOptions`
- New struct `MonitorsClientGetDefaultKeyResponse`
- New struct `MonitorsClientGetDefaultKeyResult`
- New struct `MonitorsClientGetOptions`
- New struct `MonitorsClientGetResponse`
- New struct `MonitorsClientGetResult`
- New struct `MonitorsClientListAPIKeysOptions`
- New struct `MonitorsClientListAPIKeysPager`
- New struct `MonitorsClientListAPIKeysResponse`
- New struct `MonitorsClientListAPIKeysResult`
- New struct `MonitorsClientListByResourceGroupOptions`
- New struct `MonitorsClientListByResourceGroupPager`
- New struct `MonitorsClientListByResourceGroupResponse`
- New struct `MonitorsClientListByResourceGroupResult`
- New struct `MonitorsClientListHostsOptions`
- New struct `MonitorsClientListHostsPager`
- New struct `MonitorsClientListHostsResponse`
- New struct `MonitorsClientListHostsResult`
- New struct `MonitorsClientListLinkedResourcesOptions`
- New struct `MonitorsClientListLinkedResourcesPager`
- New struct `MonitorsClientListLinkedResourcesResponse`
- New struct `MonitorsClientListLinkedResourcesResult`
- New struct `MonitorsClientListMonitoredResourcesOptions`
- New struct `MonitorsClientListMonitoredResourcesPager`
- New struct `MonitorsClientListMonitoredResourcesResponse`
- New struct `MonitorsClientListMonitoredResourcesResult`
- New struct `MonitorsClientListOptions`
- New struct `MonitorsClientListPager`
- New struct `MonitorsClientListResponse`
- New struct `MonitorsClientListResult`
- New struct `MonitorsClientRefreshSetPasswordLinkOptions`
- New struct `MonitorsClientRefreshSetPasswordLinkResponse`
- New struct `MonitorsClientRefreshSetPasswordLinkResult`
- New struct `MonitorsClientSetDefaultKeyOptions`
- New struct `MonitorsClientSetDefaultKeyResponse`
- New struct `MonitorsClientUpdatePoller`
- New struct `MonitorsClientUpdatePollerResponse`
- New struct `MonitorsClientUpdateResponse`
- New struct `MonitorsClientUpdateResult`
- New struct `OperationsClientListOptions`
- New struct `OperationsClientListPager`
- New struct `OperationsClientListResponse`
- New struct `OperationsClientListResult`
- New struct `OrganizationProperties`
- New struct `SetPasswordLink`
- New struct `SingleSignOnConfigurationsClientBeginCreateOrUpdateOptions`
- New struct `SingleSignOnConfigurationsClientCreateOrUpdatePoller`
- New struct `SingleSignOnConfigurationsClientCreateOrUpdatePollerResponse`
- New struct `SingleSignOnConfigurationsClientCreateOrUpdateResponse`
- New struct `SingleSignOnConfigurationsClientCreateOrUpdateResult`
- New struct `SingleSignOnConfigurationsClientGetOptions`
- New struct `SingleSignOnConfigurationsClientGetResponse`
- New struct `SingleSignOnConfigurationsClientGetResult`
- New struct `SingleSignOnConfigurationsClientListOptions`
- New struct `SingleSignOnConfigurationsClientListPager`
- New struct `SingleSignOnConfigurationsClientListResponse`
- New struct `SingleSignOnConfigurationsClientListResult`
- New struct `SingleSignOnProperties`
- New struct `SingleSignOnResource`
- New struct `SingleSignOnResourceListResponse`
- New struct `TagRulesClientCreateOrUpdateOptions`
- New struct `TagRulesClientCreateOrUpdateResponse`
- New struct `TagRulesClientCreateOrUpdateResult`
- New struct `TagRulesClientGetOptions`
- New struct `TagRulesClientGetResponse`
- New struct `TagRulesClientGetResult`
- New struct `TagRulesClientListOptions`
- New struct `TagRulesClientListPager`
- New struct `TagRulesClientListResponse`
- New struct `TagRulesClientListResult`
- New field `Error` in struct `ErrorResponse`


## 0.1.0 (2021-12-22)

- Init release.