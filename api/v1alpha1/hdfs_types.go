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

// HdfsSpec defines the desired state of Hdfs
type HdfsSpec struct {
	Name    string `json:"name"`
	Cluster string `json:"cluster"`
}

// HdfsStatus defines the observed state of Hdfs
type HdfsStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:path=hdfses

// Hdfs is the Schema for the hdfses API
type Hdfs struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HdfsSpec   `json:"spec,omitempty"`
	Status HdfsStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HdfsList contains a list of Hdfs
type HdfsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Hdfs `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Hdfs{}, &HdfsList{})
}
