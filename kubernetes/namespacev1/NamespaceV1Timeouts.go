// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package namespacev1


type NamespaceV1Timeouts struct {
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.3.0/docs/resources/namespace_v1#delete NamespaceV1#delete}
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

