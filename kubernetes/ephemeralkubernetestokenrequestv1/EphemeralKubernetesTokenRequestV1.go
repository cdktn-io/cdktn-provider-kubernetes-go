// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralkubernetestokenrequestv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-kubernetes-go/kubernetes/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-kubernetes-go/kubernetes/v17/ephemeralkubernetestokenrequestv1/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.2.1/docs/ephemeral-resources/token_request_v1 kubernetes_token_request_v1}.
type EphemeralKubernetesTokenRequestV1 interface {
	cdktn.TerraformEphemeralResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExpirationTimestamp() *string
	SetExpirationTimestamp(val *string)
	ExpirationTimestampInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformEphemeralResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformEphemeralResourceLifecycle)
	Metadata() EphemeralKubernetesTokenRequestV1MetadataOutputReference
	MetadataInput() interface{}
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Spec() EphemeralKubernetesTokenRequestV1SpecOutputReference
	SpecInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutMetadata(value *EphemeralKubernetesTokenRequestV1Metadata)
	PutSpec(value *EphemeralKubernetesTokenRequestV1Spec)
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	ResetExpirationTimestamp()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSpec()
	ResetToken()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this ephemeral resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for EphemeralKubernetesTokenRequestV1
type jsiiProxy_EphemeralKubernetesTokenRequestV1 struct {
	internal.Type__cdktnTerraformEphemeralResource
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1) ExpirationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1) ExpirationTimestampInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationTimestampInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1) Lifecycle() *cdktn.TerraformEphemeralResourceLifecycle {
	var returns *cdktn.TerraformEphemeralResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1) Metadata() EphemeralKubernetesTokenRequestV1MetadataOutputReference {
	var returns EphemeralKubernetesTokenRequestV1MetadataOutputReference
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1) MetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1) Spec() EphemeralKubernetesTokenRequestV1SpecOutputReference {
	var returns EphemeralKubernetesTokenRequestV1SpecOutputReference
	_jsii_.Get(
		j,
		"spec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1) SpecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"specInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.2.1/docs/ephemeral-resources/token_request_v1 kubernetes_token_request_v1} Ephemeral Resource.
func NewEphemeralKubernetesTokenRequestV1(scope constructs.Construct, id *string, config *EphemeralKubernetesTokenRequestV1Config) EphemeralKubernetesTokenRequestV1 {
	_init_.Initialize()

	if err := validateNewEphemeralKubernetesTokenRequestV1Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EphemeralKubernetesTokenRequestV1{}

	_jsii_.Create(
		"@cdktn/provider-kubernetes.ephemeralKubernetesTokenRequestV1.EphemeralKubernetesTokenRequestV1",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.2.1/docs/ephemeral-resources/token_request_v1 kubernetes_token_request_v1} Ephemeral Resource.
func NewEphemeralKubernetesTokenRequestV1_Override(e EphemeralKubernetesTokenRequestV1, scope constructs.Construct, id *string, config *EphemeralKubernetesTokenRequestV1Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-kubernetes.ephemeralKubernetesTokenRequestV1.EphemeralKubernetesTokenRequestV1",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1)SetExpirationTimestamp(val *string) {
	if err := j.validateSetExpirationTimestampParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expirationTimestamp",
		val,
	)
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1)SetLifecycle(val *cdktn.TerraformEphemeralResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EphemeralKubernetesTokenRequestV1)SetToken(val *string) {
	if err := j.validateSetTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func EphemeralKubernetesTokenRequestV1_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEphemeralKubernetesTokenRequestV1_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-kubernetes.ephemeralKubernetesTokenRequestV1.EphemeralKubernetesTokenRequestV1",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EphemeralKubernetesTokenRequestV1_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEphemeralKubernetesTokenRequestV1_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-kubernetes.ephemeralKubernetesTokenRequestV1.EphemeralKubernetesTokenRequestV1",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EphemeralKubernetesTokenRequestV1_IsTerraformEphemeralResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEphemeralKubernetesTokenRequestV1_IsTerraformEphemeralResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-kubernetes.ephemeralKubernetesTokenRequestV1.EphemeralKubernetesTokenRequestV1",
		"isTerraformEphemeralResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EphemeralKubernetesTokenRequestV1_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-kubernetes.ephemeralKubernetesTokenRequestV1.EphemeralKubernetesTokenRequestV1",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EphemeralKubernetesTokenRequestV1) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EphemeralKubernetesTokenRequestV1) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralKubernetesTokenRequestV1) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralKubernetesTokenRequestV1) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralKubernetesTokenRequestV1) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralKubernetesTokenRequestV1) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralKubernetesTokenRequestV1) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralKubernetesTokenRequestV1) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralKubernetesTokenRequestV1) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralKubernetesTokenRequestV1) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralKubernetesTokenRequestV1) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralKubernetesTokenRequestV1) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EphemeralKubernetesTokenRequestV1) PutMetadata(value *EphemeralKubernetesTokenRequestV1Metadata) {
	if err := e.validatePutMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMetadata",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EphemeralKubernetesTokenRequestV1) PutSpec(value *EphemeralKubernetesTokenRequestV1Spec) {
	if err := e.validatePutSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSpec",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EphemeralKubernetesTokenRequestV1) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := e.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (e *jsiiProxy_EphemeralKubernetesTokenRequestV1) ResetExpirationTimestamp() {
	_jsii_.InvokeVoid(
		e,
		"resetExpirationTimestamp",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralKubernetesTokenRequestV1) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralKubernetesTokenRequestV1) ResetSpec() {
	_jsii_.InvokeVoid(
		e,
		"resetSpec",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralKubernetesTokenRequestV1) ResetToken() {
	_jsii_.InvokeVoid(
		e,
		"resetToken",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralKubernetesTokenRequestV1) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralKubernetesTokenRequestV1) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralKubernetesTokenRequestV1) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralKubernetesTokenRequestV1) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralKubernetesTokenRequestV1) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralKubernetesTokenRequestV1) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralKubernetesTokenRequestV1) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		e,
		"with",
		args,
		&returns,
	)

	return returns
}

