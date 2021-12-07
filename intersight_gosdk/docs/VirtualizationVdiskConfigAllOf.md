# VirtualizationVdiskConfigAllOf

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

### NewVirtualizationVdiskConfigAllOf

`func NewVirtualizationVdiskConfigAllOf(classId string, objectType string, ) *VirtualizationVdiskConfigAllOf`

NewVirtualizationVdiskConfigAllOf instantiates a new VirtualizationVdiskConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVdiskConfigAllOfWithDefaults

`func NewVirtualizationVdiskConfigAllOfWithDefaults() *VirtualizationVdiskConfigAllOf`

NewVirtualizationVdiskConfigAllOfWithDefaults instantiates a new VirtualizationVdiskConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVdiskConfigAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVdiskConfigAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVdiskConfigAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVdiskConfigAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVdiskConfigAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVdiskConfigAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccessMode

`func (o *VirtualizationVdiskConfigAllOf) GetAccessMode() string`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *VirtualizationVdiskConfigAllOf) GetAccessModeOk() (*string, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *VirtualizationVdiskConfigAllOf) SetAccessMode(v string)`

SetAccessMode sets AccessMode field to given value.

### HasAccessMode

`func (o *VirtualizationVdiskConfigAllOf) HasAccessMode() bool`

HasAccessMode returns a boolean if a field has been set.

### GetCapacity

`func (o *VirtualizationVdiskConfigAllOf) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *VirtualizationVdiskConfigAllOf) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *VirtualizationVdiskConfigAllOf) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *VirtualizationVdiskConfigAllOf) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetMode

`func (o *VirtualizationVdiskConfigAllOf) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *VirtualizationVdiskConfigAllOf) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *VirtualizationVdiskConfigAllOf) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *VirtualizationVdiskConfigAllOf) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationVdiskConfigAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationVdiskConfigAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationVdiskConfigAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationVdiskConfigAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceFilePath

`func (o *VirtualizationVdiskConfigAllOf) GetSourceFilePath() string`

GetSourceFilePath returns the SourceFilePath field if non-nil, zero value otherwise.

### GetSourceFilePathOk

`func (o *VirtualizationVdiskConfigAllOf) GetSourceFilePathOk() (*string, bool)`

GetSourceFilePathOk returns a tuple with the SourceFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFilePath

`func (o *VirtualizationVdiskConfigAllOf) SetSourceFilePath(v string)`

SetSourceFilePath sets SourceFilePath field to given value.

### HasSourceFilePath

`func (o *VirtualizationVdiskConfigAllOf) HasSourceFilePath() bool`

HasSourceFilePath returns a boolean if a field has been set.

### GetSourceVirtualDisk

`func (o *VirtualizationVdiskConfigAllOf) GetSourceVirtualDisk() string`

GetSourceVirtualDisk returns the SourceVirtualDisk field if non-nil, zero value otherwise.

### GetSourceVirtualDiskOk

`func (o *VirtualizationVdiskConfigAllOf) GetSourceVirtualDiskOk() (*string, bool)`

GetSourceVirtualDiskOk returns a tuple with the SourceVirtualDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVirtualDisk

`func (o *VirtualizationVdiskConfigAllOf) SetSourceVirtualDisk(v string)`

SetSourceVirtualDisk sets SourceVirtualDisk field to given value.

### HasSourceVirtualDisk

`func (o *VirtualizationVdiskConfigAllOf) HasSourceVirtualDisk() bool`

HasSourceVirtualDisk returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualizationVdiskConfigAllOf) GetStatus() VirtualizationDiskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualizationVdiskConfigAllOf) GetStatusOk() (*VirtualizationDiskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualizationVdiskConfigAllOf) SetStatus(v VirtualizationDiskStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualizationVdiskConfigAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *VirtualizationVdiskConfigAllOf) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *VirtualizationVdiskConfigAllOf) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUuid

`func (o *VirtualizationVdiskConfigAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VirtualizationVdiskConfigAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VirtualizationVdiskConfigAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VirtualizationVdiskConfigAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


