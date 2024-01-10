// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cloudinitconfig

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudinitConfigPartList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CloudinitConfigPartList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudinitConfigPartList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudinitConfigPartList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CloudinitConfigPartList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudinitConfigPartList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudinitConfigPartList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudinitConfigPartListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

