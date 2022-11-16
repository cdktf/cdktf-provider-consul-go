//go:build no_runtime_type_checking

package node

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_Node) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (n *jsiiProxy_Node) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_Node) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_Node) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_Node) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_Node) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_Node) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_Node) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_Node) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_Node) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_Node) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_Node) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateNode_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNode_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateNode_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Node) validateSetAddressParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Node) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Node) validateSetDatacenterParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Node) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Node) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Node) validateSetMetaParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Node) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Node) validateSetPartitionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Node) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Node) validateSetTokenParameters(val *string) error {
	return nil
}

func validateNewNodeParameters(scope constructs.Construct, id *string, config *NodeConfig) error {
	return nil
}

