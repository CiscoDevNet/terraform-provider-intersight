# VirtualizationVmwareVmDiskCommitInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareVmDiskCommitInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareVmDiskCommitInfo"]
**CommittedDisk** | Pointer to **int64** | Disk committed in bytes on this virtual machine (disk space used up). | [optional] 
**UnCommittedDisk** | Pointer to **int64** | Total uncommitted disk space that is available for use (in bytes). | [optional] 
**UnsharedDisk** | Pointer to **int64** | Total unshared disk space (in bytes). | [optional] 

## Methods

### NewVirtualizationVmwareVmDiskCommitInfo

`func NewVirtualizationVmwareVmDiskCommitInfo(classId string, objectType string, ) *VirtualizationVmwareVmDiskCommitInfo`

NewVirtualizationVmwareVmDiskCommitInfo instantiates a new VirtualizationVmwareVmDiskCommitInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareVmDiskCommitInfoWithDefaults

`func NewVirtualizationVmwareVmDiskCommitInfoWithDefaults() *VirtualizationVmwareVmDiskCommitInfo`

NewVirtualizationVmwareVmDiskCommitInfoWithDefaults instantiates a new VirtualizationVmwareVmDiskCommitInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareVmDiskCommitInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareVmDiskCommitInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareVmDiskCommitInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareVmDiskCommitInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareVmDiskCommitInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareVmDiskCommitInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCommittedDisk

`func (o *VirtualizationVmwareVmDiskCommitInfo) GetCommittedDisk() int64`

GetCommittedDisk returns the CommittedDisk field if non-nil, zero value otherwise.

### GetCommittedDiskOk

`func (o *VirtualizationVmwareVmDiskCommitInfo) GetCommittedDiskOk() (*int64, bool)`

GetCommittedDiskOk returns a tuple with the CommittedDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittedDisk

`func (o *VirtualizationVmwareVmDiskCommitInfo) SetCommittedDisk(v int64)`

SetCommittedDisk sets CommittedDisk field to given value.

### HasCommittedDisk

`func (o *VirtualizationVmwareVmDiskCommitInfo) HasCommittedDisk() bool`

HasCommittedDisk returns a boolean if a field has been set.

### GetUnCommittedDisk

`func (o *VirtualizationVmwareVmDiskCommitInfo) GetUnCommittedDisk() int64`

GetUnCommittedDisk returns the UnCommittedDisk field if non-nil, zero value otherwise.

### GetUnCommittedDiskOk

`func (o *VirtualizationVmwareVmDiskCommitInfo) GetUnCommittedDiskOk() (*int64, bool)`

GetUnCommittedDiskOk returns a tuple with the UnCommittedDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnCommittedDisk

`func (o *VirtualizationVmwareVmDiskCommitInfo) SetUnCommittedDisk(v int64)`

SetUnCommittedDisk sets UnCommittedDisk field to given value.

### HasUnCommittedDisk

`func (o *VirtualizationVmwareVmDiskCommitInfo) HasUnCommittedDisk() bool`

HasUnCommittedDisk returns a boolean if a field has been set.

### GetUnsharedDisk

`func (o *VirtualizationVmwareVmDiskCommitInfo) GetUnsharedDisk() int64`

GetUnsharedDisk returns the UnsharedDisk field if non-nil, zero value otherwise.

### GetUnsharedDiskOk

`func (o *VirtualizationVmwareVmDiskCommitInfo) GetUnsharedDiskOk() (*int64, bool)`

GetUnsharedDiskOk returns a tuple with the UnsharedDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsharedDisk

`func (o *VirtualizationVmwareVmDiskCommitInfo) SetUnsharedDisk(v int64)`

SetUnsharedDisk sets UnsharedDisk field to given value.

### HasUnsharedDisk

`func (o *VirtualizationVmwareVmDiskCommitInfo) HasUnsharedDisk() bool`

HasUnsharedDisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


