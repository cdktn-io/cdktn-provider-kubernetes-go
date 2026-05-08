// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package validatingadmissionpolicyv1


type ValidatingAdmissionPolicyV1Metadata struct {
	// Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.
	//
	// They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#annotations ValidatingAdmissionPolicyV1#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided.
	//
	// If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server.
	// If this field is specified and the generated name exists, the server will return a 409.
	// Applied only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#idempotency
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#generate_name ValidatingAdmissionPolicyV1#generate_name}
	GenerateName *string `field:"optional" json:"generateName" yaml:"generateName"`
	// A sequence number representing a specific generation of the desired state. Populated by the system. Read-only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#generation ValidatingAdmissionPolicyV1#generation}
	Generation *float64 `field:"optional" json:"generation" yaml:"generation"`
	// Map of string keys and values that can be used to organize and categorize (scope and select) objects.
	//
	// May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#labels ValidatingAdmissionPolicyV1#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Name must be unique within a namespace.
	//
	// Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#name ValidatingAdmissionPolicyV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Namespace defines the space within which each name must be unique.
	//
	// An empty namespace is equivalent to the "default" namespace, but "default" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.
	// Must be a DNS_LABEL. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#namespace ValidatingAdmissionPolicyV1#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed.
	//
	// May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources.
	// Populated by the system. Read-only. Value must be treated as opaque by clients and . More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#resource_version ValidatingAdmissionPolicyV1#resource_version}
	ResourceVersion *string `field:"optional" json:"resourceVersion" yaml:"resourceVersion"`
	// UID is the unique in time and space value for this object.
	//
	// It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.
	// Populated by the system. Read-only. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#uids
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.1.0/docs/resources/validating_admission_policy_v1#uid ValidatingAdmissionPolicyV1#uid}
	Uid *string `field:"optional" json:"uid" yaml:"uid"`
}

