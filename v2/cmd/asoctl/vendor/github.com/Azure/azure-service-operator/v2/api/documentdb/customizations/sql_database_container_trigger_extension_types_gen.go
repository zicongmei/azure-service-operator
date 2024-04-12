// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20210515 "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20210515"
	v20210515s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20210515/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type SqlDatabaseContainerTriggerExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *SqlDatabaseContainerTriggerExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20210515.SqlDatabaseContainerTrigger{},
		&v20210515s.SqlDatabaseContainerTrigger{}}
}
