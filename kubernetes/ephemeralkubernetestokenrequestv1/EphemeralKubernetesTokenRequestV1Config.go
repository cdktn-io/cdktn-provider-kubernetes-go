// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralkubernetestokenrequestv1

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralKubernetesTokenRequestV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.2.1/docs/ephemeral-resources/token_request_v1#metadata EphemeralKubernetesTokenRequestV1#metadata}
	Metadata *EphemeralKubernetesTokenRequestV1Metadata `field:"required" json:"metadata" yaml:"metadata"`
	// ExpirationTimestamp is the time of expiration of the returned token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.2.1/docs/ephemeral-resources/token_request_v1#expiration_timestamp EphemeralKubernetesTokenRequestV1#expiration_timestamp}
	ExpirationTimestamp *string `field:"optional" json:"expirationTimestamp" yaml:"expirationTimestamp"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.2.1/docs/ephemeral-resources/token_request_v1#spec EphemeralKubernetesTokenRequestV1#spec}
	Spec *EphemeralKubernetesTokenRequestV1Spec `field:"optional" json:"spec" yaml:"spec"`
	// Token is the opaque bearer token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.2.1/docs/ephemeral-resources/token_request_v1#token EphemeralKubernetesTokenRequestV1#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
}

