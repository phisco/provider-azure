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

type SpringCloudDevToolPortalObservation struct {

	// The ID of the Spring Cloud Dev Tool Portal.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SpringCloudDevToolPortalParameters struct {

	// Should the Accelerator plugin be enabled?
	// +kubebuilder:validation:Optional
	ApplicationAcceleratorEnabled *bool `json:"applicationAcceleratorEnabled,omitempty" tf:"application_accelerator_enabled,omitempty"`

	// Should the Application Live View be enabled?
	// +kubebuilder:validation:Optional
	ApplicationLiveViewEnabled *bool `json:"applicationLiveViewEnabled,omitempty" tf:"application_live_view_enabled,omitempty"`

	// The name which should be used for this Spring Cloud Dev Tool Portal. The only possible value is default. Changing this forces a new Spring Cloud Dev Tool Portal to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Is public network access enabled?
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The ID of the Spring Cloud Service. Changing this forces a new Spring Cloud Dev Tool Portal to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/appplatform/v1beta1.SpringCloudService
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SpringCloudServiceID *string `json:"springCloudServiceId,omitempty" tf:"spring_cloud_service_id,omitempty"`

	// Reference to a SpringCloudService in appplatform to populate springCloudServiceId.
	// +kubebuilder:validation:Optional
	SpringCloudServiceIDRef *v1.Reference `json:"springCloudServiceIdRef,omitempty" tf:"-"`

	// Selector for a SpringCloudService in appplatform to populate springCloudServiceId.
	// +kubebuilder:validation:Optional
	SpringCloudServiceIDSelector *v1.Selector `json:"springCloudServiceIdSelector,omitempty" tf:"-"`

	// A sso block as defined below.
	// +kubebuilder:validation:Optional
	Sso []SsoParameters `json:"sso,omitempty" tf:"sso,omitempty"`
}

type SsoObservation struct {
}

type SsoParameters struct {

	// Specifies the public identifier for the application.
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Specifies the secret known only to the application and the authorization server.
	// +kubebuilder:validation:Optional
	ClientSecret *string `json:"clientSecret,omitempty" tf:"client_secret,omitempty"`

	// Specifies the URI of a JSON file with generic OIDC provider configuration.
	// +kubebuilder:validation:Optional
	MetadataURL *string `json:"metadataUrl,omitempty" tf:"metadata_url,omitempty"`

	// Specifies a list of specific actions applications can be allowed to do on a user's behalf.
	// +kubebuilder:validation:Optional
	Scope []*string `json:"scope,omitempty" tf:"scope,omitempty"`
}

// SpringCloudDevToolPortalSpec defines the desired state of SpringCloudDevToolPortal
type SpringCloudDevToolPortalSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpringCloudDevToolPortalParameters `json:"forProvider"`
}

// SpringCloudDevToolPortalStatus defines the observed state of SpringCloudDevToolPortal.
type SpringCloudDevToolPortalStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpringCloudDevToolPortalObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudDevToolPortal is the Schema for the SpringCloudDevToolPortals API. Manages a Spring Cloud Dev Tool Portal.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SpringCloudDevToolPortal struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpringCloudDevToolPortalSpec   `json:"spec"`
	Status            SpringCloudDevToolPortalStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudDevToolPortalList contains a list of SpringCloudDevToolPortals
type SpringCloudDevToolPortalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpringCloudDevToolPortal `json:"items"`
}

// Repository type metadata.
var (
	SpringCloudDevToolPortal_Kind             = "SpringCloudDevToolPortal"
	SpringCloudDevToolPortal_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpringCloudDevToolPortal_Kind}.String()
	SpringCloudDevToolPortal_KindAPIVersion   = SpringCloudDevToolPortal_Kind + "." + CRDGroupVersion.String()
	SpringCloudDevToolPortal_GroupVersionKind = CRDGroupVersion.WithKind(SpringCloudDevToolPortal_Kind)
)

func init() {
	SchemeBuilder.Register(&SpringCloudDevToolPortal{}, &SpringCloudDevToolPortalList{})
}