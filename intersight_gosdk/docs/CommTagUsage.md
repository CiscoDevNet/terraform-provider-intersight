# CommTagUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "comm.TagUsage"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "comm.TagUsage"]
**ValueUsages** | Pointer to [**[]CommTagValueUsage**](CommTagValueUsage.md) |  | [optional] 

## Methods

### NewCommTagUsage

`func NewCommTagUsage(classId string, objectType string, ) *CommTagUsage`

NewCommTagUsage instantiates a new CommTagUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommTagUsageWithDefaults

`func NewCommTagUsageWithDefaults() *CommTagUsage`

NewCommTagUsageWithDefaults instantiates a new CommTagUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CommTagUsage) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CommTagUsage) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CommTagUsage) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CommTagUsage) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CommTagUsage) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CommTagUsage) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetValueUsages

`func (o *CommTagUsage) GetValueUsages() []CommTagValueUsage`

GetValueUsages returns the ValueUsages field if non-nil, zero value otherwise.

### GetValueUsagesOk

`func (o *CommTagUsage) GetValueUsagesOk() (*[]CommTagValueUsage, bool)`

GetValueUsagesOk returns a tuple with the ValueUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueUsages

`func (o *CommTagUsage) SetValueUsages(v []CommTagValueUsage)`

SetValueUsages sets ValueUsages field to given value.

### HasValueUsages

`func (o *CommTagUsage) HasValueUsages() bool`

HasValueUsages returns a boolean if a field has been set.

### SetValueUsagesNil

`func (o *CommTagUsage) SetValueUsagesNil(b bool)`

 SetValueUsagesNil sets the value for ValueUsages to be an explicit nil

### UnsetValueUsages
`func (o *CommTagUsage) UnsetValueUsages()`

UnsetValueUsages ensures that no value is present for ValueUsages, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


