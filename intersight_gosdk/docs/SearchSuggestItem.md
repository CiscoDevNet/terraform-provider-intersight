# SearchSuggestItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "search.SuggestItem"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "search.SuggestItem"]
**CompleteMo** | Pointer to **bool** | Flag for returning complete objects that matched the global search criteria. | [optional] 
**Rawquery** | Pointer to **string** | Additional filter parameters for global search. You can also specify OData select fields here. Maximum Query Length is limited to 10000. | [optional] 
**Skip** | Pointer to **int64** | Starting offset for the results to be returned from external search engine. | [optional] 
**SuggestTerm** | Pointer to **string** | Main search term used for global search across all Managed Objects that has search enabled. Search Term can be up to 200 characters long. | [optional] 
**Top** | Pointer to **int64** | Maximum number of results to be returned from external search engine. | [optional] 
**Type** | Pointer to **string** | Object type filter of a Managed Object. Search will be restricted only on the specified object types.  Do not provide IndexMoTypes filter in the rawquery, if you specify values in this field. | [optional] 

## Methods

### NewSearchSuggestItem

`func NewSearchSuggestItem(classId string, objectType string, ) *SearchSuggestItem`

NewSearchSuggestItem instantiates a new SearchSuggestItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSuggestItemWithDefaults

`func NewSearchSuggestItemWithDefaults() *SearchSuggestItem`

NewSearchSuggestItemWithDefaults instantiates a new SearchSuggestItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SearchSuggestItem) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SearchSuggestItem) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SearchSuggestItem) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SearchSuggestItem) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SearchSuggestItem) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SearchSuggestItem) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCompleteMo

`func (o *SearchSuggestItem) GetCompleteMo() bool`

GetCompleteMo returns the CompleteMo field if non-nil, zero value otherwise.

### GetCompleteMoOk

`func (o *SearchSuggestItem) GetCompleteMoOk() (*bool, bool)`

GetCompleteMoOk returns a tuple with the CompleteMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteMo

`func (o *SearchSuggestItem) SetCompleteMo(v bool)`

SetCompleteMo sets CompleteMo field to given value.

### HasCompleteMo

`func (o *SearchSuggestItem) HasCompleteMo() bool`

HasCompleteMo returns a boolean if a field has been set.

### GetRawquery

`func (o *SearchSuggestItem) GetRawquery() string`

GetRawquery returns the Rawquery field if non-nil, zero value otherwise.

### GetRawqueryOk

`func (o *SearchSuggestItem) GetRawqueryOk() (*string, bool)`

GetRawqueryOk returns a tuple with the Rawquery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawquery

`func (o *SearchSuggestItem) SetRawquery(v string)`

SetRawquery sets Rawquery field to given value.

### HasRawquery

`func (o *SearchSuggestItem) HasRawquery() bool`

HasRawquery returns a boolean if a field has been set.

### GetSkip

`func (o *SearchSuggestItem) GetSkip() int64`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *SearchSuggestItem) GetSkipOk() (*int64, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *SearchSuggestItem) SetSkip(v int64)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *SearchSuggestItem) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetSuggestTerm

`func (o *SearchSuggestItem) GetSuggestTerm() string`

GetSuggestTerm returns the SuggestTerm field if non-nil, zero value otherwise.

### GetSuggestTermOk

`func (o *SearchSuggestItem) GetSuggestTermOk() (*string, bool)`

GetSuggestTermOk returns a tuple with the SuggestTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestTerm

`func (o *SearchSuggestItem) SetSuggestTerm(v string)`

SetSuggestTerm sets SuggestTerm field to given value.

### HasSuggestTerm

`func (o *SearchSuggestItem) HasSuggestTerm() bool`

HasSuggestTerm returns a boolean if a field has been set.

### GetTop

`func (o *SearchSuggestItem) GetTop() int64`

GetTop returns the Top field if non-nil, zero value otherwise.

### GetTopOk

`func (o *SearchSuggestItem) GetTopOk() (*int64, bool)`

GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop

`func (o *SearchSuggestItem) SetTop(v int64)`

SetTop sets Top field to given value.

### HasTop

`func (o *SearchSuggestItem) HasTop() bool`

HasTop returns a boolean if a field has been set.

### GetType

`func (o *SearchSuggestItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchSuggestItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchSuggestItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SearchSuggestItem) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


