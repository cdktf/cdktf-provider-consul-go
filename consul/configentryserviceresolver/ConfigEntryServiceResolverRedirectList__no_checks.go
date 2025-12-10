// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package configentryserviceresolver

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConfigEntryServiceResolverRedirectList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ConfigEntryServiceResolverRedirectList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ConfigEntryServiceResolverRedirectList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ConfigEntryServiceResolverRedirectList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConfigEntryServiceResolverRedirectList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConfigEntryServiceResolverRedirectList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ConfigEntryServiceResolverRedirectList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewConfigEntryServiceResolverRedirectListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

