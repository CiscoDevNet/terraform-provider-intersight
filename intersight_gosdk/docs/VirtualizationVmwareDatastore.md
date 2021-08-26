# VirtualizationVmwareDatastore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareDatastore"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareDatastore"]
**Accessible** | Pointer to **bool** | Shows if this datastore is accessible. | [optional] 
**InventoryPath** | Pointer to **string** | Inventory path of the Datastore. | [optional] 
**MaintenanceMode** | Pointer to **bool** | Indicates if the datastore is in maintenance mode. Will be set to True, when in maintenance mode. | [optional] 
**MultipleHostAccess** | Pointer to **bool** | Indicates if this datastore is connected to multiple hosts. | [optional] 
**Status** | Pointer to **string** | Datastore health status, as reported by the hypervisor platform. * &#x60;Unknown&#x60; - Entity status is unknown. * &#x60;Degraded&#x60; - State is degraded, and might impact normal operation of the entity. * &#x60;Critical&#x60; - Entity is in a critical state, impacting operations. * &#x60;Ok&#x60; - Entity status is in a stable state, operating normally. | [optional] [default to "Unknown"]
**ThinProvisioningSupported** | Pointer to **bool** | Indicates if this datastore supports thin provisioning for files. | [optional] 
**UnCommitted** | Pointer to **int64** | Space uncommitted in this datastore in bytes. | [optional] 
**Url** | Pointer to **string** | The URL to access this datastore (example - &#39;ds:///vmfs/volumes/562a4e8a-0eeb5372-dd61-78baf9cb9afa/&#39;). | [optional] 
**Cluster** | Pointer to [**VirtualizationVmwareClusterRelationship**](VirtualizationVmwareClusterRelationship.md) |  | [optional] 
**Datacenter** | Pointer to [**VirtualizationVmwareDatacenterRelationship**](VirtualizationVmwareDatacenterRelationship.md) |  | [optional] 
**Hosts** | Pointer to [**[]VirtualizationVmwareHostRelationship**](VirtualizationVmwareHostRelationship.md) | An array of relationships to virtualizationVmwareHost resources. | [optional] [readonly] 

## Methods

### NewVirtualizationVmwareDatastore

`func NewVirtualizationVmwareDatastore(classId string, objectType string, ) *VirtualizationVmwareDatastore`

NewVirtualizationVmwareDatastore instantiates a new VirtualizationVmwareDatastore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareDatastoreWithDefaults

`func NewVirtualizationVmwareDatastoreWithDefaults() *VirtualizationVmwareDatastore`

NewVirtualizationVmwareDatastoreWithDefaults instantiates a new VirtualizationVmwareDatastore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareDatastore) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareDatastore) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareDatastore) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareDatastore) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareDatastore) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareDatastore) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccessible

`func (o *VirtualizationVmwareDatastore) GetAccessible() bool`

GetAccessible returns the Accessible field if non-nil, zero value otherwise.

### GetAccessibleOk

`func (o *VirtualizationVmwareDatastore) GetAccessibleOk() (*bool, bool)`

GetAccessibleOk returns a tuple with the Accessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessible

`func (o *VirtualizationVmwareDatastore) SetAccessible(v bool)`

SetAccessible sets Accessible field to given value.

### HasAccessible

`func (o *VirtualizationVmwareDatastore) HasAccessible() bool`

HasAccessible returns a boolean if a field has been set.

### GetInventoryPath

`func (o *VirtualizationVmwareDatastore) GetInventoryPath() string`

GetInventoryPath returns the InventoryPath field if non-nil, zero value otherwise.

### GetInventoryPathOk

`func (o *VirtualizationVmwareDatastore) GetInventoryPathOk() (*string, bool)`

GetInventoryPathOk returns a tuple with the InventoryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryPath

`func (o *VirtualizationVmwareDatastore) SetInventoryPath(v string)`

SetInventoryPath sets InventoryPath field to given value.

### HasInventoryPath

`func (o *VirtualizationVmwareDatastore) HasInventoryPath() bool`

HasInventoryPath returns a boolean if a field has been set.

### GetMaintenanceMode

`func (o *VirtualizationVmwareDatastore) GetMaintenanceMode() bool`

GetMaintenanceMode returns the MaintenanceMode field if non-nil, zero value otherwise.

### GetMaintenanceModeOk

`func (o *VirtualizationVmwareDatastore) GetMaintenanceModeOk() (*bool, bool)`

GetMaintenanceModeOk returns a tuple with the MaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceMode

`func (o *VirtualizationVmwareDatastore) SetMaintenanceMode(v bool)`

SetMaintenanceMode sets MaintenanceMode field to given value.

### HasMaintenanceMode

`func (o *VirtualizationVmwareDatastore) HasMaintenanceMode() bool`

HasMaintenanceMode returns a boolean if a field has been set.

### GetMultipleHostAccess

`func (o *VirtualizationVmwareDatastore) GetMultipleHostAccess() bool`

GetMultipleHostAccess returns the MultipleHostAccess field if non-nil, zero value otherwise.

### GetMultipleHostAccessOk

`func (o *VirtualizationVmwareDatastore) GetMultipleHostAccessOk() (*bool, bool)`

GetMultipleHostAccessOk returns a tuple with the MultipleHostAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleHostAccess

`func (o *VirtualizationVmwareDatastore) SetMultipleHostAccess(v bool)`

SetMultipleHostAccess sets MultipleHostAccess field to given value.

### HasMultipleHostAccess

`func (o *VirtualizationVmwareDatastore) HasMultipleHostAccess() bool`

HasMultipleHostAccess returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualizationVmwareDatastore) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualizationVmwareDatastore) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualizationVmwareDatastore) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualizationVmwareDatastore) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetThinProvisioningSupported

`func (o *VirtualizationVmwareDatastore) GetThinProvisioningSupported() bool`

GetThinProvisioningSupported returns the ThinProvisioningSupported field if non-nil, zero value otherwise.

### GetThinProvisioningSupportedOk

`func (o *VirtualizationVmwareDatastore) GetThinProvisioningSupportedOk() (*bool, bool)`

GetThinProvisioningSupportedOk returns a tuple with the ThinProvisioningSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinProvisioningSupported

`func (o *VirtualizationVmwareDatastore) SetThinProvisioningSupported(v bool)`

SetThinProvisioningSupported sets ThinProvisioningSupported field to given value.

### HasThinProvisioningSupported

`func (o *VirtualizationVmwareDatastore) HasThinProvisioningSupported() bool`

HasThinProvisioningSupported returns a boolean if a field has been set.

### GetUnCommitted

`func (o *VirtualizationVmwareDatastore) GetUnCommitted() int64`

GetUnCommitted returns the UnCommitted field if non-nil, zero value otherwise.

### GetUnCommittedOk

`func (o *VirtualizationVmwareDatastore) GetUnCommittedOk() (*int64, bool)`

GetUnCommittedOk returns a tuple with the UnCommitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnCommitted

`func (o *VirtualizationVmwareDatastore) SetUnCommitted(v int64)`

SetUnCommitted sets UnCommitted field to given value.

### HasUnCommitted

`func (o *VirtualizationVmwareDatastore) HasUnCommitted() bool`

HasUnCommitted returns a boolean if a field has been set.

### GetUrl

`func (o *VirtualizationVmwareDatastore) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VirtualizationVmwareDatastore) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VirtualizationVmwareDatastore) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VirtualizationVmwareDatastore) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCluster

`func (o *VirtualizationVmwareDatastore) GetCluster() VirtualizationVmwareClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualizationVmwareDatastore) GetClusterOk() (*VirtualizationVmwareClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualizationVmwareDatastore) SetCluster(v VirtualizationVmwareClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualizationVmwareDatastore) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDatacenter

`func (o *VirtualizationVmwareDatastore) GetDatacenter() VirtualizationVmwareDatacenterRelationship`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *VirtualizationVmwareDatastore) GetDatacenterOk() (*VirtualizationVmwareDatacenterRelationship, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *VirtualizationVmwareDatastore) SetDatacenter(v VirtualizationVmwareDatacenterRelationship)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *VirtualizationVmwareDatastore) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetHosts

`func (o *VirtualizationVmwareDatastore) GetHosts() []VirtualizationVmwareHostRelationship`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *VirtualizationVmwareDatastore) GetHostsOk() (*[]VirtualizationVmwareHostRelationship, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *VirtualizationVmwareDatastore) SetHosts(v []VirtualizationVmwareHostRelationship)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *VirtualizationVmwareDatastore) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### SetHostsNil

`func (o *VirtualizationVmwareDatastore) SetHostsNil(b bool)`

 SetHostsNil sets the value for Hosts to be an explicit nil

### UnsetHosts
`func (o *VirtualizationVmwareDatastore) UnsetHosts()`

UnsetHosts ensures that no value is present for Hosts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


