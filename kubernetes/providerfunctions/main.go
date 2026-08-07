// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package providerfunctions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-kubernetes.providerFunctions.KubernetesProviderFunctions",
		reflect.TypeOf((*KubernetesProviderFunctions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "manifestDecode", GoMethod: "ManifestDecode"},
			_jsii_.MemberMethod{JsiiMethod: "manifestDecodeMulti", GoMethod: "ManifestDecodeMulti"},
			_jsii_.MemberMethod{JsiiMethod: "manifestEncode", GoMethod: "ManifestEncode"},
		},
		func() interface{} {
			return &jsiiProxy_KubernetesProviderFunctions{}
		},
	)
}
