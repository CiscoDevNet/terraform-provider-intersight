# HyperflexTrackedDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.TrackedDisk"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.TrackedDisk"]
**DiskFiles** | Pointer to [**[]HyperflexTrackedFile**](HyperflexTrackedFile.md) |  | [optional] 
**DiskType** | Pointer to **string** | Disk type for a vm virtual disk. * &#x60;NONE&#x60; - The disk type for this VM is None. * &#x60;NATIVE&#x60; - The disk type for this VM is Native. * &#x60;NONNATIVE&#x60; - The disk type for this VM is Non-Native. | [optional] [readonly] [default to "NONE"]

## Methods

### NewHyperflexTrackedDisk

`func NewHyperflexTrackedDisk(classId string, objectType string, ) *HyperflexTrackedDisk`

NewHyperflexTrackedDisk instantiates a new HyperflexTrackedDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexTrackedDiskWithDefaults

`func NewHyperflexTrackedDiskWithDefaults() *HyperflexTrackedDisk`

NewHyperflexTrackedDiskWithDefaults instantiates a new HyperflexTrackedDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexTrackedDisk) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexTrackedDisk) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexTrackedDisk) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexTrackedDisk) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexTrackedDisk) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexTrackedDisk) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDiskFiles

`func (o *HyperflexTrackedDisk) GetDiskFiles() []HyperflexTrackedFile`

GetDiskFiles returns the DiskFiles field if non-nil, zero value otherwise.

### GetDiskFilesOk

`func (o *HyperflexTrackedDisk) GetDiskFilesOk() (*[]HyperflexTrackedFile, bool)`

GetDiskFilesOk returns a tuple with the DiskFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskFiles

`func (o *HyperflexTrackedDisk) SetDiskFiles(v []HyperflexTrackedFile)`

SetDiskFiles sets DiskFiles field to given value.

### HasDiskFiles

`func (o *HyperflexTrackedDisk) HasDiskFiles() bool`

HasDiskFiles returns a boolean if a field has been set.

### SetDiskFilesNil

`func (o *HyperflexTrackedDisk) SetDiskFilesNil(b bool)`

 SetDiskFilesNil sets the value for DiskFiles to be an explicit nil

### UnsetDiskFiles
`func (o *HyperflexTrackedDisk) UnsetDiskFiles()`

UnsetDiskFiles ensures that no value is present for DiskFiles, not even an explicit nil
### GetDiskType

`func (o *HyperflexTrackedDisk) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *HyperflexTrackedDisk) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *HyperflexTrackedDisk) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *HyperflexTrackedDisk) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


