// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401storage

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_StorageAccountsBlobService_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsBlobService via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsBlobService, StorageAccountsBlobServiceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsBlobService runs a test to see if a specific instance of StorageAccountsBlobService round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsBlobService(subject StorageAccountsBlobService) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual StorageAccountsBlobService
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of StorageAccountsBlobService instances for property testing - lazily
//instantiated by StorageAccountsBlobServiceGenerator()
var storageAccountsBlobServiceGenerator gopter.Gen

// StorageAccountsBlobServiceGenerator returns a generator of StorageAccountsBlobService instances for property testing.
func StorageAccountsBlobServiceGenerator() gopter.Gen {
	if storageAccountsBlobServiceGenerator != nil {
		return storageAccountsBlobServiceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForStorageAccountsBlobService(generators)
	storageAccountsBlobServiceGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobService{}), generators)

	return storageAccountsBlobServiceGenerator
}

// AddRelatedPropertyGeneratorsForStorageAccountsBlobService is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsBlobService(gens map[string]gopter.Gen) {
	gens["Spec"] = StorageAccountsBlobServicesSpecGenerator()
	gens["Status"] = BlobServicePropertiesStatusGenerator()
}

func Test_BlobServiceProperties_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BlobServiceProperties_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBlobServicePropertiesStatus, BlobServicePropertiesStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBlobServicePropertiesStatus runs a test to see if a specific instance of BlobServiceProperties_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForBlobServicePropertiesStatus(subject BlobServiceProperties_Status) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual BlobServiceProperties_Status
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of BlobServiceProperties_Status instances for property testing -
//lazily instantiated by BlobServicePropertiesStatusGenerator()
var blobServicePropertiesStatusGenerator gopter.Gen

// BlobServicePropertiesStatusGenerator returns a generator of BlobServiceProperties_Status instances for property testing.
// We first initialize blobServicePropertiesStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BlobServicePropertiesStatusGenerator() gopter.Gen {
	if blobServicePropertiesStatusGenerator != nil {
		return blobServicePropertiesStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobServicePropertiesStatus(generators)
	blobServicePropertiesStatusGenerator = gen.Struct(reflect.TypeOf(BlobServiceProperties_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobServicePropertiesStatus(generators)
	AddRelatedPropertyGeneratorsForBlobServicePropertiesStatus(generators)
	blobServicePropertiesStatusGenerator = gen.Struct(reflect.TypeOf(BlobServiceProperties_Status{}), generators)

	return blobServicePropertiesStatusGenerator
}

// AddIndependentPropertyGeneratorsForBlobServicePropertiesStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBlobServicePropertiesStatus(gens map[string]gopter.Gen) {
	gens["AutomaticSnapshotPolicyEnabled"] = gen.PtrOf(gen.Bool())
	gens["DefaultServiceVersion"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IsVersioningEnabled"] = gen.PtrOf(gen.Bool())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBlobServicePropertiesStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBlobServicePropertiesStatus(gens map[string]gopter.Gen) {
	gens["ChangeFeed"] = gen.PtrOf(ChangeFeedStatusGenerator())
	gens["ContainerDeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicyStatusGenerator())
	gens["Cors"] = gen.PtrOf(CorsRulesStatusGenerator())
	gens["DeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicyStatusGenerator())
	gens["LastAccessTimeTrackingPolicy"] = gen.PtrOf(LastAccessTimeTrackingPolicyStatusGenerator())
	gens["RestorePolicy"] = gen.PtrOf(RestorePolicyPropertiesStatusGenerator())
	gens["Sku"] = gen.PtrOf(SkuStatusGenerator())
}

func Test_StorageAccountsBlobServices_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsBlobServices_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsBlobServicesSpec, StorageAccountsBlobServicesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsBlobServicesSpec runs a test to see if a specific instance of StorageAccountsBlobServices_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsBlobServicesSpec(subject StorageAccountsBlobServices_Spec) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual StorageAccountsBlobServices_Spec
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of StorageAccountsBlobServices_Spec instances for property testing -
//lazily instantiated by StorageAccountsBlobServicesSpecGenerator()
var storageAccountsBlobServicesSpecGenerator gopter.Gen

// StorageAccountsBlobServicesSpecGenerator returns a generator of StorageAccountsBlobServices_Spec instances for property testing.
// We first initialize storageAccountsBlobServicesSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsBlobServicesSpecGenerator() gopter.Gen {
	if storageAccountsBlobServicesSpecGenerator != nil {
		return storageAccountsBlobServicesSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesSpec(generators)
	storageAccountsBlobServicesSpecGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobServices_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesSpec(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesSpec(generators)
	storageAccountsBlobServicesSpecGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobServices_Spec{}), generators)

	return storageAccountsBlobServicesSpecGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesSpec(gens map[string]gopter.Gen) {
	gens["AutomaticSnapshotPolicyEnabled"] = gen.PtrOf(gen.Bool())
	gens["DefaultServiceVersion"] = gen.PtrOf(gen.AlphaString())
	gens["IsVersioningEnabled"] = gen.PtrOf(gen.Bool())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesSpec(gens map[string]gopter.Gen) {
	gens["ChangeFeed"] = gen.PtrOf(ChangeFeedGenerator())
	gens["ContainerDeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicyGenerator())
	gens["Cors"] = gen.PtrOf(CorsRulesGenerator())
	gens["DeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicyGenerator())
	gens["LastAccessTimeTrackingPolicy"] = gen.PtrOf(LastAccessTimeTrackingPolicyGenerator())
	gens["RestorePolicy"] = gen.PtrOf(RestorePolicyPropertiesGenerator())
}

func Test_ChangeFeed_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ChangeFeed via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForChangeFeed, ChangeFeedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForChangeFeed runs a test to see if a specific instance of ChangeFeed round trips to JSON and back losslessly
func RunJSONSerializationTestForChangeFeed(subject ChangeFeed) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual ChangeFeed
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of ChangeFeed instances for property testing - lazily instantiated by
//ChangeFeedGenerator()
var changeFeedGenerator gopter.Gen

// ChangeFeedGenerator returns a generator of ChangeFeed instances for property testing.
func ChangeFeedGenerator() gopter.Gen {
	if changeFeedGenerator != nil {
		return changeFeedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForChangeFeed(generators)
	changeFeedGenerator = gen.Struct(reflect.TypeOf(ChangeFeed{}), generators)

	return changeFeedGenerator
}

// AddIndependentPropertyGeneratorsForChangeFeed is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForChangeFeed(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
}

func Test_ChangeFeed_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ChangeFeed_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForChangeFeedStatus, ChangeFeedStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForChangeFeedStatus runs a test to see if a specific instance of ChangeFeed_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForChangeFeedStatus(subject ChangeFeed_Status) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual ChangeFeed_Status
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of ChangeFeed_Status instances for property testing - lazily
//instantiated by ChangeFeedStatusGenerator()
var changeFeedStatusGenerator gopter.Gen

// ChangeFeedStatusGenerator returns a generator of ChangeFeed_Status instances for property testing.
func ChangeFeedStatusGenerator() gopter.Gen {
	if changeFeedStatusGenerator != nil {
		return changeFeedStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForChangeFeedStatus(generators)
	changeFeedStatusGenerator = gen.Struct(reflect.TypeOf(ChangeFeed_Status{}), generators)

	return changeFeedStatusGenerator
}

// AddIndependentPropertyGeneratorsForChangeFeedStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForChangeFeedStatus(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
}

func Test_CorsRules_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorsRules via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorsRules, CorsRulesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorsRules runs a test to see if a specific instance of CorsRules round trips to JSON and back losslessly
func RunJSONSerializationTestForCorsRules(subject CorsRules) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual CorsRules
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of CorsRules instances for property testing - lazily instantiated by
//CorsRulesGenerator()
var corsRulesGenerator gopter.Gen

// CorsRulesGenerator returns a generator of CorsRules instances for property testing.
func CorsRulesGenerator() gopter.Gen {
	if corsRulesGenerator != nil {
		return corsRulesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForCorsRules(generators)
	corsRulesGenerator = gen.Struct(reflect.TypeOf(CorsRules{}), generators)

	return corsRulesGenerator
}

// AddRelatedPropertyGeneratorsForCorsRules is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCorsRules(gens map[string]gopter.Gen) {
	gens["CorsRules"] = gen.SliceOf(CorsRuleGenerator())
}

func Test_CorsRules_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorsRules_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorsRulesStatus, CorsRulesStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorsRulesStatus runs a test to see if a specific instance of CorsRules_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForCorsRulesStatus(subject CorsRules_Status) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual CorsRules_Status
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of CorsRules_Status instances for property testing - lazily
//instantiated by CorsRulesStatusGenerator()
var corsRulesStatusGenerator gopter.Gen

// CorsRulesStatusGenerator returns a generator of CorsRules_Status instances for property testing.
func CorsRulesStatusGenerator() gopter.Gen {
	if corsRulesStatusGenerator != nil {
		return corsRulesStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForCorsRulesStatus(generators)
	corsRulesStatusGenerator = gen.Struct(reflect.TypeOf(CorsRules_Status{}), generators)

	return corsRulesStatusGenerator
}

// AddRelatedPropertyGeneratorsForCorsRulesStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCorsRulesStatus(gens map[string]gopter.Gen) {
	gens["CorsRules"] = gen.SliceOf(CorsRuleStatusGenerator())
}

func Test_DeleteRetentionPolicy_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DeleteRetentionPolicy via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDeleteRetentionPolicy, DeleteRetentionPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDeleteRetentionPolicy runs a test to see if a specific instance of DeleteRetentionPolicy round trips to JSON and back losslessly
func RunJSONSerializationTestForDeleteRetentionPolicy(subject DeleteRetentionPolicy) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual DeleteRetentionPolicy
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of DeleteRetentionPolicy instances for property testing - lazily
//instantiated by DeleteRetentionPolicyGenerator()
var deleteRetentionPolicyGenerator gopter.Gen

// DeleteRetentionPolicyGenerator returns a generator of DeleteRetentionPolicy instances for property testing.
func DeleteRetentionPolicyGenerator() gopter.Gen {
	if deleteRetentionPolicyGenerator != nil {
		return deleteRetentionPolicyGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDeleteRetentionPolicy(generators)
	deleteRetentionPolicyGenerator = gen.Struct(reflect.TypeOf(DeleteRetentionPolicy{}), generators)

	return deleteRetentionPolicyGenerator
}

// AddIndependentPropertyGeneratorsForDeleteRetentionPolicy is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDeleteRetentionPolicy(gens map[string]gopter.Gen) {
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

func Test_DeleteRetentionPolicy_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DeleteRetentionPolicy_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDeleteRetentionPolicyStatus, DeleteRetentionPolicyStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDeleteRetentionPolicyStatus runs a test to see if a specific instance of DeleteRetentionPolicy_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForDeleteRetentionPolicyStatus(subject DeleteRetentionPolicy_Status) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual DeleteRetentionPolicy_Status
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of DeleteRetentionPolicy_Status instances for property testing -
//lazily instantiated by DeleteRetentionPolicyStatusGenerator()
var deleteRetentionPolicyStatusGenerator gopter.Gen

// DeleteRetentionPolicyStatusGenerator returns a generator of DeleteRetentionPolicy_Status instances for property testing.
func DeleteRetentionPolicyStatusGenerator() gopter.Gen {
	if deleteRetentionPolicyStatusGenerator != nil {
		return deleteRetentionPolicyStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDeleteRetentionPolicyStatus(generators)
	deleteRetentionPolicyStatusGenerator = gen.Struct(reflect.TypeOf(DeleteRetentionPolicy_Status{}), generators)

	return deleteRetentionPolicyStatusGenerator
}

// AddIndependentPropertyGeneratorsForDeleteRetentionPolicyStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDeleteRetentionPolicyStatus(gens map[string]gopter.Gen) {
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

func Test_LastAccessTimeTrackingPolicy_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LastAccessTimeTrackingPolicy via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLastAccessTimeTrackingPolicy, LastAccessTimeTrackingPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLastAccessTimeTrackingPolicy runs a test to see if a specific instance of LastAccessTimeTrackingPolicy round trips to JSON and back losslessly
func RunJSONSerializationTestForLastAccessTimeTrackingPolicy(subject LastAccessTimeTrackingPolicy) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual LastAccessTimeTrackingPolicy
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of LastAccessTimeTrackingPolicy instances for property testing -
//lazily instantiated by LastAccessTimeTrackingPolicyGenerator()
var lastAccessTimeTrackingPolicyGenerator gopter.Gen

// LastAccessTimeTrackingPolicyGenerator returns a generator of LastAccessTimeTrackingPolicy instances for property testing.
func LastAccessTimeTrackingPolicyGenerator() gopter.Gen {
	if lastAccessTimeTrackingPolicyGenerator != nil {
		return lastAccessTimeTrackingPolicyGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicy(generators)
	lastAccessTimeTrackingPolicyGenerator = gen.Struct(reflect.TypeOf(LastAccessTimeTrackingPolicy{}), generators)

	return lastAccessTimeTrackingPolicyGenerator
}

// AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicy is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicy(gens map[string]gopter.Gen) {
	gens["BlobType"] = gen.SliceOf(gen.AlphaString())
	gens["Enable"] = gen.PtrOf(gen.Bool())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["TrackingGranularityInDays"] = gen.PtrOf(gen.Int())
}

func Test_LastAccessTimeTrackingPolicy_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LastAccessTimeTrackingPolicy_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLastAccessTimeTrackingPolicyStatus, LastAccessTimeTrackingPolicyStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLastAccessTimeTrackingPolicyStatus runs a test to see if a specific instance of LastAccessTimeTrackingPolicy_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForLastAccessTimeTrackingPolicyStatus(subject LastAccessTimeTrackingPolicy_Status) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual LastAccessTimeTrackingPolicy_Status
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of LastAccessTimeTrackingPolicy_Status instances for property testing
//- lazily instantiated by LastAccessTimeTrackingPolicyStatusGenerator()
var lastAccessTimeTrackingPolicyStatusGenerator gopter.Gen

// LastAccessTimeTrackingPolicyStatusGenerator returns a generator of LastAccessTimeTrackingPolicy_Status instances for property testing.
func LastAccessTimeTrackingPolicyStatusGenerator() gopter.Gen {
	if lastAccessTimeTrackingPolicyStatusGenerator != nil {
		return lastAccessTimeTrackingPolicyStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicyStatus(generators)
	lastAccessTimeTrackingPolicyStatusGenerator = gen.Struct(reflect.TypeOf(LastAccessTimeTrackingPolicy_Status{}), generators)

	return lastAccessTimeTrackingPolicyStatusGenerator
}

// AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicyStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicyStatus(gens map[string]gopter.Gen) {
	gens["BlobType"] = gen.SliceOf(gen.AlphaString())
	gens["Enable"] = gen.PtrOf(gen.Bool())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["TrackingGranularityInDays"] = gen.PtrOf(gen.Int())
}

func Test_RestorePolicyProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RestorePolicyProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRestorePolicyProperties, RestorePolicyPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRestorePolicyProperties runs a test to see if a specific instance of RestorePolicyProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForRestorePolicyProperties(subject RestorePolicyProperties) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual RestorePolicyProperties
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of RestorePolicyProperties instances for property testing - lazily
//instantiated by RestorePolicyPropertiesGenerator()
var restorePolicyPropertiesGenerator gopter.Gen

// RestorePolicyPropertiesGenerator returns a generator of RestorePolicyProperties instances for property testing.
func RestorePolicyPropertiesGenerator() gopter.Gen {
	if restorePolicyPropertiesGenerator != nil {
		return restorePolicyPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRestorePolicyProperties(generators)
	restorePolicyPropertiesGenerator = gen.Struct(reflect.TypeOf(RestorePolicyProperties{}), generators)

	return restorePolicyPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForRestorePolicyProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRestorePolicyProperties(gens map[string]gopter.Gen) {
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

func Test_RestorePolicyProperties_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RestorePolicyProperties_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRestorePolicyPropertiesStatus, RestorePolicyPropertiesStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRestorePolicyPropertiesStatus runs a test to see if a specific instance of RestorePolicyProperties_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForRestorePolicyPropertiesStatus(subject RestorePolicyProperties_Status) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual RestorePolicyProperties_Status
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of RestorePolicyProperties_Status instances for property testing -
//lazily instantiated by RestorePolicyPropertiesStatusGenerator()
var restorePolicyPropertiesStatusGenerator gopter.Gen

// RestorePolicyPropertiesStatusGenerator returns a generator of RestorePolicyProperties_Status instances for property testing.
func RestorePolicyPropertiesStatusGenerator() gopter.Gen {
	if restorePolicyPropertiesStatusGenerator != nil {
		return restorePolicyPropertiesStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRestorePolicyPropertiesStatus(generators)
	restorePolicyPropertiesStatusGenerator = gen.Struct(reflect.TypeOf(RestorePolicyProperties_Status{}), generators)

	return restorePolicyPropertiesStatusGenerator
}

// AddIndependentPropertyGeneratorsForRestorePolicyPropertiesStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRestorePolicyPropertiesStatus(gens map[string]gopter.Gen) {
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["LastEnabledTime"] = gen.PtrOf(gen.AlphaString())
	gens["MinRestoreTime"] = gen.PtrOf(gen.AlphaString())
}

func Test_CorsRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorsRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorsRule, CorsRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorsRule runs a test to see if a specific instance of CorsRule round trips to JSON and back losslessly
func RunJSONSerializationTestForCorsRule(subject CorsRule) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual CorsRule
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of CorsRule instances for property testing - lazily instantiated by
//CorsRuleGenerator()
var corsRuleGenerator gopter.Gen

// CorsRuleGenerator returns a generator of CorsRule instances for property testing.
func CorsRuleGenerator() gopter.Gen {
	if corsRuleGenerator != nil {
		return corsRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCorsRule(generators)
	corsRuleGenerator = gen.Struct(reflect.TypeOf(CorsRule{}), generators)

	return corsRuleGenerator
}

// AddIndependentPropertyGeneratorsForCorsRule is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCorsRule(gens map[string]gopter.Gen) {
	gens["AllowedHeaders"] = gen.SliceOf(gen.AlphaString())
	gens["AllowedMethods"] = gen.SliceOf(gen.AlphaString())
	gens["AllowedOrigins"] = gen.SliceOf(gen.AlphaString())
	gens["ExposedHeaders"] = gen.SliceOf(gen.AlphaString())
	gens["MaxAgeInSeconds"] = gen.PtrOf(gen.Int())
}

func Test_CorsRule_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorsRule_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorsRuleStatus, CorsRuleStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorsRuleStatus runs a test to see if a specific instance of CorsRule_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForCorsRuleStatus(subject CorsRule_Status) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual CorsRule_Status
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of CorsRule_Status instances for property testing - lazily
//instantiated by CorsRuleStatusGenerator()
var corsRuleStatusGenerator gopter.Gen

// CorsRuleStatusGenerator returns a generator of CorsRule_Status instances for property testing.
func CorsRuleStatusGenerator() gopter.Gen {
	if corsRuleStatusGenerator != nil {
		return corsRuleStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCorsRuleStatus(generators)
	corsRuleStatusGenerator = gen.Struct(reflect.TypeOf(CorsRule_Status{}), generators)

	return corsRuleStatusGenerator
}

// AddIndependentPropertyGeneratorsForCorsRuleStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCorsRuleStatus(gens map[string]gopter.Gen) {
	gens["AllowedHeaders"] = gen.SliceOf(gen.AlphaString())
	gens["AllowedMethods"] = gen.SliceOf(gen.AlphaString())
	gens["AllowedOrigins"] = gen.SliceOf(gen.AlphaString())
	gens["ExposedHeaders"] = gen.SliceOf(gen.AlphaString())
	gens["MaxAgeInSeconds"] = gen.PtrOf(gen.Int())
}
