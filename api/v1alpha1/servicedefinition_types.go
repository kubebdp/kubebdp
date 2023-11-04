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

type DependencyOperator string

var (
	DependencyOpIn           DependencyOperator = "In"
	DependencyOpNotIn        DependencyOperator = "NotIn"
	DependencyOpExists       DependencyOperator = "Exists"
	DependencyOpDoesNotExist DependencyOperator = "DoesNotExist"
)

type ConfigurationDefinition struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type Configurations struct {
	Type        string                    `json:"type"`
	Format      string                    `json:"format"`
	Definitions []ConfigurationDefinition `json:"definitions"`
}

// ServiceDefinitionSpec defines the desired state of ServiceDefinition
type ServiceDefinitionSpec struct {
	Service        string           `json:"service"`
	Version        string           `json:"version"`
	Description    string           `json:"description"`
	Configurations []Configurations `json:"configurations"`
	// operator, helm, raw
	Engine string `json:"engine"`
	Helm   string `json:"helm"`
	Raw    string `json:"raw"`
}

// ServiceDefinitionStatus defines the observed state of ServiceDefinition
type ServiceDefinitionStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ServiceDefinition is the Schema for the ServiceDefinitions API
type ServiceDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceDefinitionSpec   `json:"spec,omitempty"`
	Status ServiceDefinitionStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ServiceDefinitionList contains a list of ServiceDefinition
type ServiceDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceDefinition `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ServiceDefinition{}, &ServiceDefinitionList{})
}
