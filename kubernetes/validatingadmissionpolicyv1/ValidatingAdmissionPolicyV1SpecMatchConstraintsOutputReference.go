// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package validatingadmissionpolicyv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-kubernetes-go/kubernetes/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-kubernetes-go/kubernetes/v16/validatingadmissionpolicyv1/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference interface {
	cdktn.ComplexObject
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
	ExcludeResourceRules() ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList
	ExcludeResourceRulesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MatchPolicy() *string
	SetMatchPolicy(val *string)
	MatchPolicyInput() *string
	NamespaceSelector() ValidatingAdmissionPolicyV1SpecMatchConstraintsNamespaceSelectorOutputReference
	NamespaceSelectorInput() interface{}
	ObjectSelector() ValidatingAdmissionPolicyV1SpecMatchConstraintsObjectSelectorOutputReference
	ObjectSelectorInput() interface{}
	ResourceRules() ValidatingAdmissionPolicyV1SpecMatchConstraintsResourceRulesList
	ResourceRulesInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	PutExcludeResourceRules(value interface{})
	PutNamespaceSelector(value *ValidatingAdmissionPolicyV1SpecMatchConstraintsNamespaceSelector)
	PutObjectSelector(value *ValidatingAdmissionPolicyV1SpecMatchConstraintsObjectSelector)
	PutResourceRules(value interface{})
	ResetExcludeResourceRules()
	ResetMatchPolicy()
	ResetNamespaceSelector()
	ResetObjectSelector()
	ResetResourceRules()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference
type jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) ExcludeResourceRules() ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList {
	var returns ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList
	_jsii_.Get(
		j,
		"excludeResourceRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) ExcludeResourceRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeResourceRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) MatchPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) MatchPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) NamespaceSelector() ValidatingAdmissionPolicyV1SpecMatchConstraintsNamespaceSelectorOutputReference {
	var returns ValidatingAdmissionPolicyV1SpecMatchConstraintsNamespaceSelectorOutputReference
	_jsii_.Get(
		j,
		"namespaceSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) NamespaceSelectorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"namespaceSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) ObjectSelector() ValidatingAdmissionPolicyV1SpecMatchConstraintsObjectSelectorOutputReference {
	var returns ValidatingAdmissionPolicyV1SpecMatchConstraintsObjectSelectorOutputReference
	_jsii_.Get(
		j,
		"objectSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) ObjectSelectorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"objectSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) ResourceRules() ValidatingAdmissionPolicyV1SpecMatchConstraintsResourceRulesList {
	var returns ValidatingAdmissionPolicyV1SpecMatchConstraintsResourceRulesList
	_jsii_.Get(
		j,
		"resourceRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) ResourceRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference {
	_init_.Initialize()

	if err := validateNewValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-kubernetes.validatingAdmissionPolicyV1.ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference_Override(v ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-kubernetes.validatingAdmissionPolicyV1.ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference)SetMatchPolicy(val *string) {
	if err := j.validateSetMatchPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchPolicy",
		val,
	)
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) PutExcludeResourceRules(value interface{}) {
	if err := v.validatePutExcludeResourceRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putExcludeResourceRules",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) PutNamespaceSelector(value *ValidatingAdmissionPolicyV1SpecMatchConstraintsNamespaceSelector) {
	if err := v.validatePutNamespaceSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putNamespaceSelector",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) PutObjectSelector(value *ValidatingAdmissionPolicyV1SpecMatchConstraintsObjectSelector) {
	if err := v.validatePutObjectSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putObjectSelector",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) PutResourceRules(value interface{}) {
	if err := v.validatePutResourceRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putResourceRules",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) ResetExcludeResourceRules() {
	_jsii_.InvokeVoid(
		v,
		"resetExcludeResourceRules",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) ResetMatchPolicy() {
	_jsii_.InvokeVoid(
		v,
		"resetMatchPolicy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) ResetNamespaceSelector() {
	_jsii_.InvokeVoid(
		v,
		"resetNamespaceSelector",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) ResetObjectSelector() {
	_jsii_.InvokeVoid(
		v,
		"resetObjectSelector",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) ResetResourceRules() {
	_jsii_.InvokeVoid(
		v,
		"resetResourceRules",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

