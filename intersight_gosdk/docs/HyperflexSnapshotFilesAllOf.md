# HyperflexSnapshotFilesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.SnapshotFiles"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.SnapshotFiles"]
**NameTrackedFiles** | Pointer to [**[]HyperflexFilePath**](HyperflexFilePath.md) |  | [optional] 
**UuidTrackedDisksMap** | Pointer to [**[]HyperflexMapUuidToTrackedDisk**](HyperflexMapUuidToTrackedDisk.md) |  | [optional] 

## Methods

### NewHyperflexSnapshotFilesAllOf

`func NewHyperflexSnapshotFilesAllOf(classId string, objectType string, ) *HyperflexSnapshotFilesAllOf`

NewHyperflexSnapshotFilesAllOf instantiates a new HyperflexSnapshotFilesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexSnapshotFilesAllOfWithDefaults

`func NewHyperflexSnapshotFilesAllOfWithDefaults() *HyperflexSnapshotFilesAllOf`

NewHyperflexSnapshotFilesAllOfWithDefaults instantiates a new HyperflexSnapshotFilesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexSnapshotFilesAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexSnapshotFilesAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexSnapshotFilesAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexSnapshotFilesAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexSnapshotFilesAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexSnapshotFilesAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetNameTrackedFiles

`func (o *HyperflexSnapshotFilesAllOf) GetNameTrackedFiles() []HyperflexFilePath`

GetNameTrackedFiles returns the NameTrackedFiles field if non-nil, zero value otherwise.

### GetNameTrackedFilesOk

`func (o *HyperflexSnapshotFilesAllOf) GetNameTrackedFilesOk() (*[]HyperflexFilePath, bool)`

GetNameTrackedFilesOk returns a tuple with the NameTrackedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameTrackedFiles

`func (o *HyperflexSnapshotFilesAllOf) SetNameTrackedFiles(v []HyperflexFilePath)`

SetNameTrackedFiles sets NameTrackedFiles field to given value.

### HasNameTrackedFiles

`func (o *HyperflexSnapshotFilesAllOf) HasNameTrackedFiles() bool`

HasNameTrackedFiles returns a boolean if a field has been set.

### SetNameTrackedFilesNil

`func (o *HyperflexSnapshotFilesAllOf) SetNameTrackedFilesNil(b bool)`

 SetNameTrackedFilesNil sets the value for NameTrackedFiles to be an explicit nil

### UnsetNameTrackedFiles
`func (o *HyperflexSnapshotFilesAllOf) UnsetNameTrackedFiles()`

UnsetNameTrackedFiles ensures that no value is present for NameTrackedFiles, not even an explicit nil
### GetUuidTrackedDisksMap

`func (o *HyperflexSnapshotFilesAllOf) GetUuidTrackedDisksMap() []HyperflexMapUuidToTrackedDisk`

GetUuidTrackedDisksMap returns the UuidTrackedDisksMap field if non-nil, zero value otherwise.

### GetUuidTrackedDisksMapOk

`func (o *HyperflexSnapshotFilesAllOf) GetUuidTrackedDisksMapOk() (*[]HyperflexMapUuidToTrackedDisk, bool)`

GetUuidTrackedDisksMapOk returns a tuple with the UuidTrackedDisksMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuidTrackedDisksMap

`func (o *HyperflexSnapshotFilesAllOf) SetUuidTrackedDisksMap(v []HyperflexMapUuidToTrackedDisk)`

SetUuidTrackedDisksMap sets UuidTrackedDisksMap field to given value.

### HasUuidTrackedDisksMap

`func (o *HyperflexSnapshotFilesAllOf) HasUuidTrackedDisksMap() bool`

HasUuidTrackedDisksMap returns a boolean if a field has been set.

### SetUuidTrackedDisksMapNil

`func (o *HyperflexSnapshotFilesAllOf) SetUuidTrackedDisksMapNil(b bool)`

 SetUuidTrackedDisksMapNil sets the value for UuidTrackedDisksMap to be an explicit nil

### UnsetUuidTrackedDisksMap
`func (o *HyperflexSnapshotFilesAllOf) UnsetUuidTrackedDisksMap()`

UnsetUuidTrackedDisksMap ensures that no value is present for UuidTrackedDisksMap, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


