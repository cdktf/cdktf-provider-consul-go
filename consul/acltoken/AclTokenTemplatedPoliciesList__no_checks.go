// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package acltoken

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AclTokenTemplatedPoliciesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AclTokenTemplatedPoliciesList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AclTokenTemplatedPoliciesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AclTokenTemplatedPoliciesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AclTokenTemplatedPoliciesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AclTokenTemplatedPoliciesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AclTokenTemplatedPoliciesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAclTokenTemplatedPoliciesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

