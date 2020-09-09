# VirtualizationVmwareDatastoreAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessible** | Pointer to **bool** | Shows if this datastore is accessible. | [optional] 
**MaintenanceMode** | Pointer to **bool** | Indicates if the datastore is in maintenance mode. Will be set to True, when in maintenance mode. | [optional] 
**MultipleHostAccess** | Pointer to **bool** | Indicates if this datastore is connected to multiple hosts. | [optional] 
**Status** | Pointer to **string** | Datastore health status, as reported by the hypervisor platform. * &#x60;Unknown&#x60; - Entity status is unknown. * &#x60;Degraded&#x60; - State is degraded, and might impact normal operation of the entity. * &#x60;Critical&#x60; - Entity is in a critical state, impacting operations. * &#x60;Ok&#x60; - Entity status is in a stable state, operating normally. | [optional] [default to "Unknown"]
**ThinProvisioningSupported** | Pointer to **bool** | Indicates if this datastore supports thin provisioning for files. | [optional] 
**UnCommitted** | Pointer to **int64** | Space uncommitted in this datastore in bytes. | [optional] 
**Url** | Pointer to **string** | The URL to access this datastore (example - &#39;ds:///vmfs/volumes/562a4e8a-0eeb5372-dd61-78baf9cb9afa/&#39;). | [optional] 
**Cluster** | Pointer to [**VirtualizationVmwareClusterRelationship**](virtualization.VmwareCluster.Relationship.md) |  | [optional] 
**Datacenter** | Pointer to [**VirtualizationVmwareDatacenterRelationship**](virtualization.VmwareDatacenter.Relationship.md) |  | [optional] 
**Hosts** | Pointer to [**[]VirtualizationVmwareHostRelationship**](virtualization.VmwareHost.Relationship.md) | An array of relationships to virtualizationVmwareHost resources. | [optional] [readonly] 

## Methods

### NewVirtualizationVmwareDatastoreAllOf

`func NewVirtualizationVmwareDatastoreAllOf() *VirtualizationVmwareDatastoreAllOf`

NewVirtualizationVmwareDatastoreAllOf instantiates a new VirtualizationVmwareDatastoreAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareDatastoreAllOfWithDefaults

`func NewVirtualizationVmwareDatastoreAllOfWithDefaults() *VirtualizationVmwareDatastoreAllOf`

NewVirtualizationVmwareDatastoreAllOfWithDefaults instantiates a new VirtualizationVmwareDatastoreAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessible

`func (o *VirtualizationVmwareDatastoreAllOf) GetAccessible() bool`

GetAccessible returns the Accessible field if non-nil, zero value otherwise.

### GetAccessibleOk

`func (o *VirtualizationVmwareDatastoreAllOf) GetAccessibleOk() (*bool, bool)`

GetAccessibleOk returns a tuple with the Accessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessible

`func (o *VirtualizationVmwareDatastoreAllOf) SetAccessible(v bool)`

SetAccessible sets Accessible field to given value.

### HasAccessible

`func (o *VirtualizationVmwareDatastoreAllOf) HasAccessible() bool`

HasAccessible returns a boolean if a field has been set.

### GetMaintenanceMode

`func (o *VirtualizationVmwareDatastoreAllOf) GetMaintenanceMode() bool`

GetMaintenanceMode returns the MaintenanceMode field if non-nil, zero value otherwise.

### GetMaintenanceModeOk

`func (o *VirtualizationVmwareDatastoreAllOf) GetMaintenanceModeOk() (*bool, bool)`

GetMaintenanceModeOk returns a tuple with the MaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceMode

`func (o *VirtualizationVmwareDatastoreAllOf) SetMaintenanceMode(v bool)`

SetMaintenanceMode sets MaintenanceMode field to given value.

### HasMaintenanceMode

`func (o *VirtualizationVmwareDatastoreAllOf) HasMaintenanceMode() bool`

HasMaintenanceMode returns a boolean if a field has been set.

### GetMultipleHostAccess

`func (o *VirtualizationVmwareDatastoreAllOf) GetMultipleHostAccess() bool`

GetMultipleHostAccess returns the MultipleHostAccess field if non-nil, zero value otherwise.

### GetMultipleHostAccessOk

`func (o *VirtualizationVmwareDatastoreAllOf) GetMultipleHostAccessOk() (*bool, bool)`

GetMultipleHostAccessOk returns a tuple with the MultipleHostAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleHostAccess

`func (o *VirtualizationVmwareDatastoreAllOf) SetMultipleHostAccess(v bool)`

SetMultipleHostAccess sets MultipleHostAccess field to given value.

### HasMultipleHostAccess

`func (o *VirtualizationVmwareDatastoreAllOf) HasMultipleHostAccess() bool`

HasMultipleHostAccess returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualizationVmwareDatastoreAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualizationVmwareDatastoreAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualizationVmwareDatastoreAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualizationVmwareDatastoreAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetThinProvisioningSupported

`func (o *VirtualizationVmwareDatastoreAllOf) GetThinProvisioningSupported() bool`

GetThinProvisioningSupported returns the ThinProvisioningSupported field if non-nil, zero value otherwise.

### GetThinProvisioningSupportedOk

`func (o *VirtualizationVmwareDatastoreAllOf) GetThinProvisioningSupportedOk() (*bool, bool)`

GetThinProvisioningSupportedOk returns a tuple with the ThinProvisioningSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinProvisioningSupported

`func (o *VirtualizationVmwareDatastoreAllOf) SetThinProvisioningSupported(v bool)`

SetThinProvisioningSupported sets ThinProvisioningSupported field to given value.

### HasThinProvisioningSupported

`func (o *VirtualizationVmwareDatastoreAllOf) HasThinProvisioningSupported() bool`

HasThinProvisioningSupported returns a boolean if a field has been set.

### GetUnCommitted

`func (o *VirtualizationVmwareDatastoreAllOf) GetUnCommitted() int64`

GetUnCommitted returns the UnCommitted field if non-nil, zero value otherwise.

### GetUnCommittedOk

`func (o *VirtualizationVmwareDatastoreAllOf) GetUnCommittedOk() (*int64, bool)`

GetUnCommittedOk returns a tuple with the UnCommitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnCommitted

`func (o *VirtualizationVmwareDatastoreAllOf) SetUnCommitted(v int64)`

SetUnCommitted sets UnCommitted field to given value.

### HasUnCommitted

`func (o *VirtualizationVmwareDatastoreAllOf) HasUnCommitted() bool`

HasUnCommitted returns a boolean if a field has been set.

### GetUrl

`func (o *VirtualizationVmwareDatastoreAllOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VirtualizationVmwareDatastoreAllOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VirtualizationVmwareDatastoreAllOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VirtualizationVmwareDatastoreAllOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCluster

`func (o *VirtualizationVmwareDatastoreAllOf) GetCluster() VirtualizationVmwareClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualizationVmwareDatastoreAllOf) GetClusterOk() (*VirtualizationVmwareClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualizationVmwareDatastoreAllOf) SetCluster(v VirtualizationVmwareClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualizationVmwareDatastoreAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDatacenter

`func (o *VirtualizationVmwareDatastoreAllOf) GetDatacenter() VirtualizationVmwareDatacenterRelationship`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *VirtualizationVmwareDatastoreAllOf) GetDatacenterOk() (*VirtualizationVmwareDatacenterRelationship, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *VirtualizationVmwareDatastoreAllOf) SetDatacenter(v VirtualizationVmwareDatacenterRelationship)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *VirtualizationVmwareDatastoreAllOf) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetHosts

`func (o *VirtualizationVmwareDatastoreAllOf) GetHosts() []VirtualizationVmwareHostRelationship`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *VirtualizationVmwareDatastoreAllOf) GetHostsOk() (*[]VirtualizationVmwareHostRelationship, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *VirtualizationVmwareDatastoreAllOf) SetHosts(v []VirtualizationVmwareHostRelationship)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *VirtualizationVmwareDatastoreAllOf) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### SetHostsNil

`func (o *VirtualizationVmwareDatastoreAllOf) SetHostsNil(b bool)`

 SetHostsNil sets the value for Hosts to be an explicit nil

### UnsetHosts
`func (o *VirtualizationVmwareDatastoreAllOf) UnsetHosts()`

UnsetHosts ensures that no value is present for Hosts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


