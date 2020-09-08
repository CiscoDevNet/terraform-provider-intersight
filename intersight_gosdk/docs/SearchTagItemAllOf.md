# SearchTagItemAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | The number of times this tag key has been set across all resources. | [optional] [readonly] 
**Key** | Pointer to **string** | Key of the Tag from all the resources in Intersight. | [optional] [readonly] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSearchTagItemAllOf

`func NewSearchTagItemAllOf() *SearchTagItemAllOf`

NewSearchTagItemAllOf instantiates a new SearchTagItemAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchTagItemAllOfWithDefaults

`func NewSearchTagItemAllOfWithDefaults() *SearchTagItemAllOf`

NewSearchTagItemAllOfWithDefaults instantiates a new SearchTagItemAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


