// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package validatingadmissionpolicyv1


type ValidatingAdmissionPolicyV1SpecMatchConditions struct {
	// Expression represents the expression which will be evaluated by CEL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#expression ValidatingAdmissionPolicyV1#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Name is an identifier for this match condition, used for strategic merging of MatchConditions, as well as providing an identifier for logging purposes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#name ValidatingAdmissionPolicyV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

