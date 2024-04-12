/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	search "github.com/Azure/azure-service-operator/v2/api/search/v1api20220901/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
)

var _ genruntime.KubernetesExporter = &SearchServiceExtension{}

func (ext *SearchServiceExtension) ExportKubernetesResources(
	ctx context.Context,
	obj genruntime.MetaObject,
	armClient *genericarmclient.GenericClient,
	log logr.Logger) ([]client.Object, error) {

	// This has to be the current hub devices version. It will need to be updated
	// if the hub devices version changes.
	typedObj, ok := obj.(*search.SearchService)
	if !ok {
		return nil, errors.Errorf("cannot run on unknown resource type %T, expected *devices.IotHub", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not
	var _ conversion.Hub = typedObj

	hasSecrets := secretsSpecified(typedObj)
	if !hasSecrets {
		log.V(Debug).Info("No secrets retrieval to perform as operatorSpec is empty")
		return nil, nil
	}

	id, err := genruntime.GetAndParseResourceID(typedObj)
	if err != nil {
		return nil, err
	}

	queryKeys := make(map[string]armsearch.QueryKey)
	var adminKeys armsearch.AdminKeysClientGetResponse
	// Only bother calling ListKeys if there are secrets to retrieve
	if hasSecrets {
		subscription := id.SubscriptionID
		// Using armClient.ClientOptions() here ensures we share the same HTTP connection, so this is not opening a new
		// connection each time through
		var queryKeysClient *armsearch.QueryKeysClient
		queryKeysClient, err = armsearch.NewQueryKeysClient(subscription, armClient.Creds(), armClient.ClientOptions())
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create new SeachServiceQueryClient")
		}

		var pager *runtime.Pager[armsearch.QueryKeysClientListBySearchServiceResponse]
		var resp armsearch.QueryKeysClientListBySearchServiceResponse
		pager = queryKeysClient.NewListBySearchServicePager(id.ResourceGroupName, typedObj.AzureName(), nil, nil)
		for pager.More() {
			resp, err = pager.NextPage(ctx)
			addSecretsToMap(resp.Value, queryKeys)
		}

		if err != nil {
			return nil, errors.Wrapf(err, "failed listing query keys")
		}

		var adminKeysClient *armsearch.AdminKeysClient
		adminKeysClient, err = armsearch.NewAdminKeysClient(subscription, armClient.Creds(), armClient.ClientOptions())
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create new SeachServiceAdminClient")
		}

		adminKeys, err = adminKeysClient.Get(ctx, id.ResourceGroupName, typedObj.AzureName(), nil, nil)
		if err != nil {
			return nil, err
		}
	}

	secretSlice, err := secretsToWrite(typedObj, queryKeys, adminKeys)
	if err != nil {
		return nil, err
	}

	return secrets.SliceToClientObjectSlice(secretSlice), nil
}

func secretsSpecified(obj *search.SearchService) bool {
	if obj.Spec.OperatorSpec == nil || obj.Spec.OperatorSpec.Secrets == nil {
		return false
	}

	secrets := obj.Spec.OperatorSpec.Secrets

	if secrets.AdminPrimaryKey != nil ||
		secrets.AdminSecondaryKey != nil ||
		secrets.QueryKey != nil {
		return true
	}

	return false
}

func addSecretsToMap(keys []*armsearch.QueryKey, result map[string]armsearch.QueryKey) {
	for _, key := range keys {
		if key == nil {
			continue
		}

		// We have to do it this way, since the autogenerated query key has key.Name == nil. See screenshot in https://learn.microsoft.com/en-us/azure/search/search-security-api-keys?tabs=portal-use%2Cportal-find%2Cportal-query#find-existing-keys
		if key.Name == nil && key.Key != nil {
			result["default"] = *key
			continue
		}

		result[*key.Name] = *key
	}
}

func secretsToWrite(obj *search.SearchService, queryKeys map[string]armsearch.QueryKey, adminKeys armsearch.AdminKeysClientGetResponse) ([]*v1.Secret, error) {
	operatorSpecSecrets := obj.Spec.OperatorSpec.Secrets
	if operatorSpecSecrets == nil {
		return nil, errors.Errorf("unexpected nil operatorspec")
	}

	collector := secrets.NewCollector(obj.Namespace)
	defaultQueryKey, ok := queryKeys["default"]
	if ok {
		collector.AddValue(operatorSpecSecrets.QueryKey, to.Value(defaultQueryKey.Key))
	}

	collector.AddValue(operatorSpecSecrets.AdminPrimaryKey, to.Value(adminKeys.PrimaryKey))
	collector.AddValue(operatorSpecSecrets.AdminSecondaryKey, to.Value(adminKeys.SecondaryKey))

	return collector.Values()
}
