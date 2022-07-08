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

type ExpressRouteCircuitConnectionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ExpressRouteCircuitConnectionParameters struct {

	// +kubebuilder:validation:Required
	AddressPrefixIPv4 *string `json:"addressPrefixIpv4" tf:"address_prefix_ipv4,omitempty"`

	// +kubebuilder:validation:Optional
	AddressPrefixIPv6 *string `json:"addressPrefixIpv6,omitempty" tf:"address_prefix_ipv6,omitempty"`

	// +kubebuilder:validation:Optional
	AuthorizationKeySecretRef *v1.SecretKeySelector `json:"authorizationKeySecretRef,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=ExpressRouteCircuitPeering
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PeerPeeringID *string `json:"peerPeeringId,omitempty" tf:"peer_peering_id,omitempty"`

	// +kubebuilder:validation:Optional
	PeerPeeringIDRef *v1.Reference `json:"peerPeeringIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PeerPeeringIDSelector *v1.Selector `json:"peerPeeringIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=ExpressRouteCircuitPeering
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PeeringID *string `json:"peeringId,omitempty" tf:"peering_id,omitempty"`

	// +kubebuilder:validation:Optional
	PeeringIDRef *v1.Reference `json:"peeringIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PeeringIDSelector *v1.Selector `json:"peeringIdSelector,omitempty" tf:"-"`
}

// ExpressRouteCircuitConnectionSpec defines the desired state of ExpressRouteCircuitConnection
type ExpressRouteCircuitConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExpressRouteCircuitConnectionParameters `json:"forProvider"`
}

// ExpressRouteCircuitConnectionStatus defines the observed state of ExpressRouteCircuitConnection.
type ExpressRouteCircuitConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExpressRouteCircuitConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRouteCircuitConnection is the Schema for the ExpressRouteCircuitConnections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ExpressRouteCircuitConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExpressRouteCircuitConnectionSpec   `json:"spec"`
	Status            ExpressRouteCircuitConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRouteCircuitConnectionList contains a list of ExpressRouteCircuitConnections
type ExpressRouteCircuitConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExpressRouteCircuitConnection `json:"items"`
}

// Repository type metadata.
var (
	ExpressRouteCircuitConnection_Kind             = "ExpressRouteCircuitConnection"
	ExpressRouteCircuitConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ExpressRouteCircuitConnection_Kind}.String()
	ExpressRouteCircuitConnection_KindAPIVersion   = ExpressRouteCircuitConnection_Kind + "." + CRDGroupVersion.String()
	ExpressRouteCircuitConnection_GroupVersionKind = CRDGroupVersion.WithKind(ExpressRouteCircuitConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&ExpressRouteCircuitConnection{}, &ExpressRouteCircuitConnectionList{})
}