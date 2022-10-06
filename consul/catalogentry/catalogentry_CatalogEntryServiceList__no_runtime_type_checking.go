//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package catalogentry

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CatalogEntryServiceList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CatalogEntryServiceList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CatalogEntryServiceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CatalogEntryServiceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CatalogEntryServiceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CatalogEntryServiceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCatalogEntryServiceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

