# TamEolAdvisoryDetailsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "tam.EolAdvisoryDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "tam.EolAdvisoryDetails"]
**AllMilestones** | Pointer to [**[]TamMilestone**](TamMilestone.md) |  | [optional] 
**Milestone** | Pointer to [**NullableTamMilestone**](TamMilestone.md) |  | [optional] 
**Release** | Pointer to **string** | The name of the impacted release this advisory milestone intends to address, e.g. \&quot;3.5 (2x)\&quot;. | [optional] 

## Methods

### NewTamEolAdvisoryDetailsAllOf

`func NewTamEolAdvisoryDetailsAllOf(classId string, objectType string, ) *TamEolAdvisoryDetailsAllOf`

NewTamEolAdvisoryDetailsAllOf instantiates a new TamEolAdvisoryDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamEolAdvisoryDetailsAllOfWithDefaults

`func NewTamEolAdvisoryDetailsAllOfWithDefaults() *TamEolAdvisoryDetailsAllOf`

NewTamEolAdvisoryDetailsAllOfWithDefaults instantiates a new TamEolAdvisoryDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TamEolAdvisoryDetailsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TamEolAdvisoryDetailsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TamEolAdvisoryDetailsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TamEolAdvisoryDetailsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TamEolAdvisoryDetailsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TamEolAdvisoryDetailsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAllMilestones

`func (o *TamEolAdvisoryDetailsAllOf) GetAllMilestones() []TamMilestone`

GetAllMilestones returns the AllMilestones field if non-nil, zero value otherwise.

### GetAllMilestonesOk

`func (o *TamEolAdvisoryDetailsAllOf) GetAllMilestonesOk() (*[]TamMilestone, bool)`

GetAllMilestonesOk returns a tuple with the AllMilestones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllMilestones

`func (o *TamEolAdvisoryDetailsAllOf) SetAllMilestones(v []TamMilestone)`

SetAllMilestones sets AllMilestones field to given value.

### HasAllMilestones

`func (o *TamEolAdvisoryDetailsAllOf) HasAllMilestones() bool`

HasAllMilestones returns a boolean if a field has been set.

### SetAllMilestonesNil

`func (o *TamEolAdvisoryDetailsAllOf) SetAllMilestonesNil(b bool)`

 SetAllMilestonesNil sets the value for AllMilestones to be an explicit nil

### UnsetAllMilestones
`func (o *TamEolAdvisoryDetailsAllOf) UnsetAllMilestones()`

UnsetAllMilestones ensures that no value is present for AllMilestones, not even an explicit nil
### GetMilestone

`func (o *TamEolAdvisoryDetailsAllOf) GetMilestone() TamMilestone`

GetMilestone returns the Milestone field if non-nil, zero value otherwise.

### GetMilestoneOk

`func (o *TamEolAdvisoryDetailsAllOf) GetMilestoneOk() (*TamMilestone, bool)`

GetMilestoneOk returns a tuple with the Milestone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestone

`func (o *TamEolAdvisoryDetailsAllOf) SetMilestone(v TamMilestone)`

SetMilestone sets Milestone field to given value.

### HasMilestone

`func (o *TamEolAdvisoryDetailsAllOf) HasMilestone() bool`

HasMilestone returns a boolean if a field has been set.

### SetMilestoneNil

`func (o *TamEolAdvisoryDetailsAllOf) SetMilestoneNil(b bool)`

 SetMilestoneNil sets the value for Milestone to be an explicit nil

### UnsetMilestone
`func (o *TamEolAdvisoryDetailsAllOf) UnsetMilestone()`

UnsetMilestone ensures that no value is present for Milestone, not even an explicit nil
### GetRelease

`func (o *TamEolAdvisoryDetailsAllOf) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *TamEolAdvisoryDetailsAllOf) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *TamEolAdvisoryDetailsAllOf) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *TamEolAdvisoryDetailsAllOf) HasRelease() bool`

HasRelease returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


