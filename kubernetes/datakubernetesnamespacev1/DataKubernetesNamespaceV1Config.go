// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datakubernetesnamespacev1

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataKubernetesNamespaceV1Config struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/3.3.0/docs/data-sources/namespace_v1#metadata DataKubernetesNamespaceV1#metadata}
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
}

