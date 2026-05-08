// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package validatingadmissionpolicyv1


type ValidatingAdmissionPolicyV1Spec struct {
	// auditAnnotations contains CEL expressions which are used to produce audit annotations for the audit event of the API request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#audit_annotations ValidatingAdmissionPolicyV1#audit_annotations}
	AuditAnnotations interface{} `field:"required" json:"auditAnnotations" yaml:"auditAnnotations"`
	// failurePolicy defines how to handle failures for the admission policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#failure_policy ValidatingAdmissionPolicyV1#failure_policy}
	FailurePolicy *string `field:"required" json:"failurePolicy" yaml:"failurePolicy"`
	// MatchConstraints specifies what resources this policy is designed to validate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#match_constraints ValidatingAdmissionPolicyV1#match_constraints}
	MatchConstraints *ValidatingAdmissionPolicyV1SpecMatchConstraints `field:"required" json:"matchConstraints" yaml:"matchConstraints"`
	// MatchConditions is a list of conditions that must be met for a request to be validated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#match_conditions ValidatingAdmissionPolicyV1#match_conditions}
	MatchConditions interface{} `field:"optional" json:"matchConditions" yaml:"matchConditions"`
	// ParamKind specifies the kind of resources used to parameterize this policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#param_kind ValidatingAdmissionPolicyV1#param_kind}
	ParamKind *ValidatingAdmissionPolicyV1SpecParamKind `field:"optional" json:"paramKind" yaml:"paramKind"`
	// Validations contain CEL expressions which is used to apply the validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#validations ValidatingAdmissionPolicyV1#validations}
	Validations interface{} `field:"optional" json:"validations" yaml:"validations"`
	// Variables contain definitions of variables that can be used in composition of other expressions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#variables ValidatingAdmissionPolicyV1#variables}
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
}

