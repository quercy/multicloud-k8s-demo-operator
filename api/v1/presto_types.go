/*


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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PrestoSpec defines the desired state of Presto
type PrestoSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Workers          int32        `json:"workers"`
	Node             PrestoNode   `json:"node"`
	Config           PrestoConfig `json:"config"`
	LogLevel         string       `json:"logLevel"`
	MaxMemory        string       `json:"maxMemory"`
	MaxMemoryPerNode string       `json:"maxMemoryPerNode"`
	JvmConfig        JvmConfig    `json:"jvmConfig"`
	Image            Image        `json:"image"`
}

// Image defines the Presto image to pull
type Image struct {
	Repository      string          `json:"repository"`
	Tag             string          `json:"tag"`
	PullPolicy      string          `json:"pullPolicy"`
	SecurityContext SecurityContext `json:"securityContext"`
}

// SecurityContext defines the SecurityContext
type SecurityContext struct {
	RunAsUser  int `json:"runAsUser"`
	RunAsGroup int `json:"runAsGroup"`
}

// PrestoNode defines the Presto Node config
type PrestoNode struct {
	Environment string `json:"environment"`
	DataDir     string `json:"dataDir"`
	PluginDir   string `json:"pluginDir"`
}

// PrestoConfig defines the Presto config
type PrestoConfig struct {
	Path     string `json:"path"`
	HTTPPort int32  `json:"httpPort"`
}

// JvmConfig configures the JVM
type JvmConfig struct {
	MaxHeapSize string   `json:"maxHeapSize"`
	GcMethod    GcMethod `json:"gcMethod"`
}

// GcMethod defines the garbage collection method
type GcMethod struct {
	Type string `json:"type"`
	G1   G1     `json:"g1"`
}

// G1 does something
type G1 struct {
	HeapRegionSize string `json:"heapRegionSize"`
}

// PrestoStatus defines the observed state of Presto
type PrestoStatus struct {
	Controller string `json:"controller,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Presto is the Schema for the prestoes API
type Presto struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PrestoSpec   `json:"spec,omitempty"`
	Status PrestoStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrestoList contains a list of Presto
type PrestoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Presto `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Presto{}, &PrestoList{})
}
