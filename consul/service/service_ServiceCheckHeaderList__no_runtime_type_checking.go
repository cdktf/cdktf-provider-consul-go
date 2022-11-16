//go:build no_runtime_type_checking

package service

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceCheckHeaderList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceCheckHeaderList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceCheckHeaderList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServiceCheckHeaderList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceCheckHeaderList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceCheckHeaderList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceCheckHeaderListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

