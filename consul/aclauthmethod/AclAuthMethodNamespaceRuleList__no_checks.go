// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package aclauthmethod

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AclAuthMethodNamespaceRuleList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AclAuthMethodNamespaceRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AclAuthMethodNamespaceRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AclAuthMethodNamespaceRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AclAuthMethodNamespaceRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AclAuthMethodNamespaceRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AclAuthMethodNamespaceRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAclAuthMethodNamespaceRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

