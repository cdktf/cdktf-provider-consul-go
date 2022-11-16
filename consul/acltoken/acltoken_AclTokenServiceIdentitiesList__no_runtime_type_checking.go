//go:build no_runtime_type_checking

package acltoken

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AclTokenServiceIdentitiesList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AclTokenServiceIdentitiesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AclTokenServiceIdentitiesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AclTokenServiceIdentitiesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AclTokenServiceIdentitiesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AclTokenServiceIdentitiesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAclTokenServiceIdentitiesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

