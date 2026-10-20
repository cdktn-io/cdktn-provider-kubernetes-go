// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package rolev1

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RoleV1MetadataList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RoleV1MetadataList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RoleV1MetadataList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RoleV1MetadataList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RoleV1MetadataList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RoleV1MetadataList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RoleV1MetadataList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRoleV1MetadataListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

