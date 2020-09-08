# TamApiDataSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoType** | Pointer to **string** | Type of Intersight managed object used as data source. | [optional] 
**Queries** | Pointer to [**[]TamQueryEntry**](tam.QueryEntry.md) |  | [optional] 

## Methods

### NewTamApiDataSource

`func NewTamApiDataSource() *TamApiDataSource`

NewTamApiDataSource instantiates a new TamApiDataSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamApiDataSourceWithDefaults

`func NewTamApiDataSourceWithDefaults() *TamApiDataSource`

NewTamApiDataSourceWithDefaults instantiates a new TamApiDataSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoType

`func (o *TamApiDataSource) GetMoType() string`

GetMoType returns the MoType field if non-nil, zero value otherwise.

### GetMoTypeOk

`func (o *TamApiDataSource) GetMoTypeOk() (*string, bool)`

GetMoTypeOk returns a tuple with the MoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoType

`func (o *TamApiDataSource) SetMoType(v string)`

SetMoType sets MoType field to given value.

### HasMoType

`func (o *TamApiDataSource) HasMoType() bool`

HasMoType returns a boolean if a field has been set.

### GetQueries

`func (o *TamApiDataSource) GetQueries() []TamQueryEntry`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *TamApiDataSource) GetQueriesOk() (*[]TamQueryEntry, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *TamApiDataSource) SetQueries(v []TamQueryEntry)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *TamApiDataSource) HasQueries() bool`

HasQueries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


