# HyperflexHxUuIdDtAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HxUuIdDt"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HxUuIdDt"]
**Links** | Pointer to [**[]HyperflexHxLinkDt**](HyperflexHxLinkDt.md) |  | [optional] 
**Uuid** | Pointer to **string** | The unique identifier string of an entity. | [optional] [readonly] 

## Methods

### NewHyperflexHxUuIdDtAllOf

`func NewHyperflexHxUuIdDtAllOf(classId string, objectType string, ) *HyperflexHxUuIdDtAllOf`

NewHyperflexHxUuIdDtAllOf instantiates a new HyperflexHxUuIdDtAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxUuIdDtAllOfWithDefaults

`func NewHyperflexHxUuIdDtAllOfWithDefaults() *HyperflexHxUuIdDtAllOf`

NewHyperflexHxUuIdDtAllOfWithDefaults instantiates a new HyperflexHxUuIdDtAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxUuIdDtAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxUuIdDtAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxUuIdDtAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxUuIdDtAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxUuIdDtAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxUuIdDtAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLinks

`func (o *HyperflexHxUuIdDtAllOf) GetLinks() []HyperflexHxLinkDt`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *HyperflexHxUuIdDtAllOf) GetLinksOk() (*[]HyperflexHxLinkDt, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *HyperflexHxUuIdDtAllOf) SetLinks(v []HyperflexHxLinkDt)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *HyperflexHxUuIdDtAllOf) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *HyperflexHxUuIdDtAllOf) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *HyperflexHxUuIdDtAllOf) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetUuid

`func (o *HyperflexHxUuIdDtAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HyperflexHxUuIdDtAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HyperflexHxUuIdDtAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HyperflexHxUuIdDtAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


