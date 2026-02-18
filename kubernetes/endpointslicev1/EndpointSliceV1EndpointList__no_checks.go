// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package endpointslicev1

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EndpointSliceV1EndpointList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EndpointSliceV1EndpointList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EndpointSliceV1EndpointList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EndpointSliceV1EndpointList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EndpointSliceV1EndpointList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EndpointSliceV1EndpointList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EndpointSliceV1EndpointList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEndpointSliceV1EndpointListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

