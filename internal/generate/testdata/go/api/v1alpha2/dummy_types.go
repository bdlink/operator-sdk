// Copyright 2020 The Operator-SDK Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen=false
// +k8s:openapi-gen=false
type NoKindSpec struct {
	// Not included in anything, no kind type
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// +operator-sdk:csv:customresourcedefinitions:type=status
	Size int32 `json:"size"`
	// Not included in anything, no kind type
	Boss Hog `json:"hog"`
}

// +k8s:deepcopy-gen=false
// +k8s:openapi-gen=false
type NoKindStatus struct {
	// Not included in anything, no kind type
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// +operator-sdk:csv:customresourcedefinitions:type=status
	Nodes []string `json:"nodes"`
}

// +k8s:deepcopy-gen=false
// +k8s:openapi-gen=false
type DummySpec struct {
	// Should be in spec
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="dummy-size",xDescriptors="urn:alm:descriptor:com.tectonic.ui:podCount"
	Size int32 `json:"size"`
	// Should be in spec, but should not have array index in path
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Wheels",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	Wheels []Wheel `json:"wheels"`
}

// +k8s:deepcopy-gen=false
// +k8s:openapi-gen=false
type DummyStatus struct {
	// Should be in status but not spec, since DummyStatus isn't in DummySpec
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// +operator-sdk:csv:customresourcedefinitions:type=status
	Nodes []string `json:"nodes"`
	// Not included in status but children should be
	Boss Hog `json:"hog"`
}

// +k8s:deepcopy-gen=false
// +k8s:openapi-gen=false
type Hog struct {
	// Should be in status but not spec, since Hog isn't in DummySpec
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// +operator-sdk:csv:customresourcedefinitions:type=status,displayName="boss-hog-engine"
	Engine Engine `json:"engine"`
	// Not in spec or status, no boolean annotation
	// +operator-sdk:csv:customresourcedefinitions:displayName="doesnt-matter"
	Brand string `json:"brand"`
	// Not in spec or status
	Helmet string `json:"helmet"`
	// Fields should be inlined
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// +operator-sdk:csv:customresourcedefinitions:type=status
	Inlined InlinedComponent `json:",inline"`
	// Fields should be inlined
	InlinedComponent `json:",inline"`
	// Should be ignored
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// +operator-sdk:csv:customresourcedefinitions:type=status
	Ignored IgnoredComponent `json:"-"`
	// Should be ignored, but exported children should not be
	notExported `json:",inline"`
}

type notExported struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// +operator-sdk:csv:customresourcedefinitions:type=status
	Public string `json:"foo"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// +operator-sdk:csv:customresourcedefinitions:type=status
	private string `json:"-"`
}

// +k8s:deepcopy-gen=false
// +k8s:openapi-gen=false
type Engine struct {
	// Should not be included, no annotations.
	Pistons []string `json:"pistons"`
}

// +k8s:deepcopy-gen=false
// +k8s:openapi-gen=false
type Wheel struct {
	// Type should be in spec with path equal to wheels[0].type
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Wheel Type",xDescriptors="urn:alm:descriptor:com.tectonic.ui:arrayFieldGroup:wheels" ; "urn:alm:descriptor:com.tectonic.ui:text"
	Type string `json:"type"`
}

// +k8s:deepcopy-gen=false
// +k8s:openapi-gen=false
type InlinedComponent struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// +operator-sdk:csv:customresourcedefinitions:type=status
	SeatMaterial string `json:"seatMaterial"`
}

// +k8s:deepcopy-gen=false
// +k8s:openapi-gen=false
type IgnoredComponent struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// +operator-sdk:csv:customresourcedefinitions:type=status
	TrunkSpace string `json:"trunkSpace"`
}

// +k8s:deepcopy-gen=false
// +k8s:openapi-gen=false
type OtherDummyStatus struct {
	// Should be in status but not spec, since this isn't a spec type
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// +operator-sdk:csv:customresourcedefinitions:type=status
	Nothing string `json:"nothing"`
}

// Dummy is the Schema for the dummy API
// +k8s:deepcopy-gen=false
// +k8s:openapi-gen=false
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=dummys,scope=Namespaced
// +operator-sdk:csv:customresourcedefinitions:displayName="Dummy App"
// +operator-sdk:csv:customresourcedefinitions:resources={{Deployment,v1,"dummy-deployment"},{ReplicaSet,v1beta2,"dummy-replicaset"},{Pod,v1,"dummy-pod"}}
type Dummy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DummySpec   `json:"spec,omitempty"`
	Status DummyStatus `json:"status,omitempty"`
}

// OtherDummy is the Schema for the other dummy API
// +k8s:deepcopy-gen=false
// +k8s:openapi-gen=false
// +operator-sdk:csv:customresourcedefinitions:displayName="Other Dummy App"
// +operator-sdk:csv:customresourcedefinitions:resources={{Service,v1,"other-dummy-service"},{Pod,v1,"other-dummy-pod"}}
type OtherDummy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Hog              `json:"spec,omitempty"`
	Status OtherDummyStatus `json:"status,omitempty"`
}

// DummyList contains a list of Dummy
// +k8s:deepcopy-gen=false
// +k8s:openapi-gen=false
type DummyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dummy `json:"items"`
}
