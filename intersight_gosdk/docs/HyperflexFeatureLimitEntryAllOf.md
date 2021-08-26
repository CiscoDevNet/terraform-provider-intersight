# HyperflexFeatureLimitEntryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.FeatureLimitEntry"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.FeatureLimitEntry"]
**Constraint** | Pointer to [**NullableHyperflexAppSettingConstraint**](HyperflexAppSettingConstraint.md) |  | [optional] 

## Methods

### NewHyperflexFeatureLimitEntryAllOf

`func NewHyperflexFeatureLimitEntryAllOf(classId string, objectType string, ) *HyperflexFeatureLimitEntryAllOf`

NewHyperflexFeatureLimitEntryAllOf instantiates a new HyperflexFeatureLimitEntryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexFeatureLimitEntryAllOfWithDefaults

`func NewHyperflexFeatureLimitEntryAllOfWithDefaults() *HyperflexFeatureLimitEntryAllOf`

NewHyperflexFeatureLimitEntryAllOfWithDefaults instantiates a new HyperflexFeatureLimitEntryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexFeatureLimitEntryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexFeatureLimitEntryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexFeatureLimitEntryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexFeatureLimitEntryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexFeatureLimitEntryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexFeatureLimitEntryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConstraint

`func (o *HyperflexFeatureLimitEntryAllOf) GetConstraint() HyperflexAppSettingConstraint`

GetConstraint returns the Constraint field if non-nil, zero value otherwise.

### GetConstraintOk

`func (o *HyperflexFeatureLimitEntryAllOf) GetConstraintOk() (*HyperflexAppSettingConstraint, bool)`

GetConstraintOk returns a tuple with the Constraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraint

`func (o *HyperflexFeatureLimitEntryAllOf) SetConstraint(v HyperflexAppSettingConstraint)`

SetConstraint sets Constraint field to given value.

### HasConstraint

`func (o *HyperflexFeatureLimitEntryAllOf) HasConstraint() bool`

HasConstraint returns a boolean if a field has been set.

### SetConstraintNil

`func (o *HyperflexFeatureLimitEntryAllOf) SetConstraintNil(b bool)`

 SetConstraintNil sets the value for Constraint to be an explicit nil

### UnsetConstraint
`func (o *HyperflexFeatureLimitEntryAllOf) UnsetConstraint()`

UnsetConstraint ensures that no value is present for Constraint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


