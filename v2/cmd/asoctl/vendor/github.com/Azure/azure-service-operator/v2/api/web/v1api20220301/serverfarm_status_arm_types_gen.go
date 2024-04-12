// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220301

type Serverfarm_STATUS_ARM struct {
	// ExtendedLocation: Extended Location.
	ExtendedLocation *ExtendedLocation_STATUS_ARM `json:"extendedLocation,omitempty"`

	// Id: Resource Id.
	Id *string `json:"id,omitempty"`

	// Kind: Kind of resource.
	Kind *string `json:"kind,omitempty"`

	// Location: Resource Location.
	Location *string `json:"location,omitempty"`

	// Name: Resource Name.
	Name *string `json:"name,omitempty"`

	// Properties: AppServicePlan resource specific properties
	Properties *Serverfarm_Properties_STATUS_ARM `json:"properties,omitempty"`

	// Sku: Description of a SKU for a scalable resource.
	Sku *SkuDescription_STATUS_ARM `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// Extended Location.
type ExtendedLocation_STATUS_ARM struct {
	// Name: Name of extended location.
	Name *string `json:"name,omitempty"`

	// Type: Type of extended location.
	Type *string `json:"type,omitempty"`
}

type Serverfarm_Properties_STATUS_ARM struct {
	// ElasticScaleEnabled: ServerFarm supports ElasticScale. Apps in this plan will scale as if the ServerFarm was
	// ElasticPremium sku
	ElasticScaleEnabled *bool `json:"elasticScaleEnabled,omitempty"`

	// FreeOfferExpirationTime: The time when the server farm free offer expires.
	FreeOfferExpirationTime *string `json:"freeOfferExpirationTime,omitempty"`

	// GeoRegion: Geographical location for the App Service plan.
	GeoRegion *string `json:"geoRegion,omitempty"`

	// HostingEnvironmentProfile: Specification for the App Service Environment to use for the App Service plan.
	HostingEnvironmentProfile *HostingEnvironmentProfile_STATUS_ARM `json:"hostingEnvironmentProfile,omitempty"`

	// HyperV: If Hyper-V container app service plan <code>true</code>, <code>false</code> otherwise.
	HyperV *bool `json:"hyperV,omitempty"`

	// IsSpot: If <code>true</code>, this App Service Plan owns spot instances.
	IsSpot *bool `json:"isSpot,omitempty"`

	// IsXenon: Obsolete: If Hyper-V container app service plan <code>true</code>, <code>false</code> otherwise.
	IsXenon *bool `json:"isXenon,omitempty"`

	// KubeEnvironmentProfile: Specification for the Kubernetes Environment to use for the App Service plan.
	KubeEnvironmentProfile *KubeEnvironmentProfile_STATUS_ARM `json:"kubeEnvironmentProfile,omitempty"`

	// MaximumElasticWorkerCount: Maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan
	MaximumElasticWorkerCount *int `json:"maximumElasticWorkerCount,omitempty"`

	// MaximumNumberOfWorkers: Maximum number of instances that can be assigned to this App Service plan.
	MaximumNumberOfWorkers *int `json:"maximumNumberOfWorkers,omitempty"`

	// NumberOfSites: Number of apps assigned to this App Service plan.
	NumberOfSites *int `json:"numberOfSites,omitempty"`

	// NumberOfWorkers: The number of instances that are assigned to this App Service plan.
	NumberOfWorkers *int `json:"numberOfWorkers,omitempty"`

	// PerSiteScaling: If <code>true</code>, apps assigned to this App Service plan can be scaled independently.
	// If <code>false</code>, apps assigned to this App Service plan will scale to all instances of the plan.
	PerSiteScaling *bool `json:"perSiteScaling,omitempty"`

	// ProvisioningState: Provisioning state of the App Service Plan.
	ProvisioningState *Serverfarm_Properties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// Reserved: If Linux app service plan <code>true</code>, <code>false</code> otherwise.
	Reserved *bool `json:"reserved,omitempty"`

	// ResourceGroup: Resource group of the App Service plan.
	ResourceGroup *string `json:"resourceGroup,omitempty"`

	// SpotExpirationTime: The time when the server farm expires. Valid only if it is a spot server farm.
	SpotExpirationTime *string `json:"spotExpirationTime,omitempty"`

	// Status: App Service plan status.
	Status *Serverfarm_Properties_Status_STATUS `json:"status,omitempty"`

	// Subscription: App Service plan subscription.
	Subscription *string `json:"subscription,omitempty"`

	// TargetWorkerCount: Scaling worker count.
	TargetWorkerCount *int `json:"targetWorkerCount,omitempty"`

	// TargetWorkerSizeId: Scaling worker size ID.
	TargetWorkerSizeId *int `json:"targetWorkerSizeId,omitempty"`

	// WorkerTierName: Target worker tier assigned to the App Service plan.
	WorkerTierName *string `json:"workerTierName,omitempty"`

	// ZoneRedundant: If <code>true</code>, this App Service Plan will perform availability zone balancing.
	// If <code>false</code>, this App Service Plan will not perform availability zone balancing.
	ZoneRedundant *bool `json:"zoneRedundant,omitempty"`
}

// Description of a SKU for a scalable resource.
type SkuDescription_STATUS_ARM struct {
	// Capabilities: Capabilities of the SKU, e.g., is traffic manager enabled?
	Capabilities []Capability_STATUS_ARM `json:"capabilities,omitempty"`

	// Capacity: Current number of instances assigned to the resource.
	Capacity *int `json:"capacity,omitempty"`

	// Family: Family code of the resource SKU.
	Family *string `json:"family,omitempty"`

	// Locations: Locations of the SKU.
	Locations []string `json:"locations,omitempty"`

	// Name: Name of the resource SKU.
	Name *string `json:"name,omitempty"`

	// Size: Size specifier of the resource SKU.
	Size *string `json:"size,omitempty"`

	// SkuCapacity: Min, max, and default scale values of the SKU.
	SkuCapacity *SkuCapacity_STATUS_ARM `json:"skuCapacity,omitempty"`

	// Tier: Service tier of the resource SKU.
	Tier *string `json:"tier,omitempty"`
}

// Describes the capabilities/features allowed for a specific SKU.
type Capability_STATUS_ARM struct {
	// Name: Name of the SKU capability.
	Name *string `json:"name,omitempty"`

	// Reason: Reason of the SKU capability.
	Reason *string `json:"reason,omitempty"`

	// Value: Value of the SKU capability.
	Value *string `json:"value,omitempty"`
}

// Specification for an App Service Environment to use for this resource.
type HostingEnvironmentProfile_STATUS_ARM struct {
	// Id: Resource ID of the App Service Environment.
	Id *string `json:"id,omitempty"`

	// Name: Name of the App Service Environment.
	Name *string `json:"name,omitempty"`

	// Type: Resource type of the App Service Environment.
	Type *string `json:"type,omitempty"`
}

// Specification for a Kubernetes Environment to use for this resource.
type KubeEnvironmentProfile_STATUS_ARM struct {
	// Id: Resource ID of the Kubernetes Environment.
	Id *string `json:"id,omitempty"`

	// Name: Name of the Kubernetes Environment.
	Name *string `json:"name,omitempty"`

	// Type: Resource type of the Kubernetes Environment.
	Type *string `json:"type,omitempty"`
}

// Description of the App Service plan scale options.
type SkuCapacity_STATUS_ARM struct {
	// Default: Default number of workers for this App Service plan SKU.
	Default *int `json:"default,omitempty"`

	// ElasticMaximum: Maximum number of Elastic workers for this App Service plan SKU.
	ElasticMaximum *int `json:"elasticMaximum,omitempty"`

	// Maximum: Maximum number of workers for this App Service plan SKU.
	Maximum *int `json:"maximum,omitempty"`

	// Minimum: Minimum number of workers for this App Service plan SKU.
	Minimum *int `json:"minimum,omitempty"`

	// ScaleType: Available scale configurations for an App Service plan.
	ScaleType *string `json:"scaleType,omitempty"`
}
