# VirtualizationBaseDatastoreAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "virtualization.VmwareDatastore"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "virtualization.VmwareDatastore"]
**Capacity** | Pointer to [**NullableVirtualizationStorageCapacity**](VirtualizationStorageCapacity.md) |  | [optional] 
**HostCount** | Pointer to **int64** | Number of hosts attached to or supported-by this datastore. | [optional] 
**Identity** | Pointer to **string** | The internally generated identity of this datastore. This entity is not manipulated by users. It aids in uniquely identifying the datastore object. For VMware, this is a MOR (managed object reference). | [optional] 
**Name** | Pointer to **string** | Name of this datastore supplied by user. It is not the identity of the datastore. The name is subject to user manipulations. | [optional] 
**Type** | Pointer to **string** | A string indicating the type of the datastore (VMFS, NFS, etc). * &#x60;Unknown&#x60; - The nature of the file system is unknown. * &#x60;VMFS&#x60; - It is a Virtual Machine Filesystem. * &#x60;NFS&#x60; - It is a Network File System. * &#x60;vSAN&#x60; - It is a virtual Storage Area Network file system. * &#x60;VirtualVolume&#x60; - A Virtual Volume datastore represents a storage container in a hypervisor server. | [optional] [default to "Unknown"]
**VmCount** | Pointer to **int64** | Number of virtual machines relying on (using) this datastore. | [optional] 

## Methods

### NewVirtualizationBaseDatastoreAllOf

`func NewVirtualizationBaseDatastoreAllOf(classId string, objectType string, ) *VirtualizationBaseDatastoreAllOf`

NewVirtualizationBaseDatastoreAllOf instantiates a new VirtualizationBaseDatastoreAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationBaseDatastoreAllOfWithDefaults

`func NewVirtualizationBaseDatastoreAllOfWithDefaults() *VirtualizationBaseDatastoreAllOf`

NewVirtualizationBaseDatastoreAllOfWithDefaults instantiates a new VirtualizationBaseDatastoreAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationBaseDatastoreAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationBaseDatastoreAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationBaseDatastoreAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationBaseDatastoreAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationBaseDatastoreAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationBaseDatastoreAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapacity

`func (o *VirtualizationBaseDatastoreAllOf) GetCapacity() VirtualizationStorageCapacity`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *VirtualizationBaseDatastoreAllOf) GetCapacityOk() (*VirtualizationStorageCapacity, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *VirtualizationBaseDatastoreAllOf) SetCapacity(v VirtualizationStorageCapacity)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *VirtualizationBaseDatastoreAllOf) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### SetCapacityNil

`func (o *VirtualizationBaseDatastoreAllOf) SetCapacityNil(b bool)`

 SetCapacityNil sets the value for Capacity to be an explicit nil

### UnsetCapacity
`func (o *VirtualizationBaseDatastoreAllOf) UnsetCapacity()`

UnsetCapacity ensures that no value is present for Capacity, not even an explicit nil
### GetHostCount

`func (o *VirtualizationBaseDatastoreAllOf) GetHostCount() int64`

GetHostCount returns the HostCount field if non-nil, zero value otherwise.

### GetHostCountOk

`func (o *VirtualizationBaseDatastoreAllOf) GetHostCountOk() (*int64, bool)`

GetHostCountOk returns a tuple with the HostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostCount

`func (o *VirtualizationBaseDatastoreAllOf) SetHostCount(v int64)`

SetHostCount sets HostCount field to given value.

### HasHostCount

`func (o *VirtualizationBaseDatastoreAllOf) HasHostCount() bool`

HasHostCount returns a boolean if a field has been set.

### GetIdentity

`func (o *VirtualizationBaseDatastoreAllOf) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *VirtualizationBaseDatastoreAllOf) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *VirtualizationBaseDatastoreAllOf) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *VirtualizationBaseDatastoreAllOf) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationBaseDatastoreAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationBaseDatastoreAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationBaseDatastoreAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationBaseDatastoreAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *VirtualizationBaseDatastoreAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VirtualizationBaseDatastoreAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VirtualizationBaseDatastoreAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VirtualizationBaseDatastoreAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVmCount

`func (o *VirtualizationBaseDatastoreAllOf) GetVmCount() int64`

GetVmCount returns the VmCount field if non-nil, zero value otherwise.

### GetVmCountOk

`func (o *VirtualizationBaseDatastoreAllOf) GetVmCountOk() (*int64, bool)`

GetVmCountOk returns a tuple with the VmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCount

`func (o *VirtualizationBaseDatastoreAllOf) SetVmCount(v int64)`

SetVmCount sets VmCount field to given value.

### HasVmCount

`func (o *VirtualizationBaseDatastoreAllOf) HasVmCount() bool`

HasVmCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


