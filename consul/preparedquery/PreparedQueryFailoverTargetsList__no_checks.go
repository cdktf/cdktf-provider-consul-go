//go:build no_runtime_type_checking

package preparedquery

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PreparedQueryFailoverTargetsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PreparedQueryFailoverTargetsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PreparedQueryFailoverTargetsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PreparedQueryFailoverTargetsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PreparedQueryFailoverTargetsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PreparedQueryFailoverTargetsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPreparedQueryFailoverTargetsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

