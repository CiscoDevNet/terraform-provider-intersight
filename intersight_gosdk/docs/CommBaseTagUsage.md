# CommBaseTagUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Count** | Pointer to **int64** | Total count of managed objects associated with this tag usage. | [optional] [readonly] 
**TaggedObjectTypes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCommBaseTagUsage

`func NewCommBaseTagUsage(classId string, objectType string, ) *CommBaseTagUsage`

NewCommBaseTagUsage instantiates a new CommBaseTagUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommBaseTagUsageWithDefaults

`func NewCommBaseTagUsageWithDefaults() *CommBaseTagUsage`

NewCommBaseTagUsageWithDefaults instantiates a new CommBaseTagUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CommBaseTagUsage) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CommBaseTagUsage) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CommBaseTagUsage) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CommBaseTagUsage) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CommBaseTagUsage) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CommBaseTagUsage) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCount

`func (o *CommBaseTagUsage) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CommBaseTagUsage) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CommBaseTagUsage) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *CommBaseTagUsage) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTaggedObjectTypes

`func (o *CommBaseTagUsage) GetTaggedObjectTypes() []string`

GetTaggedObjectTypes returns the TaggedObjectTypes field if non-nil, zero value otherwise.

### GetTaggedObjectTypesOk

`func (o *CommBaseTagUsage) GetTaggedObjectTypesOk() (*[]string, bool)`

GetTaggedObjectTypesOk returns a tuple with the TaggedObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedObjectTypes

`func (o *CommBaseTagUsage) SetTaggedObjectTypes(v []string)`

SetTaggedObjectTypes sets TaggedObjectTypes field to given value.

### HasTaggedObjectTypes

`func (o *CommBaseTagUsage) HasTaggedObjectTypes() bool`

HasTaggedObjectTypes returns a boolean if a field has been set.

### SetTaggedObjectTypesNil

`func (o *CommBaseTagUsage) SetTaggedObjectTypesNil(b bool)`

 SetTaggedObjectTypesNil sets the value for TaggedObjectTypes to be an explicit nil

### UnsetTaggedObjectTypes
`func (o *CommBaseTagUsage) UnsetTaggedObjectTypes()`

UnsetTaggedObjectTypes ensures that no value is present for TaggedObjectTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


