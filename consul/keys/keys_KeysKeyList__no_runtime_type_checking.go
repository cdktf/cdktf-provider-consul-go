//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package keys

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KeysKeyList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KeysKeyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KeysKeyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KeysKeyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KeysKeyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KeysKeyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKeysKeyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

