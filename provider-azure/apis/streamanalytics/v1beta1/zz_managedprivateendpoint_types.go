/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ManagedPrivateEndpointObservation struct {

	// The ID of the Stream Analytics.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ManagedPrivateEndpointParameters struct {

	// The name of the Resource Group where the Stream Analytics Managed Private Endpoint should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The name of the Stream Analytics Cluster where the Managed Private Endpoint should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/streamanalytics/v1beta1.Cluster
	// +kubebuilder:validation:Optional
	StreamAnalyticsClusterName *string `json:"streamAnalyticsClusterName,omitempty" tf:"stream_analytics_cluster_name,omitempty"`

	// +kubebuilder:validation:Optional
	StreamAnalyticsClusterNameRef *v1.Reference `json:"streamAnalyticsClusterNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	StreamAnalyticsClusterNameSelector *v1.Selector `json:"streamAnalyticsClusterNameSelector,omitempty" tf:"-"`

	// Specifies the sub resource name which the Stream Analytics Private Endpoint is able to connect to. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	SubresourceName *string `json:"subresourceName" tf:"subresource_name,omitempty"`

	// The ID of the Private Link Enabled Remote Resource which this Stream Analytics Private endpoint should be connected to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`

	// +kubebuilder:validation:Optional
	TargetResourceIDRef *v1.Reference `json:"targetResourceIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TargetResourceIDSelector *v1.Selector `json:"targetResourceIdSelector,omitempty" tf:"-"`
}

// ManagedPrivateEndpointSpec defines the desired state of ManagedPrivateEndpoint
type ManagedPrivateEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagedPrivateEndpointParameters `json:"forProvider"`
}

// ManagedPrivateEndpointStatus defines the observed state of ManagedPrivateEndpoint.
type ManagedPrivateEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagedPrivateEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedPrivateEndpoint is the Schema for the ManagedPrivateEndpoints API. Manages a Stream Analytics Managed Private Endpoint.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ManagedPrivateEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedPrivateEndpointSpec   `json:"spec"`
	Status            ManagedPrivateEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedPrivateEndpointList contains a list of ManagedPrivateEndpoints
type ManagedPrivateEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedPrivateEndpoint `json:"items"`
}

// Repository type metadata.
var (
	ManagedPrivateEndpoint_Kind             = "ManagedPrivateEndpoint"
	ManagedPrivateEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagedPrivateEndpoint_Kind}.String()
	ManagedPrivateEndpoint_KindAPIVersion   = ManagedPrivateEndpoint_Kind + "." + CRDGroupVersion.String()
	ManagedPrivateEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(ManagedPrivateEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagedPrivateEndpoint{}, &ManagedPrivateEndpointList{})
}