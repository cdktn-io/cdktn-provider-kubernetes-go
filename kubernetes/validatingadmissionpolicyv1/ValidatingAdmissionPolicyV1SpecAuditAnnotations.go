// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package validatingadmissionpolicyv1


type ValidatingAdmissionPolicyV1SpecAuditAnnotations struct {
	// key specifies the audit annotation key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#key ValidatingAdmissionPolicyV1#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// valueExpression represents the expression which is evaluated by CEL to produce an audit annotation value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#value_expression ValidatingAdmissionPolicyV1#value_expression}
	ValueExpression *string `field:"required" json:"valueExpression" yaml:"valueExpression"`
}

