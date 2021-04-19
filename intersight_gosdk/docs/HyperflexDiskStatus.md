# HyperflexDiskStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.DiskStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.DiskStatus"]
**DownloadPercentage** | Pointer to **string** | Percentage of download completed. | [optional] [readonly] 
**State** | Pointer to **string** | Current state of the virtual disk. * &#x60;Unknown&#x60; - No details available on the disk state. * &#x60;Succeeded&#x60; - Last operation on the disk has been successful. * &#x60;ImportInProgress&#x60; - Import operation on the disk is in progress. * &#x60;ImportFailed&#x60; - Import operation on the disk has failed. * &#x60;CloneInProgress&#x60; - Disk clone operation on the disk is in progress. * &#x60;CloneFailed&#x60; - Clone operation on the disk has failed. * &#x60;CloneScheduled&#x60; - Clone operation on the disk has been scheduled. * &#x60;ImportScheduled&#x60; - Import operation on the disk has been scheduled. * &#x60;Pending&#x60; - Submitted operation on the disk is currently pending. * &#x60;&#x60; - Disk state is not available. | [optional] [readonly] [default to "Unknown"]
**VolumeHandle** | Pointer to **string** | Identity of the Volume associated with virtual machine disk. | [optional] [readonly] 
**VolumeName** | Pointer to **string** | Name of the Volume associated with virtual machine disk. | [optional] [readonly] 

## Methods

### NewHyperflexDiskStatus

`func NewHyperflexDiskStatus(classId string, objectType string, ) *HyperflexDiskStatus`

NewHyperflexDiskStatus instantiates a new HyperflexDiskStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexDiskStatusWithDefaults

`func NewHyperflexDiskStatusWithDefaults() *HyperflexDiskStatus`

NewHyperflexDiskStatusWithDefaults instantiates a new HyperflexDiskStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexDiskStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexDiskStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexDiskStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexDiskStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexDiskStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexDiskStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDownloadPercentage

`func (o *HyperflexDiskStatus) GetDownloadPercentage() string`

GetDownloadPercentage returns the DownloadPercentage field if non-nil, zero value otherwise.

### GetDownloadPercentageOk

`func (o *HyperflexDiskStatus) GetDownloadPercentageOk() (*string, bool)`

GetDownloadPercentageOk returns a tuple with the DownloadPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadPercentage

`func (o *HyperflexDiskStatus) SetDownloadPercentage(v string)`

SetDownloadPercentage sets DownloadPercentage field to given value.

### HasDownloadPercentage

`func (o *HyperflexDiskStatus) HasDownloadPercentage() bool`

HasDownloadPercentage returns a boolean if a field has been set.

### GetState

`func (o *HyperflexDiskStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HyperflexDiskStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HyperflexDiskStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HyperflexDiskStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVolumeHandle

`func (o *HyperflexDiskStatus) GetVolumeHandle() string`

GetVolumeHandle returns the VolumeHandle field if non-nil, zero value otherwise.

### GetVolumeHandleOk

`func (o *HyperflexDiskStatus) GetVolumeHandleOk() (*string, bool)`

GetVolumeHandleOk returns a tuple with the VolumeHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeHandle

`func (o *HyperflexDiskStatus) SetVolumeHandle(v string)`

SetVolumeHandle sets VolumeHandle field to given value.

### HasVolumeHandle

`func (o *HyperflexDiskStatus) HasVolumeHandle() bool`

HasVolumeHandle returns a boolean if a field has been set.

### GetVolumeName

`func (o *HyperflexDiskStatus) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *HyperflexDiskStatus) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *HyperflexDiskStatus) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *HyperflexDiskStatus) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


