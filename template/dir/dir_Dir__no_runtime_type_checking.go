//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package dir

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_Dir) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (d *jsiiProxy_Dir) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Dir) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Dir) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Dir) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Dir) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Dir) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Dir) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Dir) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Dir) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Dir) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Dir) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateDir_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Dir) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Dir) validateSetDestinationDirParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Dir) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Dir) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Dir) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Dir) validateSetSourceDirParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Dir) validateSetVarsParameters(val *map[string]*string) error {
	return nil
}

func validateNewDirParameters(scope constructs.Construct, id *string, config *DirConfig) error {
	return nil
}

