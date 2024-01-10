// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataconsulservice

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataConsulServiceQueryOptionsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataConsulServiceQueryOptionsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataConsulServiceQueryOptionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataConsulServiceQueryOptionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataConsulServiceQueryOptionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataConsulServiceQueryOptionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataConsulServiceQueryOptionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataConsulServiceQueryOptionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

