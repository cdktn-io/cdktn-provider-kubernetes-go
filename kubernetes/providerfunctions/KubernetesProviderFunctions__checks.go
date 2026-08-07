// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package providerfunctions

import (
	"fmt"
)

func (k *jsiiProxy_KubernetesProviderFunctions) validateManifestDecodeParameters(manifest *string) error {
	if manifest == nil {
		return fmt.Errorf("parameter manifest is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KubernetesProviderFunctions) validateManifestDecodeMultiParameters(manifest *string) error {
	if manifest == nil {
		return fmt.Errorf("parameter manifest is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KubernetesProviderFunctions) validateManifestEncodeParameters(manifest interface{}) error {
	if manifest == nil {
		return fmt.Errorf("parameter manifest is required, but nil was provided")
	}

	return nil
}

func validateNewKubernetesProviderFunctionsParameters(providerLocalName *string) error {
	if providerLocalName == nil {
		return fmt.Errorf("parameter providerLocalName is required, but nil was provided")
	}

	return nil
}

