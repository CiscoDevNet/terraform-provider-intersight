# HyperflexHxSiteDtAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HxSiteDt"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HxSiteDt"]
**Name** | Pointer to **string** | Name of the site for this HyperFlex cluster. | [optional] [readonly] 
**Zone** | Pointer to [**NullableHyperflexHxZoneInfoDt**](HyperflexHxZoneInfoDt.md) |  | [optional] 

## Methods

### NewHyperflexHxSiteDtAllOf

`func NewHyperflexHxSiteDtAllOf(classId string, objectType string, ) *HyperflexHxSiteDtAllOf`

NewHyperflexHxSiteDtAllOf instantiates a new HyperflexHxSiteDtAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxSiteDtAllOfWithDefaults

`func NewHyperflexHxSiteDtAllOfWithDefaults() *HyperflexHxSiteDtAllOf`

NewHyperflexHxSiteDtAllOfWithDefaults instantiates a new HyperflexHxSiteDtAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxSiteDtAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxSiteDtAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxSiteDtAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxSiteDtAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxSiteDtAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxSiteDtAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *HyperflexHxSiteDtAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexHxSiteDtAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexHxSiteDtAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexHxSiteDtAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetZone

`func (o *HyperflexHxSiteDtAllOf) GetZone() HyperflexHxZoneInfoDt`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *HyperflexHxSiteDtAllOf) GetZoneOk() (*HyperflexHxZoneInfoDt, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *HyperflexHxSiteDtAllOf) SetZone(v HyperflexHxZoneInfoDt)`

SetZone sets Zone field to given value.

### HasZone

`func (o *HyperflexHxSiteDtAllOf) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *HyperflexHxSiteDtAllOf) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *HyperflexHxSiteDtAllOf) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


