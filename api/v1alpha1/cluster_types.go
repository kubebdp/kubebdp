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

type ConfigurationItem struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ConfigurationGroup struct {
	Name  string              `json:"name"`
	Items []ConfigurationItem `json:"items"`
}

type ClusterServiceDependency struct {
	Service string `json:"service"`
	Name    string `json:"name"`
}

type ClusterService struct {
	Name           string                     `json:"name"`
	Dependencies   []ClusterServiceDependency `json:"dependencies"`
	ServiceSpec    string                     `json:"serviceSpec"`
	Configurations []ConfigurationGroup       `json:"configurations"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	Name     string           `json:"name"`
	Stack    string           `json:"stack"`
	Version  string           `json:"version"`
	Services []ClusterService `json:"services"`
	//Configurations []ConfigurationGroup `json:"configurations"`
	// CertConfigs
	Environments map[string]string `json:"environments"`
}

type ServiceStatus struct {
}

// ClusterStatus defines the observed state of Cluster
type ClusterStatus struct {
	ServiceStatuses []ServiceStatus `json:"serviceStatuses"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Cluster is the Schema for the clusters API
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterSpec   `json:"spec,omitempty"`
	Status ClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ClusterList contains a list of Cluster
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
