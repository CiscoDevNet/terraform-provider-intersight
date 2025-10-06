# SearchTagItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "search.TagItem"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "search.TagItem"]
**Count** | Pointer to **int64** | The number of times this tag key has been set across all resources. | [optional] [readonly] 
**Key** | Pointer to **string** | Key of the tag from all the resources in Intersight. | [optional] [readonly] 
**ExtMoid** | Pointer to **string** | The unique identifier of the tag definition managed object. Refer to the comm.TagDefinition API. | [optional] [readonly] 
**TaggedObjectTypes** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** | Type of the tag definition. Refer to comm.TagDefinitionType API. | [optional] [readonly] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSearchTagItem

`func NewSearchTagItem(classId string, objectType string, ) *SearchTagItem`

NewSearchTagItem instantiates a new SearchTagItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchTagItemWithDefaults

`func NewSearchTagItemWithDefaults() *SearchTagItem`

NewSearchTagItemWithDefaults instantiates a new SearchTagItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SearchTagItem) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SearchTagItem) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SearchTagItem) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SearchTagItem) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SearchTagItem) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SearchTagItem) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCount

`func (o *SearchTagItem) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SearchTagItem) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SearchTagItem) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *SearchTagItem) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetKey

`func (o *SearchTagItem) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SearchTagItem) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SearchTagItem) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SearchTagItem) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetExtMoid

`func (o *SearchTagItem) GetExtMoid() string`

GetExtMoid returns the ExtMoid field if non-nil, zero value otherwise.

### GetExtMoidOk

`func (o *SearchTagItem) GetExtMoidOk() (*string, bool)`

GetExtMoidOk returns a tuple with the ExtMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtMoid

`func (o *SearchTagItem) SetExtMoid(v string)`

SetExtMoid sets ExtMoid field to given value.

### HasExtMoid

`func (o *SearchTagItem) HasExtMoid() bool`

HasExtMoid returns a boolean if a field has been set.

### GetTaggedObjectTypes

`func (o *SearchTagItem) GetTaggedObjectTypes() []string`

GetTaggedObjectTypes returns the TaggedObjectTypes field if non-nil, zero value otherwise.

### GetTaggedObjectTypesOk

`func (o *SearchTagItem) GetTaggedObjectTypesOk() (*[]string, bool)`

GetTaggedObjectTypesOk returns a tuple with the TaggedObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedObjectTypes

`func (o *SearchTagItem) SetTaggedObjectTypes(v []string)`

SetTaggedObjectTypes sets TaggedObjectTypes field to given value.

### HasTaggedObjectTypes

`func (o *SearchTagItem) HasTaggedObjectTypes() bool`

HasTaggedObjectTypes returns a boolean if a field has been set.

### SetTaggedObjectTypesNil

`func (o *SearchTagItem) SetTaggedObjectTypesNil(b bool)`

 SetTaggedObjectTypesNil sets the value for TaggedObjectTypes to be an explicit nil

### UnsetTaggedObjectTypes
`func (o *SearchTagItem) UnsetTaggedObjectTypes()`

UnsetTaggedObjectTypes ensures that no value is present for TaggedObjectTypes, not even an explicit nil
### GetType

`func (o *SearchTagItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchTagItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchTagItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SearchTagItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValues

`func (o *SearchTagItem) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *SearchTagItem) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *SearchTagItem) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *SearchTagItem) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *SearchTagItem) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *SearchTagItem) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


