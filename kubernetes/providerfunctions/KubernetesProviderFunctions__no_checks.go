// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package providerfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KubernetesProviderFunctions) validateManifestDecodeParameters(manifest *string) error {
	return nil
}

func (k *jsiiProxy_KubernetesProviderFunctions) validateManifestDecodeMultiParameters(manifest *string) error {
	return nil
}

func (k *jsiiProxy_KubernetesProviderFunctions) validateManifestEncodeParameters(manifest interface{}) error {
	return nil
}

func validateNewKubernetesProviderFunctionsParameters(providerLocalName *string) error {
	return nil
}

