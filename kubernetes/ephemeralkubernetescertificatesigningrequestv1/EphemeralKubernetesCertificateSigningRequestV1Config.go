// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralkubernetescertificatesigningrequestv1

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralKubernetesCertificateSigningRequestV1Config struct {
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformEphemeralResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.2.1/docs/ephemeral-resources/certificate_signing_request_v1#metadata EphemeralKubernetesCertificateSigningRequestV1#metadata}
	Metadata *EphemeralKubernetesCertificateSigningRequestV1Metadata `field:"required" json:"metadata" yaml:"metadata"`
	// Automatically approve the Certificate Signing Request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.2.1/docs/ephemeral-resources/certificate_signing_request_v1#auto_approve EphemeralKubernetesCertificateSigningRequestV1#auto_approve}
	AutoApprove interface{} `field:"optional" json:"autoApprove" yaml:"autoApprove"`
	// certificate is populated with an issued certificate by the signer after an Approved condition is present.
	//
	// This field is set via the /status subresource. Once populated, this field is immutable.
	//
	// If the certificate signing request is denied, a condition of type "Denied" is added and this field remains empty. If the signer cannot issue the certificate, a condition of type "Failed" is added and this field remains empty.
	//
	// Validation requirements:
	//  1. certificate must contain one or more PEM blocks.
	//  2. All PEM blocks must have the "CERTIFICATE" label, contain no headers, and the encoded data
	//   must be a BER-encoded ASN.1 Certificate structure as described in section 4 of RFC5280.
	//  3. Non-PEM content may appear before or after the "CERTIFICATE" PEM blocks and is unvalidated,
	//   to allow for explanatory text as described in section 5.2 of RFC7468.
	//
	// If more than one PEM block is present, and the definition of the requested spec.signerName does not indicate otherwise, the first block is the issued certificate, and subsequent blocks should be treated as intermediate certificates and presented in TLS handshakes.
	//
	// The certificate is encoded in PEM format.
	//
	// When serialized as JSON or YAML, the data is additionally base64-encoded, so it consists of:
	//
	//     base64(
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.2.1/docs/ephemeral-resources/certificate_signing_request_v1#certificate EphemeralKubernetesCertificateSigningRequestV1#certificate}
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.2.1/docs/ephemeral-resources/certificate_signing_request_v1#spec EphemeralKubernetesCertificateSigningRequestV1#spec}
	Spec *EphemeralKubernetesCertificateSigningRequestV1Spec `field:"optional" json:"spec" yaml:"spec"`
}

