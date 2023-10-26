// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package configentryservicedefaults

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewConfigEntryServiceDefaultsUpstreamConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

