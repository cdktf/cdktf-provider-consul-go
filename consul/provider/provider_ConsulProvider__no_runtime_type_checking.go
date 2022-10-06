//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConsulProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_ConsulProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateConsulProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ConsulProvider) validateSetHeaderParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConsulProvider) validateSetInsecureHttpsParameters(val interface{}) error {
	return nil
}

func validateNewConsulProviderParameters(scope constructs.Construct, id *string, config *ConsulProviderConfig) error {
	return nil
}

