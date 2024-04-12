// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"context"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// +kubebuilder:rbac:groups=dbforpostgresql.azure.com,resources=flexibleservers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=dbforpostgresql.azure.com,resources={flexibleservers/status,flexibleservers/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20221201.FlexibleServer
// Generator information:
// - Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-12-01/FlexibleServers.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}
type FlexibleServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServer_Spec   `json:"spec,omitempty"`
	Status            FlexibleServer_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &FlexibleServer{}

// GetConditions returns the conditions of the resource
func (server *FlexibleServer) GetConditions() conditions.Conditions {
	return server.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (server *FlexibleServer) SetConditions(conditions conditions.Conditions) {
	server.Status.Conditions = conditions
}

var _ genruntime.KubernetesExporter = &FlexibleServer{}

// ExportKubernetesResources defines a resource which can create other resources in Kubernetes.
func (server *FlexibleServer) ExportKubernetesResources(_ context.Context, _ genruntime.MetaObject, _ *genericarmclient.GenericClient, _ logr.Logger) ([]client.Object, error) {
	collector := configmaps.NewCollector(server.Namespace)
	if server.Spec.OperatorSpec != nil && server.Spec.OperatorSpec.ConfigMaps != nil {
		if server.Status.FullyQualifiedDomainName != nil {
			collector.AddValue(server.Spec.OperatorSpec.ConfigMaps.FullyQualifiedDomainName, *server.Status.FullyQualifiedDomainName)
		}
	}
	result, err := collector.Values()
	if err != nil {
		return nil, err
	}
	return configmaps.SliceToClientObjectSlice(result), nil
}

var _ genruntime.KubernetesResource = &FlexibleServer{}

// AzureName returns the Azure name of the resource
func (server *FlexibleServer) AzureName() string {
	return server.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-12-01"
func (server FlexibleServer) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (server *FlexibleServer) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (server *FlexibleServer) GetSpec() genruntime.ConvertibleSpec {
	return &server.Spec
}

// GetStatus returns the status of this resource
func (server *FlexibleServer) GetStatus() genruntime.ConvertibleStatus {
	return &server.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (server *FlexibleServer) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers"
func (server *FlexibleServer) GetType() string {
	return "Microsoft.DBforPostgreSQL/flexibleServers"
}

// NewEmptyStatus returns a new empty (blank) status
func (server *FlexibleServer) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &FlexibleServer_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (server *FlexibleServer) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(server.Spec)
	return server.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (server *FlexibleServer) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*FlexibleServer_STATUS); ok {
		server.Status = *st
		return nil
	}

	// Convert status to required version
	var st FlexibleServer_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	server.Status = st
	return nil
}

// Hub marks that this FlexibleServer is the hub type for conversion
func (server *FlexibleServer) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (server *FlexibleServer) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: server.Spec.OriginalVersion,
		Kind:    "FlexibleServer",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20221201.FlexibleServer
// Generator information:
// - Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-12-01/FlexibleServers.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}
type FlexibleServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServer `json:"items"`
}

// Storage version of v1api20221201.APIVersion
// +kubebuilder:validation:Enum={"2022-12-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2022-12-01")

// Storage version of v1api20221201.FlexibleServer_Spec
type FlexibleServer_Spec struct {
	AdministratorLogin         *string                     `json:"administratorLogin,omitempty"`
	AdministratorLoginPassword *genruntime.SecretReference `json:"administratorLoginPassword,omitempty"`
	AuthConfig                 *AuthConfig                 `json:"authConfig,omitempty"`
	AvailabilityZone           *string                     `json:"availabilityZone,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName         string                      `json:"azureName,omitempty"`
	Backup            *Backup                     `json:"backup,omitempty"`
	CreateMode        *string                     `json:"createMode,omitempty"`
	DataEncryption    *DataEncryption             `json:"dataEncryption,omitempty"`
	HighAvailability  *HighAvailability           `json:"highAvailability,omitempty"`
	Identity          *UserAssignedIdentity       `json:"identity,omitempty"`
	Location          *string                     `json:"location,omitempty"`
	MaintenanceWindow *MaintenanceWindow          `json:"maintenanceWindow,omitempty"`
	Network           *Network                    `json:"network,omitempty"`
	OperatorSpec      *FlexibleServerOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion   string                      `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner           *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PointInTimeUTC  *string                            `json:"pointInTimeUTC,omitempty"`
	PropertyBag     genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	ReplicationRole *string                            `json:"replicationRole,omitempty"`
	Sku             *Sku                               `json:"sku,omitempty"`

	// SourceServerResourceReference: The source server resource ID to restore from. It's required when 'createMode' is
	// 'PointInTimeRestore' or 'GeoRestore' or 'Replica'. This property is returned only for Replica server
	SourceServerResourceReference *genruntime.ResourceReference `armReference:"SourceServerResourceId" json:"sourceServerResourceReference,omitempty"`
	Storage                       *Storage                      `json:"storage,omitempty"`
	Tags                          map[string]string             `json:"tags,omitempty"`
	Version                       *string                       `json:"version,omitempty"`
}

var _ genruntime.ConvertibleSpec = &FlexibleServer_Spec{}

// ConvertSpecFrom populates our FlexibleServer_Spec from the provided source
func (server *FlexibleServer_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == server {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(server)
}

// ConvertSpecTo populates the provided destination from our FlexibleServer_Spec
func (server *FlexibleServer_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == server {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(server)
}

// Storage version of v1api20221201.FlexibleServer_STATUS
type FlexibleServer_STATUS struct {
	AdministratorLogin       *string                      `json:"administratorLogin,omitempty"`
	AuthConfig               *AuthConfig_STATUS           `json:"authConfig,omitempty"`
	AvailabilityZone         *string                      `json:"availabilityZone,omitempty"`
	Backup                   *Backup_STATUS               `json:"backup,omitempty"`
	Conditions               []conditions.Condition       `json:"conditions,omitempty"`
	CreateMode               *string                      `json:"createMode,omitempty"`
	DataEncryption           *DataEncryption_STATUS       `json:"dataEncryption,omitempty"`
	FullyQualifiedDomainName *string                      `json:"fullyQualifiedDomainName,omitempty"`
	HighAvailability         *HighAvailability_STATUS     `json:"highAvailability,omitempty"`
	Id                       *string                      `json:"id,omitempty"`
	Identity                 *UserAssignedIdentity_STATUS `json:"identity,omitempty"`
	Location                 *string                      `json:"location,omitempty"`
	MaintenanceWindow        *MaintenanceWindow_STATUS    `json:"maintenanceWindow,omitempty"`
	MinorVersion             *string                      `json:"minorVersion,omitempty"`
	Name                     *string                      `json:"name,omitempty"`
	Network                  *Network_STATUS              `json:"network,omitempty"`
	PointInTimeUTC           *string                      `json:"pointInTimeUTC,omitempty"`
	PropertyBag              genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	ReplicaCapacity          *int                         `json:"replicaCapacity,omitempty"`
	ReplicationRole          *string                      `json:"replicationRole,omitempty"`
	Sku                      *Sku_STATUS                  `json:"sku,omitempty"`
	SourceServerResourceId   *string                      `json:"sourceServerResourceId,omitempty"`
	State                    *string                      `json:"state,omitempty"`
	Storage                  *Storage_STATUS              `json:"storage,omitempty"`
	SystemData               *SystemData_STATUS           `json:"systemData,omitempty"`
	Tags                     map[string]string            `json:"tags,omitempty"`
	Type                     *string                      `json:"type,omitempty"`
	Version                  *string                      `json:"version,omitempty"`
}

var _ genruntime.ConvertibleStatus = &FlexibleServer_STATUS{}

// ConvertStatusFrom populates our FlexibleServer_STATUS from the provided source
func (server *FlexibleServer_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == server {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(server)
}

// ConvertStatusTo populates the provided destination from our FlexibleServer_STATUS
func (server *FlexibleServer_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == server {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(server)
}

// Storage version of v1api20221201.AuthConfig
// Authentication configuration properties of a server
type AuthConfig struct {
	ActiveDirectoryAuth *string                `json:"activeDirectoryAuth,omitempty"`
	PasswordAuth        *string                `json:"passwordAuth,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TenantId            *string                `json:"tenantId,omitempty"`
}

// Storage version of v1api20221201.AuthConfig_STATUS
// Authentication configuration properties of a server
type AuthConfig_STATUS struct {
	ActiveDirectoryAuth *string                `json:"activeDirectoryAuth,omitempty"`
	PasswordAuth        *string                `json:"passwordAuth,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TenantId            *string                `json:"tenantId,omitempty"`
}

// Storage version of v1api20221201.Backup
// Backup properties of a server
type Backup struct {
	BackupRetentionDays *int                   `json:"backupRetentionDays,omitempty"`
	GeoRedundantBackup  *string                `json:"geoRedundantBackup,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20221201.Backup_STATUS
// Backup properties of a server
type Backup_STATUS struct {
	BackupRetentionDays *int                   `json:"backupRetentionDays,omitempty"`
	EarliestRestoreDate *string                `json:"earliestRestoreDate,omitempty"`
	GeoRedundantBackup  *string                `json:"geoRedundantBackup,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20221201.DataEncryption
// Data encryption properties of a server
type DataEncryption struct {
	PrimaryKeyURI *string `json:"primaryKeyURI,omitempty"`

	// PrimaryUserAssignedIdentityReference: Resource Id for the User assigned identity to be used for data encryption for
	// primary server.
	PrimaryUserAssignedIdentityReference *genruntime.ResourceReference `armReference:"PrimaryUserAssignedIdentityId" json:"primaryUserAssignedIdentityReference,omitempty"`
	PropertyBag                          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	Type                                 *string                       `json:"type,omitempty"`
}

// Storage version of v1api20221201.DataEncryption_STATUS
// Data encryption properties of a server
type DataEncryption_STATUS struct {
	PrimaryKeyURI                 *string                `json:"primaryKeyURI,omitempty"`
	PrimaryUserAssignedIdentityId *string                `json:"primaryUserAssignedIdentityId,omitempty"`
	PropertyBag                   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type                          *string                `json:"type,omitempty"`
}

// Storage version of v1api20221201.FlexibleServerOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type FlexibleServerOperatorSpec struct {
	ConfigMaps  *FlexibleServerOperatorConfigMaps `json:"configMaps,omitempty"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Secrets     *FlexibleServerOperatorSecrets    `json:"secrets,omitempty"`
}

// Storage version of v1api20221201.HighAvailability
// High availability properties of a server
type HighAvailability struct {
	Mode                    *string                `json:"mode,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StandbyAvailabilityZone *string                `json:"standbyAvailabilityZone,omitempty"`
}

// Storage version of v1api20221201.HighAvailability_STATUS
// High availability properties of a server
type HighAvailability_STATUS struct {
	Mode                    *string                `json:"mode,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StandbyAvailabilityZone *string                `json:"standbyAvailabilityZone,omitempty"`
	State                   *string                `json:"state,omitempty"`
}

// Storage version of v1api20221201.MaintenanceWindow
// Maintenance window properties of a server.
type MaintenanceWindow struct {
	CustomWindow *string                `json:"customWindow,omitempty"`
	DayOfWeek    *int                   `json:"dayOfWeek,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StartHour    *int                   `json:"startHour,omitempty"`
	StartMinute  *int                   `json:"startMinute,omitempty"`
}

// Storage version of v1api20221201.MaintenanceWindow_STATUS
// Maintenance window properties of a server.
type MaintenanceWindow_STATUS struct {
	CustomWindow *string                `json:"customWindow,omitempty"`
	DayOfWeek    *int                   `json:"dayOfWeek,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StartHour    *int                   `json:"startHour,omitempty"`
	StartMinute  *int                   `json:"startMinute,omitempty"`
}

// Storage version of v1api20221201.Network
// Network properties of a server.
type Network struct {
	// DelegatedSubnetResourceReference: Delegated subnet arm resource id. This is required to be passed during create, in case
	// we want the server to be VNET injected, i.e. Private access server. During update, pass this only if we want to update
	// the value for Private DNS zone.
	DelegatedSubnetResourceReference *genruntime.ResourceReference `armReference:"DelegatedSubnetResourceId" json:"delegatedSubnetResourceReference,omitempty"`

	// PrivateDnsZoneArmResourceReference: Private dns zone arm resource id. This is required to be passed during create, in
	// case we want the server to be VNET injected, i.e. Private access server. During update, pass this only if we want to
	// update the value for Private DNS zone.
	PrivateDnsZoneArmResourceReference *genruntime.ResourceReference `armReference:"PrivateDnsZoneArmResourceId" json:"privateDnsZoneArmResourceReference,omitempty"`
	PropertyBag                        genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20221201.Network_STATUS
// Network properties of a server.
type Network_STATUS struct {
	DelegatedSubnetResourceId   *string                `json:"delegatedSubnetResourceId,omitempty"`
	PrivateDnsZoneArmResourceId *string                `json:"privateDnsZoneArmResourceId,omitempty"`
	PropertyBag                 genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	PublicNetworkAccess         *string                `json:"publicNetworkAccess,omitempty"`
}

// Storage version of v1api20221201.Sku
// Sku information related properties of a server.
type Sku struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1api20221201.Sku_STATUS
// Sku information related properties of a server.
type Sku_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1api20221201.Storage
// Storage properties of a server
type Storage struct {
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StorageSizeGB *int                   `json:"storageSizeGB,omitempty"`
}

// Storage version of v1api20221201.Storage_STATUS
// Storage properties of a server
type Storage_STATUS struct {
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StorageSizeGB *int                   `json:"storageSizeGB,omitempty"`
}

// Storage version of v1api20221201.SystemData_STATUS
// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS struct {
	CreatedAt          *string                `json:"createdAt,omitempty"`
	CreatedBy          *string                `json:"createdBy,omitempty"`
	CreatedByType      *string                `json:"createdByType,omitempty"`
	LastModifiedAt     *string                `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *string                `json:"lastModifiedByType,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20221201.UserAssignedIdentity
// Information describing the identities associated with this application.
type UserAssignedIdentity struct {
	PropertyBag            genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	Type                   *string                       `json:"type,omitempty"`
	UserAssignedIdentities []UserAssignedIdentityDetails `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1api20221201.UserAssignedIdentity_STATUS
// Information describing the identities associated with this application.
type UserAssignedIdentity_STATUS struct {
	PropertyBag            genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
	TenantId               *string                        `json:"tenantId,omitempty"`
	Type                   *string                        `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserIdentity_STATUS `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1api20221201.FlexibleServerOperatorConfigMaps
type FlexibleServerOperatorConfigMaps struct {
	FullyQualifiedDomainName *genruntime.ConfigMapDestination `json:"fullyQualifiedDomainName,omitempty"`
	PropertyBag              genruntime.PropertyBag           `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20221201.FlexibleServerOperatorSecrets
type FlexibleServerOperatorSecrets struct {
	FullyQualifiedDomainName *genruntime.SecretDestination `json:"fullyQualifiedDomainName,omitempty"`
	PropertyBag              genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20221201.UserAssignedIdentityDetails
// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails struct {
	PropertyBag genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	Reference   genruntime.ResourceReference `armReference:"Reference" json:"reference,omitempty"`
}

// Storage version of v1api20221201.UserIdentity_STATUS
// Describes a single user-assigned identity associated with the application.
type UserIdentity_STATUS struct {
	ClientId    *string                `json:"clientId,omitempty"`
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&FlexibleServer{}, &FlexibleServerList{})
}
