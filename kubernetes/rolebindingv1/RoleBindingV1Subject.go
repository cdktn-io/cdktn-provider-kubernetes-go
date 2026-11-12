// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rolebindingv1


type RoleBindingV1Subject struct {
	// The kind of the subject. One of `User`, `ServiceAccount`, or `Group`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.5.0/docs/resources/role_binding_v1#kind RoleBindingV1#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// The name of the subject.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.5.0/docs/resources/role_binding_v1#name RoleBindingV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The API group of the subject. For User and Group subjects use `rbac.authorization.k8s.io`. For ServiceAccount subjects use `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.5.0/docs/resources/role_binding_v1#api_group RoleBindingV1#api_group}
	ApiGroup *string `field:"optional" json:"apiGroup" yaml:"apiGroup"`
	// The namespace of the referenced ServiceAccount. Defaults to `default`. Used only for `ServiceAccount` subjects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.5.0/docs/resources/role_binding_v1#namespace RoleBindingV1#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

