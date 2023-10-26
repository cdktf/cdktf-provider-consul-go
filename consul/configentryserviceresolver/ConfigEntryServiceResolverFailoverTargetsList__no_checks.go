// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package configentryserviceresolver

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConfigEntryServiceResolverFailoverTargetsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ConfigEntryServiceResolverFailoverTargetsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverTargetsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverTargetsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverTargetsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverTargetsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewConfigEntryServiceResolverFailoverTargetsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

