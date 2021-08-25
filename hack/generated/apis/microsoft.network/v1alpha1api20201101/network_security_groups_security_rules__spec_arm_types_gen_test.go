// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

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

func Test_NetworkSecurityGroupsSecurityRules_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroupsSecurityRules_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupsSecurityRulesSpecARM, NetworkSecurityGroupsSecurityRulesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupsSecurityRulesSpecARM runs a test to see if a specific instance of NetworkSecurityGroupsSecurityRules_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupsSecurityRulesSpecARM(subject NetworkSecurityGroupsSecurityRules_SpecARM) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual NetworkSecurityGroupsSecurityRules_SpecARM
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

//Generator of NetworkSecurityGroupsSecurityRules_SpecARM instances for property
//testing - lazily instantiated by
//NetworkSecurityGroupsSecurityRulesSpecARMGenerator()
var networkSecurityGroupsSecurityRulesSpecARMGenerator gopter.Gen

// NetworkSecurityGroupsSecurityRulesSpecARMGenerator returns a generator of NetworkSecurityGroupsSecurityRules_SpecARM instances for property testing.
// We first initialize networkSecurityGroupsSecurityRulesSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroupsSecurityRulesSpecARMGenerator() gopter.Gen {
	if networkSecurityGroupsSecurityRulesSpecARMGenerator != nil {
		return networkSecurityGroupsSecurityRulesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRulesSpecARM(generators)
	networkSecurityGroupsSecurityRulesSpecARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupsSecurityRules_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRulesSpecARM(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRulesSpecARM(generators)
	networkSecurityGroupsSecurityRulesSpecARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupsSecurityRules_SpecARM{}), generators)

	return networkSecurityGroupsSecurityRulesSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRulesSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRulesSpecARM(gens map[string]gopter.Gen) {
	gens["APIVersion"] = gen.OneConstOf(NetworkSecurityGroupsSecurityRulesSpecAPIVersion20201101)
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.OneConstOf(NetworkSecurityGroupsSecurityRulesSpecTypeMicrosoftNetworkNetworkSecurityGroupsSecurityRules)
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRulesSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRulesSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = SecurityRulePropertiesFormatARMGenerator()
}

func Test_SecurityRulePropertiesFormatARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityRulePropertiesFormatARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityRulePropertiesFormatARM, SecurityRulePropertiesFormatARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityRulePropertiesFormatARM runs a test to see if a specific instance of SecurityRulePropertiesFormatARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityRulePropertiesFormatARM(subject SecurityRulePropertiesFormatARM) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual SecurityRulePropertiesFormatARM
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

//Generator of SecurityRulePropertiesFormatARM instances for property testing -
//lazily instantiated by SecurityRulePropertiesFormatARMGenerator()
var securityRulePropertiesFormatARMGenerator gopter.Gen

// SecurityRulePropertiesFormatARMGenerator returns a generator of SecurityRulePropertiesFormatARM instances for property testing.
// We first initialize securityRulePropertiesFormatARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SecurityRulePropertiesFormatARMGenerator() gopter.Gen {
	if securityRulePropertiesFormatARMGenerator != nil {
		return securityRulePropertiesFormatARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityRulePropertiesFormatARM(generators)
	securityRulePropertiesFormatARMGenerator = gen.Struct(reflect.TypeOf(SecurityRulePropertiesFormatARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityRulePropertiesFormatARM(generators)
	AddRelatedPropertyGeneratorsForSecurityRulePropertiesFormatARM(generators)
	securityRulePropertiesFormatARMGenerator = gen.Struct(reflect.TypeOf(SecurityRulePropertiesFormatARM{}), generators)

	return securityRulePropertiesFormatARMGenerator
}

// AddIndependentPropertyGeneratorsForSecurityRulePropertiesFormatARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecurityRulePropertiesFormatARM(gens map[string]gopter.Gen) {
	gens["Access"] = gen.OneConstOf(SecurityRulePropertiesFormatAccessAllow, SecurityRulePropertiesFormatAccessDeny)
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["DestinationPortRange"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationPortRanges"] = gen.SliceOf(gen.AlphaString())
	gens["Direction"] = gen.OneConstOf(SecurityRulePropertiesFormatDirectionInbound, SecurityRulePropertiesFormatDirectionOutbound)
	gens["Priority"] = gen.Int()
	gens["Protocol"] = gen.OneConstOf(SecurityRulePropertiesFormatProtocolAh, SecurityRulePropertiesFormatProtocolEsp, SecurityRulePropertiesFormatProtocolIcmp, SecurityRulePropertiesFormatProtocolStar, SecurityRulePropertiesFormatProtocolTcp, SecurityRulePropertiesFormatProtocolUdp)
	gens["SourceAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["SourceAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["SourcePortRange"] = gen.PtrOf(gen.AlphaString())
	gens["SourcePortRanges"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSecurityRulePropertiesFormatARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecurityRulePropertiesFormatARM(gens map[string]gopter.Gen) {
	gens["DestinationApplicationSecurityGroups"] = gen.SliceOf(SubResourceARMGenerator())
	gens["SourceApplicationSecurityGroups"] = gen.SliceOf(SubResourceARMGenerator())
}

func Test_SubResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubResourceARM, SubResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubResourceARM runs a test to see if a specific instance of SubResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSubResourceARM(subject SubResourceARM) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual SubResourceARM
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

//Generator of SubResourceARM instances for property testing - lazily instantiated
//by SubResourceARMGenerator()
var subResourceARMGenerator gopter.Gen

// SubResourceARMGenerator returns a generator of SubResourceARM instances for property testing.
func SubResourceARMGenerator() gopter.Gen {
	if subResourceARMGenerator != nil {
		return subResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubResourceARM(generators)
	subResourceARMGenerator = gen.Struct(reflect.TypeOf(SubResourceARM{}), generators)

	return subResourceARMGenerator
}

// AddIndependentPropertyGeneratorsForSubResourceARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubResourceARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.AlphaString()
}
