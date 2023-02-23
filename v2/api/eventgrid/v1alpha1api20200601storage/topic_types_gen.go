// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601storage

import (
	"fmt"
	v20200601s "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1beta20200601storage"
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
// Storage version of v1alpha1api20200601.Topic
// Deprecated version of Topic. Use v1beta20200601.Topic instead
type Topic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Topic_Spec   `json:"spec,omitempty"`
	Status            Topic_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Topic{}

// GetConditions returns the conditions of the resource
func (topic *Topic) GetConditions() conditions.Conditions {
	return topic.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (topic *Topic) SetConditions(conditions conditions.Conditions) {
	topic.Status.Conditions = conditions
}

var _ conversion.Convertible = &Topic{}

// ConvertFrom populates our Topic from the provided hub Topic
func (topic *Topic) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20200601s.Topic)
	if !ok {
		return fmt.Errorf("expected eventgrid/v1beta20200601storage/Topic but received %T instead", hub)
	}

	return topic.AssignProperties_From_Topic(source)
}

// ConvertTo populates the provided hub Topic from our Topic
func (topic *Topic) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20200601s.Topic)
	if !ok {
		return fmt.Errorf("expected eventgrid/v1beta20200601storage/Topic but received %T instead", hub)
	}

	return topic.AssignProperties_To_Topic(destination)
}

var _ genruntime.KubernetesResource = &Topic{}

// AzureName returns the Azure name of the resource
func (topic *Topic) AzureName() string {
	return topic.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (topic Topic) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (topic *Topic) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (topic *Topic) GetSpec() genruntime.ConvertibleSpec {
	return &topic.Spec
}

// GetStatus returns the status of this resource
func (topic *Topic) GetStatus() genruntime.ConvertibleStatus {
	return &topic.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventGrid/topics"
func (topic *Topic) GetType() string {
	return "Microsoft.EventGrid/topics"
}

// NewEmptyStatus returns a new empty (blank) status
func (topic *Topic) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Topic_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (topic *Topic) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(topic.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  topic.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (topic *Topic) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Topic_STATUS); ok {
		topic.Status = *st
		return nil
	}

	// Convert status to required version
	var st Topic_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	topic.Status = st
	return nil
}

// AssignProperties_From_Topic populates our Topic from the provided source Topic
func (topic *Topic) AssignProperties_From_Topic(source *v20200601s.Topic) error {

	// ObjectMeta
	topic.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Topic_Spec
	err := spec.AssignProperties_From_Topic_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Topic_Spec() to populate field Spec")
	}
	topic.Spec = spec

	// Status
	var status Topic_STATUS
	err = status.AssignProperties_From_Topic_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Topic_STATUS() to populate field Status")
	}
	topic.Status = status

	// Invoke the augmentConversionForTopic interface (if implemented) to customize the conversion
	var topicAsAny any = topic
	if augmentedTopic, ok := topicAsAny.(augmentConversionForTopic); ok {
		err := augmentedTopic.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Topic populates the provided destination Topic from our Topic
func (topic *Topic) AssignProperties_To_Topic(destination *v20200601s.Topic) error {

	// ObjectMeta
	destination.ObjectMeta = *topic.ObjectMeta.DeepCopy()

	// Spec
	var spec v20200601s.Topic_Spec
	err := topic.Spec.AssignProperties_To_Topic_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Topic_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20200601s.Topic_STATUS
	err = topic.Status.AssignProperties_To_Topic_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Topic_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForTopic interface (if implemented) to customize the conversion
	var topicAsAny any = topic
	if augmentedTopic, ok := topicAsAny.(augmentConversionForTopic); ok {
		err := augmentedTopic.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (topic *Topic) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: topic.Spec.OriginalVersion,
		Kind:    "Topic",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1alpha1api20200601.Topic
// Deprecated version of Topic. Use v1beta20200601.Topic instead
type TopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Topic `json:"items"`
}

type augmentConversionForTopic interface {
	AssignPropertiesFrom(src *v20200601s.Topic) error
	AssignPropertiesTo(dst *v20200601s.Topic) error
}

// Storage version of v1alpha1api20200601.Topic_Spec
type Topic_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName          string              `json:"azureName,omitempty"`
	InboundIpRules     []InboundIpRule     `json:"inboundIpRules,omitempty"`
	InputSchema        *string             `json:"inputSchema,omitempty"`
	InputSchemaMapping *InputSchemaMapping `json:"inputSchemaMapping,omitempty"`
	Location           *string             `json:"location,omitempty"`
	OriginalVersion    string              `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner               *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag         genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	PublicNetworkAccess *string                            `json:"publicNetworkAccess,omitempty"`
	Tags                map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Topic_Spec{}

// ConvertSpecFrom populates our Topic_Spec from the provided source
func (topic *Topic_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20200601s.Topic_Spec)
	if ok {
		// Populate our instance from source
		return topic.AssignProperties_From_Topic_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20200601s.Topic_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = topic.AssignProperties_From_Topic_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Topic_Spec
func (topic *Topic_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20200601s.Topic_Spec)
	if ok {
		// Populate destination from our instance
		return topic.AssignProperties_To_Topic_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20200601s.Topic_Spec{}
	err := topic.AssignProperties_To_Topic_Spec(dst)
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

// AssignProperties_From_Topic_Spec populates our Topic_Spec from the provided source Topic_Spec
func (topic *Topic_Spec) AssignProperties_From_Topic_Spec(source *v20200601s.Topic_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	topic.AzureName = source.AzureName

	// InboundIpRules
	if source.InboundIpRules != nil {
		inboundIpRuleList := make([]InboundIpRule, len(source.InboundIpRules))
		for inboundIpRuleIndex, inboundIpRuleItem := range source.InboundIpRules {
			// Shadow the loop variable to avoid aliasing
			inboundIpRuleItem := inboundIpRuleItem
			var inboundIpRule InboundIpRule
			err := inboundIpRule.AssignProperties_From_InboundIpRule(&inboundIpRuleItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignProperties_From_InboundIpRule() to populate field InboundIpRules")
			}
			inboundIpRuleList[inboundIpRuleIndex] = inboundIpRule
		}
		topic.InboundIpRules = inboundIpRuleList
	} else {
		topic.InboundIpRules = nil
	}

	// InputSchema
	topic.InputSchema = genruntime.ClonePointerToString(source.InputSchema)

	// InputSchemaMapping
	if source.InputSchemaMapping != nil {
		var inputSchemaMapping InputSchemaMapping
		err := inputSchemaMapping.AssignProperties_From_InputSchemaMapping(source.InputSchemaMapping)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_InputSchemaMapping() to populate field InputSchemaMapping")
		}
		topic.InputSchemaMapping = &inputSchemaMapping
	} else {
		topic.InputSchemaMapping = nil
	}

	// Location
	topic.Location = genruntime.ClonePointerToString(source.Location)

	// OriginalVersion
	topic.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		topic.Owner = &owner
	} else {
		topic.Owner = nil
	}

	// PublicNetworkAccess
	topic.PublicNetworkAccess = genruntime.ClonePointerToString(source.PublicNetworkAccess)

	// Tags
	topic.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		topic.PropertyBag = propertyBag
	} else {
		topic.PropertyBag = nil
	}

	// Invoke the augmentConversionForTopic_Spec interface (if implemented) to customize the conversion
	var topicAsAny any = topic
	if augmentedTopic, ok := topicAsAny.(augmentConversionForTopic_Spec); ok {
		err := augmentedTopic.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Topic_Spec populates the provided destination Topic_Spec from our Topic_Spec
func (topic *Topic_Spec) AssignProperties_To_Topic_Spec(destination *v20200601s.Topic_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(topic.PropertyBag)

	// AzureName
	destination.AzureName = topic.AzureName

	// InboundIpRules
	if topic.InboundIpRules != nil {
		inboundIpRuleList := make([]v20200601s.InboundIpRule, len(topic.InboundIpRules))
		for inboundIpRuleIndex, inboundIpRuleItem := range topic.InboundIpRules {
			// Shadow the loop variable to avoid aliasing
			inboundIpRuleItem := inboundIpRuleItem
			var inboundIpRule v20200601s.InboundIpRule
			err := inboundIpRuleItem.AssignProperties_To_InboundIpRule(&inboundIpRule)
			if err != nil {
				return errors.Wrap(err, "calling AssignProperties_To_InboundIpRule() to populate field InboundIpRules")
			}
			inboundIpRuleList[inboundIpRuleIndex] = inboundIpRule
		}
		destination.InboundIpRules = inboundIpRuleList
	} else {
		destination.InboundIpRules = nil
	}

	// InputSchema
	destination.InputSchema = genruntime.ClonePointerToString(topic.InputSchema)

	// InputSchemaMapping
	if topic.InputSchemaMapping != nil {
		var inputSchemaMapping v20200601s.InputSchemaMapping
		err := topic.InputSchemaMapping.AssignProperties_To_InputSchemaMapping(&inputSchemaMapping)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_InputSchemaMapping() to populate field InputSchemaMapping")
		}
		destination.InputSchemaMapping = &inputSchemaMapping
	} else {
		destination.InputSchemaMapping = nil
	}

	// Location
	destination.Location = genruntime.ClonePointerToString(topic.Location)

	// OriginalVersion
	destination.OriginalVersion = topic.OriginalVersion

	// Owner
	if topic.Owner != nil {
		owner := topic.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// PublicNetworkAccess
	destination.PublicNetworkAccess = genruntime.ClonePointerToString(topic.PublicNetworkAccess)

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(topic.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForTopic_Spec interface (if implemented) to customize the conversion
	var topicAsAny any = topic
	if augmentedTopic, ok := topicAsAny.(augmentConversionForTopic_Spec); ok {
		err := augmentedTopic.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20200601.Topic_STATUS
// Deprecated version of Topic_STATUS. Use v1beta20200601.Topic_STATUS instead
type Topic_STATUS struct {
	Conditions                 []conditions.Condition                                       `json:"conditions,omitempty"`
	Endpoint                   *string                                                      `json:"endpoint,omitempty"`
	Id                         *string                                                      `json:"id,omitempty"`
	InboundIpRules             []InboundIpRule_STATUS                                       `json:"inboundIpRules,omitempty"`
	InputSchema                *string                                                      `json:"inputSchema,omitempty"`
	InputSchemaMapping         *InputSchemaMapping_STATUS                                   `json:"inputSchemaMapping,omitempty"`
	Location                   *string                                                      `json:"location,omitempty"`
	MetricResourceId           *string                                                      `json:"metricResourceId,omitempty"`
	Name                       *string                                                      `json:"name,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded `json:"privateEndpointConnections,omitempty"`
	PropertyBag                genruntime.PropertyBag                                       `json:"$propertyBag,omitempty"`
	ProvisioningState          *string                                                      `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *string                                                      `json:"publicNetworkAccess,omitempty"`
	SystemData                 *SystemData_STATUS                                           `json:"systemData,omitempty"`
	Tags                       map[string]string                                            `json:"tags,omitempty"`
	Type                       *string                                                      `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Topic_STATUS{}

// ConvertStatusFrom populates our Topic_STATUS from the provided source
func (topic *Topic_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20200601s.Topic_STATUS)
	if ok {
		// Populate our instance from source
		return topic.AssignProperties_From_Topic_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20200601s.Topic_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = topic.AssignProperties_From_Topic_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Topic_STATUS
func (topic *Topic_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20200601s.Topic_STATUS)
	if ok {
		// Populate destination from our instance
		return topic.AssignProperties_To_Topic_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20200601s.Topic_STATUS{}
	err := topic.AssignProperties_To_Topic_STATUS(dst)
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

// AssignProperties_From_Topic_STATUS populates our Topic_STATUS from the provided source Topic_STATUS
func (topic *Topic_STATUS) AssignProperties_From_Topic_STATUS(source *v20200601s.Topic_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	topic.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Endpoint
	topic.Endpoint = genruntime.ClonePointerToString(source.Endpoint)

	// Id
	topic.Id = genruntime.ClonePointerToString(source.Id)

	// InboundIpRules
	if source.InboundIpRules != nil {
		inboundIpRuleList := make([]InboundIpRule_STATUS, len(source.InboundIpRules))
		for inboundIpRuleIndex, inboundIpRuleItem := range source.InboundIpRules {
			// Shadow the loop variable to avoid aliasing
			inboundIpRuleItem := inboundIpRuleItem
			var inboundIpRule InboundIpRule_STATUS
			err := inboundIpRule.AssignProperties_From_InboundIpRule_STATUS(&inboundIpRuleItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignProperties_From_InboundIpRule_STATUS() to populate field InboundIpRules")
			}
			inboundIpRuleList[inboundIpRuleIndex] = inboundIpRule
		}
		topic.InboundIpRules = inboundIpRuleList
	} else {
		topic.InboundIpRules = nil
	}

	// InputSchema
	topic.InputSchema = genruntime.ClonePointerToString(source.InputSchema)

	// InputSchemaMapping
	if source.InputSchemaMapping != nil {
		var inputSchemaMapping InputSchemaMapping_STATUS
		err := inputSchemaMapping.AssignProperties_From_InputSchemaMapping_STATUS(source.InputSchemaMapping)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_InputSchemaMapping_STATUS() to populate field InputSchemaMapping")
		}
		topic.InputSchemaMapping = &inputSchemaMapping
	} else {
		topic.InputSchemaMapping = nil
	}

	// Location
	topic.Location = genruntime.ClonePointerToString(source.Location)

	// MetricResourceId
	topic.MetricResourceId = genruntime.ClonePointerToString(source.MetricResourceId)

	// Name
	topic.Name = genruntime.ClonePointerToString(source.Name)

	// PrivateEndpointConnections
	if source.PrivateEndpointConnections != nil {
		privateEndpointConnectionList := make([]PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded, len(source.PrivateEndpointConnections))
		for privateEndpointConnectionIndex, privateEndpointConnectionItem := range source.PrivateEndpointConnections {
			// Shadow the loop variable to avoid aliasing
			privateEndpointConnectionItem := privateEndpointConnectionItem
			var privateEndpointConnection PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded
			err := privateEndpointConnection.AssignProperties_From_PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded(&privateEndpointConnectionItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignProperties_From_PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded() to populate field PrivateEndpointConnections")
			}
			privateEndpointConnectionList[privateEndpointConnectionIndex] = privateEndpointConnection
		}
		topic.PrivateEndpointConnections = privateEndpointConnectionList
	} else {
		topic.PrivateEndpointConnections = nil
	}

	// ProvisioningState
	topic.ProvisioningState = genruntime.ClonePointerToString(source.ProvisioningState)

	// PublicNetworkAccess
	topic.PublicNetworkAccess = genruntime.ClonePointerToString(source.PublicNetworkAccess)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
		}
		topic.SystemData = &systemDatum
	} else {
		topic.SystemData = nil
	}

	// Tags
	topic.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	topic.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		topic.PropertyBag = propertyBag
	} else {
		topic.PropertyBag = nil
	}

	// Invoke the augmentConversionForTopic_STATUS interface (if implemented) to customize the conversion
	var topicAsAny any = topic
	if augmentedTopic, ok := topicAsAny.(augmentConversionForTopic_STATUS); ok {
		err := augmentedTopic.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Topic_STATUS populates the provided destination Topic_STATUS from our Topic_STATUS
func (topic *Topic_STATUS) AssignProperties_To_Topic_STATUS(destination *v20200601s.Topic_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(topic.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(topic.Conditions)

	// Endpoint
	destination.Endpoint = genruntime.ClonePointerToString(topic.Endpoint)

	// Id
	destination.Id = genruntime.ClonePointerToString(topic.Id)

	// InboundIpRules
	if topic.InboundIpRules != nil {
		inboundIpRuleList := make([]v20200601s.InboundIpRule_STATUS, len(topic.InboundIpRules))
		for inboundIpRuleIndex, inboundIpRuleItem := range topic.InboundIpRules {
			// Shadow the loop variable to avoid aliasing
			inboundIpRuleItem := inboundIpRuleItem
			var inboundIpRule v20200601s.InboundIpRule_STATUS
			err := inboundIpRuleItem.AssignProperties_To_InboundIpRule_STATUS(&inboundIpRule)
			if err != nil {
				return errors.Wrap(err, "calling AssignProperties_To_InboundIpRule_STATUS() to populate field InboundIpRules")
			}
			inboundIpRuleList[inboundIpRuleIndex] = inboundIpRule
		}
		destination.InboundIpRules = inboundIpRuleList
	} else {
		destination.InboundIpRules = nil
	}

	// InputSchema
	destination.InputSchema = genruntime.ClonePointerToString(topic.InputSchema)

	// InputSchemaMapping
	if topic.InputSchemaMapping != nil {
		var inputSchemaMapping v20200601s.InputSchemaMapping_STATUS
		err := topic.InputSchemaMapping.AssignProperties_To_InputSchemaMapping_STATUS(&inputSchemaMapping)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_InputSchemaMapping_STATUS() to populate field InputSchemaMapping")
		}
		destination.InputSchemaMapping = &inputSchemaMapping
	} else {
		destination.InputSchemaMapping = nil
	}

	// Location
	destination.Location = genruntime.ClonePointerToString(topic.Location)

	// MetricResourceId
	destination.MetricResourceId = genruntime.ClonePointerToString(topic.MetricResourceId)

	// Name
	destination.Name = genruntime.ClonePointerToString(topic.Name)

	// PrivateEndpointConnections
	if topic.PrivateEndpointConnections != nil {
		privateEndpointConnectionList := make([]v20200601s.PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded, len(topic.PrivateEndpointConnections))
		for privateEndpointConnectionIndex, privateEndpointConnectionItem := range topic.PrivateEndpointConnections {
			// Shadow the loop variable to avoid aliasing
			privateEndpointConnectionItem := privateEndpointConnectionItem
			var privateEndpointConnection v20200601s.PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded
			err := privateEndpointConnectionItem.AssignProperties_To_PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded(&privateEndpointConnection)
			if err != nil {
				return errors.Wrap(err, "calling AssignProperties_To_PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded() to populate field PrivateEndpointConnections")
			}
			privateEndpointConnectionList[privateEndpointConnectionIndex] = privateEndpointConnection
		}
		destination.PrivateEndpointConnections = privateEndpointConnectionList
	} else {
		destination.PrivateEndpointConnections = nil
	}

	// ProvisioningState
	destination.ProvisioningState = genruntime.ClonePointerToString(topic.ProvisioningState)

	// PublicNetworkAccess
	destination.PublicNetworkAccess = genruntime.ClonePointerToString(topic.PublicNetworkAccess)

	// SystemData
	if topic.SystemData != nil {
		var systemDatum v20200601s.SystemData_STATUS
		err := topic.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(topic.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(topic.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForTopic_STATUS interface (if implemented) to customize the conversion
	var topicAsAny any = topic
	if augmentedTopic, ok := topicAsAny.(augmentConversionForTopic_STATUS); ok {
		err := augmentedTopic.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForTopic_Spec interface {
	AssignPropertiesFrom(src *v20200601s.Topic_Spec) error
	AssignPropertiesTo(dst *v20200601s.Topic_Spec) error
}

type augmentConversionForTopic_STATUS interface {
	AssignPropertiesFrom(src *v20200601s.Topic_STATUS) error
	AssignPropertiesTo(dst *v20200601s.Topic_STATUS) error
}

// Storage version of v1alpha1api20200601.PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded
// Deprecated version of PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded. Use v1beta20200601.PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded instead
type PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded populates our PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded from the provided source PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded
func (embedded *PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded) AssignProperties_From_PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded(source *v20200601s.PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Id
	embedded.Id = genruntime.ClonePointerToString(source.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		embedded.PropertyBag = propertyBag
	} else {
		embedded.PropertyBag = nil
	}

	// Invoke the augmentConversionForPrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded interface (if implemented) to customize the conversion
	var embeddedAsAny any = embedded
	if augmentedEmbedded, ok := embeddedAsAny.(augmentConversionForPrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded); ok {
		err := augmentedEmbedded.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded populates the provided destination PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded from our PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded
func (embedded *PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded) AssignProperties_To_PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded(destination *v20200601s.PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(embedded.PropertyBag)

	// Id
	destination.Id = genruntime.ClonePointerToString(embedded.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForPrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded interface (if implemented) to customize the conversion
	var embeddedAsAny any = embedded
	if augmentedEmbedded, ok := embeddedAsAny.(augmentConversionForPrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded); ok {
		err := augmentedEmbedded.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForPrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded interface {
	AssignPropertiesFrom(src *v20200601s.PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded) error
	AssignPropertiesTo(dst *v20200601s.PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded) error
}

func init() {
	SchemeBuilder.Register(&Topic{}, &TopicList{})
}
