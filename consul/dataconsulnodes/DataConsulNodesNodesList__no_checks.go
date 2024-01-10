// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataconsulnodes

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataConsulNodesNodesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataConsulNodesNodesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataConsulNodesNodesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataConsulNodesNodesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataConsulNodesNodesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataConsulNodesNodesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataConsulNodesNodesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

