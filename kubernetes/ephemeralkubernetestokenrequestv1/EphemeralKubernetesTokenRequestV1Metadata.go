// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralkubernetestokenrequestv1


type EphemeralKubernetesTokenRequestV1Metadata struct {
	// Name must be unique within a namespace.
	//
	// Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.2.1/docs/ephemeral-resources/token_request_v1#name EphemeralKubernetesTokenRequestV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Namespace defines the space within which each name must be unique.
	//
	// An empty namespace is equivalent to the "default" namespace, but "default" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.
	//
	// Must be a DNS_LABEL. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.2.1/docs/ephemeral-resources/token_request_v1#namespace EphemeralKubernetesTokenRequestV1#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

