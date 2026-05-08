// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package validatingadmissionpolicyv1


type ValidatingAdmissionPolicyV1SpecValidations struct {
	// Expression represents the expression which will be evaluated by CEL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#expression ValidatingAdmissionPolicyV1#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Message represents the message displayed when validation fails.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#message ValidatingAdmissionPolicyV1#message}
	Message *string `field:"required" json:"message" yaml:"message"`
	// Message Expression declares a CEL expression that evaluates to the validation failure message that is returned when this rule fails.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#message_expression ValidatingAdmissionPolicyV1#message_expression}
	MessageExpression *string `field:"optional" json:"messageExpression" yaml:"messageExpression"`
	// Reason represents a machine-readable description of why this validation failed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#reason ValidatingAdmissionPolicyV1#reason}
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
}

