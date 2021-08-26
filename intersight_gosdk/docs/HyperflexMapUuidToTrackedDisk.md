# HyperflexMapUuidToTrackedDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.MapUuidToTrackedDisk"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.MapUuidToTrackedDisk"]
**TrackedDisk** | Pointer to [**NullableHyperflexTrackedDisk**](HyperflexTrackedDisk.md) |  | [optional] 
**Uuid** | Pointer to **string** | Disk unique id for a snapshot. | [optional] [readonly] 

## Methods

### NewHyperflexMapUuidToTrackedDisk

`func NewHyperflexMapUuidToTrackedDisk(classId string, objectType string, ) *HyperflexMapUuidToTrackedDisk`

NewHyperflexMapUuidToTrackedDisk instantiates a new HyperflexMapUuidToTrackedDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexMapUuidToTrackedDiskWithDefaults

`func NewHyperflexMapUuidToTrackedDiskWithDefaults() *HyperflexMapUuidToTrackedDisk`

NewHyperflexMapUuidToTrackedDiskWithDefaults instantiates a new HyperflexMapUuidToTrackedDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexMapUuidToTrackedDisk) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexMapUuidToTrackedDisk) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexMapUuidToTrackedDisk) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexMapUuidToTrackedDisk) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexMapUuidToTrackedDisk) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexMapUuidToTrackedDisk) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTrackedDisk

`func (o *HyperflexMapUuidToTrackedDisk) GetTrackedDisk() HyperflexTrackedDisk`

GetTrackedDisk returns the TrackedDisk field if non-nil, zero value otherwise.

### GetTrackedDiskOk

`func (o *HyperflexMapUuidToTrackedDisk) GetTrackedDiskOk() (*HyperflexTrackedDisk, bool)`

GetTrackedDiskOk returns a tuple with the TrackedDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackedDisk

`func (o *HyperflexMapUuidToTrackedDisk) SetTrackedDisk(v HyperflexTrackedDisk)`

SetTrackedDisk sets TrackedDisk field to given value.

### HasTrackedDisk

`func (o *HyperflexMapUuidToTrackedDisk) HasTrackedDisk() bool`

HasTrackedDisk returns a boolean if a field has been set.

### SetTrackedDiskNil

`func (o *HyperflexMapUuidToTrackedDisk) SetTrackedDiskNil(b bool)`

 SetTrackedDiskNil sets the value for TrackedDisk to be an explicit nil

### UnsetTrackedDisk
`func (o *HyperflexMapUuidToTrackedDisk) UnsetTrackedDisk()`

UnsetTrackedDisk ensures that no value is present for TrackedDisk, not even an explicit nil
### GetUuid

`func (o *HyperflexMapUuidToTrackedDisk) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HyperflexMapUuidToTrackedDisk) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HyperflexMapUuidToTrackedDisk) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HyperflexMapUuidToTrackedDisk) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


