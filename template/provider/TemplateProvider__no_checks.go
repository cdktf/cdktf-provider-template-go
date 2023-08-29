// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TemplateProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (t *jsiiProxy_TemplateProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateTemplateProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTemplateProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateTemplateProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func validateNewTemplateProviderParameters(scope constructs.Construct, id *string, config *TemplateProviderConfig) error {
	return nil
}

