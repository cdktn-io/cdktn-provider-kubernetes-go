// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package validatingadmissionpolicyv1


type ValidatingAdmissionPolicyV1SpecMatchConstraintsNamespaceSelectorMatchExpressions struct {
	// key is the label key that the selector applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.2.1/docs/resources/validating_admission_policy_v1#key ValidatingAdmissionPolicyV1#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.2.1/docs/resources/validating_admission_policy_v1#operator ValidatingAdmissionPolicyV1#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// values is an array of string values.
	//
	// If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.2.1/docs/resources/validating_admission_policy_v1#values ValidatingAdmissionPolicyV1#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

