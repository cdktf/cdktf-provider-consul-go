// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package keyprefix

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KeyPrefixSubkeyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (k *jsiiProxy_KeyPrefixSubkeyList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KeyPrefixSubkeyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KeyPrefixSubkeyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KeyPrefixSubkeyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KeyPrefixSubkeyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KeyPrefixSubkeyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKeyPrefixSubkeyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

