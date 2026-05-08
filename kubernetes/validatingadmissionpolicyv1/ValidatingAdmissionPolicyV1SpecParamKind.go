// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package validatingadmissionpolicyv1


type ValidatingAdmissionPolicyV1SpecParamKind struct {
	// APIVersion is the API group version the resources belong to. In format of "group/version".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#api_version ValidatingAdmissionPolicyV1#api_version}
	ApiVersion *string `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// Kind is the API kind the resources belong to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#kind ValidatingAdmissionPolicyV1#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
}

