# VirtualizationVmEsxiDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmEsxiDisk"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmEsxiDisk"]
**Capacity** | Pointer to **string** | Disk capacity to be provided with units example - 10Gi. | [optional] 
**Datastore** | Pointer to **string** | Datastore of the disk from the space is allocated. | [optional] 
**Diskmode** | Pointer to **string** | Mode of the disk, like persistent, independent persistent. | [optional] 
**Name** | Pointer to **string** | Name of the virtual disk. | [optional] 
**StorageAllocation** | Pointer to **string** | Specify the allocation type as thin or thick. | [optional] 
**StorageController** | Pointer to **string** | Controller name of the disk, if not specified uses the default one. | [optional] 

## Methods

### NewVirtualizationVmEsxiDisk

`func NewVirtualizationVmEsxiDisk(classId string, objectType string, ) *VirtualizationVmEsxiDisk`

NewVirtualizationVmEsxiDisk instantiates a new VirtualizationVmEsxiDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmEsxiDiskWithDefaults

`func NewVirtualizationVmEsxiDiskWithDefaults() *VirtualizationVmEsxiDisk`

NewVirtualizationVmEsxiDiskWithDefaults instantiates a new VirtualizationVmEsxiDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmEsxiDisk) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmEsxiDisk) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmEsxiDisk) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmEsxiDisk) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmEsxiDisk) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmEsxiDisk) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapacity

`func (o *VirtualizationVmEsxiDisk) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *VirtualizationVmEsxiDisk) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *VirtualizationVmEsxiDisk) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *VirtualizationVmEsxiDisk) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetDatastore

`func (o *VirtualizationVmEsxiDisk) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *VirtualizationVmEsxiDisk) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *VirtualizationVmEsxiDisk) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *VirtualizationVmEsxiDisk) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetDiskmode

`func (o *VirtualizationVmEsxiDisk) GetDiskmode() string`

GetDiskmode returns the Diskmode field if non-nil, zero value otherwise.

### GetDiskmodeOk

`func (o *VirtualizationVmEsxiDisk) GetDiskmodeOk() (*string, bool)`

GetDiskmodeOk returns a tuple with the Diskmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskmode

`func (o *VirtualizationVmEsxiDisk) SetDiskmode(v string)`

SetDiskmode sets Diskmode field to given value.

### HasDiskmode

`func (o *VirtualizationVmEsxiDisk) HasDiskmode() bool`

HasDiskmode returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationVmEsxiDisk) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationVmEsxiDisk) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationVmEsxiDisk) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationVmEsxiDisk) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStorageAllocation

`func (o *VirtualizationVmEsxiDisk) GetStorageAllocation() string`

GetStorageAllocation returns the StorageAllocation field if non-nil, zero value otherwise.

### GetStorageAllocationOk

`func (o *VirtualizationVmEsxiDisk) GetStorageAllocationOk() (*string, bool)`

GetStorageAllocationOk returns a tuple with the StorageAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAllocation

`func (o *VirtualizationVmEsxiDisk) SetStorageAllocation(v string)`

SetStorageAllocation sets StorageAllocation field to given value.

### HasStorageAllocation

`func (o *VirtualizationVmEsxiDisk) HasStorageAllocation() bool`

HasStorageAllocation returns a boolean if a field has been set.

### GetStorageController

`func (o *VirtualizationVmEsxiDisk) GetStorageController() string`

GetStorageController returns the StorageController field if non-nil, zero value otherwise.

### GetStorageControllerOk

`func (o *VirtualizationVmEsxiDisk) GetStorageControllerOk() (*string, bool)`

GetStorageControllerOk returns a tuple with the StorageController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageController

`func (o *VirtualizationVmEsxiDisk) SetStorageController(v string)`

SetStorageController sets StorageController field to given value.

### HasStorageController

`func (o *VirtualizationVmEsxiDisk) HasStorageController() bool`

HasStorageController returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


