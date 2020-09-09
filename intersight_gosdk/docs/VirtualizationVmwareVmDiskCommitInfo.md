# VirtualizationVmwareVmDiskCommitInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommittedDisk** | Pointer to **int64** | Disk committed in bytes on this virtual machine (disk space used up). | [optional] 
**UnCommittedDisk** | Pointer to **int64** | Total uncommitted disk space that is available for use (in bytes). | [optional] 
**UnsharedDisk** | Pointer to **int64** | Total unshared disk space (in bytes). | [optional] 

## Methods

### NewVirtualizationVmwareVmDiskCommitInfo

`func NewVirtualizationVmwareVmDiskCommitInfo() *VirtualizationVmwareVmDiskCommitInfo`

NewVirtualizationVmwareVmDiskCommitInfo instantiates a new VirtualizationVmwareVmDiskCommitInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareVmDiskCommitInfoWithDefaults

`func NewVirtualizationVmwareVmDiskCommitInfoWithDefaults() *VirtualizationVmwareVmDiskCommitInfo`

NewVirtualizationVmwareVmDiskCommitInfoWithDefaults instantiates a new VirtualizationVmwareVmDiskCommitInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


