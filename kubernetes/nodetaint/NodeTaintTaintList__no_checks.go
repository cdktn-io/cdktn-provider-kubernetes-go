// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package nodetaint

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NodeTaintTaintList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NodeTaintTaintList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NodeTaintTaintList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NodeTaintTaintList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NodeTaintTaintList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NodeTaintTaintList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NodeTaintTaintList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNodeTaintTaintListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

