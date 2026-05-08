// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package validatingadmissionpolicyv1


type ValidatingAdmissionPolicyV1SpecMatchConstraintsResourceRules struct {
	// APIGroups is the API groups the resources belong to.
	//
	// '\*' is all groups. If '\*' is present, the length of the slice must be one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#api_groups ValidatingAdmissionPolicyV1#api_groups}
	ApiGroups *[]*string `field:"required" json:"apiGroups" yaml:"apiGroups"`
	// APIVersions is the API versions the resources belong to.
	//
	// '\*' is all versions. If '\*' is present, the length of the slice must be one. Required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#api_versions ValidatingAdmissionPolicyV1#api_versions}
	ApiVersions *[]*string `field:"required" json:"apiVersions" yaml:"apiVersions"`
	// Operations is the operations the admission hook cares about - CREATE, UPDATE, DELETE, CONNECT or * for all of those operations and any future admission operations that are added.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#operations ValidatingAdmissionPolicyV1#operations}
	Operations *[]*string `field:"required" json:"operations" yaml:"operations"`
	// Resources is a list of resources this rule applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#resources ValidatingAdmissionPolicyV1#resources}
	Resources *[]*string `field:"required" json:"resources" yaml:"resources"`
	// ResourceNames is an optional white list of names that the rule applies to.
	//
	// An empty set means that everything is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#resource_names ValidatingAdmissionPolicyV1#resource_names}
	ResourceNames *[]*string `field:"optional" json:"resourceNames" yaml:"resourceNames"`
	// scope specifies the scope of this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#scope ValidatingAdmissionPolicyV1#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

