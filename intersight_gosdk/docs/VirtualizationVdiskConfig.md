# VirtualizationVdiskConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VdiskConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VdiskConfig"]
**AccessMode** | Pointer to **string** | Access mode of the virtual disk. * &#x60;ReadWriteOnce&#x60; - Read write permisisons to a Virtual disk by a single virtual machine. * &#x60;ReadWriteMany&#x60; - Read write permisisons to a Virtual disk by multiple virtual machines. * &#x60;ReadOnlyMany&#x60; - Read only permisisons to a Virtual disk by multiple virtual machines. * &#x60;&#x60; - Unknown disk access mode. | [optional] [readonly] [default to "ReadWriteOnce"]
**Capacity** | Pointer to **int64** | Disk capacity represented in bytes. | [optional] [readonly] 
**Mode** | Pointer to **string** | File mode of the disk, example - Filesystem, Block. * &#x60;Block&#x60; - It is a Block virtual disk. * &#x60;Filesystem&#x60; - It is a File system virtual disk. * &#x60;&#x60; - Disk mode is either unknown or not supported. | [optional] [readonly] [default to "Block"]
**Name** | Pointer to **string** | Name of the virtual disk. | [optional] [readonly] 
**SourceFilePath** | Pointer to **string** | Source file path associated with virtual machine disk. | [optional] [readonly] 
**SourceVirtualDisk** | Pointer to **string** | Source disk name from where the clone is done. | [optional] [readonly] 
**Status** | Pointer to [**NullableVirtualizationDiskStatus**](VirtualizationDiskStatus.md) |  | [optional] 
**Uuid** | Pointer to **string** | UUID of the virtual disk. | [optional] [readonly] 

## Methods

### NewVirtualizationVdiskConfig

`func NewVirtualizationVdiskConfig(classId string, objectType string, ) *VirtualizationVdiskConfig`

NewVirtualizationVdiskConfig instantiates a new VirtualizationVdiskConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVdiskConfigWithDefaults

`func NewVirtualizationVdiskConfigWithDefaults() *VirtualizationVdiskConfig`

NewVirtualizationVdiskConfigWithDefaults instantiates a new VirtualizationVdiskConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVdiskConfig) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVdiskConfig) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVdiskConfig) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVdiskConfig) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVdiskConfig) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVdiskConfig) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccessMode

`func (o *VirtualizationVdiskConfig) GetAccessMode() string`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *VirtualizationVdiskConfig) GetAccessModeOk() (*string, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *VirtualizationVdiskConfig) SetAccessMode(v string)`

SetAccessMode sets AccessMode field to given value.

### HasAccessMode

`func (o *VirtualizationVdiskConfig) HasAccessMode() bool`

HasAccessMode returns a boolean if a field has been set.

### GetCapacity

`func (o *VirtualizationVdiskConfig) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *VirtualizationVdiskConfig) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *VirtualizationVdiskConfig) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *VirtualizationVdiskConfig) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetMode

`func (o *VirtualizationVdiskConfig) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *VirtualizationVdiskConfig) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *VirtualizationVdiskConfig) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *VirtualizationVdiskConfig) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationVdiskConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationVdiskConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationVdiskConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationVdiskConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceFilePath

`func (o *VirtualizationVdiskConfig) GetSourceFilePath() string`

GetSourceFilePath returns the SourceFilePath field if non-nil, zero value otherwise.

### GetSourceFilePathOk

`func (o *VirtualizationVdiskConfig) GetSourceFilePathOk() (*string, bool)`

GetSourceFilePathOk returns a tuple with the SourceFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFilePath

`func (o *VirtualizationVdiskConfig) SetSourceFilePath(v string)`

SetSourceFilePath sets SourceFilePath field to given value.

### HasSourceFilePath

`func (o *VirtualizationVdiskConfig) HasSourceFilePath() bool`

HasSourceFilePath returns a boolean if a field has been set.

### GetSourceVirtualDisk

`func (o *VirtualizationVdiskConfig) GetSourceVirtualDisk() string`

GetSourceVirtualDisk returns the SourceVirtualDisk field if non-nil, zero value otherwise.

### GetSourceVirtualDiskOk

`func (o *VirtualizationVdiskConfig) GetSourceVirtualDiskOk() (*string, bool)`

GetSourceVirtualDiskOk returns a tuple with the SourceVirtualDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVirtualDisk

`func (o *VirtualizationVdiskConfig) SetSourceVirtualDisk(v string)`

SetSourceVirtualDisk sets SourceVirtualDisk field to given value.

### HasSourceVirtualDisk

`func (o *VirtualizationVdiskConfig) HasSourceVirtualDisk() bool`

HasSourceVirtualDisk returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualizationVdiskConfig) GetStatus() VirtualizationDiskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualizationVdiskConfig) GetStatusOk() (*VirtualizationDiskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualizationVdiskConfig) SetStatus(v VirtualizationDiskStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualizationVdiskConfig) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *VirtualizationVdiskConfig) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *VirtualizationVdiskConfig) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUuid

`func (o *VirtualizationVdiskConfig) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VirtualizationVdiskConfig) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VirtualizationVdiskConfig) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VirtualizationVdiskConfig) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


