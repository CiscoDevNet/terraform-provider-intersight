# HyperflexMapClusterIdToProtectionInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.MapClusterIdToProtectionInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.MapClusterIdToProtectionInfo"]
**ClusterId** | Pointer to **string** | The Cluster Id we are using to map to ProtectionInfo. | [optional] [readonly] 
**ProtectionInfo** | Pointer to [**NullableHyperflexProtectionInfo**](HyperflexProtectionInfo.md) |  | [optional] 

## Methods

### NewHyperflexMapClusterIdToProtectionInfoAllOf

`func NewHyperflexMapClusterIdToProtectionInfoAllOf(classId string, objectType string, ) *HyperflexMapClusterIdToProtectionInfoAllOf`

NewHyperflexMapClusterIdToProtectionInfoAllOf instantiates a new HyperflexMapClusterIdToProtectionInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexMapClusterIdToProtectionInfoAllOfWithDefaults

`func NewHyperflexMapClusterIdToProtectionInfoAllOfWithDefaults() *HyperflexMapClusterIdToProtectionInfoAllOf`

NewHyperflexMapClusterIdToProtectionInfoAllOfWithDefaults instantiates a new HyperflexMapClusterIdToProtectionInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexMapClusterIdToProtectionInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexMapClusterIdToProtectionInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexMapClusterIdToProtectionInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexMapClusterIdToProtectionInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexMapClusterIdToProtectionInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexMapClusterIdToProtectionInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterId

`func (o *HyperflexMapClusterIdToProtectionInfoAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *HyperflexMapClusterIdToProtectionInfoAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *HyperflexMapClusterIdToProtectionInfoAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *HyperflexMapClusterIdToProtectionInfoAllOf) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetProtectionInfo

`func (o *HyperflexMapClusterIdToProtectionInfoAllOf) GetProtectionInfo() HyperflexProtectionInfo`

GetProtectionInfo returns the ProtectionInfo field if non-nil, zero value otherwise.

### GetProtectionInfoOk

`func (o *HyperflexMapClusterIdToProtectionInfoAllOf) GetProtectionInfoOk() (*HyperflexProtectionInfo, bool)`

GetProtectionInfoOk returns a tuple with the ProtectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionInfo

`func (o *HyperflexMapClusterIdToProtectionInfoAllOf) SetProtectionInfo(v HyperflexProtectionInfo)`

SetProtectionInfo sets ProtectionInfo field to given value.

### HasProtectionInfo

`func (o *HyperflexMapClusterIdToProtectionInfoAllOf) HasProtectionInfo() bool`

HasProtectionInfo returns a boolean if a field has been set.

### SetProtectionInfoNil

`func (o *HyperflexMapClusterIdToProtectionInfoAllOf) SetProtectionInfoNil(b bool)`

 SetProtectionInfoNil sets the value for ProtectionInfo to be an explicit nil

### UnsetProtectionInfo
`func (o *HyperflexMapClusterIdToProtectionInfoAllOf) UnsetProtectionInfo()`

UnsetProtectionInfo ensures that no value is present for ProtectionInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


