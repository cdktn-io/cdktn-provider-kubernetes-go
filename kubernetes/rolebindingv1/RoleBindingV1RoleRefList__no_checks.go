// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package rolebindingv1

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RoleBindingV1RoleRefList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RoleBindingV1RoleRefList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RoleBindingV1RoleRefList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RoleBindingV1RoleRefList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RoleBindingV1RoleRefList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RoleBindingV1RoleRefList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RoleBindingV1RoleRefList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRoleBindingV1RoleRefListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

