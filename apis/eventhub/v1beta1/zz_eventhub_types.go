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

type CaptureDescriptionObservation struct {
}

type CaptureDescriptionParameters struct {

	// A destination block as defined below.
	// +kubebuilder:validation:Required
	Destination []DestinationParameters `json:"destination" tf:"destination,omitempty"`

	// Specifies if the Capture Description is Enabled.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// Specifies the Encoding used for the Capture Description. Possible values are Avro and AvroDeflate.
	// +kubebuilder:validation:Required
	Encoding *string `json:"encoding" tf:"encoding,omitempty"`

	// Specifies the time interval in seconds at which the capture will happen. Values can be between 60 and 900 seconds. Defaults to 300 seconds.
	// +kubebuilder:validation:Optional
	IntervalInSeconds *float64 `json:"intervalInSeconds,omitempty" tf:"interval_in_seconds,omitempty"`

	// Specifies the amount of data built up in your EventHub before a Capture Operation occurs. Value should be between 10485760 and 524288000  bytes. Defaults to 314572800 bytes.
	// +kubebuilder:validation:Optional
	SizeLimitInBytes *float64 `json:"sizeLimitInBytes,omitempty" tf:"size_limit_in_bytes,omitempty"`

	// Specifies if empty files should not be emitted if no events occur during the Capture time window.  Defaults to false.
	// +kubebuilder:validation:Optional
	SkipEmptyArchives *bool `json:"skipEmptyArchives,omitempty" tf:"skip_empty_archives,omitempty"`
}

type DestinationObservation struct {
}

type DestinationParameters struct {

	// The Blob naming convention for archiving. e.g. {Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}. Here all the parameters (Namespace,EventHub .. etc) are mandatory irrespective of order
	// +kubebuilder:validation:Required
	ArchiveNameFormat *string `json:"archiveNameFormat" tf:"archive_name_format,omitempty"`

	// The name of the Container within the Blob Storage Account where messages should be archived.
	// +kubebuilder:validation:Required
	BlobContainerName *string `json:"blobContainerName" tf:"blob_container_name,omitempty"`

	// Specifies the name of the EventHub resource. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The ID of the Blob Storage Account where messages should be archived.
	// +kubebuilder:validation:Required
	StorageAccountID *string `json:"storageAccountId" tf:"storage_account_id,omitempty"`
}

type EventHubObservation struct {

	// The ID of the EventHub.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The identifiers for partitions created for Event Hubs.
	PartitionIds []*string `json:"partitionIds,omitempty" tf:"partition_ids,omitempty"`
}

type EventHubParameters struct {

	// A capture_description block as defined below.
	// +kubebuilder:validation:Optional
	CaptureDescription []CaptureDescriptionParameters `json:"captureDescription,omitempty" tf:"capture_description,omitempty"`

	// Specifies the number of days to retain the events for this Event Hub.
	// +kubebuilder:validation:Required
	MessageRetention *float64 `json:"messageRetention" tf:"message_retention,omitempty"`

	// Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=EventHubNamespace
	// +kubebuilder:validation:Optional
	NamespaceName *string `json:"namespaceName,omitempty" tf:"namespace_name,omitempty"`

	// Reference to a EventHubNamespace to populate namespaceName.
	// +kubebuilder:validation:Optional
	NamespaceNameRef *v1.Reference `json:"namespaceNameRef,omitempty" tf:"-"`

	// Selector for a EventHubNamespace to populate namespaceName.
	// +kubebuilder:validation:Optional
	NamespaceNameSelector *v1.Selector `json:"namespaceNameSelector,omitempty" tf:"-"`

	// Specifies the current number of shards on the Event Hub.
	// +kubebuilder:validation:Required
	PartitionCount *float64 `json:"partitionCount" tf:"partition_count,omitempty"`

	// The name of the resource group in which the EventHub's parent Namespace exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the status of the Event Hub resource. Possible values are Active, Disabled and SendDisabled. Defaults to Active.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

// EventHubSpec defines the desired state of EventHub
type EventHubSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EventHubParameters `json:"forProvider"`
}

// EventHubStatus defines the observed state of EventHub.
type EventHubStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EventHubObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EventHub is the Schema for the EventHubs API. Manages a Event Hubs as a nested resource within an Event Hubs namespace.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type EventHub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventHubSpec   `json:"spec"`
	Status            EventHubStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventHubList contains a list of EventHubs
type EventHubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventHub `json:"items"`
}

// Repository type metadata.
var (
	EventHub_Kind             = "EventHub"
	EventHub_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EventHub_Kind}.String()
	EventHub_KindAPIVersion   = EventHub_Kind + "." + CRDGroupVersion.String()
	EventHub_GroupVersionKind = CRDGroupVersion.WithKind(EventHub_Kind)
)

func init() {
	SchemeBuilder.Register(&EventHub{}, &EventHubList{})
}