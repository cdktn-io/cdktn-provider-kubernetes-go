// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package validatingadmissionpolicyv1


type ValidatingAdmissionPolicyV1SpecMatchConstraintsObjectSelector struct {
	// A label query over a set of resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.2.0/docs/resources/validating_admission_policy_v1#label_selector ValidatingAdmissionPolicyV1#label_selector}
	LabelSelector *ValidatingAdmissionPolicyV1SpecMatchConstraintsObjectSelectorLabelSelector `field:"optional" json:"labelSelector" yaml:"labelSelector"`
}

