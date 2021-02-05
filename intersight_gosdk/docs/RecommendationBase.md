# RecommendationBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "recommendation.CapacityRunway"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "recommendation.CapacityRunway"]
**LastUpdatedTime** | Pointer to **time.Time** | The time when the recommendation was last updated. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the recommendation. | [optional] [readonly] 
**RequirementMet** | Pointer to **bool** | Indicates if the recommendation requirement is met by the existing setup by adding hardware components to it or it needs expansion. For example if the recommendation is to add 16 disks to a HyperFlex cluster but the cluster is already alost full and only 8 disks can be added, then this property is set to false. | [optional] [readonly] 

## Methods

### NewRecommendationBase

`func NewRecommendationBase(classId string, objectType string, ) *RecommendationBase`

NewRecommendationBase instantiates a new RecommendationBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationBaseWithDefaults

`func NewRecommendationBaseWithDefaults() *RecommendationBase`

NewRecommendationBaseWithDefaults instantiates a new RecommendationBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *RecommendationBase) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *RecommendationBase) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *RecommendationBase) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *RecommendationBase) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RecommendationBase) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RecommendationBase) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLastUpdatedTime

`func (o *RecommendationBase) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *RecommendationBase) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *RecommendationBase) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *RecommendationBase) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetName

`func (o *RecommendationBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecommendationBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecommendationBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecommendationBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequirementMet

`func (o *RecommendationBase) GetRequirementMet() bool`

GetRequirementMet returns the RequirementMet field if non-nil, zero value otherwise.

### GetRequirementMetOk

`func (o *RecommendationBase) GetRequirementMetOk() (*bool, bool)`

GetRequirementMetOk returns a tuple with the RequirementMet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementMet

`func (o *RecommendationBase) SetRequirementMet(v bool)`

SetRequirementMet sets RequirementMet field to given value.

### HasRequirementMet

`func (o *RecommendationBase) HasRequirementMet() bool`

HasRequirementMet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


