// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"fmt"
	v20230401s "github.com/Azure/azure-service-operator/v2/api/cache/v1api20230401/storage"
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
// Storage version of v1api20201201.RedisLinkedServer
// Generator information:
// - Generated from: /redis/resource-manager/Microsoft.Cache/stable/2020-12-01/redis.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/linkedServers/{linkedServerName}
type RedisLinkedServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Redis_LinkedServer_Spec   `json:"spec,omitempty"`
	Status            Redis_LinkedServer_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RedisLinkedServer{}

// GetConditions returns the conditions of the resource
func (server *RedisLinkedServer) GetConditions() conditions.Conditions {
	return server.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (server *RedisLinkedServer) SetConditions(conditions conditions.Conditions) {
	server.Status.Conditions = conditions
}

var _ conversion.Convertible = &RedisLinkedServer{}

// ConvertFrom populates our RedisLinkedServer from the provided hub RedisLinkedServer
func (server *RedisLinkedServer) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20230401s.RedisLinkedServer)
	if !ok {
		return fmt.Errorf("expected cache/v1api20230401/storage/RedisLinkedServer but received %T instead", hub)
	}

	return server.AssignProperties_From_RedisLinkedServer(source)
}

// ConvertTo populates the provided hub RedisLinkedServer from our RedisLinkedServer
func (server *RedisLinkedServer) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20230401s.RedisLinkedServer)
	if !ok {
		return fmt.Errorf("expected cache/v1api20230401/storage/RedisLinkedServer but received %T instead", hub)
	}

	return server.AssignProperties_To_RedisLinkedServer(destination)
}

var _ genruntime.KubernetesResource = &RedisLinkedServer{}

// AzureName returns the Azure name of the resource
func (server *RedisLinkedServer) AzureName() string {
	return server.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (server RedisLinkedServer) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (server *RedisLinkedServer) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (server *RedisLinkedServer) GetSpec() genruntime.ConvertibleSpec {
	return &server.Spec
}

// GetStatus returns the status of this resource
func (server *RedisLinkedServer) GetStatus() genruntime.ConvertibleStatus {
	return &server.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (server *RedisLinkedServer) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/linkedServers"
func (server *RedisLinkedServer) GetType() string {
	return "Microsoft.Cache/redis/linkedServers"
}

// NewEmptyStatus returns a new empty (blank) status
func (server *RedisLinkedServer) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Redis_LinkedServer_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (server *RedisLinkedServer) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(server.Spec)
	return server.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (server *RedisLinkedServer) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Redis_LinkedServer_STATUS); ok {
		server.Status = *st
		return nil
	}

	// Convert status to required version
	var st Redis_LinkedServer_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	server.Status = st
	return nil
}

// AssignProperties_From_RedisLinkedServer populates our RedisLinkedServer from the provided source RedisLinkedServer
func (server *RedisLinkedServer) AssignProperties_From_RedisLinkedServer(source *v20230401s.RedisLinkedServer) error {

	// ObjectMeta
	server.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Redis_LinkedServer_Spec
	err := spec.AssignProperties_From_Redis_LinkedServer_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Redis_LinkedServer_Spec() to populate field Spec")
	}
	server.Spec = spec

	// Status
	var status Redis_LinkedServer_STATUS
	err = status.AssignProperties_From_Redis_LinkedServer_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Redis_LinkedServer_STATUS() to populate field Status")
	}
	server.Status = status

	// Invoke the augmentConversionForRedisLinkedServer interface (if implemented) to customize the conversion
	var serverAsAny any = server
	if augmentedServer, ok := serverAsAny.(augmentConversionForRedisLinkedServer); ok {
		err := augmentedServer.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_RedisLinkedServer populates the provided destination RedisLinkedServer from our RedisLinkedServer
func (server *RedisLinkedServer) AssignProperties_To_RedisLinkedServer(destination *v20230401s.RedisLinkedServer) error {

	// ObjectMeta
	destination.ObjectMeta = *server.ObjectMeta.DeepCopy()

	// Spec
	var spec v20230401s.Redis_LinkedServer_Spec
	err := server.Spec.AssignProperties_To_Redis_LinkedServer_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Redis_LinkedServer_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20230401s.Redis_LinkedServer_STATUS
	err = server.Status.AssignProperties_To_Redis_LinkedServer_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Redis_LinkedServer_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForRedisLinkedServer interface (if implemented) to customize the conversion
	var serverAsAny any = server
	if augmentedServer, ok := serverAsAny.(augmentConversionForRedisLinkedServer); ok {
		err := augmentedServer.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (server *RedisLinkedServer) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: server.Spec.OriginalVersion,
		Kind:    "RedisLinkedServer",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20201201.RedisLinkedServer
// Generator information:
// - Generated from: /redis/resource-manager/Microsoft.Cache/stable/2020-12-01/redis.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/linkedServers/{linkedServerName}
type RedisLinkedServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisLinkedServer `json:"items"`
}

type augmentConversionForRedisLinkedServer interface {
	AssignPropertiesFrom(src *v20230401s.RedisLinkedServer) error
	AssignPropertiesTo(dst *v20230401s.RedisLinkedServer) error
}

// Storage version of v1api20201201.Redis_LinkedServer_Spec
type Redis_LinkedServer_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                string  `json:"azureName,omitempty"`
	LinkedRedisCacheLocation *string `json:"linkedRedisCacheLocation,omitempty"`

	// +kubebuilder:validation:Required
	// LinkedRedisCacheReference: Fully qualified resourceId of the linked redis cache.
	LinkedRedisCacheReference *genruntime.ResourceReference `armReference:"LinkedRedisCacheId" json:"linkedRedisCacheReference,omitempty"`
	OriginalVersion           string                        `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a cache.azure.com/Redis resource
	Owner       *genruntime.KnownResourceReference `group:"cache.azure.com" json:"owner,omitempty" kind:"Redis"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	ServerRole  *string                            `json:"serverRole,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Redis_LinkedServer_Spec{}

// ConvertSpecFrom populates our Redis_LinkedServer_Spec from the provided source
func (server *Redis_LinkedServer_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20230401s.Redis_LinkedServer_Spec)
	if ok {
		// Populate our instance from source
		return server.AssignProperties_From_Redis_LinkedServer_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20230401s.Redis_LinkedServer_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = server.AssignProperties_From_Redis_LinkedServer_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Redis_LinkedServer_Spec
func (server *Redis_LinkedServer_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20230401s.Redis_LinkedServer_Spec)
	if ok {
		// Populate destination from our instance
		return server.AssignProperties_To_Redis_LinkedServer_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20230401s.Redis_LinkedServer_Spec{}
	err := server.AssignProperties_To_Redis_LinkedServer_Spec(dst)
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

// AssignProperties_From_Redis_LinkedServer_Spec populates our Redis_LinkedServer_Spec from the provided source Redis_LinkedServer_Spec
func (server *Redis_LinkedServer_Spec) AssignProperties_From_Redis_LinkedServer_Spec(source *v20230401s.Redis_LinkedServer_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	server.AzureName = source.AzureName

	// LinkedRedisCacheLocation
	server.LinkedRedisCacheLocation = genruntime.ClonePointerToString(source.LinkedRedisCacheLocation)

	// LinkedRedisCacheReference
	if source.LinkedRedisCacheReference != nil {
		linkedRedisCacheReference := source.LinkedRedisCacheReference.Copy()
		server.LinkedRedisCacheReference = &linkedRedisCacheReference
	} else {
		server.LinkedRedisCacheReference = nil
	}

	// OriginalVersion
	server.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		server.Owner = &owner
	} else {
		server.Owner = nil
	}

	// ServerRole
	server.ServerRole = genruntime.ClonePointerToString(source.ServerRole)

	// Update the property bag
	if len(propertyBag) > 0 {
		server.PropertyBag = propertyBag
	} else {
		server.PropertyBag = nil
	}

	// Invoke the augmentConversionForRedis_LinkedServer_Spec interface (if implemented) to customize the conversion
	var serverAsAny any = server
	if augmentedServer, ok := serverAsAny.(augmentConversionForRedis_LinkedServer_Spec); ok {
		err := augmentedServer.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Redis_LinkedServer_Spec populates the provided destination Redis_LinkedServer_Spec from our Redis_LinkedServer_Spec
func (server *Redis_LinkedServer_Spec) AssignProperties_To_Redis_LinkedServer_Spec(destination *v20230401s.Redis_LinkedServer_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(server.PropertyBag)

	// AzureName
	destination.AzureName = server.AzureName

	// LinkedRedisCacheLocation
	destination.LinkedRedisCacheLocation = genruntime.ClonePointerToString(server.LinkedRedisCacheLocation)

	// LinkedRedisCacheReference
	if server.LinkedRedisCacheReference != nil {
		linkedRedisCacheReference := server.LinkedRedisCacheReference.Copy()
		destination.LinkedRedisCacheReference = &linkedRedisCacheReference
	} else {
		destination.LinkedRedisCacheReference = nil
	}

	// OriginalVersion
	destination.OriginalVersion = server.OriginalVersion

	// Owner
	if server.Owner != nil {
		owner := server.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// ServerRole
	destination.ServerRole = genruntime.ClonePointerToString(server.ServerRole)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForRedis_LinkedServer_Spec interface (if implemented) to customize the conversion
	var serverAsAny any = server
	if augmentedServer, ok := serverAsAny.(augmentConversionForRedis_LinkedServer_Spec); ok {
		err := augmentedServer.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20201201.Redis_LinkedServer_STATUS
type Redis_LinkedServer_STATUS struct {
	Conditions               []conditions.Condition `json:"conditions,omitempty"`
	Id                       *string                `json:"id,omitempty"`
	LinkedRedisCacheId       *string                `json:"linkedRedisCacheId,omitempty"`
	LinkedRedisCacheLocation *string                `json:"linkedRedisCacheLocation,omitempty"`
	Name                     *string                `json:"name,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState        *string                `json:"provisioningState,omitempty"`
	ServerRole               *string                `json:"serverRole,omitempty"`
	Type                     *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Redis_LinkedServer_STATUS{}

// ConvertStatusFrom populates our Redis_LinkedServer_STATUS from the provided source
func (server *Redis_LinkedServer_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20230401s.Redis_LinkedServer_STATUS)
	if ok {
		// Populate our instance from source
		return server.AssignProperties_From_Redis_LinkedServer_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20230401s.Redis_LinkedServer_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = server.AssignProperties_From_Redis_LinkedServer_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Redis_LinkedServer_STATUS
func (server *Redis_LinkedServer_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20230401s.Redis_LinkedServer_STATUS)
	if ok {
		// Populate destination from our instance
		return server.AssignProperties_To_Redis_LinkedServer_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20230401s.Redis_LinkedServer_STATUS{}
	err := server.AssignProperties_To_Redis_LinkedServer_STATUS(dst)
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

// AssignProperties_From_Redis_LinkedServer_STATUS populates our Redis_LinkedServer_STATUS from the provided source Redis_LinkedServer_STATUS
func (server *Redis_LinkedServer_STATUS) AssignProperties_From_Redis_LinkedServer_STATUS(source *v20230401s.Redis_LinkedServer_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	server.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// GeoReplicatedPrimaryHostName
	if source.GeoReplicatedPrimaryHostName != nil {
		propertyBag.Add("GeoReplicatedPrimaryHostName", *source.GeoReplicatedPrimaryHostName)
	} else {
		propertyBag.Remove("GeoReplicatedPrimaryHostName")
	}

	// Id
	server.Id = genruntime.ClonePointerToString(source.Id)

	// LinkedRedisCacheId
	server.LinkedRedisCacheId = genruntime.ClonePointerToString(source.LinkedRedisCacheId)

	// LinkedRedisCacheLocation
	server.LinkedRedisCacheLocation = genruntime.ClonePointerToString(source.LinkedRedisCacheLocation)

	// Name
	server.Name = genruntime.ClonePointerToString(source.Name)

	// PrimaryHostName
	if source.PrimaryHostName != nil {
		propertyBag.Add("PrimaryHostName", *source.PrimaryHostName)
	} else {
		propertyBag.Remove("PrimaryHostName")
	}

	// ProvisioningState
	server.ProvisioningState = genruntime.ClonePointerToString(source.ProvisioningState)

	// ServerRole
	server.ServerRole = genruntime.ClonePointerToString(source.ServerRole)

	// Type
	server.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		server.PropertyBag = propertyBag
	} else {
		server.PropertyBag = nil
	}

	// Invoke the augmentConversionForRedis_LinkedServer_STATUS interface (if implemented) to customize the conversion
	var serverAsAny any = server
	if augmentedServer, ok := serverAsAny.(augmentConversionForRedis_LinkedServer_STATUS); ok {
		err := augmentedServer.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Redis_LinkedServer_STATUS populates the provided destination Redis_LinkedServer_STATUS from our Redis_LinkedServer_STATUS
func (server *Redis_LinkedServer_STATUS) AssignProperties_To_Redis_LinkedServer_STATUS(destination *v20230401s.Redis_LinkedServer_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(server.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(server.Conditions)

	// GeoReplicatedPrimaryHostName
	if propertyBag.Contains("GeoReplicatedPrimaryHostName") {
		var geoReplicatedPrimaryHostName string
		err := propertyBag.Pull("GeoReplicatedPrimaryHostName", &geoReplicatedPrimaryHostName)
		if err != nil {
			return errors.Wrap(err, "pulling 'GeoReplicatedPrimaryHostName' from propertyBag")
		}

		destination.GeoReplicatedPrimaryHostName = &geoReplicatedPrimaryHostName
	} else {
		destination.GeoReplicatedPrimaryHostName = nil
	}

	// Id
	destination.Id = genruntime.ClonePointerToString(server.Id)

	// LinkedRedisCacheId
	destination.LinkedRedisCacheId = genruntime.ClonePointerToString(server.LinkedRedisCacheId)

	// LinkedRedisCacheLocation
	destination.LinkedRedisCacheLocation = genruntime.ClonePointerToString(server.LinkedRedisCacheLocation)

	// Name
	destination.Name = genruntime.ClonePointerToString(server.Name)

	// PrimaryHostName
	if propertyBag.Contains("PrimaryHostName") {
		var primaryHostName string
		err := propertyBag.Pull("PrimaryHostName", &primaryHostName)
		if err != nil {
			return errors.Wrap(err, "pulling 'PrimaryHostName' from propertyBag")
		}

		destination.PrimaryHostName = &primaryHostName
	} else {
		destination.PrimaryHostName = nil
	}

	// ProvisioningState
	destination.ProvisioningState = genruntime.ClonePointerToString(server.ProvisioningState)

	// ServerRole
	destination.ServerRole = genruntime.ClonePointerToString(server.ServerRole)

	// Type
	destination.Type = genruntime.ClonePointerToString(server.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForRedis_LinkedServer_STATUS interface (if implemented) to customize the conversion
	var serverAsAny any = server
	if augmentedServer, ok := serverAsAny.(augmentConversionForRedis_LinkedServer_STATUS); ok {
		err := augmentedServer.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForRedis_LinkedServer_Spec interface {
	AssignPropertiesFrom(src *v20230401s.Redis_LinkedServer_Spec) error
	AssignPropertiesTo(dst *v20230401s.Redis_LinkedServer_Spec) error
}

type augmentConversionForRedis_LinkedServer_STATUS interface {
	AssignPropertiesFrom(src *v20230401s.Redis_LinkedServer_STATUS) error
	AssignPropertiesTo(dst *v20230401s.Redis_LinkedServer_STATUS) error
}

func init() {
	SchemeBuilder.Register(&RedisLinkedServer{}, &RedisLinkedServerList{})
}
