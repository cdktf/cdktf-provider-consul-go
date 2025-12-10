// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataconsulpeerings

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataConsulPeeringsPeersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataConsulPeeringsPeersList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataConsulPeeringsPeersList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataConsulPeeringsPeersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataConsulPeeringsPeersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataConsulPeeringsPeersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataConsulPeeringsPeersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

