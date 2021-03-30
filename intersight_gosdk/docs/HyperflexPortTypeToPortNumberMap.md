# HyperflexPortTypeToPortNumberMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.PortTypeToPortNumberMap"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.PortTypeToPortNumberMap"]
**I16** | Pointer to **int64** | Integer describing port type to port number map. | [optional] [readonly] 
**String** | Pointer to **string** | String describing port type to port number map. | [optional] [readonly] 

## Methods

### NewHyperflexPortTypeToPortNumberMap

`func NewHyperflexPortTypeToPortNumberMap(classId string, objectType string, ) *HyperflexPortTypeToPortNumberMap`

NewHyperflexPortTypeToPortNumberMap instantiates a new HyperflexPortTypeToPortNumberMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexPortTypeToPortNumberMapWithDefaults

`func NewHyperflexPortTypeToPortNumberMapWithDefaults() *HyperflexPortTypeToPortNumberMap`

NewHyperflexPortTypeToPortNumberMapWithDefaults instantiates a new HyperflexPortTypeToPortNumberMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexPortTypeToPortNumberMap) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexPortTypeToPortNumberMap) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexPortTypeToPortNumberMap) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexPortTypeToPortNumberMap) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexPortTypeToPortNumberMap) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexPortTypeToPortNumberMap) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetI16

`func (o *HyperflexPortTypeToPortNumberMap) GetI16() int64`

GetI16 returns the I16 field if non-nil, zero value otherwise.

### GetI16Ok

`func (o *HyperflexPortTypeToPortNumberMap) GetI16Ok() (*int64, bool)`

GetI16Ok returns a tuple with the I16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI16

`func (o *HyperflexPortTypeToPortNumberMap) SetI16(v int64)`

SetI16 sets I16 field to given value.

### HasI16

`func (o *HyperflexPortTypeToPortNumberMap) HasI16() bool`

HasI16 returns a boolean if a field has been set.

### GetString

`func (o *HyperflexPortTypeToPortNumberMap) GetString() string`

GetString returns the String field if non-nil, zero value otherwise.

### GetStringOk

`func (o *HyperflexPortTypeToPortNumberMap) GetStringOk() (*string, bool)`

GetStringOk returns a tuple with the String field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetString

`func (o *HyperflexPortTypeToPortNumberMap) SetString(v string)`

SetString sets String field to given value.

### HasString

`func (o *HyperflexPortTypeToPortNumberMap) HasString() bool`

HasString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


