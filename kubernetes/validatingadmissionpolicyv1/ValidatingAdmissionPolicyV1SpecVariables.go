// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package validatingadmissionpolicyv1


type ValidatingAdmissionPolicyV1SpecVariables struct {
	// Expression is the expression that will be evaluated as the value of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#expression ValidatingAdmissionPolicyV1#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// Name is the name of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#name ValidatingAdmissionPolicyV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

