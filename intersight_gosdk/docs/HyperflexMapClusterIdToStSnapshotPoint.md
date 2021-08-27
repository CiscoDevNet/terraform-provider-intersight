# HyperflexMapClusterIdToStSnapshotPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.MapClusterIdToStSnapshotPoint"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.MapClusterIdToStSnapshotPoint"]
**ClusterId** | Pointer to **string** | ClusterId of the snapshot point. | [optional] [readonly] 
**SnapshotPoint** | Pointer to [**NullableHyperflexSnapshotPoint**](HyperflexSnapshotPoint.md) |  | [optional] 

## Methods

### NewHyperflexMapClusterIdToStSnapshotPoint

`func NewHyperflexMapClusterIdToStSnapshotPoint(classId string, objectType string, ) *HyperflexMapClusterIdToStSnapshotPoint`

NewHyperflexMapClusterIdToStSnapshotPoint instantiates a new HyperflexMapClusterIdToStSnapshotPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexMapClusterIdToStSnapshotPointWithDefaults

`func NewHyperflexMapClusterIdToStSnapshotPointWithDefaults() *HyperflexMapClusterIdToStSnapshotPoint`

NewHyperflexMapClusterIdToStSnapshotPointWithDefaults instantiates a new HyperflexMapClusterIdToStSnapshotPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexMapClusterIdToStSnapshotPoint) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexMapClusterIdToStSnapshotPoint) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexMapClusterIdToStSnapshotPoint) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexMapClusterIdToStSnapshotPoint) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexMapClusterIdToStSnapshotPoint) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexMapClusterIdToStSnapshotPoint) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterId

`func (o *HyperflexMapClusterIdToStSnapshotPoint) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *HyperflexMapClusterIdToStSnapshotPoint) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *HyperflexMapClusterIdToStSnapshotPoint) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *HyperflexMapClusterIdToStSnapshotPoint) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetSnapshotPoint

`func (o *HyperflexMapClusterIdToStSnapshotPoint) GetSnapshotPoint() HyperflexSnapshotPoint`

GetSnapshotPoint returns the SnapshotPoint field if non-nil, zero value otherwise.

### GetSnapshotPointOk

`func (o *HyperflexMapClusterIdToStSnapshotPoint) GetSnapshotPointOk() (*HyperflexSnapshotPoint, bool)`

GetSnapshotPointOk returns a tuple with the SnapshotPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotPoint

`func (o *HyperflexMapClusterIdToStSnapshotPoint) SetSnapshotPoint(v HyperflexSnapshotPoint)`

SetSnapshotPoint sets SnapshotPoint field to given value.

### HasSnapshotPoint

`func (o *HyperflexMapClusterIdToStSnapshotPoint) HasSnapshotPoint() bool`

HasSnapshotPoint returns a boolean if a field has been set.

### SetSnapshotPointNil

`func (o *HyperflexMapClusterIdToStSnapshotPoint) SetSnapshotPointNil(b bool)`

 SetSnapshotPointNil sets the value for SnapshotPoint to be an explicit nil

### UnsetSnapshotPoint
`func (o *HyperflexMapClusterIdToStSnapshotPoint) UnsetSnapshotPoint()`

UnsetSnapshotPoint ensures that no value is present for SnapshotPoint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


