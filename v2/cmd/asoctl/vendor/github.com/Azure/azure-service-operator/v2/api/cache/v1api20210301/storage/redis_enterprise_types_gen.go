// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"fmt"
	v20230401s "github.com/Azure/azure-service-operator/v2/api/cache/v1api20230401/storage"
	v20230701s "github.com/Azure/azure-service-operator/v2/api/cache/v1api20230701/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20210301.RedisEnterprise
// Generator information:
// - Generated from: /redisenterprise/resource-manager/Microsoft.Cache/stable/2021-03-01/redisenterprise.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}
type RedisEnterprise struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisEnterprise_Spec   `json:"spec,omitempty"`
	Status            RedisEnterprise_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RedisEnterprise{}

// GetConditions returns the conditions of the resource
func (enterprise *RedisEnterprise) GetConditions() conditions.Conditions {
	return enterprise.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (enterprise *RedisEnterprise) SetConditions(conditions conditions.Conditions) {
	enterprise.Status.Conditions = conditions
}

var _ conversion.Convertible = &RedisEnterprise{}

// ConvertFrom populates our RedisEnterprise from the provided hub RedisEnterprise
func (enterprise *RedisEnterprise) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20230701s.RedisEnterprise)
	if !ok {
		return fmt.Errorf("expected cache/v1api20230701/storage/RedisEnterprise but received %T instead", hub)
	}

	return enterprise.AssignProperties_From_RedisEnterprise(source)
}

// ConvertTo populates the provided hub RedisEnterprise from our RedisEnterprise
func (enterprise *RedisEnterprise) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20230701s.RedisEnterprise)
	if !ok {
		return fmt.Errorf("expected cache/v1api20230701/storage/RedisEnterprise but received %T instead", hub)
	}

	return enterprise.AssignProperties_To_RedisEnterprise(destination)
}

var _ genruntime.KubernetesResource = &RedisEnterprise{}

// AzureName returns the Azure name of the resource
func (enterprise *RedisEnterprise) AzureName() string {
	return enterprise.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-03-01"
func (enterprise RedisEnterprise) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (enterprise *RedisEnterprise) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (enterprise *RedisEnterprise) GetSpec() genruntime.ConvertibleSpec {
	return &enterprise.Spec
}

// GetStatus returns the status of this resource
func (enterprise *RedisEnterprise) GetStatus() genruntime.ConvertibleStatus {
	return &enterprise.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (enterprise *RedisEnterprise) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redisEnterprise"
func (enterprise *RedisEnterprise) GetType() string {
	return "Microsoft.Cache/redisEnterprise"
}

// NewEmptyStatus returns a new empty (blank) status
func (enterprise *RedisEnterprise) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RedisEnterprise_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (enterprise *RedisEnterprise) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(enterprise.Spec)
	return enterprise.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (enterprise *RedisEnterprise) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RedisEnterprise_STATUS); ok {
		enterprise.Status = *st
		return nil
	}

	// Convert status to required version
	var st RedisEnterprise_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	enterprise.Status = st
	return nil
}

// AssignProperties_From_RedisEnterprise populates our RedisEnterprise from the provided source RedisEnterprise
func (enterprise *RedisEnterprise) AssignProperties_From_RedisEnterprise(source *v20230701s.RedisEnterprise) error {

	// ObjectMeta
	enterprise.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec RedisEnterprise_Spec
	err := spec.AssignProperties_From_RedisEnterprise_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_RedisEnterprise_Spec() to populate field Spec")
	}
	enterprise.Spec = spec

	// Status
	var status RedisEnterprise_STATUS
	err = status.AssignProperties_From_RedisEnterprise_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_RedisEnterprise_STATUS() to populate field Status")
	}
	enterprise.Status = status

	// Invoke the augmentConversionForRedisEnterprise interface (if implemented) to customize the conversion
	var enterpriseAsAny any = enterprise
	if augmentedEnterprise, ok := enterpriseAsAny.(augmentConversionForRedisEnterprise); ok {
		err := augmentedEnterprise.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_RedisEnterprise populates the provided destination RedisEnterprise from our RedisEnterprise
func (enterprise *RedisEnterprise) AssignProperties_To_RedisEnterprise(destination *v20230701s.RedisEnterprise) error {

	// ObjectMeta
	destination.ObjectMeta = *enterprise.ObjectMeta.DeepCopy()

	// Spec
	var spec v20230701s.RedisEnterprise_Spec
	err := enterprise.Spec.AssignProperties_To_RedisEnterprise_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_RedisEnterprise_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20230701s.RedisEnterprise_STATUS
	err = enterprise.Status.AssignProperties_To_RedisEnterprise_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_RedisEnterprise_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForRedisEnterprise interface (if implemented) to customize the conversion
	var enterpriseAsAny any = enterprise
	if augmentedEnterprise, ok := enterpriseAsAny.(augmentConversionForRedisEnterprise); ok {
		err := augmentedEnterprise.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (enterprise *RedisEnterprise) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: enterprise.Spec.OriginalVersion,
		Kind:    "RedisEnterprise",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20210301.RedisEnterprise
// Generator information:
// - Generated from: /redisenterprise/resource-manager/Microsoft.Cache/stable/2021-03-01/redisenterprise.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}
type RedisEnterpriseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisEnterprise `json:"items"`
}

// Storage version of v1api20210301.APIVersion
// +kubebuilder:validation:Enum={"2021-03-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2021-03-01")

type augmentConversionForRedisEnterprise interface {
	AssignPropertiesFrom(src *v20230701s.RedisEnterprise) error
	AssignPropertiesTo(dst *v20230701s.RedisEnterprise) error
}

// Storage version of v1api20210301.RedisEnterprise_Spec
type RedisEnterprise_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName         string  `json:"azureName,omitempty"`
	Location          *string `json:"location,omitempty"`
	MinimumTlsVersion *string `json:"minimumTlsVersion,omitempty"`
	OriginalVersion   string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner       *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Sku         *Sku                               `json:"sku,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
	Zones       []string                           `json:"zones,omitempty"`
}

var _ genruntime.ConvertibleSpec = &RedisEnterprise_Spec{}

// ConvertSpecFrom populates our RedisEnterprise_Spec from the provided source
func (enterprise *RedisEnterprise_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20230701s.RedisEnterprise_Spec)
	if ok {
		// Populate our instance from source
		return enterprise.AssignProperties_From_RedisEnterprise_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20230701s.RedisEnterprise_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = enterprise.AssignProperties_From_RedisEnterprise_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RedisEnterprise_Spec
func (enterprise *RedisEnterprise_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20230701s.RedisEnterprise_Spec)
	if ok {
		// Populate destination from our instance
		return enterprise.AssignProperties_To_RedisEnterprise_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20230701s.RedisEnterprise_Spec{}
	err := enterprise.AssignProperties_To_RedisEnterprise_Spec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_RedisEnterprise_Spec populates our RedisEnterprise_Spec from the provided source RedisEnterprise_Spec
func (enterprise *RedisEnterprise_Spec) AssignProperties_From_RedisEnterprise_Spec(source *v20230701s.RedisEnterprise_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	enterprise.AzureName = source.AzureName

	// Location
	enterprise.Location = genruntime.ClonePointerToString(source.Location)

	// MinimumTlsVersion
	enterprise.MinimumTlsVersion = genruntime.ClonePointerToString(source.MinimumTlsVersion)

	// OriginalVersion
	enterprise.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		enterprise.Owner = &owner
	} else {
		enterprise.Owner = nil
	}

	// Sku
	if source.Sku != nil {
		var skuStash v20230401s.Sku
		err := skuStash.AssignProperties_From_Sku(source.Sku)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_Sku() to populate field SkuStash from Sku")
		}
		var sku Sku
		err = sku.AssignProperties_From_Sku(&skuStash)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_Sku() to populate field Sku from SkuStash")
		}
		enterprise.Sku = &sku
	} else {
		enterprise.Sku = nil
	}

	// Tags
	enterprise.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Zones
	enterprise.Zones = genruntime.CloneSliceOfString(source.Zones)

	// Update the property bag
	if len(propertyBag) > 0 {
		enterprise.PropertyBag = propertyBag
	} else {
		enterprise.PropertyBag = nil
	}

	// Invoke the augmentConversionForRedisEnterprise_Spec interface (if implemented) to customize the conversion
	var enterpriseAsAny any = enterprise
	if augmentedEnterprise, ok := enterpriseAsAny.(augmentConversionForRedisEnterprise_Spec); ok {
		err := augmentedEnterprise.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_RedisEnterprise_Spec populates the provided destination RedisEnterprise_Spec from our RedisEnterprise_Spec
func (enterprise *RedisEnterprise_Spec) AssignProperties_To_RedisEnterprise_Spec(destination *v20230701s.RedisEnterprise_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(enterprise.PropertyBag)

	// AzureName
	destination.AzureName = enterprise.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(enterprise.Location)

	// MinimumTlsVersion
	destination.MinimumTlsVersion = genruntime.ClonePointerToString(enterprise.MinimumTlsVersion)

	// OriginalVersion
	destination.OriginalVersion = enterprise.OriginalVersion

	// Owner
	if enterprise.Owner != nil {
		owner := enterprise.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Sku
	if enterprise.Sku != nil {
		var skuStash v20230401s.Sku
		err := enterprise.Sku.AssignProperties_To_Sku(&skuStash)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_Sku() to populate field SkuStash from Sku")
		}
		var sku v20230701s.Sku
		err = skuStash.AssignProperties_To_Sku(&sku)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_Sku() to populate field Sku from SkuStash")
		}
		destination.Sku = &sku
	} else {
		destination.Sku = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(enterprise.Tags)

	// Zones
	destination.Zones = genruntime.CloneSliceOfString(enterprise.Zones)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForRedisEnterprise_Spec interface (if implemented) to customize the conversion
	var enterpriseAsAny any = enterprise
	if augmentedEnterprise, ok := enterpriseAsAny.(augmentConversionForRedisEnterprise_Spec); ok {
		err := augmentedEnterprise.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20210301.RedisEnterprise_STATUS
type RedisEnterprise_STATUS struct {
	Conditions                 []conditions.Condition             `json:"conditions,omitempty"`
	HostName                   *string                            `json:"hostName,omitempty"`
	Id                         *string                            `json:"id,omitempty"`
	Location                   *string                            `json:"location,omitempty"`
	MinimumTlsVersion          *string                            `json:"minimumTlsVersion,omitempty"`
	Name                       *string                            `json:"name,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS `json:"privateEndpointConnections,omitempty"`
	PropertyBag                genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	ProvisioningState          *string                            `json:"provisioningState,omitempty"`
	RedisVersion               *string                            `json:"redisVersion,omitempty"`
	ResourceState              *string                            `json:"resourceState,omitempty"`
	Sku                        *Sku_STATUS                        `json:"sku,omitempty"`
	Tags                       map[string]string                  `json:"tags,omitempty"`
	Type                       *string                            `json:"type,omitempty"`
	Zones                      []string                           `json:"zones,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RedisEnterprise_STATUS{}

// ConvertStatusFrom populates our RedisEnterprise_STATUS from the provided source
func (enterprise *RedisEnterprise_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20230701s.RedisEnterprise_STATUS)
	if ok {
		// Populate our instance from source
		return enterprise.AssignProperties_From_RedisEnterprise_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20230701s.RedisEnterprise_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = enterprise.AssignProperties_From_RedisEnterprise_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our RedisEnterprise_STATUS
func (enterprise *RedisEnterprise_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20230701s.RedisEnterprise_STATUS)
	if ok {
		// Populate destination from our instance
		return enterprise.AssignProperties_To_RedisEnterprise_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20230701s.RedisEnterprise_STATUS{}
	err := enterprise.AssignProperties_To_RedisEnterprise_STATUS(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

// AssignProperties_From_RedisEnterprise_STATUS populates our RedisEnterprise_STATUS from the provided source RedisEnterprise_STATUS
func (enterprise *RedisEnterprise_STATUS) AssignProperties_From_RedisEnterprise_STATUS(source *v20230701s.RedisEnterprise_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	enterprise.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// HostName
	enterprise.HostName = genruntime.ClonePointerToString(source.HostName)

	// Id
	enterprise.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	enterprise.Location = genruntime.ClonePointerToString(source.Location)

	// MinimumTlsVersion
	enterprise.MinimumTlsVersion = genruntime.ClonePointerToString(source.MinimumTlsVersion)

	// Name
	enterprise.Name = genruntime.ClonePointerToString(source.Name)

	// PrivateEndpointConnections
	if source.PrivateEndpointConnections != nil {
		privateEndpointConnectionList := make([]PrivateEndpointConnection_STATUS, len(source.PrivateEndpointConnections))
		for privateEndpointConnectionIndex, privateEndpointConnectionItem := range source.PrivateEndpointConnections {
			// Shadow the loop variable to avoid aliasing
			privateEndpointConnectionItem := privateEndpointConnectionItem
			var privateEndpointConnectionSTATUSStash v20230401s.PrivateEndpointConnection_STATUS
			err := privateEndpointConnectionSTATUSStash.AssignProperties_From_PrivateEndpointConnection_STATUS(&privateEndpointConnectionItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignProperties_From_PrivateEndpointConnection_STATUS() to populate field PrivateEndpointConnection_STATUSStash from PrivateEndpointConnections")
			}
			var privateEndpointConnection PrivateEndpointConnection_STATUS
			err = privateEndpointConnection.AssignProperties_From_PrivateEndpointConnection_STATUS(&privateEndpointConnectionSTATUSStash)
			if err != nil {
				return errors.Wrap(err, "calling AssignProperties_From_PrivateEndpointConnection_STATUS() to populate field PrivateEndpointConnections from PrivateEndpointConnection_STATUSStash")
			}
			privateEndpointConnectionList[privateEndpointConnectionIndex] = privateEndpointConnection
		}
		enterprise.PrivateEndpointConnections = privateEndpointConnectionList
	} else {
		enterprise.PrivateEndpointConnections = nil
	}

	// ProvisioningState
	enterprise.ProvisioningState = genruntime.ClonePointerToString(source.ProvisioningState)

	// RedisVersion
	enterprise.RedisVersion = genruntime.ClonePointerToString(source.RedisVersion)

	// ResourceState
	enterprise.ResourceState = genruntime.ClonePointerToString(source.ResourceState)

	// Sku
	if source.Sku != nil {
		var skuSTATUSStash v20230401s.Sku_STATUS
		err := skuSTATUSStash.AssignProperties_From_Sku_STATUS(source.Sku)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_Sku_STATUS() to populate field Sku_STATUSStash from Sku")
		}
		var sku Sku_STATUS
		err = sku.AssignProperties_From_Sku_STATUS(&skuSTATUSStash)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_Sku_STATUS() to populate field Sku from Sku_STATUSStash")
		}
		enterprise.Sku = &sku
	} else {
		enterprise.Sku = nil
	}

	// Tags
	enterprise.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	enterprise.Type = genruntime.ClonePointerToString(source.Type)

	// Zones
	enterprise.Zones = genruntime.CloneSliceOfString(source.Zones)

	// Update the property bag
	if len(propertyBag) > 0 {
		enterprise.PropertyBag = propertyBag
	} else {
		enterprise.PropertyBag = nil
	}

	// Invoke the augmentConversionForRedisEnterprise_STATUS interface (if implemented) to customize the conversion
	var enterpriseAsAny any = enterprise
	if augmentedEnterprise, ok := enterpriseAsAny.(augmentConversionForRedisEnterprise_STATUS); ok {
		err := augmentedEnterprise.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_RedisEnterprise_STATUS populates the provided destination RedisEnterprise_STATUS from our RedisEnterprise_STATUS
func (enterprise *RedisEnterprise_STATUS) AssignProperties_To_RedisEnterprise_STATUS(destination *v20230701s.RedisEnterprise_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(enterprise.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(enterprise.Conditions)

	// HostName
	destination.HostName = genruntime.ClonePointerToString(enterprise.HostName)

	// Id
	destination.Id = genruntime.ClonePointerToString(enterprise.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(enterprise.Location)

	// MinimumTlsVersion
	destination.MinimumTlsVersion = genruntime.ClonePointerToString(enterprise.MinimumTlsVersion)

	// Name
	destination.Name = genruntime.ClonePointerToString(enterprise.Name)

	// PrivateEndpointConnections
	if enterprise.PrivateEndpointConnections != nil {
		privateEndpointConnectionList := make([]v20230701s.PrivateEndpointConnection_STATUS, len(enterprise.PrivateEndpointConnections))
		for privateEndpointConnectionIndex, privateEndpointConnectionItem := range enterprise.PrivateEndpointConnections {
			// Shadow the loop variable to avoid aliasing
			privateEndpointConnectionItem := privateEndpointConnectionItem
			var privateEndpointConnectionSTATUSStash v20230401s.PrivateEndpointConnection_STATUS
			err := privateEndpointConnectionItem.AssignProperties_To_PrivateEndpointConnection_STATUS(&privateEndpointConnectionSTATUSStash)
			if err != nil {
				return errors.Wrap(err, "calling AssignProperties_To_PrivateEndpointConnection_STATUS() to populate field PrivateEndpointConnection_STATUSStash from PrivateEndpointConnections")
			}
			var privateEndpointConnection v20230701s.PrivateEndpointConnection_STATUS
			err = privateEndpointConnectionSTATUSStash.AssignProperties_To_PrivateEndpointConnection_STATUS(&privateEndpointConnection)
			if err != nil {
				return errors.Wrap(err, "calling AssignProperties_To_PrivateEndpointConnection_STATUS() to populate field PrivateEndpointConnections from PrivateEndpointConnection_STATUSStash")
			}
			privateEndpointConnectionList[privateEndpointConnectionIndex] = privateEndpointConnection
		}
		destination.PrivateEndpointConnections = privateEndpointConnectionList
	} else {
		destination.PrivateEndpointConnections = nil
	}

	// ProvisioningState
	destination.ProvisioningState = genruntime.ClonePointerToString(enterprise.ProvisioningState)

	// RedisVersion
	destination.RedisVersion = genruntime.ClonePointerToString(enterprise.RedisVersion)

	// ResourceState
	destination.ResourceState = genruntime.ClonePointerToString(enterprise.ResourceState)

	// Sku
	if enterprise.Sku != nil {
		var skuSTATUSStash v20230401s.Sku_STATUS
		err := enterprise.Sku.AssignProperties_To_Sku_STATUS(&skuSTATUSStash)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_Sku_STATUS() to populate field Sku_STATUSStash from Sku")
		}
		var sku v20230701s.Sku_STATUS
		err = skuSTATUSStash.AssignProperties_To_Sku_STATUS(&sku)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_Sku_STATUS() to populate field Sku from Sku_STATUSStash")
		}
		destination.Sku = &sku
	} else {
		destination.Sku = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(enterprise.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(enterprise.Type)

	// Zones
	destination.Zones = genruntime.CloneSliceOfString(enterprise.Zones)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForRedisEnterprise_STATUS interface (if implemented) to customize the conversion
	var enterpriseAsAny any = enterprise
	if augmentedEnterprise, ok := enterpriseAsAny.(augmentConversionForRedisEnterprise_STATUS); ok {
		err := augmentedEnterprise.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForRedisEnterprise_Spec interface {
	AssignPropertiesFrom(src *v20230701s.RedisEnterprise_Spec) error
	AssignPropertiesTo(dst *v20230701s.RedisEnterprise_Spec) error
}

type augmentConversionForRedisEnterprise_STATUS interface {
	AssignPropertiesFrom(src *v20230701s.RedisEnterprise_STATUS) error
	AssignPropertiesTo(dst *v20230701s.RedisEnterprise_STATUS) error
}

// Storage version of v1api20210301.PrivateEndpointConnection_STATUS
// The Private Endpoint Connection resource.
type PrivateEndpointConnection_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_PrivateEndpointConnection_STATUS populates our PrivateEndpointConnection_STATUS from the provided source PrivateEndpointConnection_STATUS
func (connection *PrivateEndpointConnection_STATUS) AssignProperties_From_PrivateEndpointConnection_STATUS(source *v20230401s.PrivateEndpointConnection_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Id
	connection.Id = genruntime.ClonePointerToString(source.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		connection.PropertyBag = propertyBag
	} else {
		connection.PropertyBag = nil
	}

	// Invoke the augmentConversionForPrivateEndpointConnection_STATUS interface (if implemented) to customize the conversion
	var connectionAsAny any = connection
	if augmentedConnection, ok := connectionAsAny.(augmentConversionForPrivateEndpointConnection_STATUS); ok {
		err := augmentedConnection.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_PrivateEndpointConnection_STATUS populates the provided destination PrivateEndpointConnection_STATUS from our PrivateEndpointConnection_STATUS
func (connection *PrivateEndpointConnection_STATUS) AssignProperties_To_PrivateEndpointConnection_STATUS(destination *v20230401s.PrivateEndpointConnection_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(connection.PropertyBag)

	// Id
	destination.Id = genruntime.ClonePointerToString(connection.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForPrivateEndpointConnection_STATUS interface (if implemented) to customize the conversion
	var connectionAsAny any = connection
	if augmentedConnection, ok := connectionAsAny.(augmentConversionForPrivateEndpointConnection_STATUS); ok {
		err := augmentedConnection.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20210301.Sku
// SKU parameters supplied to the create RedisEnterprise operation.
type Sku struct {
	Capacity    *int                   `json:"capacity,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_Sku populates our Sku from the provided source Sku
func (sku *Sku) AssignProperties_From_Sku(source *v20230401s.Sku) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Capacity
	sku.Capacity = genruntime.ClonePointerToInt(source.Capacity)

	// Family
	if source.Family != nil {
		propertyBag.Add("Family", *source.Family)
	} else {
		propertyBag.Remove("Family")
	}

	// Name
	sku.Name = genruntime.ClonePointerToString(source.Name)

	// Update the property bag
	if len(propertyBag) > 0 {
		sku.PropertyBag = propertyBag
	} else {
		sku.PropertyBag = nil
	}

	// Invoke the augmentConversionForSku interface (if implemented) to customize the conversion
	var skuAsAny any = sku
	if augmentedSku, ok := skuAsAny.(augmentConversionForSku); ok {
		err := augmentedSku.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Sku populates the provided destination Sku from our Sku
func (sku *Sku) AssignProperties_To_Sku(destination *v20230401s.Sku) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(sku.PropertyBag)

	// Capacity
	destination.Capacity = genruntime.ClonePointerToInt(sku.Capacity)

	// Family
	if propertyBag.Contains("Family") {
		var family string
		err := propertyBag.Pull("Family", &family)
		if err != nil {
			return errors.Wrap(err, "pulling 'Family' from propertyBag")
		}

		destination.Family = &family
	} else {
		destination.Family = nil
	}

	// Name
	destination.Name = genruntime.ClonePointerToString(sku.Name)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForSku interface (if implemented) to customize the conversion
	var skuAsAny any = sku
	if augmentedSku, ok := skuAsAny.(augmentConversionForSku); ok {
		err := augmentedSku.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20210301.Sku_STATUS
// SKU parameters supplied to the create RedisEnterprise operation.
type Sku_STATUS struct {
	Capacity    *int                   `json:"capacity,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_Sku_STATUS populates our Sku_STATUS from the provided source Sku_STATUS
func (sku *Sku_STATUS) AssignProperties_From_Sku_STATUS(source *v20230401s.Sku_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Capacity
	sku.Capacity = genruntime.ClonePointerToInt(source.Capacity)

	// Family
	if source.Family != nil {
		propertyBag.Add("Family", *source.Family)
	} else {
		propertyBag.Remove("Family")
	}

	// Name
	sku.Name = genruntime.ClonePointerToString(source.Name)

	// Update the property bag
	if len(propertyBag) > 0 {
		sku.PropertyBag = propertyBag
	} else {
		sku.PropertyBag = nil
	}

	// Invoke the augmentConversionForSku_STATUS interface (if implemented) to customize the conversion
	var skuAsAny any = sku
	if augmentedSku, ok := skuAsAny.(augmentConversionForSku_STATUS); ok {
		err := augmentedSku.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Sku_STATUS populates the provided destination Sku_STATUS from our Sku_STATUS
func (sku *Sku_STATUS) AssignProperties_To_Sku_STATUS(destination *v20230401s.Sku_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(sku.PropertyBag)

	// Capacity
	destination.Capacity = genruntime.ClonePointerToInt(sku.Capacity)

	// Family
	if propertyBag.Contains("Family") {
		var family string
		err := propertyBag.Pull("Family", &family)
		if err != nil {
			return errors.Wrap(err, "pulling 'Family' from propertyBag")
		}

		destination.Family = &family
	} else {
		destination.Family = nil
	}

	// Name
	destination.Name = genruntime.ClonePointerToString(sku.Name)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForSku_STATUS interface (if implemented) to customize the conversion
	var skuAsAny any = sku
	if augmentedSku, ok := skuAsAny.(augmentConversionForSku_STATUS); ok {
		err := augmentedSku.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForPrivateEndpointConnection_STATUS interface {
	AssignPropertiesFrom(src *v20230401s.PrivateEndpointConnection_STATUS) error
	AssignPropertiesTo(dst *v20230401s.PrivateEndpointConnection_STATUS) error
}

type augmentConversionForSku interface {
	AssignPropertiesFrom(src *v20230401s.Sku) error
	AssignPropertiesTo(dst *v20230401s.Sku) error
}

type augmentConversionForSku_STATUS interface {
	AssignPropertiesFrom(src *v20230401s.Sku_STATUS) error
	AssignPropertiesTo(dst *v20230401s.Sku_STATUS) error
}

func init() {
	SchemeBuilder.Register(&RedisEnterprise{}, &RedisEnterpriseList{})
}
