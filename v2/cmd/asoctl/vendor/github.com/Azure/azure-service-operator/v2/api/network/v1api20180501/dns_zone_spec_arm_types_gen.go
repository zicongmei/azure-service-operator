// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20180501

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DnsZone_Spec_ARM struct {
	// Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: The properties of the zone.
	Properties *ZoneProperties_ARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DnsZone_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-05-01"
func (zone DnsZone_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (zone *DnsZone_Spec_ARM) GetName() string {
	return zone.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/dnsZones"
func (zone *DnsZone_Spec_ARM) GetType() string {
	return "Microsoft.Network/dnsZones"
}

// Represents the properties of the zone.
type ZoneProperties_ARM struct {
	// RegistrationVirtualNetworks: A list of references to virtual networks that register hostnames in this DNS zone. This is
	// a only when ZoneType is Private.
	RegistrationVirtualNetworks []SubResource_ARM `json:"registrationVirtualNetworks,omitempty"`

	// ResolutionVirtualNetworks: A list of references to virtual networks that resolve records in this DNS zone. This is a
	// only when ZoneType is Private.
	ResolutionVirtualNetworks []SubResource_ARM `json:"resolutionVirtualNetworks,omitempty"`

	// ZoneType: The type of this DNS zone (Public or Private).
	ZoneType *ZoneProperties_ZoneType `json:"zoneType,omitempty"`
}

// A reference to a another resource
type SubResource_ARM struct {
	Id *string `json:"id,omitempty"`
}
