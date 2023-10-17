// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package intention

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_Intention) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (i *jsiiProxy_Intention) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (i *jsiiProxy_Intention) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Intention) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Intention) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Intention) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Intention) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Intention) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Intention) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Intention) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Intention) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Intention) validateImportFromParameters(id *string) error {
	return nil
}

func (i *jsiiProxy_Intention) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Intention) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (i *jsiiProxy_Intention) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateIntention_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateIntention_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIntention_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateIntention_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Intention) validateSetActionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Intention) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Intention) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Intention) validateSetDatacenterParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Intention) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Intention) validateSetDestinationNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Intention) validateSetDestinationNamespaceParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Intention) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Intention) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Intention) validateSetMetaParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Intention) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Intention) validateSetSourceNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Intention) validateSetSourceNamespaceParameters(val *string) error {
	return nil
}

func validateNewIntentionParameters(scope constructs.Construct, id *string, config *IntentionConfig) error {
	return nil
}

