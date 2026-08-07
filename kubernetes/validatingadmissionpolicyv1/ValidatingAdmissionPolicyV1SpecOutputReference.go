// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package validatingadmissionpolicyv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-kubernetes-go/kubernetes/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-kubernetes-go/kubernetes/v17/validatingadmissionpolicyv1/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ValidatingAdmissionPolicyV1SpecOutputReference interface {
	cdktn.ComplexObject
	AuditAnnotations() ValidatingAdmissionPolicyV1SpecAuditAnnotationsList
	AuditAnnotationsInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FailurePolicy() *string
	SetFailurePolicy(val *string)
	FailurePolicyInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MatchConditions() ValidatingAdmissionPolicyV1SpecMatchConditionsList
	MatchConditionsInput() interface{}
	MatchConstraints() ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference
	MatchConstraintsInput() interface{}
	ParamKind() ValidatingAdmissionPolicyV1SpecParamKindOutputReference
	ParamKindInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Validations() ValidatingAdmissionPolicyV1SpecValidationsList
	ValidationsInput() interface{}
	Variables() ValidatingAdmissionPolicyV1SpecVariablesList
	VariablesInput() interface{}
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutAuditAnnotations(value interface{})
	PutMatchConditions(value interface{})
	PutMatchConstraints(value *ValidatingAdmissionPolicyV1SpecMatchConstraints)
	PutParamKind(value *ValidatingAdmissionPolicyV1SpecParamKind)
	PutValidations(value interface{})
	PutVariables(value interface{})
	ResetMatchConditions()
	ResetParamKind()
	ResetValidations()
	ResetVariables()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ValidatingAdmissionPolicyV1SpecOutputReference
type jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) AuditAnnotations() ValidatingAdmissionPolicyV1SpecAuditAnnotationsList {
	var returns ValidatingAdmissionPolicyV1SpecAuditAnnotationsList
	_jsii_.Get(
		j,
		"auditAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) AuditAnnotationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auditAnnotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) FailurePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failurePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) FailurePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failurePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) MatchConditions() ValidatingAdmissionPolicyV1SpecMatchConditionsList {
	var returns ValidatingAdmissionPolicyV1SpecMatchConditionsList
	_jsii_.Get(
		j,
		"matchConditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) MatchConditionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchConditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) MatchConstraints() ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference {
	var returns ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference
	_jsii_.Get(
		j,
		"matchConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) MatchConstraintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchConstraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) ParamKind() ValidatingAdmissionPolicyV1SpecParamKindOutputReference {
	var returns ValidatingAdmissionPolicyV1SpecParamKindOutputReference
	_jsii_.Get(
		j,
		"paramKind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) ParamKindInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"paramKindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) Validations() ValidatingAdmissionPolicyV1SpecValidationsList {
	var returns ValidatingAdmissionPolicyV1SpecValidationsList
	_jsii_.Get(
		j,
		"validations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) ValidationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) Variables() ValidatingAdmissionPolicyV1SpecVariablesList {
	var returns ValidatingAdmissionPolicyV1SpecVariablesList
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) VariablesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"variablesInput",
		&returns,
	)
	return returns
}


func NewValidatingAdmissionPolicyV1SpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ValidatingAdmissionPolicyV1SpecOutputReference {
	_init_.Initialize()

	if err := validateNewValidatingAdmissionPolicyV1SpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-kubernetes.validatingAdmissionPolicyV1.ValidatingAdmissionPolicyV1SpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewValidatingAdmissionPolicyV1SpecOutputReference_Override(v ValidatingAdmissionPolicyV1SpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-kubernetes.validatingAdmissionPolicyV1.ValidatingAdmissionPolicyV1SpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference)SetFailurePolicy(val *string) {
	if err := j.validateSetFailurePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failurePolicy",
		val,
	)
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) PutAuditAnnotations(value interface{}) {
	if err := v.validatePutAuditAnnotationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putAuditAnnotations",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) PutMatchConditions(value interface{}) {
	if err := v.validatePutMatchConditionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putMatchConditions",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) PutMatchConstraints(value *ValidatingAdmissionPolicyV1SpecMatchConstraints) {
	if err := v.validatePutMatchConstraintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putMatchConstraints",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) PutParamKind(value *ValidatingAdmissionPolicyV1SpecParamKind) {
	if err := v.validatePutParamKindParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putParamKind",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) PutValidations(value interface{}) {
	if err := v.validatePutValidationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putValidations",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) PutVariables(value interface{}) {
	if err := v.validatePutVariablesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putVariables",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) ResetMatchConditions() {
	_jsii_.InvokeVoid(
		v,
		"resetMatchConditions",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) ResetParamKind() {
	_jsii_.InvokeVoid(
		v,
		"resetParamKind",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) ResetValidations() {
	_jsii_.InvokeVoid(
		v,
		"resetValidations",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) ResetVariables() {
	_jsii_.InvokeVoid(
		v,
		"resetVariables",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := v.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

