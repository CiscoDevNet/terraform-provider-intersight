# HyperflexVdiskConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.VdiskConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.VdiskConfig"]
**AccessMode** | Pointer to **string** | Access mode of the virtual disk. * &#x60;ReadWriteOnce&#x60; - Read write permisisons to a Virtual disk by a single virtual machine. * &#x60;ReadWriteMany&#x60; - Read write permisisons to a Virtual disk by multiple virtual machines. * &#x60;ReadOnlyMany&#x60; - Read only permisisons to a Virtual disk by multiple virtual machines. * &#x60;&#x60; - Unknown disk access mode. | [optional] [readonly] [default to "ReadWriteOnce"]
**Capacity** | Pointer to **int64** | Disk capacity represented in bytes. | [optional] [readonly] 
**Mode** | Pointer to **string** | File mode of the disk, example - Filesystem, Block. * &#x60;Block&#x60; - It is a Block virtual disk. * &#x60;Filesystem&#x60; - It is a File system virtual disk. * &#x60;&#x60; - Disk mode is either unknown or not supported. | [optional] [readonly] [default to "Block"]
**Name** | Pointer to **string** | Name of the virtual disk. | [optional] [readonly] 
**SourceFilePath** | Pointer to **string** | Source file path associated with virtual machine disk. | [optional] [readonly] 
**SourceVirtualDisk** | Pointer to **string** | Source disk name from where the clone is done. | [optional] [readonly] 
**Status** | Pointer to [**NullableHyperflexDiskStatus**](HyperflexDiskStatus.md) |  | [optional] 
**Uuid** | Pointer to **string** | UUID of the virtual disk. | [optional] [readonly] 

## Methods

### NewHyperflexVdiskConfig

`func NewHyperflexVdiskConfig(classId string, objectType string, ) *HyperflexVdiskConfig`

NewHyperflexVdiskConfig instantiates a new HyperflexVdiskConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexVdiskConfigWithDefaults

`func NewHyperflexVdiskConfigWithDefaults() *HyperflexVdiskConfig`

NewHyperflexVdiskConfigWithDefaults instantiates a new HyperflexVdiskConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexVdiskConfig) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexVdiskConfig) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexVdiskConfig) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexVdiskConfig) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexVdiskConfig) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexVdiskConfig) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccessMode

`func (o *HyperflexVdiskConfig) GetAccessMode() string`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *HyperflexVdiskConfig) GetAccessModeOk() (*string, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *HyperflexVdiskConfig) SetAccessMode(v string)`

SetAccessMode sets AccessMode field to given value.

### HasAccessMode

`func (o *HyperflexVdiskConfig) HasAccessMode() bool`

HasAccessMode returns a boolean if a field has been set.

### GetCapacity

`func (o *HyperflexVdiskConfig) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *HyperflexVdiskConfig) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *HyperflexVdiskConfig) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *HyperflexVdiskConfig) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetMode

`func (o *HyperflexVdiskConfig) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *HyperflexVdiskConfig) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *HyperflexVdiskConfig) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *HyperflexVdiskConfig) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *HyperflexVdiskConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexVdiskConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexVdiskConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexVdiskConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceFilePath

`func (o *HyperflexVdiskConfig) GetSourceFilePath() string`

GetSourceFilePath returns the SourceFilePath field if non-nil, zero value otherwise.

### GetSourceFilePathOk

`func (o *HyperflexVdiskConfig) GetSourceFilePathOk() (*string, bool)`

GetSourceFilePathOk returns a tuple with the SourceFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFilePath

`func (o *HyperflexVdiskConfig) SetSourceFilePath(v string)`

SetSourceFilePath sets SourceFilePath field to given value.

### HasSourceFilePath

`func (o *HyperflexVdiskConfig) HasSourceFilePath() bool`

HasSourceFilePath returns a boolean if a field has been set.

### GetSourceVirtualDisk

`func (o *HyperflexVdiskConfig) GetSourceVirtualDisk() string`

GetSourceVirtualDisk returns the SourceVirtualDisk field if non-nil, zero value otherwise.

### GetSourceVirtualDiskOk

`func (o *HyperflexVdiskConfig) GetSourceVirtualDiskOk() (*string, bool)`

GetSourceVirtualDiskOk returns a tuple with the SourceVirtualDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVirtualDisk

`func (o *HyperflexVdiskConfig) SetSourceVirtualDisk(v string)`

SetSourceVirtualDisk sets SourceVirtualDisk field to given value.

### HasSourceVirtualDisk

`func (o *HyperflexVdiskConfig) HasSourceVirtualDisk() bool`

HasSourceVirtualDisk returns a boolean if a field has been set.

### GetStatus

`func (o *HyperflexVdiskConfig) GetStatus() HyperflexDiskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HyperflexVdiskConfig) GetStatusOk() (*HyperflexDiskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HyperflexVdiskConfig) SetStatus(v HyperflexDiskStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HyperflexVdiskConfig) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *HyperflexVdiskConfig) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *HyperflexVdiskConfig) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUuid

`func (o *HyperflexVdiskConfig) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HyperflexVdiskConfig) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HyperflexVdiskConfig) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HyperflexVdiskConfig) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


