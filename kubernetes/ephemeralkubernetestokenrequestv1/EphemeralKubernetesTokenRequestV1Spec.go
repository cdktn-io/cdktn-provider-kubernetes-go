// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralkubernetestokenrequestv1


type EphemeralKubernetesTokenRequestV1Spec struct {
	// Audiences are the intendend audiences of the token.
	//
	// A recipient of a token must identify themself with an identifier in the list of audiences of the token, and otherwise should reject the token. A token issued for multiple audiences may be used to authenticate against any of the audiences listed but implies a high degree of trust between the target audiences.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.2.1/docs/ephemeral-resources/token_request_v1#audiences EphemeralKubernetesTokenRequestV1#audiences}
	Audiences *[]*string `field:"optional" json:"audiences" yaml:"audiences"`
	// bound_object_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.2.1/docs/ephemeral-resources/token_request_v1#bound_object_ref EphemeralKubernetesTokenRequestV1#bound_object_ref}
	BoundObjectRef *EphemeralKubernetesTokenRequestV1SpecBoundObjectRef `field:"optional" json:"boundObjectRef" yaml:"boundObjectRef"`
	// ExpirationSeconds is the requested duration of validity of the request.
	//
	// The token issuer may return a token with a different validity duration so a client needs to check the 'expiration' field in a response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.2.1/docs/ephemeral-resources/token_request_v1#expiration_seconds EphemeralKubernetesTokenRequestV1#expiration_seconds}
	ExpirationSeconds *float64 `field:"optional" json:"expirationSeconds" yaml:"expirationSeconds"`
}

