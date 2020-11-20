# SearchTagItemAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "search.TagItem"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "search.TagItem"]
**Count** | Pointer to **int64** | The number of times this tag key has been set across all resources. | [optional] [readonly] 
**Key** | Pointer to **string** | Key of the Tag from all the resources in Intersight. | [optional] [readonly] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSearchTagItemAllOf

`func NewSearchTagItemAllOf(classId string, objectType string, ) *SearchTagItemAllOf`

NewSearchTagItemAllOf instantiates a new SearchTagItemAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchTagItemAllOfWithDefaults

`func NewSearchTagItemAllOfWithDefaults() *SearchTagItemAllOf`

NewSearchTagItemAllOfWithDefaults instantiates a new SearchTagItemAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SearchTagItemAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SearchTagItemAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SearchTagItemAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SearchTagItemAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SearchTagItemAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SearchTagItemAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCount

`func (o *SearchTagItemAllOf) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SearchTagItemAllOf) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SearchTagItemAllOf) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *SearchTagItemAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetKey

`func (o *SearchTagItemAllOf) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SearchTagItemAllOf) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SearchTagItemAllOf) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SearchTagItemAllOf) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValues

`func (o *SearchTagItemAllOf) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *SearchTagItemAllOf) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *SearchTagItemAllOf) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *SearchTagItemAllOf) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *SearchTagItemAllOf) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *SearchTagItemAllOf) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


