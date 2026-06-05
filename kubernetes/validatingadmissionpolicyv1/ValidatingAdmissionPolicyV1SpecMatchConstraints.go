// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package validatingadmissionpolicyv1


type ValidatingAdmissionPolicyV1SpecMatchConstraints struct {
	// ExcludeResourceRules describes what operations on what resources/subresources the ValidatingAdmissionPolicy should not care about.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.2.0/docs/resources/validating_admission_policy_v1#exclude_resource_rules ValidatingAdmissionPolicyV1#exclude_resource_rules}
	ExcludeResourceRules interface{} `field:"optional" json:"excludeResourceRules" yaml:"excludeResourceRules"`
	// matchPolicy defines how the MatchResources list is used to match incoming requests. Allowed values are Exact or Equivalent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.2.0/docs/resources/validating_admission_policy_v1#match_policy ValidatingAdmissionPolicyV1#match_policy}
	MatchPolicy *string `field:"optional" json:"matchPolicy" yaml:"matchPolicy"`
	// NamespaceSelector decides whether to run the admission control policy on an object based on whether the namespace for that object matches the selector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.2.0/docs/resources/validating_admission_policy_v1#namespace_selector ValidatingAdmissionPolicyV1#namespace_selector}
	NamespaceSelector *ValidatingAdmissionPolicyV1SpecMatchConstraintsNamespaceSelector `field:"optional" json:"namespaceSelector" yaml:"namespaceSelector"`
	// ObjectSelector decides whether to run the validation based on if the object has matching labels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.2.0/docs/resources/validating_admission_policy_v1#object_selector ValidatingAdmissionPolicyV1#object_selector}
	ObjectSelector *ValidatingAdmissionPolicyV1SpecMatchConstraintsObjectSelector `field:"optional" json:"objectSelector" yaml:"objectSelector"`
	// ResourceRules describes what operations on what resources/subresources the ValidatingAdmissionPolicy matches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.2.0/docs/resources/validating_admission_policy_v1#resource_rules ValidatingAdmissionPolicyV1#resource_rules}
	ResourceRules interface{} `field:"optional" json:"resourceRules" yaml:"resourceRules"`
}

