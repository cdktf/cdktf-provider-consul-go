//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package service

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Service) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_Service) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Service) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Service) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Service) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Service) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Service) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Service) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Service) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Service) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Service) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Service) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_Service) validatePutCheckParameters(value interface{}) error {
	return nil
}

func validateService_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetAddressParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetDatacenterParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetEnableTagOverrideParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetExternalParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetMetaParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetNamespaceParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetNodeAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetPartitionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetPortParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetServiceIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetTagsParameters(val *[]*string) error {
	return nil
}

func validateNewServiceParameters(scope constructs.Construct, id *string, config *ServiceConfig) error {
	return nil
}

