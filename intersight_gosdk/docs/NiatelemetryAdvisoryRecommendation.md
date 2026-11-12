# NiatelemetryAdvisoryRecommendation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.AdvisoryRecommendation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.AdvisoryRecommendation"]
**RecommendationDescription** | Pointer to **string** | Detailed guidance describing the advisory recommendation. | [optional] [readonly] 
**RecommendationTitle** | Pointer to **string** | Short title describing the advisory recommendation. | [optional] [readonly] 

## Methods

### NewNiatelemetryAdvisoryRecommendation

`func NewNiatelemetryAdvisoryRecommendation(classId string, objectType string, ) *NiatelemetryAdvisoryRecommendation`

NewNiatelemetryAdvisoryRecommendation instantiates a new NiatelemetryAdvisoryRecommendation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryAdvisoryRecommendationWithDefaults

`func NewNiatelemetryAdvisoryRecommendationWithDefaults() *NiatelemetryAdvisoryRecommendation`

NewNiatelemetryAdvisoryRecommendationWithDefaults instantiates a new NiatelemetryAdvisoryRecommendation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryAdvisoryRecommendation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryAdvisoryRecommendation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryAdvisoryRecommendation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryAdvisoryRecommendation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryAdvisoryRecommendation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryAdvisoryRecommendation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRecommendationDescription

`func (o *NiatelemetryAdvisoryRecommendation) GetRecommendationDescription() string`

GetRecommendationDescription returns the RecommendationDescription field if non-nil, zero value otherwise.

### GetRecommendationDescriptionOk

`func (o *NiatelemetryAdvisoryRecommendation) GetRecommendationDescriptionOk() (*string, bool)`

GetRecommendationDescriptionOk returns a tuple with the RecommendationDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationDescription

`func (o *NiatelemetryAdvisoryRecommendation) SetRecommendationDescription(v string)`

SetRecommendationDescription sets RecommendationDescription field to given value.

### HasRecommendationDescription

`func (o *NiatelemetryAdvisoryRecommendation) HasRecommendationDescription() bool`

HasRecommendationDescription returns a boolean if a field has been set.

### GetRecommendationTitle

`func (o *NiatelemetryAdvisoryRecommendation) GetRecommendationTitle() string`

GetRecommendationTitle returns the RecommendationTitle field if non-nil, zero value otherwise.

### GetRecommendationTitleOk

`func (o *NiatelemetryAdvisoryRecommendation) GetRecommendationTitleOk() (*string, bool)`

GetRecommendationTitleOk returns a tuple with the RecommendationTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationTitle

`func (o *NiatelemetryAdvisoryRecommendation) SetRecommendationTitle(v string)`

SetRecommendationTitle sets RecommendationTitle field to given value.

### HasRecommendationTitle

`func (o *NiatelemetryAdvisoryRecommendation) HasRecommendationTitle() bool`

HasRecommendationTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


