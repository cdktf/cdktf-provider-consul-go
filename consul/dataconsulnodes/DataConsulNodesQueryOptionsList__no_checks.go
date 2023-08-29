// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataconsulnodes

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataConsulNodesQueryOptionsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataConsulNodesQueryOptionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataConsulNodesQueryOptionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

