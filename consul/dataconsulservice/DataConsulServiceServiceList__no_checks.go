// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataconsulservice

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataConsulServiceServiceList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataConsulServiceServiceList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataConsulServiceServiceList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataConsulServiceServiceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataConsulServiceServiceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataConsulServiceServiceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataConsulServiceServiceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

