// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package aclrole

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AclRoleServiceIdentitiesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AclRoleServiceIdentitiesList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AclRoleServiceIdentitiesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AclRoleServiceIdentitiesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AclRoleServiceIdentitiesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AclRoleServiceIdentitiesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AclRoleServiceIdentitiesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAclRoleServiceIdentitiesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

