# TamQueryEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source. | [optional] 
**Priority** | Pointer to **int64** | An integer value depicting the priority of the query among the queries that are part of the same QueryEntry collection. | [optional] 
**Query** | Pointer to **string** | A SparkSQL query to be used on a given data source. | [optional] 

## Methods

### NewTamQueryEntry

`func NewTamQueryEntry() *TamQueryEntry`

NewTamQueryEntry instantiates a new TamQueryEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamQueryEntryWithDefaults

`func NewTamQueryEntryWithDefaults() *TamQueryEntry`

NewTamQueryEntryWithDefaults instantiates a new TamQueryEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TamQueryEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TamQueryEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TamQueryEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TamQueryEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *TamQueryEntry) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TamQueryEntry) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TamQueryEntry) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TamQueryEntry) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQuery

`func (o *TamQueryEntry) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *TamQueryEntry) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *TamQueryEntry) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *TamQueryEntry) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


