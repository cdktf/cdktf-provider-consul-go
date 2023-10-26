// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package configentryservicerouter

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewConfigEntryServiceRouterRoutesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

