// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	CTGroup         = "contivvpptelemetry"
	CTVGroupVersion = "v1"
	CTPlural        = "contivtelemtries"
	CNTFullName     = CTPlural + "." + CTGroup
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ContivTelemetry describes contiv telemetry custom resource
type ContivTelemetry struct {
	// TypeMeta is the metadata for the resource, like kind and apiversion
	metav1.TypeMeta `json:",inline"`
	// ObjectMeta contains the metadata for the particular object
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Spec is the custom resource spec
	Spec ContivTelemetrySpec `json:"spec"`
}

// ContivTelemetrySpec is the spec for the contiv telemetry resource
type ContivTelemetrySpec struct {
	// Message and SomeValue are example custom spec fields
	Node  []string `json:"node"`
	Agent []string `json:"component"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ContivTelemetryList is a list of ContivTelemetry resource
type ContivTelemetryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ContivTelemetry `json:"items"`
}