# TamApiDataSourceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "tam.ApiDataSource"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "tam.ApiDataSource"]
**MoType** | Pointer to **string** | Type of Intersight managed object used as data source. | [optional] 
**Queries** | Pointer to [**[]TamQueryEntry**](TamQueryEntry.md) |  | [optional] 

## Methods

### NewTamApiDataSourceAllOf

`func NewTamApiDataSourceAllOf(classId string, objectType string, ) *TamApiDataSourceAllOf`

NewTamApiDataSourceAllOf instantiates a new TamApiDataSourceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamApiDataSourceAllOfWithDefaults

`func NewTamApiDataSourceAllOfWithDefaults() *TamApiDataSourceAllOf`

NewTamApiDataSourceAllOfWithDefaults instantiates a new TamApiDataSourceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TamApiDataSourceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TamApiDataSourceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TamApiDataSourceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TamApiDataSourceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TamApiDataSourceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TamApiDataSourceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMoType

`func (o *TamApiDataSourceAllOf) GetMoType() string`

GetMoType returns the MoType field if non-nil, zero value otherwise.

### GetMoTypeOk

`func (o *TamApiDataSourceAllOf) GetMoTypeOk() (*string, bool)`

GetMoTypeOk returns a tuple with the MoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoType

`func (o *TamApiDataSourceAllOf) SetMoType(v string)`

SetMoType sets MoType field to given value.

### HasMoType

`func (o *TamApiDataSourceAllOf) HasMoType() bool`

HasMoType returns a boolean if a field has been set.

### GetQueries

`func (o *TamApiDataSourceAllOf) GetQueries() []TamQueryEntry`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *TamApiDataSourceAllOf) GetQueriesOk() (*[]TamQueryEntry, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *TamApiDataSourceAllOf) SetQueries(v []TamQueryEntry)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *TamApiDataSourceAllOf) HasQueries() bool`

HasQueries returns a boolean if a field has been set.

### SetQueriesNil

`func (o *TamApiDataSourceAllOf) SetQueriesNil(b bool)`

 SetQueriesNil sets the value for Queries to be an explicit nil

### UnsetQueries
`func (o *TamApiDataSourceAllOf) UnsetQueries()`

UnsetQueries ensures that no value is present for Queries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


