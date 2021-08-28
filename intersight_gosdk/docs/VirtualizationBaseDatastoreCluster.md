# VirtualizationBaseDatastoreCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "virtualization.VmwareDatastoreCluster"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "virtualization.VmwareDatastoreCluster"]
**Capacity** | Pointer to [**NullableVirtualizationStorageCapacity**](VirtualizationStorageCapacity.md) |  | [optional] 
**DatastoreCount** | Pointer to **int64** | Number of datastores present in this datastore cluster. | [optional] 
**HostCount** | Pointer to **int64** | Number of hosts attached to or supported-by this datastore cluster. | [optional] 
**Identity** | Pointer to **string** | The internally generated identity of this datastore cluster. This entity is not manipulated by users. It aids in uniquely identifying the datastore cluster object. For VMware, this is an MOR (managed object reference). | [optional] 
**Name** | Pointer to **string** | Name of the Datastore Cluster. | [optional] 
**Type** | Pointer to **string** | Type of the Datastore Cluster. * &#x60;Unknown&#x60; - The nature of the file system is unknown. * &#x60;VMFS&#x60; - It is a Virtual Machine Filesystem. * &#x60;NFS&#x60; - It is a Network File System. * &#x60;vSAN&#x60; - It is a virtual Storage Area Network file system. * &#x60;VirtualVolume&#x60; - A Virtual Volume datastore represents a storage container in a hypervisor server. | [optional] [default to "Unknown"]
**VmCount** | Pointer to **int64** | Number of virtual machines relying on (using) this datastore cluster. | [optional] 

## Methods

### NewVirtualizationBaseDatastoreCluster

`func NewVirtualizationBaseDatastoreCluster(classId string, objectType string, ) *VirtualizationBaseDatastoreCluster`

NewVirtualizationBaseDatastoreCluster instantiates a new VirtualizationBaseDatastoreCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationBaseDatastoreClusterWithDefaults

`func NewVirtualizationBaseDatastoreClusterWithDefaults() *VirtualizationBaseDatastoreCluster`

NewVirtualizationBaseDatastoreClusterWithDefaults instantiates a new VirtualizationBaseDatastoreCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationBaseDatastoreCluster) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationBaseDatastoreCluster) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationBaseDatastoreCluster) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationBaseDatastoreCluster) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationBaseDatastoreCluster) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationBaseDatastoreCluster) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapacity

`func (o *VirtualizationBaseDatastoreCluster) GetCapacity() VirtualizationStorageCapacity`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *VirtualizationBaseDatastoreCluster) GetCapacityOk() (*VirtualizationStorageCapacity, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *VirtualizationBaseDatastoreCluster) SetCapacity(v VirtualizationStorageCapacity)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *VirtualizationBaseDatastoreCluster) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### SetCapacityNil

`func (o *VirtualizationBaseDatastoreCluster) SetCapacityNil(b bool)`

 SetCapacityNil sets the value for Capacity to be an explicit nil

### UnsetCapacity
`func (o *VirtualizationBaseDatastoreCluster) UnsetCapacity()`

UnsetCapacity ensures that no value is present for Capacity, not even an explicit nil
### GetDatastoreCount

`func (o *VirtualizationBaseDatastoreCluster) GetDatastoreCount() int64`

GetDatastoreCount returns the DatastoreCount field if non-nil, zero value otherwise.

### GetDatastoreCountOk

`func (o *VirtualizationBaseDatastoreCluster) GetDatastoreCountOk() (*int64, bool)`

GetDatastoreCountOk returns a tuple with the DatastoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreCount

`func (o *VirtualizationBaseDatastoreCluster) SetDatastoreCount(v int64)`

SetDatastoreCount sets DatastoreCount field to given value.

### HasDatastoreCount

`func (o *VirtualizationBaseDatastoreCluster) HasDatastoreCount() bool`

HasDatastoreCount returns a boolean if a field has been set.

### GetHostCount

`func (o *VirtualizationBaseDatastoreCluster) GetHostCount() int64`

GetHostCount returns the HostCount field if non-nil, zero value otherwise.

### GetHostCountOk

`func (o *VirtualizationBaseDatastoreCluster) GetHostCountOk() (*int64, bool)`

GetHostCountOk returns a tuple with the HostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostCount

`func (o *VirtualizationBaseDatastoreCluster) SetHostCount(v int64)`

SetHostCount sets HostCount field to given value.

### HasHostCount

`func (o *VirtualizationBaseDatastoreCluster) HasHostCount() bool`

HasHostCount returns a boolean if a field has been set.

### GetIdentity

`func (o *VirtualizationBaseDatastoreCluster) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *VirtualizationBaseDatastoreCluster) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *VirtualizationBaseDatastoreCluster) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *VirtualizationBaseDatastoreCluster) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationBaseDatastoreCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationBaseDatastoreCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationBaseDatastoreCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationBaseDatastoreCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *VirtualizationBaseDatastoreCluster) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VirtualizationBaseDatastoreCluster) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VirtualizationBaseDatastoreCluster) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VirtualizationBaseDatastoreCluster) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVmCount

`func (o *VirtualizationBaseDatastoreCluster) GetVmCount() int64`

GetVmCount returns the VmCount field if non-nil, zero value otherwise.

### GetVmCountOk

`func (o *VirtualizationBaseDatastoreCluster) GetVmCountOk() (*int64, bool)`

GetVmCountOk returns a tuple with the VmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCount

`func (o *VirtualizationBaseDatastoreCluster) SetVmCount(v int64)`

SetVmCount sets VmCount field to given value.

### HasVmCount

`func (o *VirtualizationBaseDatastoreCluster) HasVmCount() bool`

HasVmCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


