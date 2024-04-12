// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220801

type Service_Policy_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the Policy.
	Properties *PolicyContractProperties_STATUS_ARM `json:"properties,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// Policy contract Properties.
type PolicyContractProperties_STATUS_ARM struct {
	// Format: Format of the policyContent.
	Format *PolicyContractProperties_Format_STATUS `json:"format,omitempty"`

	// Value: Contents of the Policy as defined by the format.
	Value *string `json:"value,omitempty"`
}
