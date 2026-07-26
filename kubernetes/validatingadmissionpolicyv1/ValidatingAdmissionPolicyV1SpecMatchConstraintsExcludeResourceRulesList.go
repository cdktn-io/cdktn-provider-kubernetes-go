// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package validatingadmissionpolicyv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-kubernetes-go/kubernetes/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-kubernetes-go/kubernetes/v16/validatingadmissionpolicyv1/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList
type jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList {
	_init_.Initialize()

	if err := validateNewValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList{}

	_jsii_.Create(
		"@cdktn/provider-kubernetes.validatingAdmissionPolicyV1.ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList_Override(v ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-kubernetes.validatingAdmissionPolicyV1.ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		v,
	)
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := v.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		v,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList) Get(index *float64) ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesOutputReference {
	if err := v.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesOutputReference

	_jsii_.Invoke(
		v,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (v *jsiiProxy_ValidatingAdmissionPolicyV1SpecMatchConstraintsExcludeResourceRulesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

