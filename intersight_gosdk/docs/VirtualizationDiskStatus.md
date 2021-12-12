# VirtualizationDiskStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.DiskStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.DiskStatus"]
**DownloadPercentage** | Pointer to **string** | Percentage of download completed. | [optional] [readonly] 
**Reason** | Pointer to **string** | Reason for virtual disk download failure. | [optional] [readonly] 
**State** | Pointer to **string** | Current state of the virtual disk. * &#x60;Unknown&#x60; - No details available on the disk state. * &#x60;Succeeded&#x60; - Last operation on the disk has been successful. * &#x60;ImportInProgress&#x60; - Import operation on the disk is in progress. * &#x60;ImportFailed&#x60; - Import operation on the disk has failed. * &#x60;CloneInProgress&#x60; - Disk clone operation on the disk is in progress. * &#x60;CloneFailed&#x60; - Clone operation on the disk has failed. * &#x60;CloneScheduled&#x60; - Clone operation on the disk has been scheduled. * &#x60;ImportScheduled&#x60; - Import operation on the disk has been scheduled. * &#x60;Pending&#x60; - Submitted operation on the disk is currently pending. * &#x60;&#x60; - Disk state is not available. * &#x60;Failed&#x60; - Last operation on the disk Failed. | [optional] [readonly] [default to "Unknown"]
**VolumeHandle** | Pointer to **string** | Identity of the Volume associated with virtual machine disk. | [optional] [readonly] 
**VolumeName** | Pointer to **string** | Name of the Volume associated with virtual machine disk. | [optional] [readonly] 
**VolumeVendor** | Pointer to **string** | Name of the Volume Vendor associated with virtual machine disk. | [optional] [readonly] 

## Methods

### NewVirtualizationDiskStatus

`func NewVirtualizationDiskStatus(classId string, objectType string, ) *VirtualizationDiskStatus`

NewVirtualizationDiskStatus instantiates a new VirtualizationDiskStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationDiskStatusWithDefaults

`func NewVirtualizationDiskStatusWithDefaults() *VirtualizationDiskStatus`

NewVirtualizationDiskStatusWithDefaults instantiates a new VirtualizationDiskStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationDiskStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationDiskStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationDiskStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationDiskStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationDiskStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationDiskStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDownloadPercentage

`func (o *VirtualizationDiskStatus) GetDownloadPercentage() string`

GetDownloadPercentage returns the DownloadPercentage field if non-nil, zero value otherwise.

### GetDownloadPercentageOk

`func (o *VirtualizationDiskStatus) GetDownloadPercentageOk() (*string, bool)`

GetDownloadPercentageOk returns a tuple with the DownloadPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadPercentage

`func (o *VirtualizationDiskStatus) SetDownloadPercentage(v string)`

SetDownloadPercentage sets DownloadPercentage field to given value.

### HasDownloadPercentage

`func (o *VirtualizationDiskStatus) HasDownloadPercentage() bool`

HasDownloadPercentage returns a boolean if a field has been set.

### GetReason

`func (o *VirtualizationDiskStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *VirtualizationDiskStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *VirtualizationDiskStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *VirtualizationDiskStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetState

`func (o *VirtualizationDiskStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VirtualizationDiskStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VirtualizationDiskStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *VirtualizationDiskStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVolumeHandle

`func (o *VirtualizationDiskStatus) GetVolumeHandle() string`

GetVolumeHandle returns the VolumeHandle field if non-nil, zero value otherwise.

### GetVolumeHandleOk

`func (o *VirtualizationDiskStatus) GetVolumeHandleOk() (*string, bool)`

GetVolumeHandleOk returns a tuple with the VolumeHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeHandle

`func (o *VirtualizationDiskStatus) SetVolumeHandle(v string)`

SetVolumeHandle sets VolumeHandle field to given value.

### HasVolumeHandle

`func (o *VirtualizationDiskStatus) HasVolumeHandle() bool`

HasVolumeHandle returns a boolean if a field has been set.

### GetVolumeName

`func (o *VirtualizationDiskStatus) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *VirtualizationDiskStatus) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *VirtualizationDiskStatus) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *VirtualizationDiskStatus) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.

### GetVolumeVendor

`func (o *VirtualizationDiskStatus) GetVolumeVendor() string`

GetVolumeVendor returns the VolumeVendor field if non-nil, zero value otherwise.

### GetVolumeVendorOk

`func (o *VirtualizationDiskStatus) GetVolumeVendorOk() (*string, bool)`

GetVolumeVendorOk returns a tuple with the VolumeVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeVendor

`func (o *VirtualizationDiskStatus) SetVolumeVendor(v string)`

SetVolumeVendor sets VolumeVendor field to given value.

### HasVolumeVendor

`func (o *VirtualizationDiskStatus) HasVolumeVendor() bool`

HasVolumeVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


