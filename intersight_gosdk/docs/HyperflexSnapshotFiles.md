# HyperflexSnapshotFiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.SnapshotFiles"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.SnapshotFiles"]
**NameTrackedFiles** | Pointer to [**[]HyperflexFilePath**](HyperflexFilePath.md) |  | [optional] 
**UuidTrackedDisksMap** | Pointer to [**[]HyperflexMapUuidToTrackedDisk**](HyperflexMapUuidToTrackedDisk.md) |  | [optional] 

## Methods

### NewHyperflexSnapshotFiles

`func NewHyperflexSnapshotFiles(classId string, objectType string, ) *HyperflexSnapshotFiles`

NewHyperflexSnapshotFiles instantiates a new HyperflexSnapshotFiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexSnapshotFilesWithDefaults

`func NewHyperflexSnapshotFilesWithDefaults() *HyperflexSnapshotFiles`

NewHyperflexSnapshotFilesWithDefaults instantiates a new HyperflexSnapshotFiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexSnapshotFiles) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexSnapshotFiles) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexSnapshotFiles) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexSnapshotFiles) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexSnapshotFiles) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexSnapshotFiles) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetNameTrackedFiles

`func (o *HyperflexSnapshotFiles) GetNameTrackedFiles() []HyperflexFilePath`

GetNameTrackedFiles returns the NameTrackedFiles field if non-nil, zero value otherwise.

### GetNameTrackedFilesOk

`func (o *HyperflexSnapshotFiles) GetNameTrackedFilesOk() (*[]HyperflexFilePath, bool)`

GetNameTrackedFilesOk returns a tuple with the NameTrackedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameTrackedFiles

`func (o *HyperflexSnapshotFiles) SetNameTrackedFiles(v []HyperflexFilePath)`

SetNameTrackedFiles sets NameTrackedFiles field to given value.

### HasNameTrackedFiles

`func (o *HyperflexSnapshotFiles) HasNameTrackedFiles() bool`

HasNameTrackedFiles returns a boolean if a field has been set.

### SetNameTrackedFilesNil

`func (o *HyperflexSnapshotFiles) SetNameTrackedFilesNil(b bool)`

 SetNameTrackedFilesNil sets the value for NameTrackedFiles to be an explicit nil

### UnsetNameTrackedFiles
`func (o *HyperflexSnapshotFiles) UnsetNameTrackedFiles()`

UnsetNameTrackedFiles ensures that no value is present for NameTrackedFiles, not even an explicit nil
### GetUuidTrackedDisksMap

`func (o *HyperflexSnapshotFiles) GetUuidTrackedDisksMap() []HyperflexMapUuidToTrackedDisk`

GetUuidTrackedDisksMap returns the UuidTrackedDisksMap field if non-nil, zero value otherwise.

### GetUuidTrackedDisksMapOk

`func (o *HyperflexSnapshotFiles) GetUuidTrackedDisksMapOk() (*[]HyperflexMapUuidToTrackedDisk, bool)`

GetUuidTrackedDisksMapOk returns a tuple with the UuidTrackedDisksMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuidTrackedDisksMap

`func (o *HyperflexSnapshotFiles) SetUuidTrackedDisksMap(v []HyperflexMapUuidToTrackedDisk)`

SetUuidTrackedDisksMap sets UuidTrackedDisksMap field to given value.

### HasUuidTrackedDisksMap

`func (o *HyperflexSnapshotFiles) HasUuidTrackedDisksMap() bool`

HasUuidTrackedDisksMap returns a boolean if a field has been set.

### SetUuidTrackedDisksMapNil

`func (o *HyperflexSnapshotFiles) SetUuidTrackedDisksMapNil(b bool)`

 SetUuidTrackedDisksMapNil sets the value for UuidTrackedDisksMap to be an explicit nil

### UnsetUuidTrackedDisksMap
`func (o *HyperflexSnapshotFiles) UnsetUuidTrackedDisksMap()`

UnsetUuidTrackedDisksMap ensures that no value is present for UuidTrackedDisksMap, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


