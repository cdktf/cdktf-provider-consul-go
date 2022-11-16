//go:build no_runtime_type_checking

package dataconsulkeys

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataConsulKeysKeyList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataConsulKeysKeyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataConsulKeysKeyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataConsulKeysKeyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataConsulKeysKeyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataConsulKeysKeyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataConsulKeysKeyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

