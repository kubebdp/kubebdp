/*
Copyright 2023 kubebdp.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type StackService struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Recommendation struct {
	Python string `json:"python"`
}

// StackDefinitionSpec defines the desired state of StackDefinition
type StackDefinitionSpec struct {
	Name           string         `json:"name"`
	Version        string         `json:"version"`
	Services       []StackService `json:"services"`
	Recommendation Recommendation `json:"recommendation"`
}

// StackDefinitionStatus defines the observed state of StackDefinition
type StackDefinitionStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// StackDefinition is the Schema for the stackdefinitions API
type StackDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StackDefinitionSpec   `json:"spec,omitempty"`
	Status StackDefinitionStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// StackDefinitionList contains a list of StackDefinition
type StackDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StackDefinition `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StackDefinition{}, &StackDefinitionList{})
}
