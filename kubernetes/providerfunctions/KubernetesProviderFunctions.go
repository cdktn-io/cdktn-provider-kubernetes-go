// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package providerfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-kubernetes-go/kubernetes/v17/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Provider-defined functions of the kubernetes provider.
type KubernetesProviderFunctions interface {
	// Given a YAML text containing a Kubernetes manifest, will decode and return an object representation of that resource.
	ManifestDecode(manifest *string) cdktn.IResolvable
	// Given a YAML text containing a Kubernetes manifest with multiple resources, will decode the manifest and return a tuple of object representations for each resource.
	ManifestDecodeMulti(manifest *string) cdktn.IResolvable
	// Given an object representation of a Kubernetes manifest, will encode and return a YAML string for that resource.
	ManifestEncode(manifest interface{}) *string
}

// The jsii proxy struct for KubernetesProviderFunctions
type jsiiProxy_KubernetesProviderFunctions struct {
	_ byte // padding
}

func NewKubernetesProviderFunctions(providerLocalName *string) KubernetesProviderFunctions {
	_init_.Initialize()

	if err := validateNewKubernetesProviderFunctionsParameters(providerLocalName); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesProviderFunctions{}

	_jsii_.Create(
		"@cdktn/provider-kubernetes.providerFunctions.KubernetesProviderFunctions",
		[]interface{}{providerLocalName},
		&j,
	)

	return &j
}

func NewKubernetesProviderFunctions_Override(k KubernetesProviderFunctions, providerLocalName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-kubernetes.providerFunctions.KubernetesProviderFunctions",
		[]interface{}{providerLocalName},
		k,
	)
}

func (k *jsiiProxy_KubernetesProviderFunctions) ManifestDecode(manifest *string) cdktn.IResolvable {
	if err := k.validateManifestDecodeParameters(manifest); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"manifestDecode",
		[]interface{}{manifest},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesProviderFunctions) ManifestDecodeMulti(manifest *string) cdktn.IResolvable {
	if err := k.validateManifestDecodeMultiParameters(manifest); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"manifestDecodeMulti",
		[]interface{}{manifest},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesProviderFunctions) ManifestEncode(manifest interface{}) *string {
	if err := k.validateManifestEncodeParameters(manifest); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"manifestEncode",
		[]interface{}{manifest},
		&returns,
	)

	return returns
}

