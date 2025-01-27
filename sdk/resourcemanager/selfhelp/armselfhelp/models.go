//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armselfhelp

import "time"

// CheckNameAvailabilityRequest - The check availability request body.
type CheckNameAvailabilityRequest struct {
	// The name of the resource for which availability needs to be checked.
	Name *string `json:"name,omitempty"`

	// The resource type.
	Type *string `json:"type,omitempty"`
}

// CheckNameAvailabilityResponse - Response for whether the requested resource name is available or not.
type CheckNameAvailabilityResponse struct {
	// Gets an error message explaining the 'reason' value with more details. This field is returned iif nameAvailable is false.
	Message *string `json:"message,omitempty"`

	// Returns true or false depending on the availability of the name
	NameAvailable *bool `json:"nameAvailable,omitempty"`

	// Reason for why value is not available. This field is returned if nameAvailable is false.
	Reason *string `json:"reason,omitempty"`
}

// Diagnostic - Properties returned with in an insight.
type Diagnostic struct {
	// Error definition.
	Error *Error `json:"error,omitempty"`

	// The problems (if any) detected by this insight.
	Insights []*Insight `json:"insights,omitempty"`

	// Solution Id
	SolutionID *string `json:"solutionId,omitempty"`

	// Denotes the status of the diagnostic resource.
	Status *Status `json:"status,omitempty"`
}

// DiagnosticInvocation - Solution Invocation with additional params needed for invocation.
type DiagnosticInvocation struct {
	// Additional parameters required to invoke the solutionId.
	AdditionalParameters map[string]*string `json:"additionalParameters,omitempty"`

	// Solution Id to invoke.
	SolutionID *string `json:"solutionId,omitempty"`
}

// DiagnosticResource - Diagnostic resource
type DiagnosticResource struct {
	// Diagnostic Resource properties.
	Properties *DiagnosticResourceProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// DiagnosticResourceProperties - Diagnostic resource properties.
type DiagnosticResourceProperties struct {
	// Global parameters that can be passed to all solutionIds.
	GlobalParameters map[string]*string `json:"globalParameters,omitempty"`

	// SolutionIds that are needed to be invoked.
	Insights []*DiagnosticInvocation `json:"insights,omitempty"`

	// READ-ONLY; Diagnostic Request Accepted time.
	AcceptedAt *string `json:"acceptedAt,omitempty" azure:"ro"`

	// READ-ONLY; Array of Diagnostics.
	Diagnostics []*Diagnostic `json:"diagnostics,omitempty" azure:"ro"`

	// READ-ONLY; Status of diagnostic provisioning.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// DiagnosticsClientBeginCreateOptions contains the optional parameters for the DiagnosticsClient.BeginCreate method.
type DiagnosticsClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DiagnosticsClientCheckNameAvailabilityOptions contains the optional parameters for the DiagnosticsClient.CheckNameAvailability
// method.
type DiagnosticsClientCheckNameAvailabilityOptions struct {
	// The required parameters for availability check.
	CheckNameAvailabilityRequest *CheckNameAvailabilityRequest
}

// DiagnosticsClientGetOptions contains the optional parameters for the DiagnosticsClient.Get method.
type DiagnosticsClientGetOptions struct {
	// placeholder for future optional parameters
}

// DiscoveryResponse - Discovery response.
type DiscoveryResponse struct {
	// The link used to get the next page of solution metadata.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of solution metadata.
	Value []*SolutionMetadataResource `json:"value,omitempty"`
}

// DiscoverySolutionClientListOptions contains the optional parameters for the DiscoverySolutionClient.NewListPager method.
type DiscoverySolutionClientListOptions struct {
	// Can be used to filter solutionIds by 'ProblemClassificationId'. The filter supports only 'and' and 'eq' operators. Example:
	// $filter=ProblemClassificationId eq '1ddda5b4-cf6c-4d4f-91ad-bc38ab0e811e'
	// and ProblemClassificationId eq '0a9673c2-7af6-4e19-90d3-4ee2461076d9'.
	Filter *string
	// Skiptoken is only used if a previous operation returned a partial result. If a previous response contains a nextLink element,
	// the value of the nextLink element will include a skiptoken parameter that
	// specifies a starting point to use for subsequent calls.
	Skiptoken *string
}

// Error definition.
type Error struct {
	// An array of additional nested error response info objects, as described by this contract.
	Details []*Error `json:"details,omitempty"`

	// READ-ONLY; Service specific error code which serves as the substatus for the HTTP error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; Description of the error.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; Service specific error type which serves as additional context for the error herein.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// Insight - Detailed insights(s) obtained via the invocation of an insight diagnostic troubleshooter.
type Insight struct {
	// Article id.
	ID *string `json:"id,omitempty"`

	// Importance level of the insight.
	ImportanceLevel *ImportanceLevel `json:"importanceLevel,omitempty"`

	// Detailed result content.
	Results *string `json:"results,omitempty"`

	// This insight's title.
	Title *string `json:"title,omitempty"`
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType `json:"actionType,omitempty" azure:"ro"`

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool `json:"isDataAction,omitempty" azure:"ro"`

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin `json:"origin,omitempty" azure:"ro"`
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// SolutionMetadataProperties - Diagnostic solution metadata.
type SolutionMetadataProperties struct {
	// A detailed description of solution.
	Description *string `json:"description,omitempty"`

	// Required parameters for invoking this particular solution.
	RequiredParameterSets [][]*string `json:"requiredParameterSets,omitempty"`

	// Solution Id.
	SolutionID *string `json:"solutionId,omitempty"`

	// Solution Type.
	SolutionType *string `json:"solutionType,omitempty"`
}

// SolutionMetadataResource - Solution Metadata resource
type SolutionMetadataResource struct {
	// Solution metadata Resource properties.
	Properties *SolutionMetadataProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}
