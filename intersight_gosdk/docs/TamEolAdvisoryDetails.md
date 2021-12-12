# TamEolAdvisoryDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "tam.EolAdvisoryDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "tam.EolAdvisoryDetails"]
**AllMilestones** | Pointer to [**[]TamMilestone**](TamMilestone.md) |  | [optional] 
**Milestone** | Pointer to [**NullableTamMilestone**](TamMilestone.md) |  | [optional] 
**Release** | Pointer to **string** | The name of the impacted release this advisory milestone intends to address, e.g. \&quot;3.5 (2x)\&quot;. | [optional] 

## Methods

### NewTamEolAdvisoryDetails

`func NewTamEolAdvisoryDetails(classId string, objectType string, ) *TamEolAdvisoryDetails`

NewTamEolAdvisoryDetails instantiates a new TamEolAdvisoryDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamEolAdvisoryDetailsWithDefaults

`func NewTamEolAdvisoryDetailsWithDefaults() *TamEolAdvisoryDetails`

NewTamEolAdvisoryDetailsWithDefaults instantiates a new TamEolAdvisoryDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TamEolAdvisoryDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TamEolAdvisoryDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TamEolAdvisoryDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TamEolAdvisoryDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TamEolAdvisoryDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TamEolAdvisoryDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAllMilestones

`func (o *TamEolAdvisoryDetails) GetAllMilestones() []TamMilestone`

GetAllMilestones returns the AllMilestones field if non-nil, zero value otherwise.

### GetAllMilestonesOk

`func (o *TamEolAdvisoryDetails) GetAllMilestonesOk() (*[]TamMilestone, bool)`

GetAllMilestonesOk returns a tuple with the AllMilestones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllMilestones

`func (o *TamEolAdvisoryDetails) SetAllMilestones(v []TamMilestone)`

SetAllMilestones sets AllMilestones field to given value.

### HasAllMilestones

`func (o *TamEolAdvisoryDetails) HasAllMilestones() bool`

HasAllMilestones returns a boolean if a field has been set.

### SetAllMilestonesNil

`func (o *TamEolAdvisoryDetails) SetAllMilestonesNil(b bool)`

 SetAllMilestonesNil sets the value for AllMilestones to be an explicit nil

### UnsetAllMilestones
`func (o *TamEolAdvisoryDetails) UnsetAllMilestones()`

UnsetAllMilestones ensures that no value is present for AllMilestones, not even an explicit nil
### GetMilestone

`func (o *TamEolAdvisoryDetails) GetMilestone() TamMilestone`

GetMilestone returns the Milestone field if non-nil, zero value otherwise.

### GetMilestoneOk

`func (o *TamEolAdvisoryDetails) GetMilestoneOk() (*TamMilestone, bool)`

GetMilestoneOk returns a tuple with the Milestone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestone

`func (o *TamEolAdvisoryDetails) SetMilestone(v TamMilestone)`

SetMilestone sets Milestone field to given value.

### HasMilestone

`func (o *TamEolAdvisoryDetails) HasMilestone() bool`

HasMilestone returns a boolean if a field has been set.

### SetMilestoneNil

`func (o *TamEolAdvisoryDetails) SetMilestoneNil(b bool)`

 SetMilestoneNil sets the value for Milestone to be an explicit nil

### UnsetMilestone
`func (o *TamEolAdvisoryDetails) UnsetMilestone()`

UnsetMilestone ensures that no value is present for Milestone, not even an explicit nil
### GetRelease

`func (o *TamEolAdvisoryDetails) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *TamEolAdvisoryDetails) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *TamEolAdvisoryDetails) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *TamEolAdvisoryDetails) HasRelease() bool`

HasRelease returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


