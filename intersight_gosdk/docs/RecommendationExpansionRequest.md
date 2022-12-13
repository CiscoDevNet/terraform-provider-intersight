# RecommendationExpansionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "recommendation.HardwareExpansionRequest"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "recommendation.HardwareExpansionRequest"]
**Name** | Pointer to **string** | The name of the expansion specification. | [optional] [readonly] 
**RequestTime** | Pointer to **time.Time** | The time when the expansion was requested. | [optional] [readonly] 

## Methods

### NewRecommendationExpansionRequest

`func NewRecommendationExpansionRequest(classId string, objectType string, ) *RecommendationExpansionRequest`

NewRecommendationExpansionRequest instantiates a new RecommendationExpansionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationExpansionRequestWithDefaults

`func NewRecommendationExpansionRequestWithDefaults() *RecommendationExpansionRequest`

NewRecommendationExpansionRequestWithDefaults instantiates a new RecommendationExpansionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *RecommendationExpansionRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *RecommendationExpansionRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *RecommendationExpansionRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *RecommendationExpansionRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RecommendationExpansionRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RecommendationExpansionRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *RecommendationExpansionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecommendationExpansionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecommendationExpansionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecommendationExpansionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequestTime

`func (o *RecommendationExpansionRequest) GetRequestTime() time.Time`

GetRequestTime returns the RequestTime field if non-nil, zero value otherwise.

### GetRequestTimeOk

`func (o *RecommendationExpansionRequest) GetRequestTimeOk() (*time.Time, bool)`

GetRequestTimeOk returns a tuple with the RequestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTime

`func (o *RecommendationExpansionRequest) SetRequestTime(v time.Time)`

SetRequestTime sets RequestTime field to given value.

### HasRequestTime

`func (o *RecommendationExpansionRequest) HasRequestTime() bool`

HasRequestTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


