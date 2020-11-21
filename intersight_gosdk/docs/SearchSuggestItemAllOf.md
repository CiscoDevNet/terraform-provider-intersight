# SearchSuggestItemAllOf

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

### NewSearchSuggestItemAllOf

`func NewSearchSuggestItemAllOf(classId string, objectType string, ) *SearchSuggestItemAllOf`

NewSearchSuggestItemAllOf instantiates a new SearchSuggestItemAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSuggestItemAllOfWithDefaults

`func NewSearchSuggestItemAllOfWithDefaults() *SearchSuggestItemAllOf`

NewSearchSuggestItemAllOfWithDefaults instantiates a new SearchSuggestItemAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SearchSuggestItemAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SearchSuggestItemAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SearchSuggestItemAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SearchSuggestItemAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SearchSuggestItemAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SearchSuggestItemAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCompleteMo

`func (o *SearchSuggestItemAllOf) GetCompleteMo() bool`

GetCompleteMo returns the CompleteMo field if non-nil, zero value otherwise.

### GetCompleteMoOk

`func (o *SearchSuggestItemAllOf) GetCompleteMoOk() (*bool, bool)`

GetCompleteMoOk returns a tuple with the CompleteMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteMo

`func (o *SearchSuggestItemAllOf) SetCompleteMo(v bool)`

SetCompleteMo sets CompleteMo field to given value.

### HasCompleteMo

`func (o *SearchSuggestItemAllOf) HasCompleteMo() bool`

HasCompleteMo returns a boolean if a field has been set.

### GetRawquery

`func (o *SearchSuggestItemAllOf) GetRawquery() string`

GetRawquery returns the Rawquery field if non-nil, zero value otherwise.

### GetRawqueryOk

`func (o *SearchSuggestItemAllOf) GetRawqueryOk() (*string, bool)`

GetRawqueryOk returns a tuple with the Rawquery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawquery

`func (o *SearchSuggestItemAllOf) SetRawquery(v string)`

SetRawquery sets Rawquery field to given value.

### HasRawquery

`func (o *SearchSuggestItemAllOf) HasRawquery() bool`

HasRawquery returns a boolean if a field has been set.

### GetSkip

`func (o *SearchSuggestItemAllOf) GetSkip() int64`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *SearchSuggestItemAllOf) GetSkipOk() (*int64, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *SearchSuggestItemAllOf) SetSkip(v int64)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *SearchSuggestItemAllOf) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetSuggestTerm

`func (o *SearchSuggestItemAllOf) GetSuggestTerm() string`

GetSuggestTerm returns the SuggestTerm field if non-nil, zero value otherwise.

### GetSuggestTermOk

`func (o *SearchSuggestItemAllOf) GetSuggestTermOk() (*string, bool)`

GetSuggestTermOk returns a tuple with the SuggestTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestTerm

`func (o *SearchSuggestItemAllOf) SetSuggestTerm(v string)`

SetSuggestTerm sets SuggestTerm field to given value.

### HasSuggestTerm

`func (o *SearchSuggestItemAllOf) HasSuggestTerm() bool`

HasSuggestTerm returns a boolean if a field has been set.

### GetTop

`func (o *SearchSuggestItemAllOf) GetTop() int64`

GetTop returns the Top field if non-nil, zero value otherwise.

### GetTopOk

`func (o *SearchSuggestItemAllOf) GetTopOk() (*int64, bool)`

GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop

`func (o *SearchSuggestItemAllOf) SetTop(v int64)`

SetTop sets Top field to given value.

### HasTop

`func (o *SearchSuggestItemAllOf) HasTop() bool`

HasTop returns a boolean if a field has been set.

### GetType

`func (o *SearchSuggestItemAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchSuggestItemAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchSuggestItemAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SearchSuggestItemAllOf) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


