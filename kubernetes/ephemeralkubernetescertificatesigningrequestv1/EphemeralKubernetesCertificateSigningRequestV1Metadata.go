// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralkubernetescertificatesigningrequestv1


type EphemeralKubernetesCertificateSigningRequestV1Metadata struct {
	// Name must be unique within a namespace.
	//
	// Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.2.1/docs/ephemeral-resources/certificate_signing_request_v1#name EphemeralKubernetesCertificateSigningRequestV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

