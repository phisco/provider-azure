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

type DataProtectionReplicationObservation struct {
}

type DataProtectionReplicationParameters struct {

	// +kubebuilder:validation:Optional
	EndpointType *string `json:"endpointType,omitempty" tf:"endpoint_type,omitempty"`

	// +kubebuilder:validation:Required
	RemoteVolumeLocation *string `json:"remoteVolumeLocation" tf:"remote_volume_location,omitempty"`

	// +crossplane:generate:reference:type=Volume
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RemoteVolumeResourceID *string `json:"remoteVolumeResourceId,omitempty" tf:"remote_volume_resource_id,omitempty"`

	// +kubebuilder:validation:Optional
	RemoteVolumeResourceIDRef *v1.Reference `json:"remoteVolumeResourceIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RemoteVolumeResourceIDSelector *v1.Selector `json:"remoteVolumeResourceIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	ReplicationFrequency *string `json:"replicationFrequency" tf:"replication_frequency,omitempty"`
}

type DataProtectionSnapshotPolicyObservation struct {
}

type DataProtectionSnapshotPolicyParameters struct {

	// +crossplane:generate:reference:type=SnapshotPolicy
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SnapshotPolicyID *string `json:"snapshotPolicyId,omitempty" tf:"snapshot_policy_id,omitempty"`

	// +kubebuilder:validation:Optional
	SnapshotPolicyIDRef *v1.Reference `json:"snapshotPolicyIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SnapshotPolicyIDSelector *v1.Selector `json:"snapshotPolicyIdSelector,omitempty" tf:"-"`
}

type ExportPolicyRuleObservation struct {
}

type ExportPolicyRuleParameters struct {

	// +kubebuilder:validation:Required
	AllowedClients []*string `json:"allowedClients" tf:"allowed_clients,omitempty"`

	// +kubebuilder:validation:Optional
	ProtocolsEnabled []*string `json:"protocolsEnabled,omitempty" tf:"protocols_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	RootAccessEnabled *bool `json:"rootAccessEnabled,omitempty" tf:"root_access_enabled,omitempty"`

	// +kubebuilder:validation:Required
	RuleIndex *float64 `json:"ruleIndex" tf:"rule_index,omitempty"`

	// +kubebuilder:validation:Optional
	UnixReadOnly *bool `json:"unixReadOnly,omitempty" tf:"unix_read_only,omitempty"`

	// +kubebuilder:validation:Optional
	UnixReadWrite *bool `json:"unixReadWrite,omitempty" tf:"unix_read_write,omitempty"`
}

type VolumeObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	MountIPAddresses []*string `json:"mountIpAddresses,omitempty" tf:"mount_ip_addresses,omitempty"`
}

type VolumeParameters struct {

	// +crossplane:generate:reference:type=Account
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// +kubebuilder:validation:Optional
	AccountNameRef *v1.Reference `json:"accountNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AccountNameSelector *v1.Selector `json:"accountNameSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=Snapshot
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CreateFromSnapshotResourceID *string `json:"createFromSnapshotResourceId,omitempty" tf:"create_from_snapshot_resource_id,omitempty"`

	// +kubebuilder:validation:Optional
	CreateFromSnapshotResourceIDRef *v1.Reference `json:"createFromSnapshotResourceIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	CreateFromSnapshotResourceIDSelector *v1.Selector `json:"createFromSnapshotResourceIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DataProtectionReplication []DataProtectionReplicationParameters `json:"dataProtectionReplication,omitempty" tf:"data_protection_replication,omitempty"`

	// +kubebuilder:validation:Optional
	DataProtectionSnapshotPolicy []DataProtectionSnapshotPolicyParameters `json:"dataProtectionSnapshotPolicy,omitempty" tf:"data_protection_snapshot_policy,omitempty"`

	// +kubebuilder:validation:Optional
	ExportPolicyRule []ExportPolicyRuleParameters `json:"exportPolicyRule,omitempty" tf:"export_policy_rule,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +crossplane:generate:reference:type=Pool
	// +kubebuilder:validation:Optional
	PoolName *string `json:"poolName,omitempty" tf:"pool_name,omitempty"`

	// +kubebuilder:validation:Optional
	PoolNameRef *v1.Reference `json:"poolNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PoolNameSelector *v1.Selector `json:"poolNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Protocols []*string `json:"protocols,omitempty" tf:"protocols,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SecurityStyle *string `json:"securityStyle,omitempty" tf:"security_style,omitempty"`

	// +kubebuilder:validation:Required
	ServiceLevel *string `json:"serviceLevel" tf:"service_level,omitempty"`

	// +kubebuilder:validation:Optional
	SnapshotDirectoryVisible *bool `json:"snapshotDirectoryVisible,omitempty" tf:"snapshot_directory_visible,omitempty"`

	// +kubebuilder:validation:Required
	StorageQuotaInGb *float64 `json:"storageQuotaInGb" tf:"storage_quota_in_gb,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	ThroughputInMibps *float64 `json:"throughputInMibps,omitempty" tf:"throughput_in_mibps,omitempty"`

	// +kubebuilder:validation:Required
	VolumePath *string `json:"volumePath" tf:"volume_path,omitempty"`
}

// VolumeSpec defines the desired state of Volume
type VolumeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VolumeParameters `json:"forProvider"`
}

// VolumeStatus defines the observed state of Volume.
type VolumeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VolumeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Volume is the Schema for the Volumes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Volume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VolumeSpec   `json:"spec"`
	Status            VolumeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeList contains a list of Volumes
type VolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Volume `json:"items"`
}

// Repository type metadata.
var (
	Volume_Kind             = "Volume"
	Volume_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Volume_Kind}.String()
	Volume_KindAPIVersion   = Volume_Kind + "." + CRDGroupVersion.String()
	Volume_GroupVersionKind = CRDGroupVersion.WithKind(Volume_Kind)
)

func init() {
	SchemeBuilder.Register(&Volume{}, &VolumeList{})
}
