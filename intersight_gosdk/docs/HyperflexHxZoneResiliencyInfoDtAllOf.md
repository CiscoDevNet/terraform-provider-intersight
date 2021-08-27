# HyperflexHxZoneResiliencyInfoDtAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HxZoneResiliencyInfoDt"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HxZoneResiliencyInfoDt"]
**Name** | Pointer to **string** | The name of the availability zone. | [optional] [readonly] 
**ResiliencyInfo** | Pointer to [**NullableHyperflexHxResiliencyInfoDt**](HyperflexHxResiliencyInfoDt.md) |  | [optional] 

## Methods

### NewHyperflexHxZoneResiliencyInfoDtAllOf

`func NewHyperflexHxZoneResiliencyInfoDtAllOf(classId string, objectType string, ) *HyperflexHxZoneResiliencyInfoDtAllOf`

NewHyperflexHxZoneResiliencyInfoDtAllOf instantiates a new HyperflexHxZoneResiliencyInfoDtAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxZoneResiliencyInfoDtAllOfWithDefaults

`func NewHyperflexHxZoneResiliencyInfoDtAllOfWithDefaults() *HyperflexHxZoneResiliencyInfoDtAllOf`

NewHyperflexHxZoneResiliencyInfoDtAllOfWithDefaults instantiates a new HyperflexHxZoneResiliencyInfoDtAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxZoneResiliencyInfoDtAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxZoneResiliencyInfoDtAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxZoneResiliencyInfoDtAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxZoneResiliencyInfoDtAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxZoneResiliencyInfoDtAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxZoneResiliencyInfoDtAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *HyperflexHxZoneResiliencyInfoDtAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexHxZoneResiliencyInfoDtAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexHxZoneResiliencyInfoDtAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexHxZoneResiliencyInfoDtAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResiliencyInfo

`func (o *HyperflexHxZoneResiliencyInfoDtAllOf) GetResiliencyInfo() HyperflexHxResiliencyInfoDt`

GetResiliencyInfo returns the ResiliencyInfo field if non-nil, zero value otherwise.

### GetResiliencyInfoOk

`func (o *HyperflexHxZoneResiliencyInfoDtAllOf) GetResiliencyInfoOk() (*HyperflexHxResiliencyInfoDt, bool)`

GetResiliencyInfoOk returns a tuple with the ResiliencyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResiliencyInfo

`func (o *HyperflexHxZoneResiliencyInfoDtAllOf) SetResiliencyInfo(v HyperflexHxResiliencyInfoDt)`

SetResiliencyInfo sets ResiliencyInfo field to given value.

### HasResiliencyInfo

`func (o *HyperflexHxZoneResiliencyInfoDtAllOf) HasResiliencyInfo() bool`

HasResiliencyInfo returns a boolean if a field has been set.

### SetResiliencyInfoNil

`func (o *HyperflexHxZoneResiliencyInfoDtAllOf) SetResiliencyInfoNil(b bool)`

 SetResiliencyInfoNil sets the value for ResiliencyInfo to be an explicit nil

### UnsetResiliencyInfo
`func (o *HyperflexHxZoneResiliencyInfoDtAllOf) UnsetResiliencyInfo()`

UnsetResiliencyInfo ensures that no value is present for ResiliencyInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


