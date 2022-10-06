//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package aclrole

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AclRoleNodeIdentitiesList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AclRoleNodeIdentitiesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AclRoleNodeIdentitiesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AclRoleNodeIdentitiesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AclRoleNodeIdentitiesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AclRoleNodeIdentitiesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAclRoleNodeIdentitiesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

